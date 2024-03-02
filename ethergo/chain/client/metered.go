package client

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/semaphore"
	"golang.org/x/time/rate"
	"math/big"
	"sync/atomic"
	"time"
)

// MeteredEVMClient is a metered client that exposes a counter metric for evm requests sent through the client.
//
//go:generate go run github.com/vektra/mockery/v2 --name MeteredEVMClient --output ./mocks --case=underscore
type MeteredEVMClient interface {
	EVMClient
	// RequestCount gets the request count from the metered evm client
	RequestCount() int64
	// ConcurrencyCount gets the request concurrency on the meterd client
	ConcurrencyCount() int32
	// ClientID is a unique identifier for the client.
	//
	// note: this is not guaranteed to be unique - it's set by the caller.
	ClientID() string
	// AttemptReconnect attempts to reconnect
	// TODO: replace with https://github.com/ethereum/go-ethereum/issues/22266
	AttemptReconnect() bool
}

// meteredEVMClientImpl is an instrumented client with a rate limiter
// it wraps EVMClient and takes the keepRateLimiter as a config
// we can't use the original keep rate limiter since it does not respect
// the context analytics.
type meteredEVMClientImpl struct {
	LifecycleClient
	// semaphore is used to manage concurrency limits. If
	// concurrency limiting is turned off, this is not used
	semaphore *semaphore.Weighted
	// limiter is used to limit the number of requests
	limiter *rate.Limiter
	// acquirePermitTimeout is the max amount of time a metered client
	// will wait to acquire a semaphore and make a request. This will default to 5 minutes
	acquirePermitTimeout time.Duration
	// counter is the counter for total number of requests since client creation
	counter uint64
	// concurrency is the number of ongoing requests at any given time
	concurrency uint32
	// chainID is the chain id used for the client
	chainID *big.Int
	// clientID is a unique identifier used for metrics across a pool.
	// this can be the wsurl in many cases (if metrics are private/authentication on rpc server)
	// is properly secured without a token based url
	clientID string
}

// AttemptReconnect attempts to reconnect
// TODO implement.
func (m meteredEVMClientImpl) AttemptReconnect() bool {
	return true
}

const requestTimeout = time.Minute

// NewMeteredClient wraps an evm client in a keepRate limiter and creates
// a metric handler with some standard metrics. It also implements a ChainConfig()
// method to get the chainconfig for a given chain by id. This will return nil if no chain config is found.
func NewMeteredClient(client EVMClient, chainID *big.Int, clientID string, config *LimiterConfig) MeteredEVMClient {
	meteredClient := getMeteredClientStub(chainID, clientID, config)
	meteredClient.LifecycleClient = NewLifecycleClient(client, chainID, meteredClient, requestTimeout)
	return &meteredClient
}

// getMeteredClientStub stub gets the metered client without setting the lifecylce using the config.
func getMeteredClientStub(chainID *big.Int, clientID string, config *LimiterConfig) meteredEVMClientImpl {
	meteredClient := meteredEVMClientImpl{
		// counter is an atomicically incrementing int used to instrument eth client
		counter:  0,
		chainID:  chainID,
		clientID: clientID,
	}

	if config == nil {
		config = &LimiterConfig{}
	}

	// setup the rate limiter
	if config.RequestsPerSecondLimit > 0 {
		meteredClient.limiter = rate.NewLimiter(
			rate.Limit(config.RequestsPerSecondLimit),
			1,
		)
	}

	if config.ConcurrencyLimit > 0 {
		meteredClient.semaphore = semaphore.NewWeighted(
			int64(config.ConcurrencyLimit),
		)
	}

	if config.AcquirePermitTimeout > 0 {
		meteredClient.acquirePermitTimeout = config.AcquirePermitTimeout
	} else {
		meteredClient.acquirePermitTimeout = 5 * time.Minute
	}

	return meteredClient
}

// ClientID gets the unique client identifier.
func (m meteredEVMClientImpl) ClientID() string {
	return m.clientID
}

// AcquirePermit attempts to acquire a permit from the keepRate limiter.
// (this needs to be called before the request).
func (m meteredEVMClientImpl) AcquirePermit(ctx context.Context) (err error) {
	ctx, cancel := context.WithTimeout(ctx, m.acquirePermitTimeout)
	defer cancel()

	if m.limiter != nil {
		err = m.limiter.Wait(ctx)
		if err != nil {
			return fmt.Errorf("cannot wait for limiter: %w", err)
		}
	}

	if m.semaphore != nil {
		err = m.semaphore.Acquire(ctx, 1)
		if err != nil {
			return fmt.Errorf("could not acquire semaphore: %w", err)
		}
	}

	atomic.AddUint64(&m.counter, 1)
	atomic.AddUint32(&m.concurrency, 1)
	return nil
}

// ReleasePermit releases a permit (this needs to be called after the request).
func (m meteredEVMClientImpl) ReleasePermit() {
	if m.semaphore != nil {
		m.semaphore.Release(1)
	}
	// decrement the concurrency counter (see the atomic.AddUint32())
	atomic.AddUint32(&m.concurrency, ^uint32(0))
}

// RequestCount gets the request count.
func (m meteredEVMClientImpl) RequestCount() int64 {
	return int64(atomic.LoadUint64(&m.counter))
}

// ConcurrencyCount gets the number of requests in progress.
func (m meteredEVMClientImpl) ConcurrencyCount() int32 {
	return int32(atomic.LoadUint32(&m.concurrency))
}

// RequestCountMetricName is the name of the metric which counts requests.
const RequestCountMetricName = "request_count_total"

// ConcurrencyGaugeMetricName is the name of the metric which gauges request concurrency.
const ConcurrencyGaugeMetricName = "request_concurrency"

// GetMetrics gets metrics associated with the network provider.
func (m meteredEVMClientImpl) GetMetrics(labels map[string]string) []prometheus.Collector {
	requestCount := prometheus.NewCounterFunc(prometheus.CounterOpts{
		Name: RequestCountMetricName,
		Help: "the number of requests sent by the client",
	}, func() float64 {
		return float64(m.RequestCount())
	})

	concurrencyCount := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: ConcurrencyGaugeMetricName,
		Help: "the number of requests sent by the client concurrently",
	}, func() float64 {
		return float64(m.ConcurrencyCount())
	})

	return []prometheus.Collector{requestCount, concurrencyCount}
}

var _ MeteredEVMClient = &meteredEVMClientImpl{}

// LimiterConfig represents the configuration of the rate limiter.
// copied from https://github.com/keep-network/keep-common/blob/v1.7.0/pkg/rate/limiter.go#L19
type LimiterConfig struct {
	// RequestsPerSecondLimit sets the maximum average number of requests
	// per second. It's important to note that in short periods of time
	// the actual average may exceed this limit slightly.
	RequestsPerSecondLimit int

	// ConcurrencyLimit sets the maximum number of concurrent requests which
	// can be executed against the target at the same time.
	ConcurrencyLimit int

	// AcquirePermitTimeout determines how long a request can wait trying
	// to acquire a permit from the rate limiter.
	AcquirePermitTimeout time.Duration
}
