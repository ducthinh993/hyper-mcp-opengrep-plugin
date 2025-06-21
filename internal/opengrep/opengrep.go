package opengrep

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

const (
	// Command name for opengrep
	commandName = "opengrep"

	// Timeout for command execution
	executionTimeout = 30 * time.Second

	// Timeout for availability check
	checkTimeout = 5 * time.Second
)

var (
	// checkOnce ensures that the opengrep availability check is only performed once.
	checkOnce sync.Once

	// checkErr stores the result of the availability check.
	checkErr error
)

// CheckAvailability verifies that the opengrep command is available and executable.
// It uses sync.Once to perform the check only on the first call.
func CheckAvailability() error {
	checkOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), checkTimeout)
		defer cancel()

		cmd := exec.CommandContext(ctx, commandName, "--version")
		if err := cmd.Run(); err != nil {
			checkErr = fmt.Errorf("%s command not found or not executable: %w. Please ensure it is installed and in your PATH", commandName, err)
		}
	})
	return checkErr
}

// Execute runs the opengrep command with the provided input arguments.
// The input is expected to be a space-separated string of arguments.
func Execute(input []byte) ([]byte, error) {
	if len(input) == 0 {
		return nil, fmt.Errorf("no input provided")
	}

	// Parse input arguments
	argString := strings.TrimSpace(string(input))
	if argString == "" {
		return nil, fmt.Errorf("empty input string")
	}

	args := strings.Fields(argString)
	if len(args) == 0 {
		return nil, fmt.Errorf("no arguments provided for %s command", commandName)
	}

	return runCommand(args)
}

// runCommand executes the opengrep command with the provided arguments.
// It uses a context with timeout to prevent hanging executions.
func runCommand(args []string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), executionTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, commandName, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		// Return both error and output, as opengrep often provides useful info on stderr
		return output, fmt.Errorf("error executing %s with args %v: %w", commandName, args, err)
	}

	return output, nil
}
