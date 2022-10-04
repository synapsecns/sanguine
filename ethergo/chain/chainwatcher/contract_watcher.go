package chainwatcher

import (
	"context"
	"github.com/synapsecns/sanguine/core/prometheus"
)

// ContractWatcher is a contract watcher.
type ContractWatcher interface {
	prometheus.Instrumentable
	ListenOnContract(ctx context.Context, contractAddress string, eventLog chan interface{}) error
}
