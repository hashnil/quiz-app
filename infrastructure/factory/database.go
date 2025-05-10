package factory

import (
	"fmt"
	"quiz-app/internal/integration/db"
	"quiz-app/internal/integration/db/postgresql"
)

// Database factory
func InitDBClient(dbType string) (db.Client, error) {
	switch dbType {
	case "postgresql":
		// Initialize PostgreSQL client for database operations.
		dbClient, err := postgresql.NewPostgresClient()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize postgres client: %w", err)
		}
		return dbClient, nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
