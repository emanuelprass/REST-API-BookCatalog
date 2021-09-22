package handler

import (
	"book-catalog-rest/usecase"

	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bu        usecase.BookUsecase
	validator *validator.Validate
}

func NewBookHandler(bu usecase.BookUsecase, validator *validator.Validate) *bookHandler {
	return &bookHandler{
		bu:        bu,
		validator: validator,
	}
}
