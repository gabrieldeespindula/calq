package repl

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
	}{
		{"2 + 2\n", "4"},
		{"5 - 2\n", "3"},
		{"4 * 3\n", "12"},
		{"10 / 2\n", "5"},
		{"3 + 5 * 2\n", "13"},
		{"2 + 10 / 2\n+3\n", "> 7\n> 10"},
		{"2 + 10 / 2\n2+3\n", "> 7\n> 5"},
		{"20*j\n", "Calculation error: invalid number: j"},
		{"\n", "Welcome to calq! Type your expressions (type 'q' + Enter to quit, or press Ctrl+C):\n> > Goodbye!\n"},
	}

	for _, tt := range tests {
		input := strings.NewReader(tt.input+"q\n")
		output := &bytes.Buffer{}

		Run(input, output)

		got := output.String()
		if !strings.Contains(got, "Welcome to calq! Type your expressions (Ctrl+C to exit):") {
			t.Errorf("expected welcome message, got %q", got)
		}
		if !strings.Contains(got, tt.expected) {
			t.Errorf("expected output to contain '%s', got %q", tt.expected, got)
		}
		if !strings.Contains(got, "Goodbye!") {
			t.Errorf("expected goodbye message, got %q", got)
		}
	}
}
