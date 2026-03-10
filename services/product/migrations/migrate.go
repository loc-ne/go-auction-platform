package migrations

import (
    "os"
    "errors"
    "log"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "github.com/joho/godotenv"
)

func RunMigrations() {
    err := godotenv.Load()
    if err != nil {
		log.Println(".env not found")
	}

    dbUrl := os.Getenv("DATABASE_URL")
    if dbUrl == "" {
        log.Fatal("DATABASE_URL is not set in .env file")
    }
    migrationPath := os.Getenv("MIGRATION_PATH")
    if migrationPath == "" {
        log.Fatal("MIGRATION_PATH is not set in .env file")
    }
    m, err := migrate.New(
        migrationPath, 
        dbUrl,
    )
    // m, err := migrate.New(
    //     "file://migrations", 
    //     dbUrl,
    // )
    if err != nil {
        log.Fatal("Unable to initiate migration: ", err)
    }

    if err := m.Up(); err != nil {
        if errors.Is(err, migrate.ErrNoChange) {
            log.Println("Database is already at the latest version, no changes.")
        } else {
            log.Fatal("Error when running migration: ", err)
        }
    } else {
        log.Println("Migration successful!")
    }
}