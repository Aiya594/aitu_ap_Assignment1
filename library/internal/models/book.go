package models

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

func NewBook(id, title, author string) *Book {
	return &Book{
		ID:         id,
		Title:      title,
		Author:     author,
		IsBorrowed: false,
	}
}
