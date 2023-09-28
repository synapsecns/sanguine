package config

import "errors"

// ErrRequiredGlobalField indicates that a required field is missing.
var ErrRequiredGlobalField = errors.New("a required global config field is empty")

// ErrRequiredChainField indicates that a required field is missing.
var ErrRequiredChainField = errors.New("a required chain config field is empty")

// ErrRequiredContractField indicates that a required field is missing.
var ErrRequiredContractField = errors.New("a required contract config field is empty")

// ErrAddressLength indicates that an invalid address length is found.
var ErrAddressLength = errors.New("invalid address length")
