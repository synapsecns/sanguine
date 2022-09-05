package client

import (
	"context"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/synapsecns/sanguine/core/metrics"
	watcher "github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
	"k8s.io/apimachinery/pkg/util/wait"
	"math/big"
	"sync"
	"time"
)

// PoolClient is a type of client that allows fallbacks to multiple rpc endpoints.
//
// in the initial version, functionality is quite simple. Each node creates a block height
// watcher and the height relative to other rpcs is used to determine if this node is in the live
// or dead pool. A "leader" is picked from the live pool and this is used to satisfy all calls to the pool client.
// For this reason this client also acts as a proxy to block height wathcher.
//
// In the case where no live client can be found, a dead client is returned. Note: subscriptions using an old client
// will not automatically switch over- the new client is only used on new method calls. Right now, *all* rpc
// servers used by a node operator are trusted (and we expect users to run them themselves). In the future
// it might make sense to have users verify against all live clients or use smart rate limiting.
type PoolClient interface {
	// MeteredEVMClient is the primary client. This is swapped among clients based on liveliness
	MeteredEVMClient
	// Watcher is the watcher interface. We expose it from here since this is the most reliable client
	// it is kept ina  separate call to allow metrics to be independently assessed
	Watcher() watcher.BlockHeightWatcher
	// Instrumentable allows a user to collect metrics against pool client
	metrics.Instrumentable
}

// NewPoolClient creates a new pool client from a list of metered evm clients.
// only errors if no clients are passed in.
func NewPoolClient(ctx context.Context, chainID uint64, clients []MeteredEVMClient) (PoolClient, error) {
	if len(clients) == 0 {
		return nil, errors.New("need at least one client to create pool client")
	}
	poolClient := poolClientImpl{
		multiBlockWatcher: watcher.NewBlockBroadcaster(ctx, chainID),
		ctx:               ctx,
	}

	// populate the initial clients
	for _, client := range clients {
		poolClient.clients = append(poolClient.clients, newObservableClient(ctx, client))
	}

	poolClient.startClientObservationLoop(time.Second * 10)

	return &poolClient, nil
}

// NewPoolClientFromURLs creates a new pool client from a list of rpc url.
//
// it differs from NewPoolClient in that it can be used even if not all rpc urls
// are up at start time.
func NewPoolClientFromURLs(ctx context.Context, chainID uint64, rpcURLs []string) (PoolClient, error) {
	var wg sync.WaitGroup
	// meteredClients are the meteredClients resulting from the potentialMeteredClient calls
	var meteredClients []MeteredEVMClient
	// mux is used to make sure we don't double write to the array
	var mux sync.Mutex
	for _, rpcURL := range rpcURLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			client := newPotentialMeteredClient(ctx, url, big.NewInt(int64(chainID)), url, nil)
			mux.Lock()
			meteredClients = append(meteredClients, client)
			mux.Unlock()
		}(rpcURL)
	}

	wg.Wait()
	return NewPoolClient(ctx, chainID, meteredClients)
}

type poolClientImpl struct {
	// MeteredEVMClient is the primary client
	*observableClient
	// multiBlockWatcher is the watcher used for tracking heights across multiple clients.
	// this allows seamless handoffs of height subscriptions.
	// TODO: generalize subscriptions to be flawless multipool handoffs
	multiBlockWatcher *watcher.BlockBroadcaster
	// ctx is the context used for the block height watcher
	// storing context in a struct is an anti-pattern, but we don't have a great alternative here.
	// it's important to note this context is *not* merged with the context in the individual requests.
	// this is used for determining live/dead clients only
	//nolint: containedctx
	ctx context.Context
	// mux is used to prevent duplicate map reads/writes.
	mux sync.RWMutex
	// clients are all available clients
	clients []*observableClient
}

// startClientObservationLoop starts the obesrvation loop that updates the active clients, height, etc.
// it ensures that *an* active client is sent even if it's offline.
//
// leaderTimeout is how long we should wait for a leader before setting the first client on startup
// it has no effect after startup.
func (p *poolClientImpl) startClientObservationLoop(leaderTimeout time.Duration) {
	for _, observable := range p.clients {
		go func(client *observableClient) {
			heightSubscription := client.watcher.Subscribe()
			defer client.watcher.Unsubscribe(heightSubscription)

			for {
				select {
				case <-p.ctx.Done():
					return
				case height := <-heightSubscription:
					p.mux.Lock()
					client.latestHeight = int(height)
					client.latestUpdate = time.Now()
					// check if we're the leader, if we are skip the comparisons
					isLeader := p.checkIsLeader(client)
					// if we are the leader update all the subscribers with the new blocks
					if isLeader {
						p.multiBlockWatcher.UpdateHeight(height)
					}
					p.mux.Unlock()
				// every 5 minutes make sure we're still connected
				case <-time.After(time.Minute):
					// client disconnected, we're going to try to reconnect
					if time.Now().Add(time.Minute).Sub(client.latestUpdate) > 0 {
						client.AttemptReconnect()
					}
				}
			}
		}(observable)
	}

	leaderCtx, cancel := context.WithTimeout(p.ctx, leaderTimeout)

	wait.UntilWithContext(leaderCtx, func(ctx context.Context) {
		if p.observableClient != nil {
			cancel()
		}
	}, leaderTimeout/100)

	// if the observable client is still null, set the first client
	if p.observableClient == nil {
		p.observableClient = p.clients[0]
	}
}

// checkIsLeader checks if the current client is the leader or should be upgraded to the leadership.
// this should be run in the context of a mutex.
func (p *poolClientImpl) checkIsLeader(client *observableClient) bool {
	isLeader := p.observableClient != client
	if !isLeader {
		// if the leader has fallen behind by 30 seconds AND 5 blocks swap in with them.
		// if no leader is set, we become the leader
		swapWithLeader := p.observableClient == nil || (client.latestHeight > p.observableClient.latestHeight-5 &&
			client.latestUpdate.Sub(p.observableClient.latestUpdate.Add(30*time.Second)) < 0)

		if swapWithLeader {
			p.observableClient = client
			isLeader = true
		}
	}
	return isLeader
}

// Watcher gets the pool client watcher.
// this just wraps poolClientImpl which implements Watcher.
func (p *poolClientImpl) Watcher() watcher.BlockHeightWatcher {
	return p.multiBlockWatcher
}

// GetMetrics gets the metrics for the pool config.
func (p *poolClientImpl) GetMetrics(labels map[string]string) (metrics []prometheus.Collector) {
	for _, client := range p.clients {
		labels["client_id"] = client.ClientID()
		metrics = append(metrics, client.GetMetrics(labels)...)
		// TODO maybe re-enable so we can get per-rpc block heights?
		// metrics = append(metrics, client.watcher.GetMetrics(labels)...)
	}
	return metrics
}

var _ PoolClient = &poolClientImpl{}
