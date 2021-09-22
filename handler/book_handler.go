package handler

import (
	"book-catalog-rest/usecase"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	usecase   usecase.BookUsecase
	validator *validator.Validate
}

func NewBookHandler(bookUseCase usecase.BookUsecase, validator *validator.Validate) *bookHandler {
	return &bookHandler{
		usecase:   bookUseCase,
		validator: validator,
	}
}

func (b *bookHandler) GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := b.usecase.GetList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
