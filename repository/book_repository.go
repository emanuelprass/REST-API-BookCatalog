package repository

import (
	"book-catalog-rest/entity"
	"database/sql"
)

type BookRepository interface {
	GetList() ([]entity.Book, error)
	AddBook(payload entity.Book) error
	GetByID(id string) (*entity.Book, error)
}

type bookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{
		DB: db,
	}
}

func (b *bookRepository) GetList() ([]entity.Book, error) {
	rows, err := b.DB.Query("select book_id, book_name, book_creator from tbl_books")

	if err != nil {
		return nil, err
	}

	var books []entity.Book

	for rows.Next() {
		var res entity.Book
		_ = rows.Scan(&res.Id, &res.Name, &res.Creator)
		books = append(books, res)
	}

	return books, nil
}

func (b *bookRepository) AddBook(payload entity.Book) error {
	_, err := b.DB.Exec("INSERT INTO tbl_books (book_name, book_creator) VALUES (?, ?)", payload.Name, payload.Creator)

	if err != nil {
		return err
	}

	return nil
}

func (b *bookRepository) GetByID(id string) (*entity.Book, error) {
	sqlStatement := "SELECT * FROM tbl_books WHERE book_id = ?"
	row := b.DB.QueryRow(sqlStatement, id)
	var book entity.Book
	err := row.Scan(&book.Id, &book.Name, &book.Creator)

	if err != nil {
		return nil, err
	}

	return &book, nil
}
