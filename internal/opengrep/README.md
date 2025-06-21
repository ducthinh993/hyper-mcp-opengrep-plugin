# OpenGrep Package

This package provides functionality to execute `opengrep` commands from within the Extism plugin.

## Overview

The `opengrep` package handles:
- Availability checking of the `opengrep` command
- Command execution with proper timeout handling
- Input validation and parsing
- Error handling and output management

## Functions

### `CheckAvailability() error`
Verifies that the `opengrep` command is available and executable. Uses `sync.Once` to perform the check only once per plugin instance.

### `Execute(input []byte) ([]byte, error)`
Executes the `opengrep` command with the provided input arguments. The input is expected to be a space-separated string of arguments.

## Types

### `Request`
Represents a request to execute an opengrep command (for future JSON-based input handling).

### `Response`
Represents the result of an opengrep command execution.

## Usage

```go
// Check if opengrep is available
if err := opengrep.CheckAvailability(); err != nil {
    return err
}

// Execute opengrep command
output, err := opengrep.Execute([]byte("search pattern file.txt"))
if err != nil {
    // Handle error
}
```

## Error Handling

The package provides comprehensive error handling:
- Availability check errors
- Input validation errors
- Command execution errors with timeout protection
- Detailed error messages for debugging 