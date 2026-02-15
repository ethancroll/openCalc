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

	// it it has the plus sign, send it to the addition file
	if strings.Contains(expr, "+") {
		fmt.Println(operations.Add(expr))
	}
}
