package database

import (
	"rest-api/internal/handlers"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// RunMigrations executes DB migrations
func RunMigrations() error {

	driver, err := postgres.WithInstance(handlers.DB, &postgres.Config{})
	if err != nil {
		return nil
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../internal/database/migrations",
		"postgres",
		driver)
	if err != nil {
		return err
	}

	m.Steps(2)

	return nil
}
