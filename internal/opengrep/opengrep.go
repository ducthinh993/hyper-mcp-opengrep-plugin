package opengrep

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"sync"
)

const opengrepCmd = "opengrep"

var (
	// checkOnce ensures that the opengrep availability check is only performed once.
	checkOnce sync.Once
	// checkErr stores the result of the availability check.
	checkErr error
)

// HandleOpenGrepRequest parses the input, runs opengrep, and returns the output or an error.
func HandleOpenGrepRequest(input []byte) ([]byte, error) {
	var req OpenGrepRequest
	if err := json.Unmarshal(input, &req); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON input: %w", err)
	}

	if len(req.Args) == 0 {
		return nil, fmt.Errorf("no arguments provided for opengrep command")
	}

	return runOpenGrep(req.Args)
}

// CheckOpenGrep verifies that the opengrep command is available and executable.
// It uses sync.Once to perform the check only on the first call.
func CheckOpenGrep() error {
	checkOnce.Do(func() {
		cmd := exec.Command(opengrepCmd, "--version")
		if err := cmd.Run(); err != nil {
			checkErr = fmt.Errorf("%s command not found or not executable: %w. Please ensure it is installed and in your PATH", opengrepCmd, err)
		}
	})
	return checkErr
}

// runOpenGrep executes the opengrep command with the provided arguments and returns the combined output.
func runOpenGrep(args []string) ([]byte, error) {
	cmd := exec.Command(opengrepCmd, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Return the error and the output, as opengrep often provides useful info on stderr.
		return output, fmt.Errorf("error executing %s: %w", opengrepCmd, err)
	}
	return output, nil
}
