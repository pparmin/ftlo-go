// Package evalulator provides a library for
// simple evaluations in Go.
package calculator

import (
	"fmt"
	"math"
	"strconv"
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
			counter++
		} else {
			result -= input
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
			counter++
		} else {
			result /= input
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

// Evaluate parses an arithmetic expression provided as a string and
// performs the calculation. It then returns the result of the expression.
func Evaluate(expression string) (float64, error) {
	type evaluation struct {
		placeholder string
		buffer      string
		mode        string
	}

	eval := evaluation{
		placeholder: "", buffer: "", mode: "",
	}

	fmt.Printf("Expression: %s\n", expression)
	for _, char := range expression {
		switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
			eval.placeholder += string(char)
		case ' ':
			continue
		case '+':
			eval.mode = "Add"
			eval.buffer = eval.placeholder
			eval.placeholder = ""
		case '-':
			eval.mode = "Subtract"
			eval.buffer = eval.placeholder
			eval.placeholder = ""
		case '*':
			eval.mode = "Multiply"
			eval.buffer = eval.placeholder
			eval.placeholder = ""
		case '/':
			eval.mode = "Divide"
			eval.buffer = eval.placeholder
			eval.placeholder = ""
		default:
			return 0, fmt.Errorf("something went wrong while parsing the expression")
		}
	}

	switch eval.mode {
	case "Add":
		firstValue, err := strconv.ParseFloat(eval.buffer, 64)
		secondValue, errTwo := strconv.ParseFloat(eval.placeholder, 64)
		if err != nil || errTwo != nil {
			return 0, fmt.Errorf("%s(): conversion from string to float failed", eval.mode)
		}
		result := Add(firstValue, secondValue)
		return result, nil

	case "Subtract":
		firstValue, err := strconv.ParseFloat(eval.buffer, 64)
		secondValue, errTwo := strconv.ParseFloat(eval.placeholder, 64)

		if err != nil || errTwo != nil {
			return 0, fmt.Errorf("conversion from string to float failed")
		}
		result := Subtract(firstValue, secondValue)
		return result, nil

	case "Multiply":
		firstValue, err := strconv.ParseFloat(eval.buffer, 64)
		secondValue, errTwo := strconv.ParseFloat(eval.placeholder, 64)

		if err != nil || errTwo != nil {
			return 0, fmt.Errorf("conversion from string to float failed")
		}
		result := Multiply(firstValue, secondValue)
		return result, nil

	case "Divide":
		firstValue, err := strconv.ParseFloat(eval.buffer, 64)
		secondValue, errTwo := strconv.ParseFloat(eval.placeholder, 64)

		if err != nil && errTwo != nil {
			return 0, fmt.Errorf("conversion from string to float failed")
		}
		result, div_err := Divide(firstValue, secondValue)
		if div_err != nil {
			return result, fmt.Errorf("ERROR: Division by zero is undefined")
		}
		return result, nil
	default:
		return 0, fmt.Errorf("something went wrong while calculating the expression")
	}
}
