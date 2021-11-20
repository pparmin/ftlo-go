package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("valid input", func(t *testing.T) {
		want := "1-2-3-4"
		card, err := creditcard.New("1-2-3-4")
		if err != nil {
			t.Fatal(err)
		}

		got := card.GetNumber()
		if got != want {
			t.Errorf("Wrong number received. Expected %q, got: %q", want, got)
		}
	})
	t.Run("invalid input", func(t *testing.T) {
		want := "1-2-3-4"
		card, err := creditcard.New("3-5-7-9")
		if err != nil {
			t.Fatal(err)
		}

		got := card.GetNumber()
		if got != want {
			t.Errorf("Wrong number received. Expected %q, got: %q", want, got)
		}
	})
	t.Run("empty input", func(t *testing.T) {
		want := "1-2-3-4"
		card, err := creditcard.New("")
		if err != nil {
			t.Fatal(err)
		}

		got := card.GetNumber()
		if got != want {
			t.Errorf("Wrong number received. Expected %q, got: %q", want, got)
		}
	})
}
