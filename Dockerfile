# Use the official Go image to create a build artifact.
# https://hub.docker.com/_/golang
FROM golang:1.23-alpine AS builder

# Set the necessary environment variables for CGO and WASI build
ENV CGO_ENABLED=0 GOOS=wasip1 GOARCH=wasm
WORKDIR /src

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the internal directory containing our packages
COPY internal/ ./internal/

# Copy the source code (for the main package)
COPY main.go .

# Build the WASM plugin
RUN go build -o /out/plugin.wasm .

# Use a minimal image to package the plugin
FROM scratch

LABEL org.opencontainers.image.description="OpenGrep MCP Plugin"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.authors="ducthinh993"
LABEL org.opencontainers.image.vendor="ducthinh993"
LABEL org.opencontainers.image.url="https://github.com/ducthinh993/hyper-mcp-opengrep-plugin"

WORKDIR /
COPY --from=builder /out/plugin.wasm /plugin.wasm
ENTRYPOINT [ "/plugin.wasm" ]