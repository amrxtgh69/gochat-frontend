#!/bin/bash
set -e  # Exit immediately if a command exits with a non-zero status

# Variables
REPO_URL="https://github.com/amrxtgh69/gochat-frontend"
TMP_DIR="/tmp/gochat_clone"
OUTPUT_BIN="/usr/local/bin/gochat"

# Clone the repository
git clone "$REPO_URL" "$TMP_DIR"

# Build the Go project
cd "$TMP_DIR"
go build -o "$OUTPUT_BIN" cmd/gochat/main.go

# Clean up
cd /
rm -rf "$TMP_DIR"

echo "GoChat built and installed at $OUTPUT_BIN"
