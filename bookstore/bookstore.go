package bookstore

import (
	"errors"
	"fmt"
)

const (
	CategoryAutobiography = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

type Book struct {
	ID              int
	Author          []string
	Title           string
	Description     string
	Copies          int
	PriceCents      int
	DiscountPercent int
	category        Category
}

type Catalog map[int]Book

type Category int

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

func (b *Book) SetCategory(category Category) error {
	if !validCategory[category] {
		return fmt.Errorf("unknown category: %q", category)
	}

	/* My solution
	/*switch category {
	case CategoryAutobiography, CategoryLargePrintRomance, CategoryParticlePhysics:
		b.category = category
	default:
		return fmt.Errorf("a non-existent category was passed")
	}*/
	return nil
}

func (b Book) GetCategory() Category {
	return b.category
}
