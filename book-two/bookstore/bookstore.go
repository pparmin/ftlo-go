package bookstore

import "fmt"

type Book struct {
	Author      []string
	Title       string
	Description string
	PriceCents  int
	Edition     int
	IsSeries    bool
	NInSeries   bool
	TitleSeries string
	Featured    bool
}

var Books = []Book{{Title: "This is book 1"}, {Title: "This is book 2"}}

func GetAllBooks() []Book {
	fmt.Println("IN FUNCTION:", Books)
	var AllBooks = []Book{{Title: "This is the function"}}
	fmt.Println()
	for _, book := range Books {
		AllBooks = append(AllBooks, book)
	}
	return AllBooks
}

var NewBooks = GetAllBooks()

//fmt.Println(newBooks)
