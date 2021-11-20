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
	category        string
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

func (b *Book) SetPriceCents(newPrice int) error {
	if newPrice < 0 {
		return fmt.Errorf("negative new price was passed. Price: %d", newPrice)
	}
	b.PriceCents = newPrice
	return nil
}

func (b *Book) SetCategory(category string) error {
	if category != "Autobiography" {
		return fmt.Errorf("a non-existent category was passed. Category name: %q", category)
	}
	b.category = category
	return nil
}

func (b Book) GetCategory() string {
	return b.category
}
