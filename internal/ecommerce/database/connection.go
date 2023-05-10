package database

import (
	"fmt"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Connect(dsn string) (*bun.DB, error) {
	// Replace with your database connection details

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Create a Bun DB instance using the sql.DB instance
	db := bun.NewDB(sqlDB, pgdialect.New())


	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	return db, nil
}

