# Use the official Go image to create a build artifact.
# https://hub.docker.com/_/golang
FROM golang:1.21-alpine AS builder

# Set the necessary environment variables for CGO and WASI build
ENV CGO_ENABLED=0 GOOS=wasip1 GOARCH=wasm
WORKDIR /src

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the WASM plugin
RUN go build -o /out/plugin.wasm .

# Use a minimal image to package the plugin
FROM scratch
WORKDIR /
COPY --from=builder /out/plugin.wasm /plugin.wasm
ENTRYPOINT [ "/plugin.wasm" ]