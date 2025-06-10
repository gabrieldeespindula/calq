package calculator

import (
	"reflect"
	"testing"
	"strings"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		expr       string
		lastResult string
		want       []string
		wantErr    bool
		errMsg     string
	}{
		{"", "", nil, true, "empty expression"},
		{"42", "", []string{"42"}, false, ""},
		{"2+2", "", []string{"2", "+", "2"}, false, ""},
		{" 3 * 4 ", "", []string{"3", "*", "4"}, false, ""},
		{"+5", "10", []string{"10", "+", "5"}, false, ""},
		{"4*-5", "", []string{"4", "*", "-5"}, false, ""},
		{"+2", "-4", []string{"-4", "+", "2"}, false, ""},
	}

	for _, tt := range tests {
		got, err := Tokenize(tt.expr, tt.lastResult)
		if (err != nil) != tt.wantErr {
			t.Errorf("Tokenize(%q, %q) error = %v, wantErr %v", tt.expr, tt.lastResult, err, tt.wantErr)
			continue
		}
		if err != nil && !strings.Contains(err.Error(), tt.errMsg) {
			t.Errorf("Tokenize(%q, %q) error = %v, want error containing %q", tt.expr, tt.lastResult, err, tt.errMsg)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Tokenize(%q, %q) = %v, want %v", tt.expr, tt.lastResult, got, tt.want)
		}
	}
}

func TestEvaluate(t *testing.T) {
	tests := []struct {
		parts    []string
		want     float64
		wantErr  bool
		errMsg   string
	}{
		{[]string{"42"}, 42, false, ""},
		{[]string{"2", "+", "2"}, 4, false, ""},
		{[]string{"2", "+", "2", "*", "3"}, 8, false, ""},
		{[]string{"2", "+", "9", "/", "3"}, 5, false, ""},
		{[]string{"2", "+", "9", "/", "3", "*", "2"}, 8, false, ""},
		{[]string{"2", "+", "x"}, 0, true, "invalid number: x"},
		{[]string{"y", "+", "2"}, 0, true, "invalid number: y"},
		{[]string{"2", "^", "3"}, 0, true, "no operator found"},
		{[]string{"2"}, 2, false, ""},
		{[]string{"2", "+", "3", "/", "0"}, 0, true, "division by zero is not allowed"},
		{[]string{"2", "3"}, 0, true, "no operator found"},
		{[]string{"+", "3"}, 0, true, "invalid expression format"},
		{[]string{"2", "+", "3", "*"}, 0, true, "invalid expression format"},
	}

	for _, tt := range tests {
		got, err := Evaluate(tt.parts)
		if (err != nil) != tt.wantErr {
			t.Errorf("Evaluate(%v) error = %v, wantErr %v", tt.parts, err, tt.wantErr)
			continue
		}
		if err != nil && !strings.Contains(err.Error(), tt.errMsg) {
			t.Errorf("Evaluate(%v) error = %v, want error containing %q", tt.parts, err, tt.errMsg)
		}
		if !tt.wantErr && got != tt.want {
			t.Errorf("Evaluate(%v) = %v, want %v", tt.parts, got, tt.want)
		}
	}
}

func TestEvaluateUnsupportedOperator(t *testing.T) {
	originalOps := Operations
	defer func() { Operations = originalOps }()

	Operations = map[string]Operation{}

	parts := []string{"2", "+", "3"} // + is a known operator, but it's not in Operations now

	_, err := Evaluate(parts)
	if err == nil || !strings.Contains(err.Error(), "unsupported operator: +") {
		t.Errorf("expected unsupported operator error, got %v", err)
	}
}
