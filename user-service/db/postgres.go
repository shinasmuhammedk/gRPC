package db

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	connStr := "user=postgres password=Shinas dbname=grpc sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
