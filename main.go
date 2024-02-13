package main

import (
	"fmt"
	"pinelabs/constants"
	"pinelabs/models"
	"pinelabs/userinterface"
)

func main() {
	bookStore := models.NewBookStore()

	library1 := models.NewLibrary("1") // Initialize library
	library2 := models.NewLibrary("2") // Initialize library
	bookStore.AddLibrary(library1)
	bookStore.AddLibrary(library2)

	for {
		userinterface.DisplayMenu()
		choice := userinterface.GetUserChoice()

		switch choice {
		case constants.AddBookOption:
			userinterface.AddBook(bookStore)
		case constants.AvailableBooksOption:
			userinterface.DisplayAvailableBooks(bookStore)
		case constants.BorrowBookOption:
			userinterface.BorrowBook(bookStore)
		case constants.ReturnBookOption:
			userinterface.ReturnBook(bookStore)
		case constants.SearchBookOption:
			userinterface.SearchBook(bookStore)
		case constants.BorrowedBookOption:
			userinterface.DisplayBorrowedBooks(bookStore)
		case constants.QuitOption:
			return
		default:
			fmt.Println("Invalid Choice, please try again...")
		}
	}
}
