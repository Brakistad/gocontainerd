#!/usr/bin/env bash
# This script installs minimal go tools into the container.
# It is run as root and the GOBIN is set to /usr/local/go/bin
# so that it is available in the PATH for non-root users.

set -e

# Install tools
# Gopls is the language server for Go
go get -u golang.org/x/tools/gopls@latest
# golangci-lint is a linter for Go
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
# goimports is a tool for organizing imports
go get -u golang.org/x/tools/cmd/goimports
# gocode is an autocompletion tool for the Go language
go get -u github.com/mdempsky/gocode
# gorename is a tool for renaming identifiers in Go source code
go get -u golang.org/x/tools/cmd/gorename
# go-delve is a debugger for the Go programming language
go get -u github.com/go-delve/delve/cmd/dlv
# air is a live reload tool for Go apps
go install github.com/cosmtrek/air@latest