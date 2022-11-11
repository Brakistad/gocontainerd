#!/usr/bin/env bash
# This script installs minimal go tools into the container.
# It is run as root and the GOBIN is set to /usr/local/go/bin
# so that it is available in the PATH for non-root users.

set -e

# Install tools
# air is a live reload tool for Go apps
go install github.com/cosmtrek/air@latest