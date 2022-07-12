package db

import "errors"

// ErrNoNonceForChain indicates that no nonces have been saved for the chain yet.
var ErrNoNonceForChain = errors.New("no nonce exists for this chain")

// ErrNotFound is a not found record standardized across db drivers.
var ErrNotFound = errors.New("record not found")
