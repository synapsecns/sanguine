package inventory

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/listener"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/rfq/contracts/l1gateway"
	"github.com/synapsecns/sanguine/services/rfq/contracts/l1scrollmessenger"
	"github.com/synapsecns/sanguine/services/rfq/contracts/l2gateway"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

type rebalanceManagerScroll struct {
	// cfg is the config
	cfg relconfig.Config
	// handler is the metrics handler
	handler metrics.Handler
	// chainClient is an omnirpc client
	chainClient submitter.ClientFetcher
	// txSubmitter is the transaction submitter
	txSubmitter submitter.TransactionSubmitter
	// relayerAddress contains the relayer address
	relayerAddress common.Address
	// boundL1Gateway is the L1GatewayRouter contract
	boundL1Gateway *l1gateway.L1GatewayRouter
	// boundL1ScrollMessenger is the L1ScrollMessenger contract
	boundL1ScrollMessenger *l1scrollmessenger.L1ScrollMessenger
	// boundL2Gateway is the L2GatewayRouter contract
	boundL2Gateway *l2gateway.L2GatewayRouter
	// l1GatewayListener is the listener for the L1GatewayRouter contract
	l1GatewayListener listener.ContractListener
	// l2GatewayListener is the listener for the L2GatewayRouter contract
	l2GatewayListener listener.ContractListener
	// l1ChainID is the chain ID for the L1 chain
	l1ChainID int
	// l2ChainID is the chain ID for the L2 chain
	l2ChainID int
	// db is the database
	db reldb.Service
	// apiURL is the URL for the scroll API
	apiURL *string
	// httpClient is the client for http requests
	httpClient *http.Client
}

func newRebalanceManagerScroll(cfg relconfig.Config, handler metrics.Handler, chainClient submitter.ClientFetcher, txSubmitter submitter.TransactionSubmitter, relayerAddress common.Address, db reldb.Service) *rebalanceManagerScroll {
	return &rebalanceManagerScroll{
		cfg:            cfg,
		handler:        handler,
		chainClient:    chainClient,
		txSubmitter:    txSubmitter,
		relayerAddress: relayerAddress,
		db:             db,
		httpClient:     &http.Client{},
	}
}

const mainnetChainID = 1
const scrollChainID = 534352
const sepoliaChainID = 11155111
const scrollSepoliaChainID = 534351

func isScrollChain(chainID int) bool {
	return chainID == scrollChainID || chainID == scrollSepoliaChainID
}

func isEthereumChain(chainID int) bool {
	return chainID == mainnetChainID || chainID == sepoliaChainID
}

func isTestnetChain(chainID int) bool {
	return chainID == scrollSepoliaChainID || chainID == sepoliaChainID
}

const claimCheckInterval = 30

func (c *rebalanceManagerScroll) Start(ctx context.Context) (err error) {
	err = c.initContracts(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize contracts: %w", err)
	}

	err = c.initListeners(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize listeners: %w", err)
	}

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		return c.listenL1Gateway(ctx)
	})
	g.Go(func() error {
		return c.listenL2Gateway(ctx)
	})
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(claimCheckInterval * time.Second):
				err := c.claimL2ToL1(ctx)
				if err != nil {
					logger.Warnf("could not claim: %v", err)
				}
			}
		}
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not listen: %w", err)
	}

	return nil
}

const mainnetScrollAPIURL = "https://mainnet-api-bridge-v2.scroll.io/api/"
const testnetScrollAPIURL = "https://sepolia-api-bridge-v2.scroll.io/api/"

