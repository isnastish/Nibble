package db

// TODO: Use interface for databese so we can switch between
// multiple databases, not only postgres

// TODO: Use pgx library

type PostgresDB struct {
}

func NewPostgresDB() (*PostgresDB, error) {
	return &PostgresDB{}, nil
}

// TODO: Close database connection
func (b *PostgresDB) Close() {

}
