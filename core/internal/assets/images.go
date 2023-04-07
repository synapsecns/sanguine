package assets

import _ "embed" // embed is used to embed images.

// Logo stores the logo
//
//go:embed synapse.svg
var Logo []byte
