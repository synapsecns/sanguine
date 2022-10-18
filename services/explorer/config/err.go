package config

import "errors"

// ErrInvalidChainID indicates that the chain ID is invalid.
var ErrInvalidChainID = errors.New("invalid chain id")

// ErrRequiredField indicates that a required field is missing.
var ErrRequiredField = errors.New("field is required")

// ErrDuplicateChainID indicates that a duplicate chain ID is found.
var ErrDuplicateChainID = errors.New("duplicate chain id")

// ErrDuplicateAddress indicates that a duplicate address is found.
var ErrDuplicateAddress = errors.New("duplicate address")

// ErrAddressLength indicates that an invalid address length is found.
var ErrAddressLength = errors.New("invalid address length")