func (c *rebalanceManagerScroll) initContracts(parentCtx context.Context) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "initContracts")
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	for chainID := range c.cfg.Chains {
		if isEthereumChain(chainID) {
			c.l1ChainID = chainID
			chainClient, err := c.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
			if err != nil {
				return fmt.Errorf("could not get chain client: %w", err)
			}
			addr, err := c.cfg.GetL1GatewayAddress(chainID)
			if err != nil {
				return fmt.Errorf("could not get l1 gateway address: %w", err)
			}
			c.boundL1Gateway, err = l1gateway.NewL1GatewayRouter(common.HexToAddress(addr), chainClient)
			if err != nil {
				return fmt.Errorf("could not get l1 gateway contract: %w", err)
			}
			addr, err = c.cfg.GetL1ScrollMessengerAddress(chainID)
			if err != nil {
				return fmt.Errorf("could not get l1 scroll messenger address: %w", err)
			}
			c.boundL1ScrollMessenger, err = l1scrollmessenger.NewL1ScrollMessenger(common.HexToAddress(addr), chainClient)
			if err != nil {
				return fmt.Errorf("could not get l1 scroll messenger contract: %w", err)
			}
		} else if isScrollChain(chainID) {
			c.l2ChainID = chainID
			addr, err := c.cfg.GetL2GatewayAddress(chainID)
			if err != nil {
				return fmt.Errorf("could not get l2 gateway address: %w", err)
			}
			chainClient, err := c.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
			if err != nil {
				return fmt.Errorf("could not get chain client: %w", err)
			}
			c.boundL2Gateway, err = l2gateway.NewL2GatewayRouter(common.HexToAddress(addr), chainClient)
			if err != nil {
				return fmt.Errorf("could not get l2 gateway contract: %w", err)
			}
		}
	}
	if c.boundL1Gateway == nil {
		return fmt.Errorf("l1 gateway contract not set")
	}
	if c.boundL2Gateway == nil {
		return fmt.Errorf("l2 gateway contract not set")
	}
	if isTestnetChain(c.l1ChainID) != isTestnetChain(c.l2ChainID) {
		return fmt.Errorf("testnet chain mismatch: %d %d", c.l1ChainID, c.l2ChainID)
	}

	// set API URL
	baseURL := mainnetScrollAPIURL
	if isTestnetChain(c.l1ChainID) {
		baseURL = testnetScrollAPIURL
	}
	url := fmt.Sprintf("%s/claimable?address=%s", baseURL, c.relayerAddress.Hex())
	c.apiURL = &url

	return nil
}

func (c *rebalanceManagerScroll) initListeners(parentCtx context.Context) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "initListeners")
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	// setup l1 listener
	l1Client, err := c.chainClient.GetClient(ctx, big.NewInt(int64(c.l1ChainID)))
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}
	l1InitialBlock, err := c.cfg.GetCCTPStartBlock(c.l1ChainID)
	if err != nil {
		return fmt.Errorf("could not get cctp start block: %w", err)
	}
	l1Addr, err := c.cfg.GetL1GatewayAddress(c.l1ChainID)
	if err != nil {
		return fmt.Errorf("could not get l1 gateway address: %w", err)
	}
	c.l1GatewayListener, err = listener.NewChainListener(l1Client, c.db, common.HexToAddress(l1Addr), l1InitialBlock, c.handler)
	if err != nil {
		return fmt.Errorf("could not get messenger listener: %w", err)
	}

	// setup l2 listener
	l2Client, err := c.chainClient.GetClient(ctx, big.NewInt(int64(c.l2ChainID)))
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}
	l2InitialBlock, err := c.cfg.GetCCTPStartBlock(c.l2ChainID)
	if err != nil {
		return fmt.Errorf("could not get cctp start block: %w", err)
	}
	l2Addr, err := c.cfg.GetL2GatewayAddress(c.l2ChainID)
	if err != nil {
		return fmt.Errorf("could not get l2 gateway address: %w", err)
	}
	c.l2GatewayListener, err = listener.NewChainListener(l2Client, c.db, common.HexToAddress(l2Addr), l2InitialBlock, c.handler)
	if err != nil {
		return fmt.Errorf("could not get messenger listener: %w", err)
	}

	return nil
}

