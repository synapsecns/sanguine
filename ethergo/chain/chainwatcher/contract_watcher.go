package chainwatcher

import (
	"context"
)

// ContractWatcher is a contract watcher.
type ContractWatcher interface {
	ListenOnContract(ctx context.Context, contractAddress string, eventLog chan interface{}) error
}
