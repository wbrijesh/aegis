package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func (s *service) developerEmailExists(ctx context.Context, email string) (bool, error) {
	query := `SELECT 1 FROM developers WHERE email = $1`
	row := s.db.QueryRowContext(ctx, query, email)

	var exists int
	if err := row.Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if email exists: %v", err)
	}

	return true, nil
}

func (s *service) CreateDeveloper(name, email string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exists, err := s.developerEmailExists(ctx, email)
	if err != nil {
		return "", fmt.Errorf("failed to check if email exists: %v", err)
	}
	if exists {
		return "", fmt.Errorf("email already in use")
	}

	id := uuid.New().String()
	query := `INSERT INTO developers (id, name, email) VALUES ($1, $2, $3)`

	_, err = s.db.ExecContext(ctx, query, id, name, email)
	if err != nil {
		return "", fmt.Errorf("failed to create developer: %v", err)
	}

	return id, nil
}

func (s *service) GetDeveloper(id string) (*Developer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, name, email, created_at FROM developers WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)

	var developer Developer
	if err := row.Scan(&developer.ID, &developer.Name, &developer.Email, &developer.CreatedAt); err != nil {
		return nil, fmt.Errorf("failed to get developer: %v", err)
	}

	return &developer, nil
}

func (s *service) UpdateDeveloper(id, name, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exists, err := s.developerEmailExists(ctx, email)
	if err != nil {
		return fmt.Errorf("failed to check if email exists: %v", err)
	}
	if exists {
		return fmt.Errorf("email already in use")
	}

	query := `UPDATE developers SET name = $2, email = $3 WHERE id = $1`
	_, err = s.db.ExecContext(ctx, query, id, name, email)
	if err != nil {
		return fmt.Errorf("failed to update developer: %v", err)
	}

	return nil
}

func (s *service) DeleteDeveloper(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM developers WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete developer: %v", err)
	}

	return nil
}
