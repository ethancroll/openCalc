package operations

import (
	"fmt"
	"strconv"
	"strings"
)

// takes the addition expression, and returns the result
func Add(expr string) (float64, error) {
	parts := strings.Split(expr, "+") // ["8","9"]

	a, err := strconv.ParseFloat(parts[0], 64) // left side of expression
	if err != nil {
		return 0, fmt.Errorf("'%s' is not a valid number", parts[0])
	}

	b, err := strconv.ParseFloat(parts[1], 64) // right side of expression
	if err != nil {
		return 0, fmt.Errorf("'%s' is not a valid number", parts[1])
	}

	return a + b, nil
}
