#!/bin/bash
set -euxo pipefail

# Update package list and install build tools
apt-get update
apt-get install -y --no-install-recommends git make curl ca-certificates golang-go

# Install Go tools required by the Makefile
export PATH="/usr/local/go/bin:$HOME/go/bin:$PATH"

# Install test and lint dependencies
go install gotest.tools/gotestsum@v1.6.3

go install github.com/vektra/mockery/v2@v2.16.0

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

