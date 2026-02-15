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
		result, err := parenthesis(expr)
		if err != nil {
			return
		}
		expr = result
	}

	result, err := evaluate(expr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}
