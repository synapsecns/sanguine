package chainmanager

import (
	"context"
	"fmt"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/omnirpc/config"
	"github.com/synapsecns/sanguine/services/omnirpc/rpcinfo"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"sort"
	"sync"
	"time"
)

var logger = log.Logger("chainmanager")

// rpcTimeout is how long to wait for a response.
const rpcTimeout = time.Second * 5

// ChainManager manages chain context.
type ChainManager interface {
	// GetChainIDs gets all chainids
	GetChainIDs() (chainIDs []uint32)
	// RefreshRPCInfo refreshes rpc info for a chain id
	RefreshRPCInfo(ctx context.Context, chainID uint32)
	// GetChain gets the chain
	GetChain(chainID uint32) Chain
	// PutChain adds chain urls. Any previous chain data is overwritten
	PutChain(chainID uint32, urls []string, confirmations uint16)
}

// NewChainManager creates a new chain manager.
func NewChainManager(handler metrics.Handler) ChainManager {
	return &chainManager{
		chainList: make(map[uint32]*chain),
		// mux is used to prevent parallel manipulations to the map
		mux: sync.RWMutex{},
		// handler is the metrics handler
		handler: handler,
	}
}

// NewChainManagerFromConfig creates a new chain manager.
func NewChainManagerFromConfig(configuration config.Config, handler metrics.Handler) ChainManager {
	cm := &chainManager{
		chainList: make(map[uint32]*chain),
		mux:       sync.RWMutex{},
		handler:   handler,
	}

	for chainID, chn := range configuration.Chains {
		// default the confirmation threshold to 1
		confThreshold := uint16(1)

		// if confirmation threshold is 1, set the checks to 1
		if chn.Checks > 0 {
			confThreshold = chn.Checks
		}

		// store all the chains w/ empty latency results
		chains := make([]rpcinfo.Result, len(chn.RPCs))
		for i := range chn.RPCs {
			chains[i] = rpcinfo.Result{
				URL: chn.RPCs[i],
			}
		}

		cm.chainList[chainID] = &chain{
			chainID:               chainID,
			confirmationThreshold: confThreshold,
			rpcs:                  chains,
		}
	}

	err := cm.setupMetrics()
	if err != nil {
		logger.Errorf("could not setup metrics: %v", err)
	}

	return cm
}

// chainManager contains a chain manager.
type chainManager struct {
	chainList map[uint32]*chain
	mux       sync.RWMutex
	handler   metrics.Handler
}

func (c *chainManager) GetChain(chainID uint32) Chain {
	c.mux.RLock()
	defer c.mux.RUnlock()

	res, ok := c.chainList[chainID]
	if !ok {
		return nil
	}

	return res
}

func (c *chainManager) GetChainIDs() (chainIDs []uint32) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	chainIDs = make([]uint32, len(c.chainList))
	i := 0
	for chainID := range c.chainList {
		chainIDs[i] = chainID
		i++
	}
	return chainIDs
}

