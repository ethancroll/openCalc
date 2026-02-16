package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ethancroll/openCalc/operations"
)

func evaluate(expr string) (float64, error) {
	// if it has the plus sign, send it to the addition file
	if strings.Contains(expr, "+") {
		result, err := operations.Add(expr)
		if err != nil {
			return 0, err
		}
		return result, nil
	}

	// if it has the multiplication sign, send it to the multiply file
	if strings.Contains(expr, "*") {
		result, err := operations.Multiply(expr)
		if err != nil {
			return 0, err
		}
		return result, nil
	}

	// if it has the division sign, send it to the divide file
	if strings.Contains(expr, "/") {
		result, err := operations.Divide(expr)
		if err != nil {
			return 0, err
		}
		return result, nil
	}

	// if it has the subtraction sign, send it to the subtract file
	if strings.Contains(expr, "-") {
		result, err := operations.Subtract(expr)
		if err != nil {
			return 0, err
		}
		return result, nil
	}

	// if it has an e, then send it to the exponenet operation
	if strings.Contains(expr, "e") {
		result, err := operations.Exponent(expr)
		if err != nil {
			return 0, err
		}
		return result, nil
	}

	// no operator found — try to parse as a plain number
	num, err := strconv.ParseFloat(expr, 64)
	if err != nil {
		return 0, fmt.Errorf("'%s' is not a valid expression", expr)
	}
	return num, nil
}
