package database

import (
	"database/sql"
	"time"
)

type Developer struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

type Service interface {
	Health() map[string]string
	Close() error
	RunMigrations() error

	CreateDeveloper(name, email string) (string, error)
	GetDeveloper(id string) (*Developer, error)
	UpdateDeveloper(id, name, email string) error
	DeleteDeveloper(id string) error
}

type service struct {
	db *sql.DB
}
