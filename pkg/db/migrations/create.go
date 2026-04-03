package migrations

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Path string
}

var migrationNamePattern = regexp.MustCompile(`[^a-z0-9]+`)

func CreateFiles(dir string, name string) ([]File, error) {
	safeName := sanitizeName(name)
	if safeName == "" {
		return nil, fmt.Errorf("invalid migration name %q", name)
	}

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}

	next, err := nextSequence(dir)
	if err != nil {
		return nil, err
	}

	base := fmt.Sprintf("%05d_%s", next, safeName)
	paths := []File{
		{Name: base + ".up.sql", Path: filepath.Join(dir, base+".up.sql")},
		{Name: base + ".down.sql", Path: filepath.Join(dir, base+".down.sql")},
	}

	for _, file := range paths {
		if err := os.WriteFile(file.Path, []byte(""), 0o644); err != nil {
			return nil, err
		}
	}

	return paths, nil
}

func nextSequence(dir string) (int, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	maxSequence := 0
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".up.sql") || len(name) < 5 {
			continue
		}

		seq, err := strconv.Atoi(name[:5])
		if err != nil {
			continue
		}
		if seq > maxSequence {
			maxSequence = seq
		}
	}

	return maxSequence + 1, nil
}

func sanitizeName(name string) string {
	lower := strings.ToLower(strings.TrimSpace(name))
	safe := migrationNamePattern.ReplaceAllString(lower, "-")
	safe = strings.Trim(safe, "-")
	if safe == "" {
		return ""
	}

	parts := strings.Split(safe, "-")
	parts = slices.DeleteFunc(parts, func(part string) bool { return part == "" })
	return strings.Join(parts, "-")
}
