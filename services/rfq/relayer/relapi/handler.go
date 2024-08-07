package relapi

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/rfq/contracts/ierc20"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/util"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

// Handler is the REST API handler.
type Handler struct {
	metrics   metrics.Handler
	db        reldb.Service
	chains    map[uint32]*chain.Chain
	cfg       relconfig.Config
	submitter submitter.TransactionSubmitter
}

// NewHandler creates a new REST API handler.
func NewHandler(metricsHelper metrics.Handler, db reldb.Service, chains map[uint32]*chain.Chain, cfg relconfig.Config, txSubmitter submitter.TransactionSubmitter) *Handler {
	return &Handler{
		metrics:   metricsHelper,
		db:        db, // Store the database connection in the handler
		chains:    chains,
		cfg:       cfg,
		submitter: txSubmitter,
	}
}

// GetHealth returns a successful response to signify the API is up and running.
func (h *Handler) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

const unspecifiedTxHash = "Must specify 'hash' (corresponding to origin tx)"

// GetQuoteRequestStatusByTxHash gets the status of a quote request, given an origin tx hash.
func (h *Handler) GetQuoteRequestStatusByTxHash(c *gin.Context) {
	txHashStr := c.Query("hash")
	if txHashStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": unspecifiedTxHash})
		return
	}

	txHash := common.HexToHash(txHashStr)
	quoteRequest, err := h.db.GetQuoteRequestByOriginTxHash(c, txHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := GetQuoteRequestStatusResponse{
		Status:        quoteRequest.Status.String(),
		TxID:          hexutil.Encode(quoteRequest.TransactionID[:]),
		OriginTxHash:  quoteRequest.OriginTxHash.String(),
		OriginChainID: quoteRequest.Transaction.OriginChainId,
		DestChainID:   quoteRequest.Transaction.DestChainId,
		DestTxHash:    quoteRequest.DestTxHash.String(),
	}
	c.JSON(http.StatusOK, resp)
}

// GetQuoteRequestStatusByTxID gets the status of a quote request, given a tx id.
func (h *Handler) GetQuoteRequestStatusByTxID(c *gin.Context) {
	txIDStr := c.Query("id")
	if txIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must specify 'id'"})
		return
	}

	txIDBytes, err := hexutil.Decode(txIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid txID"})
		return
	}
	var txID [32]byte
	copy(txID[:], txIDBytes)

	quoteRequest, err := h.db.GetQuoteRequestByID(c, txID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := GetQuoteRequestStatusResponse{
		Status:        quoteRequest.Status.String(),
		TxID:          hexutil.Encode(quoteRequest.TransactionID[:]),
		OriginTxHash:  quoteRequest.OriginTxHash.String(),
		OriginChainID: quoteRequest.Transaction.OriginChainId,
		DestChainID:   quoteRequest.Transaction.DestChainId,
		DestTxHash:    quoteRequest.DestTxHash.String(),
	}
	c.JSON(http.StatusOK, resp)
}

// GetTxRetry retries a transaction based on tx hash.
func (h *Handler) GetTxRetry(c *gin.Context) {
	txHashStr := c.Query("hash")
	if txHashStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": unspecifiedTxHash})
		return
	}

	txHash := common.HexToHash(txHashStr)
	quoteRequest, err := h.db.GetQuoteRequestByOriginTxHash(c, txHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	chainID := quoteRequest.Transaction.DestChainId
	chainHandler, ok := h.chains[chainID]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No contract found for chain: %d", chainID)})
		return
	}

	// `quoteRequest == nil` case should be handled by the db query above
	nonce, gasAmount, err := chainHandler.SubmitRelay(c, *quoteRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("could not submit relay: %s", err.Error())})
		return
	}

	resp := GetTxRetryResponse{
		TxID:      hexutil.Encode(quoteRequest.TransactionID[:]),
		ChainID:   chainID,
		Nonce:     nonce,
		GasAmount: gasAmount.String(),
	}
	c.JSON(http.StatusOK, resp)
}

// GetQuoteRequestByTxID gets the quote request by tx id.
func (h *Handler) GetQuoteRequestByTxID(c *gin.Context) {
	txIDStr := c.Query("id")
	if txIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must specify 'id'"})
		return
	}

	txIDBytes, err := hexutil.Decode(txIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid txID"})
		return
	}
	var txID [32]byte
	copy(txID[:], txIDBytes)

	quoteRequest, err := h.db.GetQuoteRequestByID(c, txID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := GetQuoteRequestResponse{
		QuoteRequestRaw: common.Bytes2Hex(quoteRequest.RawRequest),
		OriginChainID:   quoteRequest.Transaction.OriginChainId,
		DestChainID:     quoteRequest.Transaction.DestChainId,
		OriginToken:     quoteRequest.Transaction.OriginToken.Hex(),
		DestToken:       quoteRequest.Transaction.DestToken.Hex(),
	}
	c.JSON(http.StatusOK, resp)
}

