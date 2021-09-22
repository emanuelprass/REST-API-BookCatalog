package entity

type Book struct {
	Id      int    `json:"book_id"`
	Name    string `json:"book_name"`
	Creator string `json:"book_creator"`
}
