package main

import (
	"context"
	"log"

	redisclient "email-ingest-api/internal/redis"
	"email-ingest-api/internal/utils"

	_ "github.com/joho/godotenv/autoload"
)

type application struct {
	port int
}

func main() {
	rdb := redisclient.New()

	// sanity check
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Cannot connect to redis: %v", err)
	}

	app := &application{
		port: utils.GetEnv("PORT", 8080),
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
