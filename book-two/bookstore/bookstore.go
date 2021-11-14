package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	ID          int
	Author      []string
	Title       string
	Description string
	Copies      int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies -= 1
	return b, nil
}

func GetAllBooks(catalog map[int]Book) []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}
	return result
}

func GetBook(catalog map[int]Book, ID int) (Book, error) {
	book, ok := catalog[ID]
	if !ok {
		return Book{}, fmt.Errorf("could not find a book with ID %d", ID)
	}
	return book, nil
}
