package calculator

import (
	"reflect"
	"testing"
	"strings"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		lastResult string
		expr       string
		want       []string
		wantErr    bool
		errMsg     string
	}{
		{"", "", nil, true, "empty expression"},
		{"", "42", []string{"42"}, false, ""},
		{"", "2+2", []string{"2", "+", "2"}, false, ""},
		{"", "2-2", []string{"2", "-", "2"}, false, ""},
		{"", "2*2", []string{"2", "*", "2"}, false, ""},
		{"", "2/2", []string{"2", "/", "2"}, false, ""},
		{"", "  3 *     4 ", []string{"3", "*", "4"}, false, ""},
		{"10", "+5", []string{"10", "+", "5"}, false, ""},
		{"", "4*-5", []string{"4", "*", "-5"}, false, ""},
		{"-4", "+2", []string{"-4", "+", "2"}, false, ""},
	}

	for _, tt := range tests {
		got, err := Tokenize(tt.lastResult, tt.expr)
		if (err != nil) != tt.wantErr {
			t.Errorf("Tokenize(%q, %q) error = %v, wantErr %v", tt.lastResult, tt.expr, err, tt.wantErr)
			continue
		}
		if err != nil && !strings.Contains(err.Error(), tt.errMsg) {
			t.Errorf("Tokenize(%q, %q) error = %v, want error containing %q", tt.lastResult, tt.expr, err, tt.errMsg)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Tokenize(%q, %q) = %v, want %v", tt.lastResult, tt.expr, got, tt.want)
		}
	}
}
