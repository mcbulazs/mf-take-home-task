package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq" // postgres driver

	scripts "mcbulazs/mf-take-home-task/sql"
)

type config struct {
	Name     string
	User     string
	Password string
	Host     string
	SSLMode  string
}

func loadConfig() (*config, error) {
	name := os.Getenv("DB_NAME")
	if name == "" {
		return nil, errors.New("env variable \"DB_NAME\" not set")
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		return nil, errors.New("env variable \"DB_USER\" not set")
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return nil, errors.New("env variable \"DB_PASSWORD\" not set")
	}
	cfg := config{
		Name:     name,
		User:     user,
		Password: password,
		Host:     "db",
		SSLMode:  "disable",
	}

	return &cfg, nil
}

func ConnectToDB() (*sql.DB, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.Name, cfg.Host, cfg.SSLMode)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	// test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	if err := migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *sql.DB) error {
	if _, err := db.Exec(scripts.MigrateSchemaSQL); err != nil {
		return fmt.Errorf("error applying schema: %w", err)
	}

	if _, err := db.Exec(scripts.MigrateSeedSQL); err != nil {
		return fmt.Errorf("error applying seed: %w", err)
	}

	return nil
}
