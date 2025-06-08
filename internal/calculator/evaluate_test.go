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
		{[]string{"2", "+", "x"}, 0, true, "invalid number"},
		{[]string{"2", "^", "3"}, 0, true, "no operator found"},
		{[]string{"2"}, 2, false, ""},
		{[]string{"2", "+", "3", "/", "0"}, 0, true, "division by zero"},
		{[]string{"2", "3"}, 0, true, "no operator found"},
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
