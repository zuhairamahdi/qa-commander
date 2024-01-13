package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"qa_commander/config"
	migration "qa_commander/migrations"
)

func Migrate(conf config.Config) {
	appConfig := &conf

	// Connect to the database
	db, err := sql.Open("postgres", appConfig.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Retrieve the directory of the migration files
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	migrationsDir := filepath.Join(dir, "migrations")

	// Apply migrations to the database
	err = migration.ApplyMigrations(db, migrationsDir)
	if err != nil {
		log.Fatalf("Error applying migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}
