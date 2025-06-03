package main

import (
	"fmt"
	"errors"
	"strings"
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

func main() {
	const (
		ADD = "+"
		SUB = "-"
		MUL = "*"
		DIV = "/"
	)

	operations := map[string]func(float64, float64) (float64, error){
		ADD: sum,
		SUB: subtract,
		MUL: multiply,
		DIV: divide,
	}

	for {
		var firstNumber, secondNumber float64
		var operation string

		fmt.Print("Type the first number: ")
		fmt.Scanln(&firstNumber)

		fmt.Println("Choose an operation: addition(+), subtraction(-), multiplication(*), division(/)")
		fmt.Scanln(&operation)

		fmt.Print("Type the second number: ")
		fmt.Scanln(&secondNumber)

		opFunc, ok := operations[operation]

		if !ok {
			fmt.Println("Invalid operation. Please choose one of the following: +, -, *, /")
			return
		}

		result, err := opFunc(firstNumber, secondNumber)
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("Please try again with valid numbers")
			return
		}

		fmt.Printf("The result of %.2f %s %.2f is: %.2f\n", firstNumber, operation, secondNumber, result)
		
		var choice string
		fmt.Print("Do you want to calculate again? (y/n): ")
		fmt.Scanln(&choice)

		if strings.ToLower(choice) != "y" {
			fmt.Println("Thank you for using the calculator!")
			break
		}
	}
}
