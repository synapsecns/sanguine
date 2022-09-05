package chainwatcher

import (
	"context"
	"github.com/synapsecns/sanguine/core/metrics"
)

// ContractWatcher is a contract watcher.
type ContractWatcher interface {
	metrics.Instrumentable
	ListenOnContract(ctx context.Context, contractAddress string, eventLog chan interface{}) error
}
