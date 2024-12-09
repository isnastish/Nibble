package db

// TODO: Use interface for databese so we can switch between
// multiple databases, not only postgres

// TODO: Use pgx library

type PostgresDB struct {
}

func NewPostgresDB() (*PostgresDB, error) {
	postgres := &PostgresDB{}

	return postgres, nil
}

// TODO: Create tables for storing user data
func (db *PostgresDB) createTables() {

}

// TODO: Close database connection
func (db *PostgresDB) Close() {

}
