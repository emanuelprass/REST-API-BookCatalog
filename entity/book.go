package entity

type Book struct {
	Id      string `json:"book_id"`
	Name    string `json:"book_name"`
	Creator string `json:"book_creator"`
}
