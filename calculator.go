// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying the second
// from the first.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide ...
func Divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by 0")
	}
	return a / b, nil
}

// Sqrt ...
func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0.0, fmt.Errorf("error: %f, negative inputs are not allowed", x)
	}
	return math.Sqrt(x), nil
}
