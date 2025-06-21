//go:build extism || !extism

package main

import (
	"fmt"
	og "hyper-mcp-opengrep-plugin/internal/opengrep"

	"github.com/extism/go-pdk"
)

//export opengrep
func opengrep() int32 {
	if err := og.CheckOpenGrep(); err != nil {
		pdk.SetError(err)
		return 1
	}

	input := pdk.Input()
	output, err := og.HandleOpenGrepRequest(input)

	// If there was an error, report it. The output from the command (which may contain
	// useful error details from opengrep itself) is still returned.
	if err != nil {
		pdk.SetError(fmt.Errorf("%w\n--- opengrep output ---\n%s", err, string(output)))
	}

	pdk.Output(output)
	return 0
}

func main() {}
