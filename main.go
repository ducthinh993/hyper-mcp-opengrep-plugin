//go:build extism || !extism

package main

import (
	"fmt"
	"log"

	og "hyper-mcp-opengrep-plugin/internal/opengrep"

	"github.com/extism/go-pdk"
)

// opengrep is the main exported function for the Extism plugin.
// It handles opengrep command execution requests from the MCP server.
//
//export opengrep
func opengrep() int32 {
	// Validate opengrep availability
	if err := og.CheckAvailability(); err != nil {
		log.Printf("opengrep availability check failed: %v", err)
		pdk.SetError(fmt.Errorf("opengrep not available: %w", err))
		return 1
	}

	// Get input from the plugin
	input := pdk.Input()
	if len(input) == 0 {
		pdk.SetError(fmt.Errorf("no input provided"))
		return 1
	}

	// Process the opengrep request
	output, err := og.Execute(input)
	if err != nil {
		// Log the error for debugging
		log.Printf("opengrep execution failed: %v", err)

		// Set error but still return output if available
		pdk.SetError(fmt.Errorf("opengrep execution failed: %w", err))
	}

	// Always return output, even if there was an error
	// (opengrep often provides useful error details in stderr)
	if len(output) > 0 {
		pdk.Output(output)
	}

	return 0
}

// main is required by Go but not used in Extism plugins
func main() {
	// This function is intentionally empty for Extism plugins
}
