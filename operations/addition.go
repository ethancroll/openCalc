package operations

import (
	"strconv"
	"strings"
)

// takes the addition expression, and returns the result
func Add(expr string) float64 {
	parts := strings.Split(expr, "+") // ["8","9"]

	a, _ := strconv.ParseFloat(parts[0], 64) // left side of expression
	b, _ := strconv.ParseFloat(parts[1], 64) // right side of expression

	return a + b
}
