package main

import (
	"database/sql"
	"qa_commander/config"
	"qa_commander/internal/server"

	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"

	// Run migrations
	migration "qa_commander/migrations"
)

func main() {
	config := config.NewConfig()

	func() {
		conf := *config
		appConfig := &conf
		db, err := sql.Open("postgres", appConfig.DatabaseURL)
		if err != nil {
			log.Fatalf("Error connecting to the database: %v", err)
		}
		defer db.Close()
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current directory: %v", err)
		}
		//migration directory is  /qa_commander/migrations
		parentDir := filepath.Dir(dir)
		migrationsDir := filepath.Join(parentDir, "migrations")
		err = migration.ApplyMigrations(db, migrationsDir)
		if err != nil {
			log.Fatalf("Error applying migrations: %v", err)
		}
		log.Println("Migrations applied successfully")
	}()

	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		panic("Failed to connect to the database")
	}

	r := server.SetupRouter(db)
	r.Run(":8080")
}
