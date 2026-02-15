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

	result, err := evaluate(expr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}
