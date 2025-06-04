package main

import (
	"bufio"
	"os"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	ADD = "+"
	SUB = "-"
	MUL = "*"
	DIV = "/"
)

var operations = map[string]func(float64, float64) (float64, error){
	ADD: add,
	SUB: subtract,
	MUL: multiply,
	DIV: divide,
}

func add(a, b float64) (float64, error)      { return a + b, nil }
func subtract(a, b float64) (float64, error) { return a - b, nil }
func multiply(a, b float64) (float64, error) { return a * b, nil }
func divide(a, b float64) (float64, error) {
	if b == 0 { 
		return 0, errors.New("division by zero is not allowed") 
	}
	return a / b, nil
}

func indexOfFirst(slice []string, targets []string) int {
	for i, v := range slice {
		for _, t := range targets {
			if v == t {
				return i
			}
		}
	}
	return -1
}

func extractOperands(parts []string, opIndex int) ([]string, error) {
	if opIndex <= 0 || opIndex+1 >= len(parts) {
		return nil, errors.New("invalid expression format")
	}
	return parts[opIndex-1 : opIndex+2], nil
}

func tokenizeExpression(expr, lastResult string) ([]string, error) {
	expr = strings.TrimSpace(expr)
	if lastResult != "" && strings.ContainsAny(string(expr[0]), ADD+SUB+MUL+DIV) {
		expr = lastResult + expr
	}

	replacer := strings.NewReplacer(
		ADD, " "+ADD+" ",
		SUB, " "+SUB+" ",
		MUL, " "+MUL+" ",
		DIV, " "+DIV+" ",
	)
	tokens := strings.Fields(replacer.Replace(expr))
	if len(tokens) == 0 {
		return nil, errors.New("empty expression")
	}
	return tokens, nil
}

func findNextOperatorIndex(parts []string) int {
	if i := indexOfFirst(parts, []string{MUL, DIV}); i != -1 { 
		return i 
	}
	return indexOfFirst(parts, []string{ADD, SUB})
}

func evaluate(parts []string) (float64, error) {
	if len(parts) == 1 { 
		return strconv.ParseFloat(parts[0], 64) 
	}

	opIndex := findNextOperatorIndex(parts)
	if opIndex == -1 {
		return 0, errors.New("no operator found")
	}

	segment, err := extractOperands(parts, opIndex)
	if err != nil {
		return 0, err
	}

	firstNumber, err := strconv.ParseFloat(segment[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", segment[0])
	}
	secondNumber, err := strconv.ParseFloat(segment[2], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", segment[2])
	}

	operator := segment[1]
	operatorFn, ok := operations[operator]
	if !ok {
		return 0, fmt.Errorf("unsupported operator: %s", operator)
	}

	result, err := operatorFn(firstNumber, secondNumber)
	if err != nil {
		return 0, err
	}

	newParts := append(
		append(parts[:opIndex-1], fmt.Sprintf("%f", result)),
		parts[opIndex+2:]...,
	)
	return evaluate(newParts)
}

func main() {
	var lastResult string
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')

		fmt.Println("Input:", input)

		tokens, err := tokenizeExpression(input, lastResult)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		result, err := evaluate(tokens)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		lastResult = fmt.Sprintf("%f", result)
		fmt.Println(result)
	}
}
