package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
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
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()

	t.Run("existing entry", func(t *testing.T) {
		catalog := map[int]bookstore.Book{
			1: {ID: 1, Title: "For the Love of Go"},
			2: {ID: 2, Title: "The Power of Go: Tools"},
		}
		want := bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}
		got, err := bookstore.GetBook(catalog, 2)

		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(want, got))
		}
	})
	t.Run("non-existent entry", func(t *testing.T) {
		catalog := map[int]bookstore.Book{
			1: {ID: 1, Title: "For the Love of Go"},
			2: {ID: 2, Title: "The Power of Go: Tools"},
		}
		want := bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}
		got, err := bookstore.GetBook(catalog, 3)

		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(want, got))
		}
	})
}
