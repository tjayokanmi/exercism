package darts


import (
	"math"
)

func Score(x, y float64) int {
	squared := (x * x) + (y * y)
	distance := math.Sqrt(squared)

	var result int

	if distance <= 1 {
		result += 10
	} else if distance > 1 && distance <= 5 {
		result += 5
	} else if distance > 5 && distance <= 10 {
		result += 1

	} else {
		result += 0
	}
	return result
}