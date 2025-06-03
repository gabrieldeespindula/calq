package main

import (
	"fmt"
	"errors"
)

func sum(a, b int) (int, error) {
	return a + b, nil
}

func subtract(a, b int) (int, error) {
	return a - b, nil
}

func multiply(a, b int) (int, error) {
	return a * b, nil
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não permitida")
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

	operations := map[string]func(int, int) (int, error){
		ADD: sum,
		SUB: subtract,
		MUL: multiply,
		DIV: divide,
	}

	var firstNumber, secondNumber int
	var operation string

	fmt.Print("Type the first number: ")
	fmt.Scanln(&firstNumber)

	fmt.Println("Choose an operation: addition(+), subtraction(-), multiplication(*), division(/)")
	fmt.Scanln(&operation)

	fmt.Print("Type the second number: ")
	fmt.Scanln(&secondNumber)

	opFunc, ok := operations[operation]

	if !ok {
		fmt.Println("Operação inválida! Por favor escolha +, -, * ou /")
		return
	}

	result, err := opFunc(firstNumber, secondNumber)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Printf("Resultado: %d\n", result)
}
