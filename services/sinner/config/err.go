package config

import "errors"

// ErrRequiredGlobalField indicates that a required field is missing.
var ErrRequiredGlobalField = errors.New("a required global config field is empty")

// ErrRequiredContractField indicates that a required field is missing.
var ErrRequiredContractField = errors.New("a required contract config field is empty")
