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
)

// MulticallDispatcher is a dispatcher that can dispatch multicalls.
type MulticallDispatcher interface {
	Start(ctx context.Context) error
	Dispatch(ctx context.Context, tx *types.Transaction) error
}

type multicallDispatcher struct {
	cfg           relconfig.Config
	txChans       map[int]chan *types.Transaction
	metricHandler metrics.Handler
	omniClient    omniClient.RPCClient
	submitter     submitter.TransactionSubmitter
}

const chanSize = 10_000

// NewMulticallDispatcher creates a new multicall dispatcher.
func NewMulticallDispatcher(cfg relconfig.Config, metricHandler metrics.Handler, submitter submitter.TransactionSubmitter) MulticallDispatcher {
	txChans := make(map[int]chan *types.Transaction)
	for chainID := range cfg.Chains {
		txChans[chainID] = make(chan *types.Transaction, chanSize)
	}
	client := omniClient.NewOmnirpcClient(cfg.OmniRPCURL, metricHandler, omniClient.WithCaptureReqRes())
	return &multicallDispatcher{
		cfg:           cfg,
		txChans:       txChans,
		metricHandler: metricHandler,
		omniClient:    client,
		submitter:     submitter,
	}
}

func (m *multicallDispatcher) Start(ctx context.Context) error {
	return nil
}

func (m *multicallDispatcher) Dispatch(ctx context.Context, tx *types.Transaction) error {
	txChan, ok := m.txChans[int(tx.ChainId().Int64())]
	if !ok {
		return fmt.Errorf("no tx channel for chain id %d", tx.ChainId().Int64())
	}
	select {
	case txChan <- tx:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// TODO: configure
var dispatchInterval = 30 * time.Second
var maxBatchSize = 100

func (m *multicallDispatcher) runDispatch(ctx context.Context, chainID int) error {
	txChan, ok := m.txChans[chainID]
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
		case tx := <-txChan:
			callQueue = append(callQueue, tx.Data())
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
