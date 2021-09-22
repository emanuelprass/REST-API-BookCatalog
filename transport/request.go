package transport

type InsertBook struct {
	Name    string `json:"bookName" validate:"required"`
	Creator string `json:"bookCreator" validate:"required"`
}

type UpdateBook struct {
	Name    string `json:"bookName"`
	Creator string `json:"bookCreator"`
}
