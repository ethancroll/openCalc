package main

import (
	"strings"

	"github.com/ethancroll/openCalc/operations"
)

func evaluate(expr string) (float64, error) {
	// if it has the plus sign, send it to the addition file
	if strings.Contains(expr, "+") {
		result, err := operations.Add(expr)
		if err != nil {
			return 0, err
		} else {
			return result, nil
		}
	}

	// if it has the multiplication sign, send it to the multiply file
	if strings.Contains(expr, "*") {
		result, err := operations.Multiply(expr)
		if err != nil {
			return 0, err
		} else {
			return result, nil
		}
	}

	return 0, nil
}
