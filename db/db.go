package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	envFile := ".env.local"

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", os.Getenv("POSTGRESQL_NAME"), os.Getenv("POSTGRESQL_PASSWORD"), os.Getenv("POSTGRESQL_DBNAME"), os.Getenv("POSTGRESQL_HOST"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	return db
}
