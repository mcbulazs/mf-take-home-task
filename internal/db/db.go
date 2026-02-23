package db

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"os"
)

//go:migrations/*.sql
var migrationFiles embed.FS

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
		Host:     "postgres",
		SSLMode:  "diable",
	}

	return &cfg, nil
}

func ConnectToDB() (*sql.DB, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s",
		cfg.Name, cfg.Password, cfg.Name, cfg.Host, cfg.SSLMode)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}
	if err := migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *sql.DB) error {
	schema, err := migrationFiles.ReadFile("001_init.sql")
	if err != nil {
		return fmt.Errorf("error reading file 001_init.sql: %w", err)
	}
	seed, err := migrationFiles.ReadFile("002_seed.sql")
	if err != nil {
		return fmt.Errorf("error reading file 002_seed.sql: %w", err)
	}

	if _, err := db.Exec(string(schema)); err != nil {
		return fmt.Errorf("error applying schema: %w", err)
	}

	if _, err := db.Exec(string(seed)); err != nil {
		return fmt.Errorf("error applying seed: %w", err)
	}

	return nil
}