func (c *rebalanceManagerScroll) Execute(ctx context.Context, rebalance *RebalanceData) (err error) {
	switch rebalance.OriginMetadata.ChainID {
	case c.l1ChainID:
		err = c.initiateL1ToL2(ctx, rebalance)
	case c.l2ChainID:
		err = c.initiateL2ToL1(ctx, rebalance)
	default:
		return fmt.Errorf("unexpected origin: %d", rebalance.OriginMetadata.ChainID)
	}
	if err != nil {
		return fmt.Errorf("could not execute rebalance: %w", err)
	}

	// store the rebalance in the db
	rebalanceModel := reldb.Rebalance{
		Origin:       uint64(rebalance.OriginMetadata.ChainID),
		Destination:  uint64(rebalance.DestMetadata.ChainID),
		OriginAmount: rebalance.Amount,
		Status:       reldb.RebalanceInitiated,
	}
	err = c.db.StoreRebalance(ctx, rebalanceModel)
	if err != nil {
		return fmt.Errorf("could not store rebalance: %w", err)
	}
	return nil
}

// TODO: configurable?
const scrollGasLimit = 200_000

func (c *rebalanceManagerScroll) initiateL1ToL2(ctx context.Context, rebalance *RebalanceData) (err error) {
	_, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(rebalance.OriginMetadata.ChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		if transactor == nil {
			return nil, fmt.Errorf("transactor is nil")
		}
		if chain.IsGasToken(rebalance.OriginMetadata.Addr) {
			tx, err = c.boundL1Gateway.DepositETH(transactor, rebalance.Amount, big.NewInt(int64(scrollGasLimit)))
			if err != nil {
				return nil, fmt.Errorf("could not deposit gas token: %w", err)
			}
		} else {
			tx, err = c.boundL1Gateway.DepositERC20(transactor, rebalance.OriginMetadata.Addr, rebalance.Amount, big.NewInt(int64(scrollGasLimit)))
			if err != nil {
				return nil, fmt.Errorf("could not deposit erc20 token: %w", err)
			}
		}
		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit transaction: %w", err)
	}
	return nil
}

func (c *rebalanceManagerScroll) initiateL2ToL1(ctx context.Context, rebalance *RebalanceData) (err error) {
	_, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(rebalance.OriginMetadata.ChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		if transactor == nil {
			return nil, fmt.Errorf("transactor is nil")
		}
		if chain.IsGasToken(rebalance.OriginMetadata.Addr) {
			tx, err = c.boundL2Gateway.WithdrawETH(transactor, c.relayerAddress, rebalance.Amount, big.NewInt(int64(scrollGasLimit)))
			if err != nil {
				return nil, fmt.Errorf("could not withdraw gas token: %w", err)
			}
		} else {
			tx, err = c.boundL2Gateway.WithdrawERC20(transactor, rebalance.OriginMetadata.Addr, rebalance.Amount, big.NewInt(int64(scrollGasLimit)))
			if err != nil {
				return nil, fmt.Errorf("could not withdraw erc20 token: %w", err)
			}
		}
		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit transaction: %w", err)
	}
	return nil
}

func getScrollRebalanceID(eventData []byte) string {
	return common.BytesToHash(eventData).Hex()
}

