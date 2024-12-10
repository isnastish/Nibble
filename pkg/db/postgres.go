package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	connPool *pgxpool.Pool
}

func NewPostgresDB() (*PostgresDB, error) {
	postgresUrl, set := os.LookupEnv("POSTGRES_URL")
	if !set || postgresUrl == "" {
		return nil, fmt.Errorf("postgres: POSTGRES_URL is not set")
	}

	config, err := pgxpool.ParseConfig(postgresUrl)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to parse config, error: %s", err.Error())
	}

	connPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to create connection pool, error: %s", err.Error())
	}

	postgres := PostgresDB{
		connPool: connPool,
	}

	if err := postgres.createTables(); err != nil {
		return nil, err
	}

	return &postgres, nil
}

// TODO: Create tables for storing user data
func (db *PostgresDB) createTables() error {
	return nil
}

// TODO: Close database connection
func (db *PostgresDB) Close() error {
	defer db.connPool.Close()

	return nil
}
