package models

type Book struct {
	Title    string
	Author   string
	Quantity int
	Borrowed int
}

func NewBook(title, author string, quantity, borrowed int) Book {
	return Book{
		Title:    title,
		Author:   author,
		Quantity: quantity,
		Borrowed: borrowed,
	}
}
