package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)


func TestShortVersionFlag(t *testing.T) {
	runCalqVersionFlag(t, "-v")
}

func TestLongVersionFlag(t *testing.T) {
	runCalqVersionFlag(t, "--version")
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