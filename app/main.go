package main

import "fmt"

func main() {
	var a, b int
	var operation string

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&a)

	fmt.Println("Escolha uma operação: soma(+), subtração(-), multiplicação(*), divisão(/)")
	fmt.Scanln(&operation)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&b)

	switch operation {
	case "+":
		fmt.Printf("Resultado: %d\n", a+b)
	case "-":
		fmt.Printf("Resultado: %d\n", a-b)
	case "*":
		fmt.Printf("Resultado: %d\n", a*b)
	case "/":
		if b != 0 {
			fmt.Printf("Resultado: %d\n", a/b)
		} else {
			fmt.Println("Erro: Divisão por zero não é permitida.")
		}
	default:
		fmt.Println("Operação inválida. Por favor, escolha entre +, -, * ou /.")
		return
	}
}
