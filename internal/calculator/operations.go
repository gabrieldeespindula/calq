package calculator

import (
	"errors"
)

const (
	ADD = "+"
	SUB = "-"
	MUL = "*"
	DIV = "/"
)

type Operation func(float64, float64) (float64, error)

var Operations = map[string]Operation{
	ADD: add,
	SUB: subtract,
	MUL: multiply,
	DIV: divide,
}

func add(a, b float64) (float64, error) {
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
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}
