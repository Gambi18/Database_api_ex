package repo

import (
	"errors"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // File source for migrations
)

func newMigrator(dbURL, migrationsPath string) (*migrate.Migrate, error) {
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return nil, err
	}

	// Normalize to forward slashes for file URL compatibility
	absPath = filepath.ToSlash(absPath)

	m, err := migrate.New(
		"file://"+absPath,
		dbURL,
	)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Migrate function applies migrations to the database.
func Migrate(dbURL string, migrationsPath string) error {
	m, err := newMigrator(dbURL, migrationsPath)
	if err != nil {
		return err
	}
	defer m.Close()

	// Apply migrations
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

// MigrateDown function rolls back migrations from the database.
func MigrateDown(dbURL string, migrationsPath string) error {
	m, err := newMigrator(dbURL, migrationsPath)
	if err != nil {
		return err
	}
	defer m.Close()

	// Apply migrations
	err = m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
