package main

import "fmt"

func NetPrice(price int, discountPercent int) int {
	newPrice := price - (price / discountPercent)
	return newPrice
}

func main() {
	type Book struct {
		author          string
		title           string
		edition         int
		priceCents      int
		isSeries        bool
		nInSeries       int
		titleSeries     string
		onOffer         bool
		discountPercent int
		netPrice        int
	}

	var data = Book{
		author:          "John Arundel",
		title:           "Data",
		edition:         1,
		priceCents:      1000,
		isSeries:        true,
		nInSeries:       2,
		titleSeries:     "For the Love of Go",
		onOffer:         false,
		discountPercent: 10,
		netPrice:        0,
	}

	// to be used as a placeholder for the books in the series that
	// have not been added yet
	var placeholder = Book{
		author:          "empty",
		title:           "emtpy",
		edition:         0,
		priceCents:      0,
		isSeries:        false,
		nInSeries:       0,
		titleSeries:     "empty",
		onOffer:         false,
		discountPercent: 0,
		netPrice:        0,
	}

	var test = Book{
		author:  "test",
		edition: 1,
		onOffer: true,
	}

	series := []Book{placeholder, data, placeholder, placeholder}
	fmt.Println(series)

	fmt.Println("TEST:", test)

	data.netPrice = NetPrice(data.priceCents, data.discountPercent)

	fmt.Printf("Calculating new price of book %s: old price: %d euro(s), discount: %d percent \n--> new price: %d euro(s)\n",
		data.title, data.priceCents/100, data.discountPercent, data.netPrice/100)
}
