//go:build integration

package main

import (
	"os/exec"
	"testing"
)

func TestMainIntegration(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "testdata/test.txt")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("unexpected error: %v, output: %s", err, output)
	}
	expected := "18y0m\n"
	if string(output) != expected {
		t.Errorf("got %q, want %q", string(output), expected)
	}
}
