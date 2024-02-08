package main

import (
	"fmt"
	"pinelabs/constants"
	"pinelabs/models"
	"pinelabs/userinterface"
)

func main() {
	library := models.NewLibrary() // Initialize library

	for {
		userinterface.DisplayMenu()
		choice := userinterface.GetUserChoice()

		switch choice {
		case constants.AddBookOption:
			userinterface.AddBook(library)
		case constants.AvailableBooksOption:
			userinterface.DisplayAvailableBooks(library)
		case constants.BorrowBookOption:
			userinterface.BorrowBook(library)
		case constants.ReturnBookOption:
			userinterface.ReturnBook(library)
		case constants.SearchBookOption:
			userinterface.SearchBook(library)
		case constants.BorrowedBookOption:
			userinterface.DisplayBorrowedBooks(library)
		case constants.QuitOption:
			return
		default:
			fmt.Println("Invalid Choice, please try again...")
		}
	}
}
