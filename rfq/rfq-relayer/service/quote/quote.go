package quote

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	EVMClient "github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/bindings"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/config"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/service/balance"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/utils"
	"gorm.io/gorm"
)

const bridgeReqCacheSize = 10000

// IQuoter is the interface for a Quoter.
type IQuoter interface {
	GetValidQuote(quoteID string, destTokenID string, destVolume *big.Int) (*Quote, error)
	PublishQuotes() error
	UpdateQuotes(quoteID string) error
	GetQuotes(quoteID string) []*Quote
	QuoteToAPIQuote(quote *Quote) (*APIQuote, error)
	HandleUnconfirmedBridgeRequest(req *bindings.FastBridgeBridgeRequested) error
	HandleCompletedBridge(transactionID string, event *bindings.IFastBridgeBridgeTransaction) error
	HandleUncompletedBridge(transactionID string, event *bindings.IFastBridgeBridgeTransaction) error
}

type quoterImpl struct {
	mux              sync.RWMutex
	quotes           map[string][]*Quote // sorted quotes
	balance          balance.IBalanceManager
	relayer          common.Address
	bridgeReqHandler IBridgeReqHandler
	rfqURL           string
}

// Quote holds all the data for a quote.
type Quote struct {
	QuoteID       string   // Hash of relayer, sell/buy asset, sell/buy chainID
	Gas           *big.Int // Customizable by the relayer config
	OriginToken   common.Address
	DestToken     common.Address
	OriginChainID uint32
	DestChainID   uint32
	MaxVolume     *big.Int // Max volume the quote applies for
	APIID         int
	// Other fields to calculate the actual fee can go here.
}

