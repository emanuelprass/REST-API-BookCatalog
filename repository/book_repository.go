package repository

import (
	"database/sql"
)

type BookRepository interface {
}

type bookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{
		DB: db,
	}
}
