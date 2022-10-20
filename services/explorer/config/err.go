package config

import "errors"

// ErrRequiredField indicates that a required field is missing.
var ErrRequiredField = errors.New("field is required")

// ErrAddressLength indicates that an invalid address length is found.
var ErrAddressLength = errors.New("invalid address length")
