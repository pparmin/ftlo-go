// Package calculator provides a library for
// simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes a variable amount of numbers and returns the
// result of adding them together.
func Add(inputs ...float64) float64 {
	var result float64 = 0
	for _, input := range inputs {
		result += input
	}
	return result
}

// Subtract takes a variable amount of numbers, and
// returns the result of subtracting them.
func Subtract(inputs ...float64) float64 {
	counter := 0
	var result float64
	for _, input := range inputs {
		// set original value to first input
		if counter == 0 {
			result = input
			fmt.Printf("Subtract: i: %d Result is now %f\n", counter, result)
			counter++
		} else {
			result -= input
			fmt.Printf("Subtract: i: %d Result is now %f\n", counter, result)
			counter++
		}
	}
	return result
}

// Subtract takes a variable amount of numbers, and
// returns the result of multiplying them.
func Multiply(inputs ...float64) float64 {
	var result float64 = 1
	for _, input := range inputs {
		result *= input
	}
	return result
}

// Subtract takes a variable amount of numbers, and
// returns the result of dividing them.
func Divide(inputs ...float64) (float64, error) {
	counter := 0
	var result float64
	for _, input := range inputs {
		if input == 0 {
			return 0, fmt.Errorf("ERROR: Division by zero is undefined")
		}
		if counter == 0 {
			result = input
			fmt.Printf("Divide: i: %d Result is now %f\n", counter, result)
			counter++
		} else {
			result /= input
			fmt.Printf("Divide i: %d Result is now %f\n", counter, result)
			counter++
		}
	}
	return result, nil
}

// Subtract takes a number, and returns the square root.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("square root of negative value is undefined")
	}
	return math.Sqrt(a), nil
}
