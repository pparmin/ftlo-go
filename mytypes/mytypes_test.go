package mytypes_test

import (
	"mytypes"
	"strings"
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

func TestStringsBuilder(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := sb.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	wantLen := 15
	gotLen := sb.Len()
	if wantLen != gotLen {
		t.Errorf("%q want len %d, got %d", sb.String(), wantLen, gotLen)
	}
}

func TestMyBuilderHello(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	mb.WriteString("Hello, ")
	mb.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := mb.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

	wantLen := 15
	gotLen := mb.Len()
	if wantLen != gotLen {
		t.Errorf("%q want len %d, got %d", mb.String(), wantLen, gotLen)
	}
}

func TestStringUpperase(t *testing.T) {
	t.Parallel()
	var su mytypes.StringUppercaser
	su.Contents.WriteString("Hello, ")
	su.Contents.WriteString("Gophers!")
	want := "HELLO, GOPHERS!"
	got := su.ToUpper()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
