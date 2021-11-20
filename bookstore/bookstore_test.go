package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		Title:       "Spark Joy",
		Author:      []string{"Marie Kondo"},
		Description: "A tiny, cheerful Japanese woman explains tidying",
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:       "Test 1",
		Author:      []string{"Marie Kondo"},
		Description: "This is the first test",
		Copies:      2,
	}

	result, err := bookstore.Buy(b)
	want := b.Copies - 1

	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if got != want {
		t.Errorf("Buy: Wrong number of copies returned. Expected: %d, Got: %d\n",
			want, got)
	}
}

func TestBuyNoCopies(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:       "No Copies",
		Author:      []string{"Marie Kondo"},
		Description: "This is the second test",
		Copies:      0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("Expected error when buying book with zero copies left, but got nil instead")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}
	got := catalog.GetAllBooks()
	// sorting the result slice is necessary since maps are random by nature in Go
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	t.Run("existing entry", func(t *testing.T) {
		catalog := bookstore.Catalog{
			1: {ID: 1, Title: "For the Love of Go"},
			2: {ID: 2, Title: "The Power of Go: Tools"},
		}
		want := bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}
		got, err := catalog.GetBook(2)

		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
			t.Error(cmp.Diff(want, got))
		}
	})
	t.Run("non-existent entry", func(t *testing.T) {
		catalog := bookstore.Catalog{
			1: {ID: 1, Title: "For the Love of Go"},
			2: {ID: 2, Title: "The Power of Go: Tools"},
		}
		want := bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}
		got, err := catalog.GetBook(3)

		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
			t.Error(cmp.Diff(want, got))
		}
	})
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:           "For the Love of Go",
		PriceCents:      4000,
		DiscountPercent: 50,
	}
	want := 2000
	got := b.NetPriceCents()

	if got != want {
		t.Errorf("ERROR: The discounted price does not match the expectation")
		t.Errorf("Original price: %d, discount: %d%%, wanted discounted price: %d, got: %d", b.PriceCents, b.DiscountPercent, want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	t.Run("valid input", func(t *testing.T) {
		b := bookstore.Book{
			Title:      "For the Love of Go",
			PriceCents: 4000,
		}
		want := 2000
		err := b.SetPriceCents(2000)

		if err != nil {
			t.Fatal(err)
		}

		got := b.PriceCents
		if got != want {
			t.Errorf("ERROR: The newly set price does not match the expectation")
			t.Errorf("Original price: %d, wanted new price: %d, got: %d", b.PriceCents, want, got)
		}
	})

	t.Run("invalid input", func(t *testing.T) {
		b := bookstore.Book{
			Title:      "For the Love of Go",
			PriceCents: 4000,
		}
		want := 2000
		err := b.SetPriceCents(-10)

		if err != nil {
			t.Fatal(err)
		}

		got := b.PriceCents
		if got != want {
			t.Errorf("ERROR: The newly set price does not match the expectation")
			t.Errorf("Original price: %d, wanted new price: %d, got: %d", b.PriceCents, want, got)
		}
	})
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	t.Run("valid input", func(t *testing.T) {
		b := bookstore.Book{
			Title: "For the Love of Go",
		}

		categories := []bookstore.Category{
			bookstore.CategoryAutobiography,
			bookstore.CategoryLargePrintRomance,
			bookstore.CategoryParticlePhysics,
		}

		for _, cat := range categories {
			err := b.SetCategory(cat)
			if err != nil {
				t.Fatal(err)
			}
			got := b.GetCategory()

			if cat != got {
				t.Errorf("ERROR: Category mismatch. Expected %d, got: %d", cat, b.GetCategory())
			}
		}
	})

	t.Run("invalid input", func(t *testing.T) {
		b := bookstore.Book{
			Title: "For the Love of Go",
		}
		err := b.SetCategory(999)

		if err == nil {
			t.Fatal(err)
		}
	})
}
