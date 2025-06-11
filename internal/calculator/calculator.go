package calculator

func Calculate(lastResult string, expr string) (float64, error) {
	tokens, err := Tokenize(lastResult, expr)
	if err != nil {
		return 0, err
	}
	return Evaluate(tokens)
}
