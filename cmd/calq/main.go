package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"calq/internal/calculator"
)

var Version = "dev"

func main() {
	var lastResult string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to calq! Type your expressions (Ctrl+C to exit):")
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		tokens, err := calculator.Tokenize(input, lastResult)
		if err != nil {
			fmt.Println("Tokenize error:", err)
			continue
		}

		result, err := calculator.Evaluate(tokens)
		if err != nil {
			fmt.Println("Evaluation error:", err)
			continue
		}

		lastResult = fmt.Sprintf("%f", result)
		fmt.Println(result)
	}
}
