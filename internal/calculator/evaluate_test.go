package calculator

import (
	"testing"
	"strings"
)

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
		{[]string{"2", "-", "3"}, -1, false, ""},
		{[]string{"2", "*", "-3"}, -6, false, ""},
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
