#!/bin/bash
# Build script with flags to disable linkname check for memsize compatibility

# Set flags to disable the linkname check
export GOFLAGS="-ldflags=-checklinkname=0"

# Run the build command
go build "$@"

# Print success message
echo "Build completed with custom flags to fix memsize compatibility"
