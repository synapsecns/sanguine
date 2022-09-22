// Package testconsts contains constants for testing
package testconsts

import _ "embed" // embed is used to embed metadata for testing.

// AmplificationUtilsMetadata contains metadata from an amplification utils deploy
// note: this is required from a file instead of test data because we assume this stays constant for testing
//
//go:embed metadata.json
var AmplificationUtilsMetadata []byte
