package calculator_test

import (
	"calculator"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type testCase struct {
	a, b        float64
	want        float64
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 1, b: 4, want: -3},
		{a: 5, b: 10, want: -5},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 5, b: 2, want: 10},
		{a: 1, b: 1, want: 1},
		{a: 5, b: 10, want: 50},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	testCases := []testCase{
		{a: 6, b: 3, want: 2, errExpected: false},
		{a: 4, b: 2, want: 2, errExpected: false},
		{a: 2, b: 2, want: 1, errExpected: false},
		{a: 6, b: 0, want: 0, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived)
		} else if tc.errExpected == errReceived {
			if tc.want != got {
				t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
			} else {
				fmt.Printf("Test passed for Divide(%f, %f): want %f, got %f\n", tc.a, tc.b, tc.want, got)
			}
		}
	}
}

func TestAddRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var max float64 = 100.0
	var min float64 = 0

	i := 0
	for i <= 100 {
		first_value := rand.Float64() * (max - min)
		second_value := rand.Float64() * (max - min)
		sum := first_value + second_value

		testCases := []testCase{
			{a: first_value, b: second_value, want: sum},
		}

		got := calculator.Add(first_value, second_value)

		for _, tc := range testCases {
			if tc.want != got {
				t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
			} else {
				//fmt.Printf("Add(%f, %f): want %f, got %f --> Test successful!\n", tc.a, tc.b, tc.want, got)
			}
		}
		i++
	}
}

func TestSqrt(t *testing.T) {
	type testCaseSqrt struct {
		a           float64
		want        float64
		errExpected bool
	}
	testCases := []testCaseSqrt{
		{a: 9, want: 3, errExpected: false},
		{a: 81, want: 9, errExpected: false},
		//{a: -9, want: 0, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("Sqrt(%f): unexpected error: %s", tc.a, err)
		} else if tc.errExpected == errReceived {
			if err != nil {
				t.Fatalf("Expected Error received: %s", err)
			}
			if tc.want != got {
				t.Errorf("Sqrt(%f): Result doesn't match expectation. want: %f, got: %f\n", tc.a, tc.want, got)
			}
		}
	}
}
