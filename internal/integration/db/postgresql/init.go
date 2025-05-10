package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"quiz-app/internal/integration/db"
	"quiz-app/internal/integration/db/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresSQLClient struct {
	client *gorm.DB
}

func NewPostgresClient() (db.Client, error) {
	sqlDB, err := connectWithStandardDSN()
	if err != nil {
		return nil, fmt.Errorf("failed to create SQL connection: %w", err)
	}

	// Initialize GORM with the existing *sql.DB connection.
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize GORM: %w", err)
	}

	db.AutoMigrate(
		&models.School{}, &models.Classroom{}, &models.User{}, &models.Quiz{},
		&models.Question{}, &models.Response{}, &models.Event{},
	)

	log.Println("Postgres SQL client initialized successfully.")
	return &PostgresSQLClient{client: db}, nil
}

// ConnectWithStandardDSN connects to GORM Postgres using a standard DSN (for AWS or local Postgres).
func connectWithStandardDSN() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		viper.GetString("db.postgres.host"),
		viper.GetString("db.postgres.user"),
		viper.GetString("db.postgres.pass"),
		viper.GetString("db.postgres.name"),
		viper.GetString("db.postgres.port"),
		viper.GetString("db.postgres.sslmode"),
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	return db, nil
}
