// Description: This package contains the logic to apply SQL migrations to a database.
package migration

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	_ "github.com/lib/pq"
)

// MigrationsTableName is the name of the table to track migration versions.
const MigrationsTableName = "migrations"

// ApplyMigrations applies SQL migrations from the specified directory to the given database.
func ApplyMigrations(db *sql.DB, migrationsDir string) error {
	// Ensure the migrations table exists
	if err := createMigrationsTable(db); err != nil {
		return err
	}

	// Get the current migration version
	currentVersion, err := getCurrentVersion(db)
	if err != nil {
		return err
	}

	// Retrieve migration files
	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))
	if err != nil {
		return err
	}

	// Sort files by version number
	sort.Strings(files)

	// Apply migrations starting from the next version after the current version
	for _, file := range files {
		version, err := parseVersion(file)
		if err != nil {
			return err
		}

		if version > currentVersion {
			sqlBytes, err := os.ReadFile(file)
			if err != nil {
				return err
			}

			sqlQuery := string(sqlBytes)

			_, err = db.Exec(sqlQuery)
			if err != nil {
				return fmt.Errorf("error executing migration %s: %v", file, err)
			}

			log.Printf("Applied migration %s\n", file)

			// Update the current version in the migrations table
			if err := updateCurrentVersion(db, version); err != nil {
				return err
			}
		}
	}

	return nil
}

func createMigrationsTable(db *sql.DB) error {
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id SERIAL PRIMARY KEY,
			version INTEGER NOT NULL
		);
	`, MigrationsTableName)

	_, err := db.Exec(query)
	return err
}

func getCurrentVersion(db *sql.DB) (int, error) {
	var version int
	query := fmt.Sprintf("SELECT COALESCE(MAX(version), 0) FROM %s;", MigrationsTableName)
	err := db.QueryRow(query).Scan(&version)
	return version, err
}

func updateCurrentVersion(db *sql.DB, version int) error {
	query := fmt.Sprintf("INSERT INTO %s (version) VALUES ($1);", MigrationsTableName)
	_, err := db.Exec(query, version)
	return err
}

func parseVersion(filename string) (int, error) {
	base := filepath.Base(filename)
	extension := filepath.Ext(base)
	// get the version number by removing the extension and and the name 001_initial.sql -> 001
	nameStr := base[:len(base)-len(extension)]
	versionStr := nameStr[:3]
	return strconv.Atoi(versionStr)
}
