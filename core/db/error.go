package db

import "errors"

// ErrNoNonceForChain indicates that no nonces have been saved for the chain yet.
var ErrNoNonceForChain = errors.New("no nonce exists for this chain")

// ErrNotFound is a not found record standardized across db drivers.
var ErrNotFound = errors.New("record not found")

// ErrNoStoredBlockForChain indicates there are no blocks stored for this domain.
var ErrNoStoredBlockForChain = errors.New("no block exists for this chain")

// ErrNoNonceForDomain indicates there is no nonce for a domain.
var ErrNoNonceForDomain = errors.New("no nonce exists for this domain")
