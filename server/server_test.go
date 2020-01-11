package main

import (
	"testing"
	"math"
	"github.com/google/go-cmp/cmp"
)

// Define constants to use with test cases
const (
	x float64 = 7.7
	y float64 = 5.5
	tolerance = 0.0000000001
)

// Borrowed code to test floating point numbers with tolerance
func testResult(result, expect float64) bool {
	opt := cmp.Comparer(func(result, expect float64) bool {
	    diff := math.Abs(result - expect)
	    mean := math.Abs(result + expect) / 2.0
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

// Test the "AddSlice" function
func TestAdditionSlice(t *testing.T) {
	inputs := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	result := AddSlice(inputs)
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

// Test the "SubtractSlice" function
func TestSubtractionSlice(t *testing.T) {
	inputs := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	result := SubtractSlice(inputs)
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

// Test the "MultiplySlice" function
func TestMultiplicationSlice(t *testing.T) {
	inputs := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	result := MultiplySlice(inputs)
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

// Test the "DivideSlice" function
func TestDivisionSlice(t *testing.T) {
	inputs := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	result := DivideSlice(inputs)
	expect := 8.333333333333333

	if !testResult(result, expect) {
		t.Errorf("Result (%v) expected (%v) from input x(%v), y(%v)", result, expect, x, y)
	}
}