// APIQuote is the struct for the quote API.
type APIQuote struct {
	Relayer string `json:"relayer" binding:"required"`

	OriginChainID    uint    `json:"origin_chain_id" binding:"required"`
	OriginToken      string  `json:"origin_token" binding:"required"`
	OriginAmount     string  `json:"origin_amount" binding:"required"`
	OriginAmountNorm float64 `json:"origin_amount_norm" binding:"required"`
	OriginDecimals   uint8   `json:"origin_decimals" binding:"required"`

	DestChainID    uint    `json:"dest_chain_id" binding:"required"`
	DestToken      string  `json:"dest_token" binding:"required"`
	DestAmount     string  `json:"dest_amount" binding:"required"`
	DestAmountNorm float64 `json:"dest_amount_norm" binding:"required"`
	DestDecimals   uint8   `json:"dest_decimals" binding:"required"`

	Price     float64        `json:"price"` // price = destAmount <quote> / originAmount <base>
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// NewQuoter creates a new quoter.
func NewQuoter(ctx context.Context, clients map[uint32]EVMClient.EVM, assets []config.AssetConfig, relayer common.Address, rfqURL string) (IQuoter, error) {
	// Create balance manager
	balanceManager, err := balance.NewBalanceManager(clients, assets, relayer)
	if err != nil {
		return nil, fmt.Errorf("could not create balance manager: %w", err)
	}

	// Get the balances for each asset
	err = balanceManager.GetAllOnChainBalances(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get on chain balances: %w", err)
	}

	bridgeReqHandler := NewBridgeReqs(bridgeReqCacheSize)

	quotes := make(map[string][]*Quote)
	for _, asset := range assets {
		tokenAddress := common.HexToAddress(asset.Address)
		// Create quote for every other asset on every other chain
		for _, destAsset := range assets {
			if asset.ChainID == destAsset.ChainID {
				continue
			}
			destTokenAddress := common.HexToAddress(destAsset.Address)
			quoteID := utils.GenerateQuoteID(asset.ChainID, tokenAddress, destAsset.ChainID, destTokenAddress)

			// Generate quote base
			quote := &Quote{
				QuoteID:       quoteID,
				OriginToken:   tokenAddress,
				DestToken:     destTokenAddress,
				OriginChainID: asset.ChainID,
				DestChainID:   destAsset.ChainID,
				MaxVolume:     big.NewInt(0),
				APIID:         -1,
			}

			// TODO: For add any future quote logic config here (make a helper function)
			// TODO: When we have more than a single quote for a single asset/chain pair, the quotes should be sorted.

			// Add quote to map
			quotes[quoteID] = append(quotes[quoteID], quote)
		}
	}

	return &quoterImpl{
		relayer:          relayer,
		quotes:           quotes,
		balance:          balanceManager,
		bridgeReqHandler: bridgeReqHandler,
		rfqURL:           rfqURL,
	}, nil
}

// GetQuotes gets all quotes for a given quoteID.
func (q *quoterImpl) GetQuotes(quoteID string) []*Quote {
	q.mux.RLock()
	defer q.mux.RUnlock()
	// TODO: investigate wether thees should be pointers or not.
	return q.quotes[quoteID]
}

func (q *quoterImpl) PublishQuotes() error {
	q.mux.Lock()
	defer q.mux.Unlock()

	for _, quotes := range q.quotes {
		for _, quote := range quotes {
			_, err := q.QuoteToAPIQuote(quote)
			if err != nil {
				return fmt.Errorf("could not convert quote to API quote: %w", err)
			}

			// TODO: Publish quote to quote API
			// TODO: /publish quote
			// @mikeyf How should we go about this http connection
		}
	}
	return nil
}

// UpdateQuotes updates the quotes in the quote API.
func (q *quoterImpl) UpdateQuotes(quoteID string) error {
	q.mux.Lock()
	defer q.mux.Unlock()

	// Get the quote from the quote API
	for _, quote := range q.quotes[quoteID] {
		_, err := q.QuoteToAPIQuote(quote)
		if err != nil {
			return fmt.Errorf("could not convert quote to API quote: %w", err)
		}

		// TODO: make an interface for the API, would be cleaner
		// TODO: /update quote
		// @mikeyf How should we go about this http connection
	}
	return nil
}

// GetValidQuote gets a valid quote, more logic can be added here as quotes get more complex.
func (q *quoterImpl) GetValidQuote(quoteID string, destTokenID string, destVolume *big.Int) (*Quote, error) {
	q.mux.RLock()
	defer q.mux.RUnlock()

	// Check if requested volume is below the destination balance before getting quote
	currentBalance := q.balance.GetBalance(destTokenID)
	if destVolume.Cmp(currentBalance.Amount) > 0 {
		return nil, fmt.Errorf("requested volume is greater than the current balance")
	}

	// Iterate through the quotes and find the first quote that matches the volumes.
	for _, quote := range q.quotes[quoteID] {
		// Check if destination volume is less than requested volume or no MaxVolume is set. More logic would be added here.
		if quote.MaxVolume.Cmp(big.NewInt(0)) == 0 || destVolume.Cmp(quote.MaxVolume) <= 0 {
			return quote, nil
		}
	}
	return nil, fmt.Errorf("no valid quote found")
}

// QuoteToAPIQuote gets current balances and converts a quote to an APIQuote,.
func (q *quoterImpl) QuoteToAPIQuote(quote *Quote) (*APIQuote, error) {
	originTokenID := utils.GenerateTokenID(quote.OriginChainID, quote.OriginToken)
	destTokenID := utils.GenerateTokenID(quote.DestChainID, quote.DestToken)
	originBalance := q.balance.GetBalance(originTokenID)
	destBalance := q.balance.GetBalance(destTokenID)
	originNormBalance, err := originBalance.ToFloat64()
	if err != nil {
		return nil, fmt.Errorf("could not convert origin balance to float64: %w", err)
	}
	destNormBalance, err := destBalance.ToFloat64()
	if err != nil {
		return nil, fmt.Errorf("could not convert dest balance to float64: %w", err)
	}
	normBalance := destNormBalance / originNormBalance
	return &APIQuote{
		Relayer: q.relayer.String(),

		OriginChainID:    uint(quote.OriginChainID),
		OriginToken:      quote.OriginToken.String(),
		OriginDecimals:   originBalance.Decimals,
		OriginAmount:     originBalance.Amount.String(),
		OriginAmountNorm: originNormBalance,

		DestChainID:    uint(quote.DestChainID),
		DestToken:      quote.DestToken.String(),
		DestAmount:     destBalance.Amount.String(),
		DestAmountNorm: destNormBalance,
		DestDecimals:   destBalance.Decimals,

		Price: normBalance,
	}, nil
}

func (q *quoterImpl) HandleUnconfirmedBridgeRequest(req *bindings.FastBridgeBridgeRequested) error {
	bridgeReq, err := utils.Decode(req.Request)
	if err != nil {
		return fmt.Errorf("could not decode bridge request: %w", err)
	}

	// Modify token balance on destination for this asset
	destTokenID := utils.GenerateTokenID(bridgeReq.DestChainId, bridgeReq.DestToken)
	q.balance.DecrementBalance(destTokenID, bridgeReq.DestAmount)

	// Update Quote in quote API
	quoteID := utils.GenerateQuoteID(bridgeReq.OriginChainId, bridgeReq.OriginToken, bridgeReq.DestChainId, bridgeReq.DestToken)
	err = q.UpdateQuotes(quoteID)
	if err != nil {
		return fmt.Errorf("could not update quote on quoter API %w", err)
	}

	// Add to unconfirmed bridge request queue (for rolling back bad balance changes)
	transactionID := common.Bytes2Hex(req.TransactionId[:]) // keccak256 hash of the request
	q.bridgeReqHandler.Put(transactionID, bridgeReq.DestAmount)

	return nil
}

// HandleCompletedBridge removes a successful bridge from the bridge request handler and updates balance if it had not existed..
func (q *quoterImpl) HandleCompletedBridge(transactionID string, event *bindings.IFastBridgeBridgeTransaction) error {
	// Remove from unconfirmed bridge request queue
	_, existed := q.bridgeReqHandler.GetAndDelete(transactionID)

	// if not in the request cache, need to decrement balance. This case will rarely be hit
	if !existed {
		// Modify token balance on destination for this asset and republish quote.
		logger.Warnf("Bridge request not found in cache, decrementing balance and republishing quote, %s, %d, %d", transactionID, event.DestChainId, event.OriginChainId)
		destTokenID := utils.GenerateTokenID(event.DestChainId, event.DestToken)
		q.balance.DecrementBalance(destTokenID, event.DestAmount)
		quoteID := utils.GenerateQuoteID(event.OriginChainId, event.OriginToken, event.DestChainId, event.DestToken)
		err := q.UpdateQuotes(quoteID)
		if err != nil {
			return fmt.Errorf("could not update quote on quoter API %w", err)
		}
	}
	return nil
}

// HandleUncompletedBridge removes a failed bridge from the bridge request handler and increments balance if it had existed.
func (q *quoterImpl) HandleUncompletedBridge(transactionID string, event *bindings.IFastBridgeBridgeTransaction) error {
	// Remove from unconfirmed bridge request queue
	_, existed := q.bridgeReqHandler.GetAndDelete(transactionID)

	// if not in the request cache, need to increment/rollback balance
	if existed {
		// Modify token balance on destination for this asset and republish quote.
		logger.Warnf("Bridge request found in cache after failing to relay, incrementing balance and republishing quote, %s, %d, %d", transactionID, event.DestChainId, event.OriginChainId)
		destTokenID := utils.GenerateTokenID(event.DestChainId, event.DestToken)
		q.balance.IncrementBalance(destTokenID, event.DestAmount)
		quoteID := utils.GenerateQuoteID(event.OriginChainId, event.OriginToken, event.DestChainId, event.DestToken)
		err := q.UpdateQuotes(quoteID)
		if err != nil {
			return fmt.Errorf("could not update quote on quoter API %w", err)
		}
	}
	return nil
}
