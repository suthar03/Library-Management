package userinterface

import (
	"fmt"
	"pinelabs/constants"
	"pinelabs/models"
	"strings"
)

// AddBook allows the user to add a book to the library.
func AddBook(store *models.BookStore) {

	title := GetBookTitle()
	author := GetBookAuthor()
	quantity := GetBookQuantity()
	libraryID := GetLibraryID(store)
	library := store.GetLibrary(libraryID)
	book := models.NewBook(title, author, quantity, 0)
	msg := library.AddBook(&book)
	fmt.Println(msg)
}

// DisplayAvailableBooks prints the list of available books in the library.
func DisplayAvailableBooks(store *models.BookStore) {

	printAllLibriesBooks(store)

}

// BorrowBook allows the user to borrow a book from the library.
func BorrowBook(store *models.BookStore) {
	title := GetBookTitle()
	libraryID := GetLibraryID(store)
	library := store.GetLibrary(libraryID)
	msg := library.BorrowBook(title)
	fmt.Println()
	fmt.Println(msg)
}

// ReturnBook allows the user to return a borrowed book to the library.
func ReturnBook(store *models.BookStore) {
	title := GetBookTitle()
	libraryID := GetLibraryID(store)
	library := store.GetLibrary(libraryID)
	msg := library.ReturnBook(title)
	fmt.Println()
	fmt.Println(msg)
}

// SearchBook allows the user to search for books by title or author.
func SearchBook(store *models.BookStore) {
	for {
		DisplaySearchMenu()
		choice := GetUserChoice()
		switch choice {
		case constants.TitleSearchOption:
			SearchBookByTitle(store)
			return
		case constants.AuthorSearchOption:
			SearchBookByAuthor(store)
			return
		default:
			fmt.Print("Invalid search option selected, please try again...")
		}
	}
}

// DisplayBorrowedBooks displays the list of books borrowed by the user.
func DisplayBorrowedBooks(store *models.BookStore) {
	id := GetLibraryID(store)
	library := store.GetLibrary(id)
	books := library.BorrowedBooks()

	fmt.Println()
	if len(books) == 0 {
		fmt.Println("No books borrowed")
	} else {
		fmt.Println("Here is the list of books you borrowed:")
		printBooksTitleAndAuthor(books)
	}
}

// SearchBookByTitle allows the user to search for books by title.
func SearchBookByTitle(store *models.BookStore) {
	title := GetBookTitle()
	for _, library := range store.Libraries {
		books := library.SearchBookByTitle(title)
		fmt.Println("library", library.ID)
		if len(books) == 0 {
			fmt.Printf("No books available with the title: %s\n", title)
		} else {
			fmt.Println("Here is the best match to your search:")
			printBooksAllDetails(books)
		}
	}

}

// SearchBookByAuthor allows the user to search for books by author.
func SearchBookByAuthor(store *models.BookStore) {
	author := GetBookAuthor()
	for _, library := range store.Libraries {
		books := library.SearchBookByAuthor(author)
		fmt.Println(" lib", library.ID)
		if len(books) == 0 {
			fmt.Printf("No books available of the author: %s\n", author)
		} else {
			fmt.Println("Here is the best match to your search:")
			printBooksAllDetails(books)
		}
	}
}

func printAllLibriesBooks(store *models.BookStore) {
	for _, lib := range store.Libraries {
		fmt.Println("Library ID:", lib.ID)
		if len(lib.Books) == 0 {
			fmt.Print("No book in the lib ")
			continue
		}
		printBooksAllDetails(lib.Books)
	}
}

// printBookTable prints the list of books in a table format.
func printBooksAllDetails(books []*models.Book) {
	// Calculate maximum lengths for columns
	maxTitleLength := 10  // Minimum width for title column
	maxAuthorLength := 10 // Minimum width for author column
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

// printBookTable prints the list of books in a table format.
func printBooksTitleAndAuthor(books []*models.Book) {
	// Calculate maximum lengths for columns
	maxTitleLength := 10  // Minimum width for title column
	maxAuthorLength := 10 // Minimum width for author column
	for _, book := range books {
		if len(book.Title) > maxTitleLength {
			maxTitleLength = len(book.Title)
		}
		if len(book.Author) > maxAuthorLength {
			maxAuthorLength = len(book.Author)
		}
	}

	fmt.Println(strings.Repeat("-", maxTitleLength+maxAuthorLength)) // Line separator
	fmt.Printf("%-*s  |  %-*s\n", maxTitleLength, "Title", maxAuthorLength, "Author")
	fmt.Println(strings.Repeat("-", maxTitleLength+maxAuthorLength)) // Line separator

	// Print books
	for _, book := range books {
		fmt.Printf("%-*s  |  %-*s \n", maxTitleLength, book.Title, maxAuthorLength, book.Author)
	}
	fmt.Println(strings.Repeat("-", maxTitleLength+maxAuthorLength)) // Line separator

}
