package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	loadDotEnv(".env")

	// joins everything into one string
	expr := strings.Join(os.Args[1:], "")

	// gets rid of spaces (if any)
	expr = strings.ReplaceAll(expr, " ", "")

	// all to lowercase
	expr = strings.ToLower(expr)

	if expr == "" {
		fmt.Println("Welcome to openCalc. Type in an expression with the command to use.")
		return
	} else {
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
}

func loadDotEnv(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if len(value) >= 2 {
			if (value[0] == '"' && value[len(value)-1] == '"') || (value[0] == '\'' && value[len(value)-1] == '\'') {
				value = value[1 : len(value)-1]
			}
		}
		if key == "" {
			continue
		}
		if _, exists := os.LookupEnv(key); !exists {
			_ = os.Setenv(key, value)
		}
	}
}
