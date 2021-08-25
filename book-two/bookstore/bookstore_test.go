package bookstore_test

import (
	"bookstore"
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		Title:       "Spark Joy",
		Author:      []string{"Marie Kondo"},
		Description: "A tiny, cheerful Japanese woman explains tidying",
		PriceCents:  1199,
		Edition:     1,
		IsSeries:    false,
		Featured:    true,
	}
}

func TestGetAllBooks(t *testing.T) {
	book1 := bookstore.Book{Title: "This is Book 1"}
	book2 := bookstore.Book{Title: "This is Book 2"}
	bookstore.Books = []bookstore.Book{book1, book2}

	want := bookstore.Books
	fmt.Println("IN TEST", want)
	got := bookstore.GetAllBooks()
	fmt.Println("GOT: ", got)

	// if !cmp.Equal(want, got) {
	// 	t.Error(cmp.Diff(want, got))
	// }
}
