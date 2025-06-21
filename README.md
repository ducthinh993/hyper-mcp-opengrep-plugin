# Hyper-MCP OpenGrep Plugin

A `hyper-mcp` plugin that provides a direct interface to the `opengrep` static code analysis engine. This plugin allows an AI agent to execute any `opengrep` command, enabling powerful and flexible code scanning directly from an MCP-compatible client like Cursor.

## Features

*   **Direct `opengrep` Access:** Functions as a complete wrapper around the `opengrep` command-line tool.
*   **Full Argument Support:** Supports all commands, flags, and arguments of the `opengrep` CLI.
*   **Flexible Input:** Accepts a JSON object containing the arguments to pass to `opengrep`.
*   **WASM-based:** Built as a secure, sandboxed WebAssembly module for `hyper-mcp`, written in Go.

## Installation and Setup

### Prerequisites

1.  A running `hyper-mcp` server.
2.  The `opengrep` binary must be installed and accessible in the system's `PATH` on the same machine where the `hyper-mcp` server is running.

### Configuration

1.  Add the plugin to your `hyper-mcp` configuration file (e.g., `~/.config/hyper-mcp/config.json` on Linux, `~/Library/Application Support/hyper-mcp/config.json` on macOS).

    ```json
    {
      "plugins": [
        {
          "name": "opengrep",
          "path": "oci://ghcr.io/ducthinh993/hyper-mcp-opengrep-plugin:latest"
        }
      ]
    }
    ```
    > **Note:** Replace `ducthinh993/hyper-mcp-opengrep-plugin` with your GitHub repository name if you have forked and published the image under your own account.

2.  Ensure your client (e.g., Cursor) is configured to use your `hyper-mcp` server instance.

## Usage

To use the plugin, the AI agent must call the `opengrep` tool with a single JSON payload containing an `args` array. The elements of this array are the command-line arguments that will be passed directly to the `opengrep` executable.

### Example

To scan all files in the current directory (`.`) for occurrences of the word `TODO` and output the results in JSON format, the AI would invoke the tool with the following input:

```json
{
  "args": ["scan", "-e", "TODO", "--json", "."]
}
```

The plugin will execute `opengrep scan -e TODO --json .` and return the complete, raw output (both `stdout` and `stderr`) from the command.

## Development

### Building Locally

This project uses a multi-stage `Dockerfile` to build the WASM binary and package it into a minimal `scratch`-based OCI image. To build the image locally, run:

```bash
docker build -t hyper-mcp-opengrep-plugin .
```

### Releasing a New Version

The repository contains a GitHub Actions workflow (`.github/workflows/release.yml`) that automatically builds and publishes a new OCI image to the GitHub Container Registry (GHCR) whenever a new version tag is pushed.

To trigger a release:

1.  Commit your changes.
2.  Create and push a new tag:
    ```bash
    git tag v0.1.0
    git push origin v0.1.0
    ```

The action will run and publish the image, for example, `ghcr.io/ducthinh993/hyper-mcp-opengrep-plugin:v0.1.0` and `ghcr.io/ducthinh993/hyper-mcp-opengrep-plugin:latest`.

## License

This project is licensed under the Apache 2.0 License.