package utils

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConnection() *sql.DB {
	godotenv.Load(".env")
	postgres_url := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", postgres_url)

	if err != nil {
		log.Fatal("Could not connect to database")
	}
	return db
}
