package transport

type InsertBook struct {
	Name    string `json:"bookName" validate:"required"`
	Creator string `json:"bookCreator" validate:"required"`
}

type UpdateBook struct {
	Id      string `json:"id" validate:"required"`
	Name    string `json:"name"`
	Creator string `json:"creator"`
}
