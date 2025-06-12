package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)


func TestVersionFlag(t *testing.T) {
	runCalqVersionFlag(t, "--version")
	runCalqVersionFlag(t, "-v")
}

func TestEvalFlag(t *testing.T) {
	runCalqEvalFlag(t, "--eval")
	runCalqEvalFlag(t, "-e")
}


func runCalqEvalFlag(t *testing.T, flag string) {
    tests := []struct {
        expr     string
        expected string
    }{
        {"1+1", "2"},
        {"2*3", "6"},
        {"10/3", "3.33"},
        {"5-3", "2"},
        {"2+2*2", "6"},
    }

    for _, tt := range tests {
        t.Run(tt.expr, func(t *testing.T) {
            cmd := exec.Command("../calq", flag, tt.expr)

            var out bytes.Buffer
            cmd.Stdout = &out
            cmd.Stderr = &out

            err := cmd.Run()
            if err != nil {
                t.Fatalf("failed to run calq with expression %q: %v", tt.expr, err)
            }

            output := out.String()
            if !strings.Contains(output, tt.expected) {
                t.Errorf("output for %q does not contain expected result %q: got %q", tt.expr, tt.expected, output)
            }
        })
    }
}


func runCalqVersionFlag(t *testing.T, flag string) {
	t.Helper()

	cmd := exec.Command("../calq", flag)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("failed to run calq with flag %q: %v", flag, err)
	}

	output := out.String()

	if !strings.Contains(output, "calq") {
		t.Errorf("output does not contain 'calq': %q", output)
	}

	if !strings.Contains(output, "dev") {
		t.Errorf("output does not contain version 'dev': %q", output)
	}
}