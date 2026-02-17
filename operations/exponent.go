package operations

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// takes the exponentiation expression, and returns the result
func Exponent(expr string) (float64, error) {
	parts := strings.Split(expr, os.Getenv("OPENCALC_EXPONENT")) // ["8","9"]

	a, err := strconv.ParseFloat(parts[0], 64) // left side of expression
	if err != nil {
		return 0, fmt.Errorf("'%s' is not a valid number", parts[0])
	}

	b, err := strconv.ParseFloat(parts[1], 64) // right side of expression
	if err != nil {
		return 0, fmt.Errorf("'%s' is not a valid number", parts[1])
	}

	return math.Pow(a, b), nil
}
