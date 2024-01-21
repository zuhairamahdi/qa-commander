package tests

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	_ "github.com/lib/pq"
)

func setupTestDB() *sql.DB {
	// Retrieve the current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Could not retrieve current directory")
	}

	// Construct the path to the SQL migration files
	migrationsDir := filepath.Join(filepath.Dir(filename), "../migrations")

	// Generate a unique test database name
	testDBName := fmt.Sprintf("test_%d", os.Getpid())

	// Create a connection to the default PostgreSQL database
	db, err := sql.Open("postgres", "host=localhost user=postgres dbname=qa_commander sslmode=disable password=password")
	if err != nil {
		log.Fatalf("Error connecting to default database: %v", err)
	}
	// Remove the test database if it already exists
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDBName))
	// Create the test database
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", testDBName))
	if err != nil {
		log.Fatalf("Error creating test database: %v", err)
	}

	// Close the connection to the default database
	db.Close()

	// Connect to the test database
	testDB, err := sql.Open("postgres", fmt.Sprintf("host=localhost user=postgres dbname=%s sslmode=disable password=password", testDBName))
	if err != nil {
		log.Fatalf("Error connecting to test database: %v", err)
	}

	// Apply migrations to the test database
	err = applyMigrations(testDB, migrationsDir)
	if err != nil {
		log.Fatalf("Error applying migrations: %v", err)
	}

	return testDB
}

func applyMigrations(db *sql.DB, migrationsDir string) error {
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			migrationPath := filepath.Join(migrationsDir, file.Name())
			sqlBytes, err := os.ReadFile(migrationPath)
			if err != nil {
				return err
			}

			sqlQuery := string(sqlBytes)

			_, err = db.Exec(sqlQuery)
			if err != nil {
				return fmt.Errorf("error executing migration %s: %v", file.Name(), err)
			}
		}
	}

	return nil
}
