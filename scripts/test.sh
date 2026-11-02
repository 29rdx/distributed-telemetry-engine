#!/bin/bash
set -e
echo "Running Go telemetry engine test suite with race detector..."
go test -v -race ./...
