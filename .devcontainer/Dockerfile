FROM golang:1.25.9-bookworm

# Install system dependencies required for building and tool management
RUN apt-get update && apt-get install -y \
    git \
    make \
    build-essential

# Install Go tools into the global GOPATH
RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/google/wire/cmd/wire@latest

# Ensure the Go toolchain and Go bin directory are in the PATH
ENV PATH="/usr/local/go/bin:/go/bin:${PATH}"

WORKDIR /workspaces/buildergen
