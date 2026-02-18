package main

import (
	"database/sql"
	"log"
	"mailbox-api/internal/database"
	"mailbox-api/internal/env"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	DB_NAME := env.GetEnvString("DB_NAME", "kwikmail")
	DB_USERNAME := env.GetEnvString("DB_USERNAME", "kwikmail")
	DB_PASSWORD := env.GetEnvString("DB_PASSWORD", "kwikmail")
	APP_PORT := env.GetEnvInt("PORT", 8080)
	JWT_SECRET := env.GetEnvString("JWT_SECRET", "some934GK4j24")

	db, err := sql.Open("postgres", "postgres://"+DB_USERNAME+":"+DB_PASSWORD+"@localhost:5433/"+DB_NAME+"?sslmode=disable")
	if err != nil {
		log.Fatal("An error occurred while opening the database:", err)
	}
	defer db.Close()

	models := database.NewModels(db)
	app := &application{
		port:      APP_PORT,
		jwtSecret: JWT_SECRET,
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
