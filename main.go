package main

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/extism/go-pdk"
)

type OpenGrepRequest struct {
	Args []string `json:"args"`
}

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

//export opengrep
func opengrep() int32 {
	// 1. Verify that opengrep is available before proceeding.
	if err := checkOpenGrep(); err != nil {
		pdk.SetError(err)
		return 1 // A non-zero exit code indicates a fatal plugin/environment error.
	}

	// 2. Read and parse the arguments from the host.
	input := pdk.Input()
	var req OpenGrepRequest
	if err := json.Unmarshal(input, &req); err != nil {
		pdk.SetError(fmt.Errorf("error unmarshaling JSON input: %v", err))
		return 1
	}

	if len(req.Args) == 0 {
		pdk.SetError(fmt.Errorf("no arguments provided for opengrep command"))
		return 1
	}

	// 3. Execute the opengrep command.
	output, err := runOpenGrep(req.Args)

	// An error from opengrep (e.g., a non-zero exit code on no match) is not a plugin
	// failure. We report it as part of the error message but still return the output.
	if err != nil {
		pdk.SetError(fmt.Errorf("error executing opengrep: %v\nOutput: %s", err, string(output)))
	}

	// 4. Return the result to the host.
	pdk.Output(output)
	return 0
}

func main() {}
