package main

import (
	"fmt"
	"errors"
	"strings"
)

const (
	ADD = "+"
	SUB = "-"
	MUL = "*"
	DIV = "/"
)

func sum(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero is not allowed")
	}
	return a / b, nil
}

func expressionToParts(expression string) ([]string, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	expression = strings.ReplaceAll(expression, "+", " + ")
	expression = strings.ReplaceAll(expression, "-", " - ")
	expression = strings.ReplaceAll(expression, "*", " * ")
	expression = strings.ReplaceAll(expression, "/", " / ")

	parts := strings.Fields(expression)

	if len(parts) == 0 {
		return nil, errors.New("No valid parts found in the expression")
	}

	return parts, nil
}

func evaluateExpression(parts []string) (float64, error) {
	if len(parts) == 0 {
		return 0, errors.New("No valid parts found")
	}

	result := 0.0
	var currentOp string
	operations := map[string]func(float64, float64) (float64, error){
		ADD: sum,
		SUB: subtract,
		MUL: multiply,
		DIV: divide,
	}

	// [23, '+', 21]
	
	// get first 3 elements

	// const [one,two, three, ...rest] = parts

	return result, nil
}

func main() {
	for {
		var expression string

		// wait for user input
		fmt.Scanln(&expression)

		parts, err := expressionToParts(expression)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		evaluateExpression, err := evaluateExpression(parts)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println(evaluateExpression)


		break
	}
}
