package main

import (
	"fmt"
	"errors"
	"strings"
	"strconv"
)

const (
	ADD = "+"
	SUB = "-"
	MUL = "*"
	DIV = "/"
)

func indexOfAny(slice []string, targets []string) int {
	for i, v := range slice {
		for _, t := range targets {
			if v == t {
				return i
			}
		}
	}
	return -1
}

func aroundIndex(slice []string, index int) []string {
	var result []string

	if index > 0 {
		result = append(result, slice[index-1])
	}

	if index >= 0 && index < len(slice) {
		result = append(result, slice[index])
	}

	if index+1 < len(slice) {
		result = append(result, slice[index+1])
	}

	return result
}

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

func expressionToParts(expression string, lastResult string) ([]string, error) {
	if lastResult != "" {
		if len(expression) > 0 {
			firstChar := string(expression[0])
			indexOfAny := indexOfAny([]string{firstChar}, []string{ADD, SUB, MUL, DIV})

			if indexOfAny != -1 {
				expression = lastResult + expression
			}
		}
	}

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

	indexOfAny := indexOfAny(parts, []string{ADD, SUB, MUL, DIV})

	if indexOfAny == -1 {
		return 0, errors.New("No valid operator found in the expression")
	}

	aroundIndex := aroundIndex(parts, indexOfAny)

	if len(aroundIndex) != 3 {
		return 0, errors.New("Invalid expression format")
	}

	// aroundIndex is something like this: ['1', '+', '2']
	// we need to convert the first and last elements to float64
	firstNum, err := strconv.ParseFloat(aroundIndex[0], 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid number: %s", aroundIndex[0])
	}
	secondNum, err := strconv.ParseFloat(aroundIndex[2], 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid number: %s", aroundIndex[2])
	}
	currentOp = aroundIndex[1]

	// now we can perform the operation
	opFunc, exists := operations[currentOp]
	if !exists {
		return 0, fmt.Errorf("Invalid operator: %s", currentOp)
	}
	result, err = opFunc(firstNum, secondNum)
	if err != nil {
		return 0, err
	}

	// Now we need to replace the evaluated part in the original parts slice
	start := indexOfAny - 1
	end := indexOfAny + 2
	newParts := append(
		append([]string{}, parts[:start]...), // antes de b
		append([]string{fmt.Sprintf("%f", result)}, parts[end:]...)..., // p + depois de d
	)

	if len(newParts) == 1 {
		return strconv.ParseFloat(newParts[0], 64)
	}

	return evaluateExpression(newParts)
}

func main() {
	var lastResult string
	for {
		var expression string

		// wait for user input
		fmt.Scanln(&expression)

		parts, err := expressionToParts(expression, lastResult)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		evaluateExpression, err := evaluateExpression(parts)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		lastResult = fmt.Sprintf("%f", evaluateExpression)

		fmt.Println(evaluateExpression)
	}
}
