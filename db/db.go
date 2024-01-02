package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("postgres", "user=luiz password=root dbname=alura_loja host=localhost sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	return db
}
