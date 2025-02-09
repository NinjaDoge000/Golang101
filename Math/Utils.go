package mathutils

import "math"

// only functions starting with capital letter are exported in go.

// this function is exported
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * Factorial(x-1)
}

// this function is not exported
func pow(x float64, y float64) float64 {
	return math.Pow(x, y)
}
