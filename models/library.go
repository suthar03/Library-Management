package models

import (
	"strings"
)

type Library struct {
	ID    string
	Books []*Book
}

func NewLibrary(id string) *Library {
	return &Library{ID: id}
}

type BookStore struct {
	Libraries []*Library
}

func NewBookStore() *BookStore {
	return &BookStore{}
}

func (store *BookStore) AddLibrary(library *Library) string {
	if library == nil {
		return "invalid library object"
	}
	store.Libraries = append(store.Libraries, library)
	return "library added successfully"
}

func (store *BookStore) GetLibrary(id string) *Library {
	for _, lib := range store.Libraries {
		if strings.EqualFold(lib.ID, id) {
			return lib
		}
	}
	return nil
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
	return "This book doesn't belong to this library, please check the title"
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