// PutChain puts new chain urls.
func (c *chainManager) PutChain(chainID uint32, urls []string, confirmations uint16) {
	rpcs := make([]rpcinfo.Result, len(urls))
	for i, url := range urls {
		rpcs[i] = rpcinfo.Result{
			URL: url,
		}
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	c.chainList[chainID] = &chain{
		chainID:               chainID,
		rpcs:                  rpcs,
		confirmationThreshold: confirmations,
	}
}

// RefreshRPCInfo refreshes rpc info for a given chain id.
func (c *chainManager) RefreshRPCInfo(ctx context.Context, chainID uint32) {
	c.mux.RLock()
	chainList, ok := c.chainList[chainID]
	c.mux.RUnlock()

	// nothing to reorder
	if !ok {
		return
	}
	rpcURLS := chainList.URLs()

	rpcInfoList := sortInfoList(rpcinfo.GetRPCLatency(ctx, rpcTimeout, rpcURLS, c.handler))

	c.mux.Lock()
	c.chainList[chainID].rpcs = rpcInfoList
	c.mux.Unlock()
}

const (
	meter             = "github.com/synapsecns/sanguine/services/omnirpc/chainmanager"
	blockNumberMetric = "block_number"
	latencyMetric     = "latency"
	blockAgeMetric    = "block_age"
)

// records metrics for various rpcs. Should only be called once.
//
// note: because of missing support for  https://github.com/open-telemetry/opentelemetry-specification/issues/2318
// this is done from the struct rather than recorded at refresh time.
//
// in a future version, this should be a synchronous gauge.
func (c *chainManager) setupMetrics() error {
	meterMaid := c.handler.Meter(meter)
	blockGauge, err := meterMaid.Int64ObservableGauge(blockNumberMetric)
	if err != nil {
		return fmt.Errorf("could not create histogram: %w", err)
	}

	latencyGauge, err := meterMaid.Float64ObservableGauge(latencyMetric, metric.WithUnit("seconds"))
	if err != nil {
		return fmt.Errorf("could not create histogram: %w", err)
	}

	ageGauge, err := meterMaid.Float64ObservableGauge(blockAgeMetric, metric.WithUnit("seconds"))
	if err != nil {
		return fmt.Errorf("could not create histogram: %w", err)
	}

	if _, err := meterMaid.RegisterCallback(func(parentCtx context.Context, o metric.Observer) (err error) {
		c.mux.RLock()
		defer c.mux.RUnlock()

		for chainID, chainInfo := range c.chainList {
			for _, rpc := range chainInfo.rpcs {
				attributeSet := attribute.NewSet(attribute.Int64(metrics.ChainID, int64(chainID)), attribute.String("rpc_url", rpc.URL))

				if rpc.HasError {
					continue
				}

				o.ObserveInt64(blockGauge, int64(rpc.BlockNumber), metric.WithAttributeSet(attributeSet))
				o.ObserveFloat64(latencyGauge, rpc.Latency.Seconds(), metric.WithAttributeSet(attributeSet))
				o.ObserveFloat64(ageGauge, rpc.BlockAge.Seconds(), metric.WithAttributeSet(attributeSet))
			}
		}

		return nil
	}, blockGauge, latencyGauge, ageGauge); err != nil {
		return fmt.Errorf("could not register callback for gauges: %w", err)
	}
	return nil
}

func sortInfoList(rpcInfoList []rpcinfo.Result) []rpcinfo.Result {
	sort.Slice(rpcInfoList, func(i, j int) bool {
		// ignore latencies with an error
		if rpcInfoList[i].HasError {
			return false
		}

		ageDifference := rpcInfoList[i].BlockAge - rpcInfoList[j].BlockAge
		if ageDifference == 0 {
			return rpcInfoList[i].Latency < rpcInfoList[j].Latency
		} else if ageDifference > 0 {
			return false
		}

		return true
	})

	return rpcInfoList
}

var _ ChainManager = &chainManager{}

// Chain contains the context for a single chain.
//
//go:generate go run github.com/vektra/mockery/v2 --name Chain --output ./mocks --case=underscore
type Chain interface {
	// ConfirmationsThreshold gets the confirmation count
	ConfirmationsThreshold() uint16
	// URLs gets the urls
	URLs() []string
	// ID returns the id of the chain
	ID() uint32
}

// chain contains the settings for a single chain.
type chain struct {
	// chainID is the chainid
	chainID uint32
	// confirmationThreshold is the confirmation threshold of the chain
	confirmationThreshold uint16
	// rpcs contains a list of rpcs sorted by speed
	rpcs []rpcinfo.Result
}

func (c *chain) ID() uint32 {
	return c.chainID
}

func (c *chain) ConfirmationsThreshold() uint16 {
	return c.confirmationThreshold
}

// URLs gets all urls for a chain.
func (c *chain) URLs() (res []string) {
	res = make([]string, len(c.rpcs))
	for i, chainInfo := range c.rpcs {
		res[i] = chainInfo.URL
	}
	return res
}

var _ Chain = &chain{}
