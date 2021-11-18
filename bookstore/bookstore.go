package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	ID              int
	Author          []string
	Title           string
	Description     string
	Copies          int
	PriceCents      int
	DiscountPercent int
}

type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies -= 1
	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

func (c Catalog) GetBook(ID int) (Book, error) {
	book, ok := c[ID]
	if !ok {
		return Book{}, fmt.Errorf("could not find a book with ID %d", ID)
	}
	return book, nil
}

func (book Book) NetPriceCents() int {
	saving := book.PriceCents * book.DiscountPercent / 100
	return book.PriceCents - saving
}
