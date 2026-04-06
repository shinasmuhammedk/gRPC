package db

import (
	"database/sql"
	_ "github.com/lib/pq" // ✅ IMPORTANT (don't remove underscore)
	"log"
)

func ConnectDB() *sql.DB {
	connStr := "user=postgres password=Shinas dbname=grpc sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB connection error :", err)
	}

	return db
}
