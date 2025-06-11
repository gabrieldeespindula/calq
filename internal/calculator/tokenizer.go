package calculator

import (
	"errors"
	"strings"
	"regexp"
)

func Tokenize(lastResult string, expr string) ([]string, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return nil, errors.New("empty expression")
	}

	if lastResult != "" && strings.ContainsAny(string(expr[0]), ADD+SUB+MUL+DIV) {
		expr = lastResult + expr
	}

	expr = delimitNumbers(expr)
	expr = addSpaceBetweenOperators(expr)
	expr = handleNegativeNumbers(expr)
	expr = cleanNumbersDelimiters(expr)

	tokens := strings.Fields(expr)

	if tokens[0] == SUB && len(tokens) > 1 {
		tokens = handleNegativeNumbersAtStart(tokens)
	}

	return tokens, nil
}

func delimitNumbers(expr string) string {
	return regexp.MustCompile(`\d+`).ReplaceAllStringFunc(expr, func(m string) string {
		return "#" + m + "#"
	})
}

func addSpaceBetweenOperators(expr string) string {
	for _, op := range []string{ADD, SUB, MUL, DIV} {
		expr = strings.ReplaceAll(expr, op, " "+op+" ")
	}

	return expr
}

func handleNegativeNumbers(expr string) string {
	return strings.ReplaceAll(expr, "  "+SUB+" #", " "+SUB+"")
}


func cleanNumbersDelimiters(expr string) string {
	return strings.ReplaceAll(expr, "#", "")
}

func handleNegativeNumbersAtStart(tokens []string) []string {
	return append([]string{tokens[0] + tokens[1]}, tokens[2:]...)
}
