package inventory

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

// RebalanceData contains metadata for a rebalance action.
type RebalanceData struct {
	OriginMetadata *TokenMetadata
	DestMetadata   *TokenMetadata
	Amount         *big.Int
}

// RebalanceManager is the interface for the rebalance manager.
type RebalanceManager interface {
	// Start starts the rebalance manager.
	Start(ctx context.Context) (err error)
	// Execute executes a rebalance action.
	Execute(ctx context.Context, rebalance *RebalanceData) error
}

type rebalanceManagerCCTP struct {
	// cfg is the config
	cfg relconfig.Config
	// handler is the metrics handler
	handler metrics.Handler
	// chainClient is an omnirpc client
	chainClient submitter.ClientFetcher
	// txSubmitter is the transaction submitter
	txSubmitter submitter.TransactionSubmitter
	// cctpContracts is the map of cctp contracts (used for rebalancing)
	cctpContracts map[int]*cctp.SynapseCCTP
	// relayerAddress contains the relayer address
	relayerAddress common.Address
	// db is the database
	db reldb.Service
}

func newRebalanceManagerCCTP(cfg relconfig.Config, handler metrics.Handler, chainClient submitter.ClientFetcher, txSubmitter submitter.TransactionSubmitter, relayerAddress common.Address, db reldb.Service) *rebalanceManagerCCTP {
	return &rebalanceManagerCCTP{
		cfg:            cfg,
		handler:        handler,
		chainClient:    chainClient,
		txSubmitter:    txSubmitter,
		cctpContracts:  make(map[int]*cctp.SynapseCCTP),
		relayerAddress: relayerAddress,
		db:             db,
	}
}

func (c *rebalanceManagerCCTP) Start(ctx context.Context) error {
	for chainID := range c.cfg.Chains {
		contractAddr, err := c.cfg.GetCCTPAddress(chainID)
		if err != nil {
			return fmt.Errorf("could not get cctp address: %w", err)
		}
		chainClient, err := c.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}
		contract, err := cctp.NewSynapseCCTP(common.HexToAddress(contractAddr), chainClient)
		if err != nil {
			return fmt.Errorf("could not get cctp: %w", err)
		}
		c.cctpContracts[chainID] = contract
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (c *rebalanceManagerCCTP) Execute(ctx context.Context, rebalance *RebalanceData) (err error) {
	contract, ok := c.cctpContracts[rebalance.DestMetadata.ChainID]
	if !ok {
		return fmt.Errorf("could not find cctp contract for chain %d", rebalance.DestMetadata.ChainID)
	}
	// perform rebalance by calling sendCircleToken()
	_, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(rebalance.OriginMetadata.ChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		tx, err = contract.SendCircleToken(
			transactor,
			c.relayerAddress,
			big.NewInt(int64(rebalance.DestMetadata.ChainID)),
			rebalance.OriginMetadata.Addr,
			rebalance.Amount,
			0,        // TODO: inspect
			[]byte{}, // TODO: inspect
		)
		if err != nil {
			return nil, fmt.Errorf("could not send circle token: %w", err)
		}
		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit CCTP rebalance: %w", err)
	}
	return nil
}
