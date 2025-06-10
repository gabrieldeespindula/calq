package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"regexp"
)

func Tokenize(expr, lastResult string) ([]string, error) {
	expr = strings.TrimSpace(expr)
	if lastResult != "" && strings.ContainsAny(string(expr[0]), ADD+SUB+MUL+DIV) {
		expr = lastResult + expr
	}

	re := regexp.MustCompile(`\d+`)
	expr = re.ReplaceAllStringFunc(expr, func(match string) string {
		return "#" + match + "#"
	})

	replacer := strings.NewReplacer(
		ADD, " "+ADD+" ",
		MUL, " "+MUL+" ",
		DIV, " "+DIV+" ",
		SUB, " "+SUB+" ",
	)

	expr = replacer.Replace(expr)

	replacer = strings.NewReplacer(
		"  " + SUB + " #", " "+SUB,
		"#", "",
	)

	expr = replacer.Replace(expr)

	tokens := strings.Fields(expr)

	if len(tokens) == 0 {
		return nil, errors.New("empty expression")
	}

	if len(tokens) > 0 && tokens[0] == "-" {
		newToken := "-" + tokens[1]
		tokens = append([]string{newToken}, tokens[2:]...)
	}

	return tokens, nil
}

func Evaluate(parts []string) (float64, error) {
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
	operatorFn, ok := Operations[operator]
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
	return Evaluate(newParts)
}

func findNextOperatorIndex(parts []string) int {
	if i := indexOfFirst(parts, []string{MUL, DIV}); i != -1 {
		return i
	}
	return indexOfFirst(parts, []string{ADD, SUB})
}

func extractOperands(parts []string, opIndex int) ([]string, error) {
	if opIndex <= 0 || opIndex+1 >= len(parts) {
		return nil, errors.New("invalid expression format")
	}
	return parts[opIndex-1 : opIndex+2], nil
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