func (c *rebalanceManagerScroll) listenL1Gateway(ctx context.Context) (err error) {
	addr, err := c.cfg.GetL1GatewayAddress(c.l1ChainID)
	if err != nil {
		return fmt.Errorf("could not get l1 gateway address: %w", err)
	}
	parser, err := l1gateway.NewParser(common.HexToAddress(addr))
	if err != nil {
		return fmt.Errorf("could not get l1 gateway parser: %w", err)
	}
	err = c.l1GatewayListener.Listen(ctx, func(parentCtx context.Context, log types.Log) (err error) {
		et, parsedEvent, ok := parser.ParseEvent(log)
		if !ok {
			return nil
		}

		ctx, span := c.handler.Tracer().Start(parentCtx, fmt.Sprintf("handleLog-%s", et), trace.WithAttributes(
			attribute.String(metrics.TxHash, log.TxHash.String()),
			attribute.String(metrics.Contract, log.Address.String()),
			attribute.String("block_hash", log.BlockHash.String()),
			attribute.Int64("block_number", int64(log.BlockNumber)),
		))
		defer func() {
			metrics.EndSpanWithErr(span, err)
		}()

		switch event := parsedEvent.(type) {
		case *l1gateway.L1GatewayRouterDepositETH:
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID:     &rebalanceID,
				Origin:          uint64(c.l1ChainID),
				OriginTxHash:    log.TxHash,
				OriginTokenAddr: chain.EthAddress,
				Destination:     uint64(c.l2ChainID),
				Status:          reldb.RebalancePending,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case *l1gateway.L1GatewayRouterDepositERC20:
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID:     &rebalanceID,
				Origin:          uint64(c.l1ChainID),
				OriginTxHash:    log.TxHash,
				OriginTokenAddr: event.L1Token,
				Destination:     uint64(c.l2ChainID),
				Status:          reldb.RebalancePending,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case *l1gateway.L1GatewayRouterFinalizeWithdrawETH:
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID: &rebalanceID,
				DestTxHash:  log.TxHash,
				Status:      reldb.RebalanceCompleted,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case *l1gateway.L1GatewayRouterFinalizeWithdrawERC20:
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID: &rebalanceID,
				DestTxHash:  log.TxHash,
				Status:      reldb.RebalanceCompleted,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("could not listen for L1GatewayRouter events: %w", err)
	}
	return nil
}

func (c *rebalanceManagerScroll) listenL2Gateway(ctx context.Context) (err error) {
	addr, err := c.cfg.GetL2GatewayAddress(c.l1ChainID)
	if err != nil {
		return fmt.Errorf("could not get l2 gateway address: %w", err)
	}
	parser, err := l2gateway.NewParser(common.HexToAddress(addr))
	if err != nil {
		return fmt.Errorf("could not get l2 gateway parser: %w", err)
	}
	err = c.l2GatewayListener.Listen(ctx, func(parentCtx context.Context, log types.Log) (err error) {
		et, parsedEvent, ok := parser.ParseEvent(log)
		if !ok {
			return nil
		}

		ctx, span := c.handler.Tracer().Start(parentCtx, fmt.Sprintf("handleLog-%s", et), trace.WithAttributes(
			attribute.String(metrics.TxHash, log.TxHash.String()),
			attribute.String(metrics.Contract, log.Address.String()),
			attribute.String("block_hash", log.BlockHash.String()),
			attribute.Int64("block_number", int64(log.BlockNumber)),
		))
		defer func() {
			metrics.EndSpanWithErr(span, err)
		}()

		switch event := parsedEvent.(type) {
		case *l2gateway.L2GatewayRouterWithdrawETH:
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID:     &rebalanceID,
				Origin:          uint64(c.l2ChainID),
				OriginTxHash:    log.TxHash,
				OriginTokenAddr: chain.EthAddress,
				Destination:     uint64(c.l1ChainID),
				Status:          reldb.RebalancePending,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case *l2gateway.L2GatewayRouterWithdrawERC20:
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID:     &rebalanceID,
				Origin:          uint64(c.l2ChainID),
				OriginTxHash:    log.TxHash,
				OriginTokenAddr: event.L2Token,
				Destination:     uint64(c.l1ChainID),
				Status:          reldb.RebalancePending,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case *l2gateway.L2GatewayRouterFinalizeDepositETH:
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID: &rebalanceID,
				DestTxHash:  log.TxHash,
				Status:      reldb.RebalanceCompleted,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case *l2gateway.L2GatewayRouterFinalizeDepositERC20:
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID: &rebalanceID,
				DestTxHash:  log.TxHash,
				Status:      reldb.RebalanceCompleted,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("could not listen for L2GatewayRouter events: %w", err)
	}
	return nil
}

type scrollAPIResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Data    struct {
		Result []struct {
			Hash           string      `json:"hash"`
			Amount         string      `json:"amount"`
			To             string      `json:"to"`
			IsL1           bool        `json:"isL1"`
			L1Token        string      `json:"l1Token"`
			L2Token        string      `json:"l2Token"`
			BlockNumber    int         `json:"blockNumber"`
			BlockTimestamp interface{} `json:"blockTimestamp"`
			FinalizeTx     struct {
				Hash           string      `json:"hash"`
				Amount         string      `json:"amount"`
				To             string      `json:"to"`
				IsL1           bool        `json:"isL1"`
				BlockNumber    int         `json:"blockNumber"`
				BlockTimestamp interface{} `json:"blockTimestamp"`
			} `json:"finalizeTx"`
			ClaimInfo   ClaimInfo   `json:"claimInfo"`
			CreatedTime interface{} `json:"createdTime"`
		} `json:"result"`
		Total int `json:"total"`
	} `json:"data"`
}

type ClaimInfo struct {
	From       string `json:"from"`
	To         string `json:"to"`
	Value      string `json:"value"`
	Nonce      string `json:"nonce"`
	BatchHash  string `json:"batch_hash"`
	Message    string `json:"message"`
	Proof      string `json:"proof"`
	BatchIndex string `json:"batch_index"`
}

func (c *rebalanceManagerScroll) claimL2ToL1(parentCtx context.Context) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "claimL2ToL1")
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	if c.apiURL == nil {
		return fmt.Errorf("api URL not set")
	}
	span.SetAttributes(attribute.String("api_url", *c.apiURL))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, *c.apiURL, nil)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not get response: %w", err)
	}
	//nolint:errcheck
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var claimableResp scrollAPIResponse
	err = json.Unmarshal(body, &claimableResp)
	if err != nil {
		return fmt.Errorf("could not unmarshal body: %w", err)
	}

	for _, result := range claimableResp.Data.Result {
		err = c.submitClaim(ctx, result.ClaimInfo)
		if err != nil {
			return fmt.Errorf("could not submit transaction: %w", err)
		}
	}
	return nil
}

