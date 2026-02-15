package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// joins everything into one string
	expr := strings.Join(os.Args[1:], "")

	// gets rid of spaces (if any)
	expr = strings.ReplaceAll(expr, " ", "")

	// resolve all parenthesized sub-expressions (innermost first)
	for strings.Contains(expr, "(") {
		// find the last '(' to handle innermost parentheses first
		openIndex := strings.LastIndex(expr, "(")
		// find the matching ')' after it
		closeIndex := strings.Index(expr[openIndex:], ")") + openIndex

		if closeIndex < openIndex {
			fmt.Println("Error: mismatched parentheses")
			return
		}

		inner := expr[openIndex+1 : closeIndex]
		result, err := evaluate(inner)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// replace the parentheses and their contents with the result
		expr = expr[:openIndex] + fmt.Sprintf("%g", result) + expr[closeIndex+1:]
	}

	result, err := evaluate(expr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}
