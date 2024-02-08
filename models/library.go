package models

import (
	"strings"
)

type Library struct {
	Books []*Book
}

func NewLibrary() *Library {
	return &Library{}
}

func (library *Library) AddBook(newBook *Book) string {
	for _, book := range library.Books {
		if strings.EqualFold(book.Title, newBook.Title) {
			return "Book with this title already exists"
		}
	}
	library.Books = append(library.Books, newBook)
	return "Book added successfully"
}

func (library *Library) BorrowBook(title string) string {
	for _, book := range library.Books {
		if strings.EqualFold(book.Title, title) {
			if book.Quantity > book.Borrowed {
				book.Borrowed++
				return "The book was borrowed successfully"
			} else {
				return "This book is not available for borrowing"
			}
		}
	}
	return "This book is not available in the library"
}

func (library *Library) ReturnBook(title string) string {
	for _, book := range library.Books {
		if strings.EqualFold(book.Title, title) {
			if book.Borrowed > 0 {
				book.Borrowed--
				return "The book was returned successfully"
			} else {
				return "This book wasn't borrowed, please check the title"
			}
		}
	}
	return "This book is not available in the library, please check the title"
}

func (library *Library) SearchBookByAuthor(author string) []*Book {
	books := make([]*Book, 0)
	for _, book := range library.Books {
		if strings.Contains(strings.ToLower(book.Author), strings.ToLower(author)) {
			books = append(books, book)
		}
	}
	return books
}

func (library *Library) SearchBookByTitle(title string) []*Book {
	books := make([]*Book, 0)
	for _, book := range library.Books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
			books = append(books, book)
		}
	}
	return books
}

func (library *Library) BorrowedBooks() []*Book {
	books := make([]*Book, 0)
	for _, book := range library.Books {
		if book.Borrowed > 0 {
			books = append(books, book)
		}
	}
	return books
}
