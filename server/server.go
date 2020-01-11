package main

// Add takes two float64s and returns the sum of them
func Add(x, y float64) float64 {
	return x + y
}

// AddSlice takes n float64 inputs as a slice and returns the sum of them
func AddSlice(inputs []float64) float64 {
	var sum float64
	for _, input := range inputs {
		sum += input
	}
	return sum
}

// Subtract takes two float64s and returns the subtraction of the second from the first
func Subtract(x, y float64) float64 {
	return x - y
}

// SubtractSlice takes n inputs as a slice and returns the ordered subtraction from left to right
func SubtractSlice(inputs []float64) float64 {
	sum := inputs[0]
	for _, input := range inputs[1:len(inputs)] {
		sum -= input
	}
	return sum
}

// Multiply takes two float64s and returns the multiplied result of them
func Multiply(x, y float64) float64 {
	return x * y
}

// MultiplySlice takes n float64 inputs as slice and returns the ordered multiplication from left to right
func MultiplySlice(inputs []float64) float64 {
	sum := inputs[0]
	for _, input := range inputs[1:len(inputs)] {
		sum *= input
	}
	return sum
}

// Divide takes two float64s and returns the division of the first by the second
func Divide(x, y float64) float64 {
	return x / y
}

// DivideSlice takes n float64 inputs as slice and returns the ordered division from left to right
func DivideSlice(inputs []float64) float64 {
	sum := inputs[0]
	for _, input := range inputs[1:len(inputs)] {
		sum /= input
	}
	return sum
}