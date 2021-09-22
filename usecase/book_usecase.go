package usecase

import (
	"book-catalog-rest/repository"
)

type BookUsecase interface {
}

type bookUsecase struct {
	br repository.BookRepository
}

func NewBookUsecase(br repository.BookRepository) BookUsecase {
	return &bookUsecase{
		br: br,
	}
}
