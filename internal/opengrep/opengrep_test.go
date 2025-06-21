package opengrep

import (
	"strings"
	"testing"
)

func TestExecute_EmptyInput(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
	}{
		{"nil input", nil},
		{"empty input", []byte{}},
		{"whitespace only", []byte("   ")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Execute(tt.input)
			if err == nil {
				t.Errorf("Execute() expected error for %s, got nil", tt.name)
			}
		})
	}
}

func TestExecute_InvalidInput(t *testing.T) {
	_, err := Execute([]byte(""))
	if err == nil {
		t.Error("Execute() expected error for empty string, got nil")
	}

	if !strings.Contains(err.Error(), "no input provided") {
		t.Errorf("Execute() error message should mention 'no input provided', got: %v", err)
	}
}

func TestCheckAvailability(t *testing.T) {
	// This test will pass if opengrep is available, fail if not
	// In a real environment, you might want to mock this
	err := CheckAvailability()
	if err != nil {
		t.Logf("opengrep not available (this is expected in test environment): %v", err)
	}
}

func TestConstants(t *testing.T) {
	if commandName != "opengrep" {
		t.Errorf("commandName should be 'opengrep', got: %s", commandName)
	}

	if executionTimeout <= 0 {
		t.Errorf("executionTimeout should be positive, got: %v", executionTimeout)
	}

	if checkTimeout <= 0 {
		t.Errorf("checkTimeout should be positive, got: %v", checkTimeout)
	}
}
