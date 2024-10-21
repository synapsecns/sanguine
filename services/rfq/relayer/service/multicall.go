package service

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridgemulti"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"golang.org/x/sync/errgroup"
)

// MulticallDispatcher is a dispatcher that can dispatch multicalls.
type MulticallDispatcher interface {
	Start(ctx context.Context) error
	Dispatch(ctx context.Context, chainID int, callData []byte) error
}

type multicallDispatcher struct {
	cfg           relconfig.Config
	callChans     map[int]chan []byte
	metricHandler metrics.Handler
	omniClient    omniClient.RPCClient
	submitter     submitter.TransactionSubmitter
}

const chanSize = 10_000

// NewMulticallDispatcher creates a new multicall dispatcher.
func NewMulticallDispatcher(cfg relconfig.Config, metricHandler metrics.Handler, submitter submitter.TransactionSubmitter) MulticallDispatcher {
	callChans := make(map[int]chan []byte)
	for chainID := range cfg.Chains {
		callChans[chainID] = make(chan []byte, chanSize)
	}
	client := omniClient.NewOmnirpcClient(cfg.OmniRPCURL, metricHandler, omniClient.WithCaptureReqRes())
	return &multicallDispatcher{
		cfg:           cfg,
		callChans:     callChans,
		metricHandler: metricHandler,
		omniClient:    client,
		submitter:     submitter,
	}
}

func (m *multicallDispatcher) Start(ctx context.Context) error {
	g, gctx := errgroup.WithContext(ctx)
	for c := range m.cfg.Chains {
		chainID := c
		g.Go(func() error {
			err := m.runQueue(gctx, chainID)
			if err != nil {
				return fmt.Errorf("could not run dispatch: %w", err)
			}
			return nil
		})
	}

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("could not start multicall dispatcher: %w", err)
	}

	return nil
}

func (m *multicallDispatcher) Dispatch(ctx context.Context, chainID int, callData []byte) error {
	callChan, ok := m.callChans[chainID]
	if !ok {
		return fmt.Errorf("no tx channel for chain id %d", chainID)
	}
	select {
	case callChan <- callData:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// TODO: configure
var dispatchInterval = 30 * time.Second
var maxBatchSize = 100

func (m *multicallDispatcher) runQueue(ctx context.Context, chainID int) error {
	callChan, ok := m.callChans[chainID]
	if !ok {
		return fmt.Errorf("no tx channel for chain id %d", chainID)
	}

	rfqAddr, err := m.cfg.GetRFQAddress(chainID)
	if err != nil {
		return fmt.Errorf("could not get rfq address: %w", err)
	}
	chainClient, err := m.omniClient.GetChainClient(ctx, chainID)
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}
	contract, err := fastbridgemulti.NewFastBridgeRef(rfqAddr, chainClient)
	if err != nil {
		return fmt.Errorf("could not create fast bridge contract at address %s: %w", contract.Address(), err)
	}

	callQueue := [][]byte{}
	ticker := time.NewTicker(dispatchInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case call := <-callChan:
			callQueue = append(callQueue, call)
			if len(callQueue) >= maxBatchSize {
				if err := m.multicall(ctx, contract, chainID, callQueue); err != nil {
					return fmt.Errorf("could not multicall: %w", err)
				}
			}
		case <-ticker.C:
			if len(callQueue) > 0 {
				if err := m.multicall(ctx, contract, chainID, callQueue); err != nil {
					return fmt.Errorf("could not multicall: %w", err)
				}
				callQueue = [][]byte{}
			}
		}
	}
}

func (m *multicallDispatcher) multicall(ctx context.Context, contract *fastbridgemulti.FastBridgeRef, chainID int, callQueue [][]byte) error {
	_, err := m.submitter.SubmitTransaction(ctx, big.NewInt(int64(chainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		tx, err = contract.MulticallNoResults(transactor, callQueue, true)
		if err != nil {
			return nil, fmt.Errorf("could not multicall: %w", err)
		}
		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit transaction: %w", err)
	}

	// empty the call queue
	callQueue = [][]byte{}
	return nil
}
