package userinterface

import (
	"fmt"
	"pinelabs/constants"
	"pinelabs/models"
	"strings"
)

// AddBook allows the user to add a book to the library.
func AddBook(library *models.Library) {

	title := GetBookTitle()
	author := GetBookAuthor()
	quantity := GetBookQuantity()
	book := models.NewBook(title, author, quantity, 0)
	msg := library.AddBook(&book)
	fmt.Println(msg)
}

// DisplayAvailableBooks prints the list of available books in the library.
func DisplayAvailableBooks(library *models.Library) {
	if len(library.Books) == 0 {
		fmt.Println("The library is empty.")
		return
	}
	printBookTable(library.Books)
}

// BorrowBook allows the user to borrow a book from the library.
func BorrowBook(library *models.Library) {
	title := GetBookTitle()
	msg := library.BorrowBook(title)
	fmt.Println(msg)
}

// ReturnBook allows the user to return a borrowed book to the library.
func ReturnBook(library *models.Library) {
	title := GetBookTitle()
	msg := library.ReturnBook(title)
	fmt.Println(msg)
}

// SearchBook allows the user to search for books by title or author.
func SearchBook(library *models.Library) {
	for {
		DisplaySearchMenu()
		choice := GetUserChoice()
		switch choice {
		case constants.TitleSearchOption:
			SearchBookByTitle(library)
			return
		case constants.AuthorSearchOption:
			SearchBookByAuthor(library)
			return
		default:
			fmt.Print("Invalid search option selected, please try again...")
		}
	}
}

// DisplayBorrowedBooks displays the list of books borrowed by the user.
func DisplayBorrowedBooks(library *models.Library) {
	books := library.BorrowedBooks()
	fmt.Println()
	if len(books) == 0 {
		fmt.Println("No books borrowed")
	} else {
		fmt.Println("Here is the list of books you borrowed:")
		printBookTable(books)
	}
}

// SearchBookByTitle allows the user to search for books by title.
func SearchBookByTitle(library *models.Library) {
	title := GetBookTitle()
	books := library.SearchBookByTitle(title)
	fmt.Println()
	if len(books) == 0 {
		fmt.Println("No books available with the provided title")
	} else {
		fmt.Println("Here is the best match to your search:")
		printBookTable(books)
	}
}

// SearchBookByAuthor allows the user to search for books by author.
func SearchBookByAuthor(library *models.Library) {
	author := GetBookAuthor()
	books := library.SearchBookByAuthor(author)
	fmt.Println()
	if len(books) == 0 {
		fmt.Println("No books available by the provided author")
	} else {
		fmt.Println("Here is the best match to your search:")
		printBookTable(books)
	}
}

// printBookTable prints the list of books in a table format.
func printBookTable(books []*models.Book) {
	// Calculate maximum lengths for columns
	maxTitleLength := 50  // Minimum width for title column
	maxAuthorLength := 50 // Minimum width for author column
	for _, book := range books {
		if len(book.Title) > maxTitleLength {
			maxTitleLength = len(book.Title)
		}
		if len(book.Author) > maxAuthorLength {
			maxAuthorLength = len(book.Author)
		}
	}

	fmt.Println(strings.Repeat("-", maxTitleLength+maxAuthorLength+20)) // Line separator
	fmt.Printf("%-*s  |  %-*s  |  %s\n", maxTitleLength, "Title", maxAuthorLength, "Author", "Available")
	fmt.Println(strings.Repeat("-", maxTitleLength+maxAuthorLength+20)) // Line separator

	// Print books
	for _, book := range books {
		available := book.Quantity - book.Borrowed
		fmt.Printf("%-*s  |  %-*s  |  %d\n", maxTitleLength, book.Title, maxAuthorLength, book.Author, available)
	}
	fmt.Println(strings.Repeat("-", maxTitleLength+maxAuthorLength+20)) // Line separator

}
