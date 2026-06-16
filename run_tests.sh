#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/gopwdutil"

echo "Running all tests..."
go test ./... -v

echo ""
echo "Done."
