package main

import (
	"fmt"
	"strings"
)

func parenthesis(expr string) (string, error) {
	// find the last ( so innernest parentheses are resolved first
	openIndex := strings.LastIndex(expr, "(")

	// find the first ) after the last ( to get the matching pair
	closeIndex := strings.Index(expr[openIndex:], ")") + openIndex

	// if no matching ) was found then return error
	if closeIndex < openIndex {
		fmt.Println("Error: mismatched parentheses")
		return "", fmt.Errorf("mismatched parentheses")
	}

	// extract the expression inside the parentheses (eexample: "2+3")
	inner := expr[openIndex+1 : closeIndex]

	// gets the raw number for the expression
	result, err := evaluate(inner)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	// rebuilds expression with result of expression taking the place of the parentheses + expression
	// (example: "2+3" becomes "5")
	expr = expr[:openIndex] + fmt.Sprintf("%g", result) + expr[closeIndex+1:]
	return expr, nil
}
