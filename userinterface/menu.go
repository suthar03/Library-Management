package userinterface

import (
	"bufio"
	"fmt"
	"os"
	"pinelabs/constants"
	"strconv"
	"strings"
)

func DisplayMenu() {
	fmt.Println("\nLibrary Management System")
	fmt.Printf("%d. Add Book\n", constants.AddBookOption)
	fmt.Printf("%d. Display Available Books\n", constants.AvailableBooksOption)
	fmt.Printf("%d. Borrow Book\n", constants.BorrowBookOption)
	fmt.Printf("%d. Return Book\n", constants.ReturnBookOption)
	fmt.Printf("%d. Search Book\n", constants.SearchBookOption)
	fmt.Printf("%d. Display Borrowed Books\n", constants.BorrowedBookOption)
	fmt.Printf("%d. Quit\n", constants.QuitOption)
}

func GetUserChoice() int {
	// Get user choice
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}

func DisplaySearchMenu() {
	fmt.Println("\nSearch Options")
	fmt.Printf("%d. Title\n", constants.TitleSearchOption)
	fmt.Printf("%d. Author\n", constants.AuthorSearchOption)
}

func GetBookTitle() string {
	var title string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter book title: ")
	if scanner.Scan() {
		title = scanner.Text()
	}
	if err := ValidateBookTitle(title); err != nil {
		fmt.Println(err.Error(), ", please try again...")
		title = GetBookTitle()
	}
	return title
}

func GetBookAuthor() string {
	var author string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter book author: ")
	if scanner.Scan() {
		author = scanner.Text()
	}
	if err := ValidateBookAuthor(author); err != nil {
		fmt.Println(err.Error(), ", please try again...")
		author = GetBookAuthor()
	}
	return author
}

func GetBookQuantity() int {
	var quantity int
	var quantityInput string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter book quantity: ")
	if scanner.Scan() {
		quantityInput = scanner.Text()
	}
	if qty, err := ValidateBookQuantity(quantityInput); err != nil {
		fmt.Println(err.Error(), ", please try again...")
		quantity = GetBookQuantity()
	} else {
		quantity = qty
	}
	return quantity
}

func ValidateBookTitle(title string) error {
	title = strings.TrimSpace(title) // Remove leading and trailing whitespace
	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	return nil
}

func ValidateBookAuthor(author string) error {
	author = strings.TrimSpace(author) // Remove leading and trailing whitespace
	if author == "" {
		return fmt.Errorf("author cannot be empty")
	}
	return nil
}

func ValidateBookQuantity(quantityInput string) (int, error) {
	quantityInput = strings.TrimSpace(quantityInput) // Remove leading and trailing whitespace
	quantity, _ := strconv.Atoi(quantityInput)
	if quantity <= 0 {
		return 0, fmt.Errorf("quantity must be a positive number")
	}
	return quantity, nil
}
