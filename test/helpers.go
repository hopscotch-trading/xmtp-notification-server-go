package test

import (
	"context"
	"testing"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/extra/bundebug"
	database "github.com/xmtp/example-notification-server-go/pkg/db"
)

const TEST_DSN = "postgres://postgres:xmtp@localhost:25432/postgres?sslmode=disable"

func createDb() *bun.DB {
	db, _ := database.CreateBunDB(TEST_DSN, 5*time.Second)
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	return db
}

func CreateTestDb(t *testing.T) *bun.DB {
	ctx := t.Context()
	db := createDb()
	if err := database.Migrate(ctx, db); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		cleanupCtx := context.Background()
		_, _ = db.NewTruncateTable().Model((*database.Installation)(nil)).Cascade().Exec(cleanupCtx)
		_, _ = db.NewTruncateTable().Model((*database.DeviceDeliveryMechanism)(nil)).Cascade().Exec(cleanupCtx)
		_, _ = db.NewTruncateTable().Model((*database.Subscription)(nil)).Cascade().Exec(cleanupCtx)
		_, _ = db.NewTruncateTable().Model((*database.SubscriptionHmacKeys)(nil)).Cascade().Exec(cleanupCtx)
	})

	return db
}
