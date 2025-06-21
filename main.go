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

//export opengrep
func opengrep() int32 {
	// Read input from the host
	input := pdk.Input()

	// Unmarshal the JSON input into our struct
	var req OpenGrepRequest
	if err := json.Unmarshal(input, &req); err != nil {
		pdk.SetError(fmt.Errorf("Error unmarshaling JSON input: %v", err))
		return 1
	}

	// Check if any arguments were provided
	if len(req.Args) == 0 {
		pdk.SetError(fmt.Errorf("No arguments provided for opengrep command"))
		return 1
	}

	// Construct the opengrep command with user-provided arguments
	cmd := exec.Command("opengrep", req.Args...)

	// Execute the command and get the combined output (stdout and stderr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// If the command fails, output the error to the host, including the command's output
		pdk.SetError(fmt.Errorf("Error executing opengrep: %v\nOutput: %s", err, string(output)))
		// Still return 0 to send the output back, the error is in the message.
		// A non-zero exit code from opengrep is not a plugin failure.
	}

	// Return the command's output to the host
	pdk.Output(output)

	return 0
}

func main() {}
