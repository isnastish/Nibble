package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/isnastish/nibble/pkg/ipresolver"
	"github.com/isnastish/nibble/pkg/log"
	"github.com/isnastish/nibble/pkg/utils"
)

type PostgresDB struct {
	connPool *pgxpool.Pool
}

func NewPostgresDB() (*PostgresDB, error) {
	postgresUrl, set := os.LookupEnv("POSTGRES_URL")

	postgresUrl = "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
	set = true

	if !set || postgresUrl == "" {
		return nil, fmt.Errorf("postgres: postgres_url is not set")
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

	log.Logger.Info("Successfully connected to postgres database")

	return &postgres, nil
}

func (db *PostgresDB) createTables() error {
	conn, err := db.connPool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("postgres: failed to acquire db connection, error: %s", err.Error())
	}

	defer conn.Release()

	// we could have devided it into two tables,
	// but it's not necessary
	query := `CREATE TABLE IF NOT EXISTS "users" (
		"id" SERIAL, 
		"first_name" VARCHAR(64) NOT NULL, 
		"last_name" VARCHAR(64) NOT NULL,
		"password" CHAR(64) NOT NULL, 
		"email" VARCHAR(64) NOT NULL UNIQUE,
		"city" VARCHAR(64) NOT NULL,
		"country" VARCHAR(64) NOT NULL,
		PRIMARY KEY("id")
	);`

	if _, err = conn.Exec(context.Background(), query); err != nil {
		return fmt.Errorf("postgres: failed to create users table, error: %s", err.Error())
	}

	return nil
}

func (db *PostgresDB) AddUser(firstName, lastName, password, email string, ipInfo *ipresolver.IpInfo) error {
	// NOTE: User data validation should be done in a separate
	conn, err := db.connPool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("postgres: failed to acquire db connection, error: %s", err.Error())
	}

	defer conn.Release()

	query := `INSERT INTO "users" 
	("first_name", "last_name", "password", "email", "city", "country") 
	values ($1, $2, $3, $4, $5, $6);`

	// Hash the password before putting it into a database
	hashedPassword := utils.Sha256([]byte(password))

	if _, err := conn.Exec(context.Background(), query, firstName, lastName, hashedPassword, email, ipInfo.City, ipInfo.Country); err != nil {
		return fmt.Errorf("postgres: failed to insert user, error: %s", err.Error())
	}

	log.Logger.Info("Succefully added user: %s %s, email: %s, city: %s, country: %s to database", firstName, lastName, email, ipInfo.City, ipInfo.Country)

	return nil
}

func (db *PostgresDB) HasUser(email string) (bool, error) {
	conn, err := db.connPool.Acquire(context.Background())
	if err != nil {
		return false, err
	}

	defer conn.Release()

	// check if the user with specified email address already exists
	query := `SELECT "email" FROM "users" WHERE "email" = ($1);`
	row := conn.QueryRow(context.Background(), query, email)

	var result string
	if err := row.Scan(&result); err != nil {
		if err == pgx.ErrNoRows {
			// user doesn't exit
			return false, nil
		}
		return false, fmt.Errorf("postgres: failed to select user, error: %s", err.Error())
	}

	return true, nil
}

// Close database connection
func (db *PostgresDB) Close() error {
	defer db.connPool.Close()
	return nil
}
