package main

import (
	"fmt"
	"os/exec"
)

// checkOpenGrep verifies that the opengrep command is available and executable.
func checkOpenGrep() error {
	cmd := exec.Command("opengrep", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("opengrep command not found or not executable: %w. Please ensure it is installed and in your PATH", err)
	}
	return nil
}

// runOpenGrep executes the opengrep command with the provided arguments and returns the combined output.
func runOpenGrep(args []string) ([]byte, error) {
	cmd := exec.Command("opengrep", args...)
	return cmd.CombinedOutput()
}
