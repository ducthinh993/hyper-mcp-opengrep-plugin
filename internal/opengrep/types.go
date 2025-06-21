package opengrep

// OpenGrepRequest defines the structure for the JSON input from the host.
type OpenGrepRequest struct {
	Args []string `json:"args"`
}
