package consumer

import "errors"

// ErrTokenDoesNotExist indicates the queried for token does not exist.
var ErrTokenDoesNotExist = errors.New("token id does not exist")
