package db

import "errors"

// ErrNotFound is a not found record standardized across db drivers.
var ErrNotFound = errors.New("record not found")
