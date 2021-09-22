package usecase

import (
	"book-catalog-rest/entity"
	"book-catalog-rest/repository"
	"book-catalog-rest/transport"
	"net/http"
)

type BookUsecase interface {
	GetList() (*transport.GetList, *transport.ResponseError)
	AddBook(data transport.InsertBook) (*transport.GeneralResponse, *transport.ResponseError)
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

func (b *bookUsecase) AddBook(data transport.InsertBook) (*transport.GeneralResponse, *transport.ResponseError) {
	createPayload := entity.Book{
		Name:    data.Name,
		Creator: data.Creator,
	}

	err := b.repository.AddBook(createPayload)

	if err != nil {
		response := &transport.ResponseError{
			Message: "Un Processable Entity",
			Status:  http.StatusUnprocessableEntity,
		}
		return nil, response
	}

	result := &transport.GeneralResponse{
		Message: "Success to insert book : " + createPayload.Name + " into database",
	}

	return result, nil
}
