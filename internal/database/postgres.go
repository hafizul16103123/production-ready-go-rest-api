package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/config"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {

	db, err := sql.Open(
		"pgx",
		cfg.DatabaseURL,
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)

	if err := migrate(ctx, db); err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(ctx context.Context, db *sql.DB) error {
	const query = `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		age INT NOT NULL
	)
	`

	_, err := db.ExecContext(ctx, query)
	return err
}