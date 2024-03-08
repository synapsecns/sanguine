// Package relapi provides RESTful API services for the RFQ relayer
package relapi

import (
	"context"
	"fmt"

	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/ethergo/submitter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/metrics"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/ethergo/listener"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

// RelayerAPIServer is a struct that holds the configuration, database connection, gin engine, RPC client, metrics handler, and fast bridge contracts.
// It is used to initialize and run the API server.
type RelayerAPIServer struct {
	cfg     relconfig.Config
	db      reldb.Service
	engine  *gin.Engine
	handler metrics.Handler
	chains  map[uint32]*chain.Chain
}

// NewRelayerAPI holds the configuration, database connection, gin engine, RPC client, metrics handler, and fast bridge contracts.
// It is used to initialize and run the API server.
//
//nolint:cyclop
func NewRelayerAPI(
	ctx context.Context,
	cfg relconfig.Config,
	handler metrics.Handler,
	omniRPCClient omniClient.RPCClient,
	store reldb.Service,
	submitter submitter.TransactionSubmitter,
) (*RelayerAPIServer, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is nil")
	}
	if handler == nil {
		return nil, fmt.Errorf("handler is nil")
	}
	if omniRPCClient == nil {
		return nil, fmt.Errorf("omniRPCClient is nil")
	}
	if store == nil {
		return nil, fmt.Errorf("store is nil")
	}

	chains := make(map[uint32]*chain.Chain)
	for chainID, chainCfg := range cfg.Chains {
		chainClient, err := omniRPCClient.GetChainClient(ctx, chainID)
		if err != nil {
			return nil, fmt.Errorf("could not create omnirpc client: %w", err)
		}
		rfqAddr, err := cfg.GetRFQAddress(chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get rfq address: %w", err)
		}
		contract, err := fastbridge.NewFastBridgeRef(common.HexToAddress(rfqAddr), chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not create fast bridge contract: %w", err)
		}
		startBlock, err := contract.DeployBlock(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("could not get deploy block: %w", err)
		}
		chainListener, err := listener.NewChainListener(chainClient, store, common.HexToAddress(rfqAddr), uint64(startBlock.Int64()), handler)
		if err != nil {
			return nil, fmt.Errorf("could not get chain listener: %w", err)
		}
		chains[uint32(chainID)], err = chain.NewChain(ctx, chainClient, common.HexToAddress(chainCfg.RFQAddress), chainListener, submitter)
		if err != nil {
			return nil, fmt.Errorf("could not create chain: %w", err)
		}
	}

	return &RelayerAPIServer{
		cfg:     cfg,
		db:      store,
		handler: handler,
		chains:  chains,
	}, nil
}

const (
	getHealthRoute              = "/health"
	getQuoteStatusByTxHashRoute = "/status"
	getQuoteStatusByTxIDRoute   = "/status/by_tx_id"
	getRetryRoute               = "/retry"
)

var logger = log.Logger("relayer-api")

// Run runs the rest api server.
func (r *RelayerAPIServer) Run(ctx context.Context) error {
	// TODO: Use Gin Helper
	engine := ginhelper.New(logger)
	h := NewHandler(r.db, r.chains)

	// Assign GET routes
	engine.GET(getHealthRoute, h.GetHealth)
	engine.GET(getQuoteStatusByTxHashRoute, h.GetQuoteRequestStatusByTxHash)
	engine.GET(getQuoteStatusByTxIDRoute, h.GetQuoteRequestStatusByTxID)
	engine.GET(getRetryRoute, h.GetTxRetry)

	r.engine = engine

	connection := baseServer.Server{}
	fmt.Printf("starting api at http://localhost:%s\n", r.cfg.RelayerAPIPort)
	err := connection.ListenAndServe(ctx, fmt.Sprintf(":%s", r.cfg.RelayerAPIPort), r.engine)
	if err != nil {
		return fmt.Errorf("could not start relayer api server: %w", err)
	}

	return nil
}
