package usecase

import (
	"book-catalog-rest/repository"
	"book-catalog-rest/transport"
	"net/http"
)

type BookUsecase interface {
	GetList() (*transport.GetList, *transport.ResponseError)
}

type bookUsecase struct {
	repository repository.BookRepository
}

func NewBookUsecase(bookRepository repository.BookRepository) BookUsecase {
	return &bookUsecase{
		repository: bookRepository,
	}
}

func (b *bookUsecase) GetList() (*transport.GetList, *transport.ResponseError) {
	result, err := b.repository.GetList()
	if err != nil {
		response := &transport.ResponseError{
			Message: "Un Processable Entity",
			Status:  http.StatusUnprocessableEntity,
		}
		return nil, response
	}
	return &transport.GetList{
		Count:    len(result),
		ListBook: result,
	}, nil
}
