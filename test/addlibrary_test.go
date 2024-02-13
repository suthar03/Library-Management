package test

import (
	"pinelabs/models"
	"testing"
)

func TestAddLibraryWithValidData(t *testing.T) {
	id := "1"
	wantID := "library added successfully"
	bookStore := models.NewBookStore()
	library := models.NewLibrary(id)
	msg := bookStore.AddLibrary(library)
	if msg != wantID {
		t.Fatalf(`AddLibrary("1") = %q, want match for %#q, nil`, msg, wantID)
	}
}

func TestAddLibraryWithInValidData(t *testing.T) {
	wantID := "library added successfully"
	bookStore := models.NewBookStore()
	msg := bookStore.AddLibrary(nil)
	if msg != wantID {
		t.Fatalf(`AddLibrary("1") = %q, want match for %#q, nil`, msg, wantID)
	}
}
