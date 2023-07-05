package chainwatcher

import (
	"context"
	"github.com/synapsecns/sanguine/core/prom"
)

// ContractWatcher is a contract watcher.
type ContractWatcher interface {
	prom.Instrumentable
	ListenOnContract(ctx context.Context, contractAddress string, eventLog chan interface{}) error
}
