package database

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}