func (c *rebalanceManagerScroll) submitClaim(parentCtx context.Context, claimInfo ClaimInfo) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "submitClaim", trace.WithAttributes(
		attribute.String("from", claimInfo.From),
		attribute.String("to", claimInfo.To),
		attribute.String("value", claimInfo.Value),
		attribute.String("nonce", claimInfo.Nonce),
		attribute.String("batch_index", claimInfo.BatchIndex),
	))
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	_, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(c.l1ChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		if transactor == nil {
			return nil, fmt.Errorf("transactor is nil")
		}
		// Note: we hardcode the 'to' parameter as our own relayerAddress as a safety measure.
		value, ok := new(big.Int).SetString(claimInfo.Value, 10)
		if !ok {
			return nil, fmt.Errorf("could not parse value: %w", err)
		}
		nonce, ok := new(big.Int).SetString(claimInfo.Nonce, 10)
		if !ok {
			return nil, fmt.Errorf("could not parse nonce: %w", err)
		}
		batchIndex, ok := new(big.Int).SetString(claimInfo.BatchIndex, 10)
		if !ok {
			return nil, fmt.Errorf("could not parse batch index: %w", err)
		}
		message, err := hexutil.Decode(claimInfo.Message)
		if err != nil {
			return nil, fmt.Errorf("could not decode message: %w", err)
		}
		merkleProof, err := hexutil.Decode(claimInfo.Proof)
		if err != nil {
			return nil, fmt.Errorf("could not decode merkle proof: %w", err)
		}
		proof := l1scrollmessenger.IL1ScrollMessengerL2MessageProof{
			BatchIndex:  batchIndex,
			MerkleProof: merkleProof,
		}
		tx, err = c.boundL1ScrollMessenger.RelayMessageWithProof(transactor, common.HexToAddress(claimInfo.From), c.relayerAddress, value, nonce, message, proof)
		if err != nil {
			return nil, fmt.Errorf("could not relay message: %w", err)
		}
		return tx, nil
	})
	return nil
}
