package handler

import (
	"book-catalog-rest/transport"
	"book-catalog-rest/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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

func (b *bookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var requestBook transport.InsertBook
	err := decoder.Decode(&requestBook)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responses := transport.ResponseError{
			Message: "error while decode request body",
			Status:  http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(responses)
		return
	}

	// checking validation
	errorValidation := b.validator.Struct(requestBook)
	if errorValidation != nil {
		w.WriteHeader(http.StatusBadRequest)
		dataResponse := transport.ResponseError{
			Message: errorValidation.Error(),
			Status:  http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(dataResponse)
		return
	}

	result, responseError := b.usecase.AddBook(requestBook)
	if responseError != nil {
		w.WriteHeader(responseError.Status)
		json.NewEncoder(w).Encode(responseError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (b *bookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["bookID"])

	result, err := b.usecase.GetByID(id)
	if err != nil {
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (b *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["bookID"])
	var requestBook transport.UpdateBook
	err := decoder.Decode(&requestBook)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responses := transport.ResponseError{
			Message: "error while decode request body",
			Status:  http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(responses)
		return
	}

	// checking validation
	errorValidation := b.validator.Struct(requestBook)

	if errorValidation != nil {
		w.WriteHeader(http.StatusBadRequest)
		dataResponse := transport.ResponseError{
			Message: errorValidation.Error(),
			Status:  http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(dataResponse)
		return
	}

	result, responseError := b.usecase.UpdateBook(id, requestBook)

	if responseError != nil {
		w.WriteHeader(responseError.Status)
		json.NewEncoder(w).Encode(responseError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
