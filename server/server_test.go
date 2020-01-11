package main

import (
	"github.com/google/go-cmp/cmp"
	"math"
	"testing"
)

// Define constants to use with test cases
const (
	x         float64 = 7.7
	y         float64 = 5.5
	tolerance         = 0.0000000001
)

// Borrowed code to test floating point numbers with tolerance
func testResult(result, expect float64) bool {
	opt := cmp.Comparer(func(result, expect float64) bool {
		diff := math.Abs(result - expect)
		mean := math.Abs(result+expect) / 2.0
		if math.IsNaN(diff / mean) {
			return true
		}
		return (diff / mean) < tolerance
	})
	if !cmp.Equal(result, expect, opt) {
		return false
	} else {
		return true
	}
}

// Test the "Add" function
func TestAddition(t *testing.T) {
	result := Add(x, y)
	expect := 13.2

	if !testResult(result, expect) {
		t.Errorf("Result (%v) expected (%v) from input x(%v), y(%v)", result, expect, x, y)
	}
}

// Test the "AddMultiple" function
func TestAdditionMultiple(t *testing.T) {
	inputs := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	result := AddMultiple(inputs)
	expect := 1.5

	if !testResult(result, expect) {
		t.Errorf("Result (%v) expected (%v) from input x(%v), y(%v)", result, expect, x, y)
	}
}

// Test the "Subtract" function
func TestSubtraction(t *testing.T) {
	result := Subtract(x, y)
	expect := 2.2

	if !testResult(result, expect) {
		t.Errorf("Result (%v) expected (%v) from input x(%v), y(%v)", result, expect, x, y)
	}
}

// Test the "SubtractMultiple" function
func TestSubtractionMultiple(t *testing.T) {
	inputs := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	result := SubtractMultiple(inputs)
	expect := -1.3

	if !testResult(result, expect) {
		t.Errorf("Result (%v) expected (%v) from input x(%v), y(%v)", result, expect, x, y)
	}
}

// Test the "Multiply" function
func TestMultiplication(t *testing.T) {
	result := Multiply(x, y)
	expect := 42.35

	if !testResult(result, expect) {
		t.Errorf("Result (%v) expected (%v) from input x(%v), y(%v)", result, expect, x, y)
	}
}

// Test the "MultiplyMultiple" function
func TestMultiplicationMultiple(t *testing.T) {
	inputs := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	result := MultiplyMultiple(inputs)
	expect := 0.0012

	if !testResult(result, expect) {
		t.Errorf("Result (%v) expected (%v) from input x(%v), y(%v)", result, expect, x, y)
	}
}

// Test the "Divide" function
func TestDivision(t *testing.T) {
	result := Divide(x, y)
	expect := 1.4

	if !testResult(result, expect) {
		t.Errorf("Result (%v) expected (%v) from input x(%v), y(%v)", result, expect, x, y)
	}
}

// Test the "DivideMultiple" function
func TestDivisionMultiple(t *testing.T) {
	inputs := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	result := DivideMultiple(inputs)
	expect := 8.333333333333333

	if !testResult(result, expect) {
		t.Errorf("Result (%v) expected (%v) from input x(%v), y(%v)", result, expect, x, y)
	}
}