// Withdraw withdraws tokens from the relayer.
//
//nolint:cyclop
func (h *Handler) Withdraw(c *gin.Context) {
	ctx, span := h.metrics.Tracer().Start(c, "withdraw")
	var err error
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	var req WithdrawRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate the token address
	if !tokenIDExists(h.cfg, req.TokenAddress, int(req.ChainID)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid token address %s for chain %d", req.TokenAddress.Hex(), req.ChainID)})
		return
	}

	// validate the withdrawal address
	if !toAddressIsWhitelisted(h.cfg, req.To) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("withdrawal address %s is not whitelisted", req.To.Hex())})
		return
	}

	var nonce uint64

	value, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid amount %s", req.Amount)})
		return
	}

	//nolint: nestif
	if util.IsGasToken(req.TokenAddress) {
		nonce, err = h.submitter.SubmitTransaction(ctx, big.NewInt(int64(req.ChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
			bc := bind.NewBoundContract(req.To, abi.ABI{}, h.chains[req.ChainID].Client, h.chains[req.ChainID].Client, h.chains[req.ChainID].Client)
			if transactor.GasPrice != nil {
				transactor.Value = value
				// nolint: wrapcheck
				return bc.Transfer(transactor)
			}
			var signer *types.Transaction
			signer, err = transactor.Signer(h.submitter.Address(), tx)
			if err != nil {
				return nil, fmt.Errorf("could not get signer: %w", err)
			}
			return signer, nil
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("could not submit transaction: %s", err.Error())})
			return
		}
	} else {
		var erc20Contract *ierc20.IERC20
		erc20Contract, err = ierc20.NewIERC20(req.TokenAddress, h.chains[req.ChainID].Client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("could not create erc20 contract: %s", err.Error())})
			return
		}

		nonce, err = h.submitter.SubmitTransaction(ctx, big.NewInt(int64(req.ChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
			// nolint: wrapcheck
			return erc20Contract.Transfer(transactor, req.To, value)
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("could not submit transaction: %s", err.Error())})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"nonce": nonce})
}

// GetTxByNonceRequest is the request for getting a transaction hash by nonce.
type GetTxByNonceRequest struct {
	ChainID uint32 `json:"chain_id"`
	Nonce   uint64 `json:"nonce"`
}

// GetTxHashByNonce gets the transaction hash by submitter nonce.
func (h *Handler) GetTxHashByNonce(c *gin.Context) {
	ctx, span := h.metrics.Tracer().Start(c, "txByNonce")
	var err error
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	chainIDStr := c.Query("chain_id")
	nonceStr := c.Query("nonce")

	chainID, ok := new(big.Int).SetString(chainIDStr, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid chainID"})
		return
	}

	span.SetAttributes(attribute.Int("chain_id", int(chainID.Uint64())))

	nonce, ok := new(big.Int).SetString(nonceStr, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid nonce"})
		return
	}
	span.SetAttributes(attribute.Int("nonce", int(nonce.Uint64())))

	tx, err := h.submitter.GetSubmissionStatus(ctx, chainID, nonce.Uint64())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("could not get tx hash: %s", err.Error())})
		return
	}

	if tx.HasTx() {
		c.JSON(http.StatusOK, gin.H{"withdrawTxHash": tx.TxHash().String()})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
}

// tokenIDExists checks if a token ID exists in the config.
// note: this method assumes that SanitizeTokenID is a method of relconfig.Config
func tokenIDExists(cfg relconfig.Config, tokenAddress common.Address, chainID int) bool {
	for quotableToken := range cfg.QuotableTokens {
		prospectiveChainID, prospectiveAddress, err := relconfig.DecodeTokenID(quotableToken)
		if err != nil {
			continue
		}

		if prospectiveChainID == chainID && prospectiveAddress == tokenAddress {
			return true
		}
	}

	return false
}

func toAddressIsWhitelisted(cfg relconfig.Config, to common.Address) bool {
	for _, addr := range cfg.WithdrawalWhitelist {
		if common.HexToAddress(addr) == to {
			return true
		}
	}
	return false
}

// WithdrawRequest is the request to withdraw tokens from the relayer.
type WithdrawRequest struct {
	// ChainID is the chain ID of the chain to withdraw from.
	ChainID uint32 `json:"chain_id"`
	// Amount is the amount to withdraw, in wei.
	Amount string `json:"amount"`
	// TokenAddress is the address of the token to withdraw.
	TokenAddress common.Address `json:"token_address"`
	// To is the address to withdraw to.
	To common.Address `json:"to"`
}

// MarshalJSON handles JSON marshaling for WithdrawRequest.
func (wr *WithdrawRequest) MarshalJSON() ([]byte, error) {
	type Alias WithdrawRequest
	// nolint: wrapcheck
	return json.Marshal(&struct {
		TokenAddress string `json:"token_address"`
		To           string `json:"to"`
		*Alias
	}{
		TokenAddress: wr.TokenAddress.Hex(),
		To:           wr.To.Hex(),
		Alias:        (*Alias)(wr),
	})
}

// UnmarshalJSON has JSON unmarshalling for WithdrawRequest.
func (wr *WithdrawRequest) UnmarshalJSON(data []byte) error {
	type Alias WithdrawRequest
	aux := &struct {
		TokenAddress string `json:"token_address"`
		To           string `json:"to"`
		*Alias
	}{
		Alias: (*Alias)(wr),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		//nolint: wrapcheck
		return err
	}

	wr.TokenAddress = common.HexToAddress(aux.TokenAddress)
	wr.To = common.HexToAddress(aux.To)

	return nil
}
