package opengrep

// Request represents a request to execute an opengrep command.
// This structure can be used for future JSON-based input handling.
type Request struct {
	// Args contains the command line arguments for opengrep
	Args []string `json:"args"`

	// Timeout specifies the maximum execution time (optional)
	Timeout *int `json:"timeout,omitempty"`
}

// Response represents the result of an opengrep command execution.
type Response struct {
	// Output contains the command output (stdout + stderr)
	Output string `json:"output"`

	// Error contains any error message if the command failed
	Error string `json:"error,omitempty"`

	// ExitCode contains the command's exit code
	ExitCode int `json:"exit_code"`
}
