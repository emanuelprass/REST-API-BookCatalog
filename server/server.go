package server

import (
	"book-catalog-rest/handler"
	"book-catalog-rest/repository"
	"book-catalog-rest/usecase"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	DB        *sql.DB
	Router    *mux.Router
	validator *validator.Validate
}

func NewServer(db *sql.DB, validator *validator.Validate) *ApiServer {
	r := mux.NewRouter()
	return &ApiServer{
		DB:        db,
		Router:    r,
		validator: validator,
	}
}

func (server *ApiServer) ListenAndServer(port string) {
	server.registerRouter()

	http.ListenAndServe(":"+port, server.Router)
}

func (server *ApiServer) registerRouter() {
	repo := repository.NewBookRepository(server.DB)
	uCase := usecase.NewBookUsecase(repo)
	bookHandler := handler.NewBookHandler(uCase, server.validator)

	fmt.Println(bookHandler)
}
