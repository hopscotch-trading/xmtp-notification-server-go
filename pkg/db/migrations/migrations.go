package migrations

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed *.sql
var migrationFS embed.FS

func Migrate(ctx context.Context, db *sql.DB) error {
	if err := reconcileExistingBunSchema(ctx, db); err != nil {
		return err
	}

	sourceDriver, err := iofs.New(migrationFS, ".")
	if err != nil {
		return err
	}
	defer func() {
		_ = sourceDriver.Close()
	}()

	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()

	driver, err := postgres.WithConnection(ctx, conn, &postgres.Config{})
	if err != nil {
		return err
	}
	defer func() {
		_ = driver.Close()
	}()

	migrator, err := migrate.NewWithInstance("iofs", sourceDriver, "postgres", driver)
	if err != nil {
		return err
	}
	defer func() {
		_, _ = migrator.Close()
	}()

	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func reconcileExistingBunSchema(ctx context.Context, db *sql.DB) error {
	alreadyTracked, err := hasSchemaMigrationState(ctx, db)
	if err != nil {
		return err
	}
	if alreadyTracked {
		return nil
	}

	exists, err := hasLegacySchema(ctx, db)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}

	version, err := latestMigrationVersion()
	if err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version bigint NOT NULL PRIMARY KEY,
			dirty boolean NOT NULL
		)`); err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, `
		INSERT INTO schema_migrations (version, dirty)
		VALUES ($1, FALSE)
		ON CONFLICT (version) DO UPDATE SET dirty = EXCLUDED.dirty
	`, version)
	return err
}

func hasSchemaMigrationState(ctx context.Context, db *sql.DB) (bool, error) {
	var tableExists bool
	if err := db.QueryRowContext(ctx, `SELECT to_regclass('public.schema_migrations') IS NOT NULL`).Scan(&tableExists); err != nil {
		return false, err
	}
	if !tableExists {
		return false, nil
	}

	var count int
	if err := db.QueryRowContext(ctx, `SELECT COUNT(*) FROM schema_migrations`).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func hasLegacySchema(ctx context.Context, db *sql.DB) (bool, error) {
	checks := []string{
		`SELECT to_regclass('public.installations') IS NOT NULL`,
		`SELECT to_regclass('public.device_delivery_mechanisms') IS NOT NULL`,
		`SELECT to_regclass('public.subscriptions') IS NOT NULL`,
		`SELECT EXISTS (
			SELECT 1
			FROM information_schema.columns
			WHERE table_schema = 'public'
			  AND table_name = 'subscriptions'
			  AND column_name = 'is_silent'
		)`,
		`SELECT to_regclass('public.subscriptions_installation_id_topic_idx') IS NOT NULL`,
		`SELECT to_regclass('public.subscription_hmac_keys') IS NOT NULL`,
	}

	for _, query := range checks {
		var ok bool
		if err := db.QueryRowContext(ctx, query).Scan(&ok); err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}

	return true, nil
}

func latestMigrationVersion() (uint, error) {
	entries, err := migrationFS.ReadDir(".")
	if err != nil {
		return 0, err
	}

	versions := make([]uint, 0, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".up.sql") {
			continue
		}

		base := filepath.Base(name)
		parts := strings.SplitN(base, "_", 2)
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid migration filename %q", name)
		}

		version, err := strconv.ParseUint(parts[0], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("parse migration version %q: %w", name, err)
		}
		versions = append(versions, uint(version))
	}

	if len(versions) == 0 {
		return 0, errors.New("no embedded up migrations found")
	}

	slices.Sort(versions)
	return versions[len(versions)-1], nil
}
