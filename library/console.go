package library

import (
	"fmt"

	"github.com/Aiya594/aitu_ap_Assignment1/library/internal/models"
)

func RunLibrary() {
	books := initBooks()
	library := initLibrary(books)

	for {
		fmt.Println(`
This is Library Management System
Options:
	1.List all books
	2.Borrow a book
	3.Return a book
	4.List available books
	5.Stop`)
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("couldnt read the input:", err)
			continue
		}

		switch choice {
		case 1:
			allBooks := library.ListAllBooks()
			fmt.Println("All Books:")
			printBooks(allBooks)

		case 2:
			var id string
			fmt.Println("Enter book id to borrow:")
			_, err = fmt.Scan(&id)
			if err != nil {
				fmt.Println("couldnt read the input:", err)
				continue
			}
			fmt.Println("Book ID: " + id)
			message, ok := library.BorrowBook(id)
			if !ok {
				fmt.Println("Couldnt borrow the book:", message)
				continue
			}
			fmt.Println("Borrowed book:")
			book, ok := library.GetBook(id)
			if !ok {
				fmt.Println("Couldnt find the book")
				continue
			}
			fmt.Println(book)

		case 3:
			var id string
			fmt.Println("Enter book id to borrow:")
			_, err = fmt.Scan(&id)
			if err != nil {
				fmt.Println("couldnt read the input:", err)
				return
			}
			message, ok := library.ReturnBook(id)
			if !ok {
				fmt.Println("Couldnt return the book:", message)
				continue
			}
			fmt.Println("Returned book:")
			book, ok := library.GetBook(id)
			if !ok {
				fmt.Println("Couldnt find the book")
				continue
			}
			fmt.Println(book)

		case 4:
			availableBooks := library.ListAvailableBooks()
			fmt.Println("List available books:")
			printBooks(availableBooks)

		case 5:
			fmt.Println("Bye")
			return

		default:
			fmt.Println("Invalid choice:", choice)

		}

	}

}

func initLibrary(books []*models.Book) *models.Library {
	library := models.NewLibrary()
	for _, book := range books {
		library.AddBook(book)
	}
	return library
}

func initBooks() []*models.Book {
	return []*models.Book{
		models.NewBook("1", "1984", "George Orwell"),
		models.NewBook("2", "To Kill a Mockingbird", "Harper Lee"),
		models.NewBook("3", "Brave New World", "Aldous Huxley"),
		models.NewBook("4", "The Catcher in the Rye", "J.D. Salinger"),
		models.NewBook("5", "The Hobbit", "J.R.R. Tolkien"),
		models.NewBook("6", "Fahrenheit 451", "Ray Bradbury"),
		models.NewBook("7", "Moby-Dick", "Herman Melville"),
		models.NewBook("8", "Pride and Prejudice", "Jane Austen"),
		models.NewBook("9", "The Lord of the Rings", "J.R.R. Tolkien"),
		models.NewBook("10", "Crime and Punishment", "Fyodor Dostoevsky"),
		models.NewBook("11", "The Green Mile", "Stephen King"),
	}
}

func printBooks(books []*models.Book) {
	for _, book := range books {
		fmt.Println(book)
	}
}
