package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestTwice(t *testing.T) {
	t.Parallel()
	input := mytypes.MyInt(9)
	want := mytypes.MyInt(18)
	got := input.Twice()

	if want != got {
		t.Errorf("twice %d: want %d, got %d", input, want, got)
	}
}

func TestMyStringLen(t *testing.T) {
	t.Parallel()
	input := mytypes.MyString("Hello")
	want := 5
	got := input.Len()

	if want != got {
		t.Errorf("Len: Lengths returned for target string %q do not match. want: %d, got: %d", input, want, got)
	}
}