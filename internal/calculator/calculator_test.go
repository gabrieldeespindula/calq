package calculator

import (
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		lastResult string
		expr       string
		expected    float64
		wantErr     bool
		errMsg      string
	}{
		{"", "2 + 2", 4, false, ""},
		{"", "5 - 2", 3, false, ""},
		{"", "4 * 3", 12, false, ""},
		{"", "10 / 2", 5, false, ""},
		{"", "3 + 5 * 2", 13, false, ""},
		{"-5", "+ 5", 0, false, ""},
		{"3", "* -3", -9, false, ""},
		{"", "20*j\n", 0, true, "invalid number: j"},
		{"", "", 0, true, "empty expression"},
	}

	for _, tt := range tests {
		got, err := Calculate(tt.lastResult, tt.expr)
		if (err != nil) != tt.wantErr {
			t.Errorf("Calculate(%v, %v) error = %v, wantErr %v", tt.lastResult, tt.expr, err, tt.wantErr)
			continue
		}
		if err != nil && !strings.Contains(err.Error(), tt.errMsg) {
			t.Errorf("Calculate(%v, %v) error = %v, want error containing %q", tt.lastResult, tt.expr, err, tt.errMsg)
		}
		if !tt.wantErr && got != tt.expected {
			t.Errorf("Calculate(%v, %v) = %v, want %v", tt.lastResult, tt.expr, got, tt.expected)
		}
	}
}
