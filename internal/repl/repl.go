package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"calq/internal/calculator"
)

func Run(in io.Reader, out io.Writer) {
	var lastResult string
	reader := bufio.NewReader(in)

	fmt.Fprintln(out, "Welcome to calq! Type your expressions (type 'q' + Enter to quit, or press Ctrl+C):")
	for {
		fmt.Fprint(out, "> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(out, "Input error:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if strings.EqualFold(input, "q") {
			fmt.Fprintln(out, "Goodbye!")
			break
		}

		result, err := calculator.Calculate(lastResult, input)
		if err != nil {
			fmt.Fprintln(out, "Calculation error:", err)
			continue
		}

		lastResult = fmt.Sprintf("%f", result)
		fmt.Fprintln(out, result)
	}
}
