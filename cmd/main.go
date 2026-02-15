package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ethancroll/openCalc/operations"
)

func main() {
	// joins everything into one string
	expr := strings.Join(os.Args[1:], "")

	// gets rid of spaces (if any)
	expr = strings.ReplaceAll(expr, " ", "")

	// if it has the plus sign, send it to the addition file
	if strings.Contains(expr, "+") {
		result, err := operations.Add(expr)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(result)
		}
	}
}
