package calculator

import (
	"testing"
)

func TestOperationsMap(t *testing.T) {
	tests := []struct {
		op       string
		a, b     float64
		expected float64
	}{
		{op: ADD, a: 2, b: 3, expected: 5},
		{op: SUB, a: 5, b: 2, expected: 3},
		{op: MUL, a: 4, b: 3, expected: 12},
		{op: DIV, a: 10, b: 2, expected: 5},
	}

	for _, tt := range tests {
		fn, ok := Operations[tt.op]
		if !ok {
			t.Errorf("operation %q not found in map", tt.op)
			continue
		}
		result, err := fn(tt.a, tt.b)
		if err != nil {
			t.Errorf("unexpected error for op %q: %v", tt.op, err)
			continue
		}
		if result != tt.expected {
			t.Errorf("expected result of %v %q %v to be %v, got %v", tt.a, tt.op, tt.b, tt.expected, result)
		}
	}
}

func TestOperationsMap_DivisionByZero(t *testing.T) {
	fn, ok := Operations[DIV]
	if !ok {
		t.Fatalf("division operator not found in map")
	}
	_, err := fn(10, 0)
	if err == nil {
		t.Error("expected error when dividing by zero, got nil")
	} else {
		t.Logf("correctly received error for division by zero: %v", err)
	}
}