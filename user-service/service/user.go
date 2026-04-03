package service

import (
	"database/sql"
	"errors"
	"user-service/repository"
)

func CreateUserService(db *sql.DB, name string) (int, error) {
	if name == "" {
		return 0, errors.New("name required")
	}
	return repository.CreateUser(db, name)
}
