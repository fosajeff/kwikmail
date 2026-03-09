package main

import (
	"log"
	"mailbox-api/internal/env"
	"os"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: 'up' or 'down'")
	}

	direction := os.Args[1]

	DB_NAME := env.GetEnvString("DB_NAME", "kwikmail")
	DB_USERNAME := env.GetEnvString("DB_USERNAME", "kwikmail")
	DB_PASSWORD := env.GetEnvString("DB_PASSWORD", "kwikmail")

	db, err := sql.Open("postgres", "postgres://"+DB_USERNAME+":"+DB_PASSWORD+"@localhost:5433/"+DB_NAME+"?sslmode=disable")
	if err != nil {
		log.Fatal("An error occurred while opening the database:", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("An error occurred while creating the migration driver: ", err)
	}

	defer db.Close()

	m, err := migrate.NewWithDatabaseInstance("file:///home/jeffrey/Desktop/apps/practice-projects/kwikmail/apps/mailbox-api/cmd/migrate/migrations", "postgres", driver)
	if err != nil {
		log.Fatal("An error occurred while creating the migration instance: ", err)
	}

	switch direction {
	case "up":
		if err = m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("An error occurred while running the migration up: ", err)
		}
	case "down":
		if err = m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("An error occurred while running the migration down: ", err)
		}
	default:
		log.Fatal("Invalid migration direction. Please use 'up' or 'down'")
	}

}
