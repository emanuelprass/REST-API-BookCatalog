package transport

import (
	"book-catalog-rest/entity"
)

type GeneralResponse struct {
	Message string `json:"message"`
}

type GetList struct {
	Count    int           `json:"count"`
	ListBook []entity.Book `json:"listBook"`
}

type GetBookResponse struct {
	Data entity.Book `json:"data"`
}

type ResponseError struct {
	Message string
	Status  int
}
