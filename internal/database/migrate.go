package database

import (
	"fmt"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/pressly/goose/v3"
)

func (s *service) RunMigrations() error {
	databaseURL := os.Getenv("DB_CONN_STR")

	migrationPath := "internal/database/migrations"

	db, err := goose.OpenDBWithDriver("postgres", databaseURL)
	if err != nil {
		return fmt.Errorf("could not open database: %v", err)
	}
	defer db.Close()

	if err := goose.Up(db, migrationPath); err != nil {
		return fmt.Errorf("could not apply migrations: %v", err)
	}

	fmt.Println("Migrations applied successfully.")
	return nil
}
