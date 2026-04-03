package repository

import "database/sql"

func CreateUser(db *sql.DB, name string) (int, error) {
	var id int
	err := db.QueryRow(
		"INSERT INTO users(name) VALUES($1) RETURNING id",
		name,
	).Scan(&id)

	return id, err
}
