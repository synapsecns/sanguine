package client

import (
	"context"
	"github.com/ethereum/go-ethereum/params"
	backoffHandler "github.com/jpillora/backoff"
	keepRate "github.com/keep-network/keep-common/pkg/rate"
	"github.com/pkg/errors"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
	"github.com/synapsecns/sanguine/ethergo/chain/watcher"
	"math/big"
	"sync"
	"time"
)

var errClientNotStarted = errors.New("client has not been able to start yet")

// observableClient is a metered evm client that exposes a status as to it's liveliness and a block height watcher subscription.
type observableClient struct {
	// MeteredEVMClient is the client we're observing
	MeteredEVMClient
	// latestHeight is the latest height observed on this client
	latestHeight int
	// latestUpdate is the last update time
	latestUpdate time.Time
	// alive indicates whether the client is currently dead or alive
	alive bool
	// watcher.BlockHeightWatcher is the block height watcher for the client
	watcher chainwatcher.BlockHeightWatcher
}

// newObservableClient creates a new observable client.
func newObservableClient(ctx context.Context, client MeteredEVMClient) *observableClient {
	return &observableClient{
		MeteredEVMClient: client,
		latestHeight:     0,
		alive:            false,
		watcher:          watcher.NewBlockHeightWatcher(ctx, client.ChainConfig().ChainID.Uint64(), client),
	}
}

// nonConnectedPermitter is a permitter which always returns an error.
// this allows us to instantiate a disconnected client and try to reconnect after the node has already started.
type nonConnectedPermitter struct{}

// AcquirePermit returns an error that the client is not connnected.
func (n nonConnectedPermitter) AcquirePermit(_ context.Context) (err error) {
	//nolint: wrapcheck
	return errClientNotStarted
}

func (n nonConnectedPermitter) ReleasePermit() {}

var _ Permitter = &nonConnectedPermitter{}

// potentialMeteredClient is a metered client that may not have been able to establish an initial connection and keeps retrying.
type potentialMeteredClient struct {
	meteredEVMClientImpl
	// ctx is the context for the underlying eth client/retry loop
	// nolint: containedctx
	ctx context.Context
	// wsURL is the websocket url to try to connect to
	wsURL string
	// logOnce ensures an error on an attempted connection is only logged wonce
	logOnce sync.Once
}

// attemptDuration is how long we should try to connect before returning a client with a nonConnectedPermitter.
var attemptDuration = time.Second * 10

// newPotentialMeteredClient attempts to create a metered evm client.
// If it cannot it keeps trying until it can, but does not return a blank client.
// Instead a client which errors on every request is returned. This should not be used
// outside of the context of pool where clients are constantly checked for liveness.
func newPotentialMeteredClient(ctx context.Context, wsurl string, chainID *big.Int, clientID string, config *keepRate.LimiterConfig) *potentialMeteredClient {
	meteredClient := potentialMeteredClient{
		meteredEVMClientImpl: getMeteredClientStub(chainID, clientID, config),
		wsURL:                wsurl,
		ctx:                  ctx,
	}
	// attempt to connect for the first time for 10 seconds. If we can't, return the nonConnectedClient and keep trying
	connected := meteredClient.attemptConnect(attemptDuration)
	// we couldn't connect so we're going to keep retrying until we can. In the meantime, this client is always going to return error
	if !connected {
		meteredClient.LifecycleClient = NewLifecycleClient(nil, chainID, nonConnectedPermitter{}, requestTimeout)
		// keep trying on a backoff up to 5 minutes
		backoff := &backoffHandler.Backoff{
			Factor: 1.3,
			Jitter: true,
			Min:    time.Second * 15,
			Max:    time.Minute * 5,
		}

		var waitTime time.Duration
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case <-time.After(waitTime):
					waitTime = backoff.Duration()
					success := meteredClient.attemptConnect(attemptDuration)
					if success {
						return
					}
				}
			}
		}()
	}
	return &meteredClient
}

// AttemptReconnect attempts to reconnect.
func (p *potentialMeteredClient) AttemptReconnect() bool {
	return p.attemptConnect(attemptDuration)
}

// attemptConnect attempts to connect before timeout/context cancellation. Returns true if connected
// and sets lifecycles to the resulting client.
func (p *potentialMeteredClient) attemptConnect(duration time.Duration) bool {
	ctx, cancel := context.WithTimeout(p.ctx, duration)
	defer cancel()
	// as per eth docs, the context here only applies to initial connection establishment
	client, err := NewClient(ctx, p.wsURL)
	if err != nil {
		p.logOnce.Do(func() {
			logger.Warnf("could not connect to ws url %s, subsequent warnings about this issue will be suppressed", err)
		})
		return false
	}

	p.LifecycleClient = NewLifecycleClient(client, p.meteredEVMClientImpl.chainID, p, requestTimeout)
	return true
}

func (p *potentialMeteredClient) ChainConfig() *params.ChainConfig {
	return ConfigFromID(p.chainID)
}
