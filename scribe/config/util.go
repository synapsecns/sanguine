package config

import "errors"

// ErrInvalidChainId indicates that the chain ID is invalid.
var ErrInvalidChainId = errors.New("invalid chain id")

// ErrRequiredField indicates that a required field is missing.
var ErrRequiredField = errors.New("field is required")

// ErrDuplicateChainID indicates that a duplicate chain ID is found.
var ErrDuplicateChainId = errors.New("duplicate chain id")

// ErrDuplicateAddress indicates that a duplicate address is found.
var ErrDuplicateAddress = errors.New("duplicate address")
