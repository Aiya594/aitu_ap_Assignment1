package models

type Library struct {
	Books map[string]*Book
}

func NewLibrary() *Library {
	return &Library{make(map[string]*Book)}
}

func (library *Library) AddBook(book *Book) {
	library.Books[book.ID] = book
}

func (library *Library) BorrowBook(id string) (string, bool) {
	book, exists := library.Books[id]
	if !exists {
		return " no such book in library", false //no such book in library
	}
	if book.IsBorrowed {
		return "book is already borrowed", false //book is already borrowed
	}

	book.IsBorrowed = true
	return "", true //successfully borrowed
}

func (library *Library) GetBook(id string) (*Book, bool) {
	book, exists := library.Books[id]
	if !exists {
		return book, false //no such book in library
	}
	return book, true
}

func (library *Library) ReturnBook(id string) (string, bool) {
	book, exists := library.Books[id]
	if !exists {
		return "have never had such boook in library", false //have never had such boook in library
	}
	if !book.IsBorrowed {
		return "this book wasnt borrowed", false //this wasnt borrowed
	}
	book.IsBorrowed = false
	return "", true //returned book
}

func (library *Library) ListAvailableBooks() []*Book {
	var result []*Book
	for _, book := range library.Books {
		if !book.IsBorrowed {
			result = append(result, book)
		}
	}
	return result
}

func (library *Library) ListAllBooks() []*Book {
	var result []*Book
	for _, book := range library.Books {
		result = append(result, book)

	}
	return result
}
