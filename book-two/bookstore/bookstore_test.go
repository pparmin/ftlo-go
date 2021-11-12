package bookstore_test

import (
	"bookstore"
	"testing"
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
		Copies:      8,
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
