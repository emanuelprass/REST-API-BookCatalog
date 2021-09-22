package main

import (
	"book-catalog-rest/config"
	"book-catalog-rest/server"

	"github.com/go-playground/validator/v10"
)

func main() {
	validation := validator.New()
	cfg := config.LoadConfig()
	dbInit, err := config.MySQL(cfg)
	if err != nil {
		panic(err)
	}

	server := server.NewServer(dbInit, validation)
	server.ListenAndServer("8080")
}
