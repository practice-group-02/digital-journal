package dal

import (
	"database/sql"
	"digital-journal/internal/config"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func NewStore(DBconfig config.DBConfig) error {
	// Validate required fields
	if DBconfig.DBhost == "" || DBconfig.DBuser == "" || DBconfig.DBname == "" {
		return fmt.Errorf("missing required database configuration")
	}

	// Default to port 5432 if not specified
	if DBconfig.DBport == "" {
		DBconfig.DBport = "5432"
	}

	// Build connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DBconfig.DBhost,
		DBconfig.DBport,
		DBconfig.DBuser,
		DBconfig.DBpassword,
		DBconfig.DBname)

	var err error
	DB, err = sql.Open("postgres", connStr) // Always use "postgres" as driver name
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Verify connection works
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}