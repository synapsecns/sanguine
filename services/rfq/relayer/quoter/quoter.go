// Package quoter submits quotes to the RFQ API for which assets the relayer is willing to relay.
package quoter

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/relayer/pricer"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"golang.org/x/exp/slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	rfqAPIClient "github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
)

var logger = log.Logger("quoter")

// Quoter submits quotes to the RFQ API.
type Quoter interface {
	// SubmitQuote submits a quote to the RFQ API.
	SubmitQuote() error
	// SubmitAllQuotes submits all quotes to the RFQ API.
	SubmitAllQuotes() error
	// GetSelfQuotes gets relayer's live quote from the RFQ API.
	GetSelfQuotes() ([]*db.Quote, error)
	// ShouldProcess determines if a quote should be processed.
	// We do this by either saving all quotes in-memory, and refreshing via GetSelfQuotes() through the API
	// The first comparison is does bridge transaction OriginChainID+TokenAddr match with a quote + DestChainID+DestTokenAddr, then we look to see if we have enough amount to relay it + if the price fits our bounds (based on that the Relayer is relaying the destination token for the origin)
	// validateQuote(BridgeEvent)
	ShouldProcess(quote reldb.QuoteRequest) bool
}

// Manager submits quotes to the RFQ API.
type Manager struct {
	// inventoryManager is used to get the relayer's inventory.
	inventoryManager inventory.Manager
	// rfqClient is used to communicate with the RFQ API.
	rfqClient rfqAPIClient.AuthenticatedClient
	// relayerSigner is the signer used by the relayer to interact on chain
	relayerSigner signer.Signer
	// quotableTokens are tokens that the relayer is willing to relay to & from
	quotableTokens map[string][]string
	// feePricer is used to price fees.
	feePricer pricer.FeePricer
}

// NewQuoterManager creates a new QuoterManager.
func NewQuoterManager(metricsHandler metrics.Handler, quotableTokens map[string][]string, inventoryManager inventory.Manager, rfqAPIUrl string, relayerSigner signer.Signer, feePricer pricer.FeePricer) (*Manager, error) {
	apiClient, err := rfqAPIClient.NewAuthenticatedClient(metricsHandler, rfqAPIUrl, relayerSigner)
	if err != nil {
		return nil, fmt.Errorf("error creating RFQ API client: %w", err)
	}

	return &Manager{
		quotableTokens:   quotableTokens,
		inventoryManager: inventoryManager,
		rfqClient:        apiClient,
		relayerSigner:    relayerSigner,
		feePricer:        feePricer,
	}, nil
}

// SetQuotableTokens sets the quotable tokens.
func (m *Manager) SetQuotableTokens(tokens map[string][]string) {
	m.quotableTokens = tokens
}

// ShouldProcess determines if a quote should be processed.
func (m *Manager) ShouldProcess(quote reldb.QuoteRequest) bool {
	// allowed pairs for this origin token on the destination
	destPairs := m.quotableTokens[quote.GetOriginIDPair()]
	if slices.Contains(destPairs, quote.GetDestIDPair()) {
		return true
	}

	// handle decimals.
	// this will never get hit if we're operating correctly.
	if quote.OriginTokenDecimals != quote.DestTokenDecimals {
		logger.Errorf("Pairing tokens with two different decimals is disabled as a safety feature right now.")
		return false
	}

	// then check if we'll make money on it
	// note: this check is not comprehensive. We still need to check gas, min fees, etc.
	// It does keep us in range though.
	if quote.Transaction.OriginAmount.Cmp(quote.Transaction.DestAmount) < 0 {
		// we're not making money on this
		return false
	}

	return false
}

// SubmitAllQuotes submits all quotes to the RFQ API.
func (m *Manager) SubmitAllQuotes(ctx context.Context) error {
	inv, err := m.inventoryManager.GetCommitableBalances(context.Background())
	if err != nil {
		return fmt.Errorf("error getting commitable balances: %w", err)
	}
	return m.prepareAndSubmitQuotes(ctx, inv)
}

// Prepares and submits quotes based on inventory.
func (m *Manager) prepareAndSubmitQuotes(ctx context.Context, inv map[int]map[common.Address]*big.Int) error {
	var allQuotes []rfqAPIClient.APIQuotePutRequest

	// First, generate all quotes
	for chainID, balances := range inv {
		for address, balance := range balances {
			quotes, err := m.GenerateQuotes(ctx, chainID, address, balance)
			if err != nil {
				return err
			}
			allQuotes = append(allQuotes, quotes...)
		}
	}

	// Now, submit all the generated quotes
	for _, quote := range allQuotes {
		if err := m.submitQuote(quote); err != nil {
			return err
		}
	}

	return nil
}

// GenerateQuotes TODO: THIS LOOP IS BROKEN
// Essentially, if we know a destination chain token balance, then we just need to find which tokens are bridgeable to it.
// We can do this by looking at the quotableTokens map, and finding the key that matches the destination chain token.
// Generates quotes for a given chain ID, address, and balance.
func (m *Manager) GenerateQuotes(ctx context.Context, chainID int, address common.Address, balance *big.Int) ([]rfqAPIClient.APIQuotePutRequest, error) {
	destTokenID := fmt.Sprintf("%d-%s", chainID, address.Hex())
	var quotes []rfqAPIClient.APIQuotePutRequest
	for keyTokenID, itemTokenIDs := range m.quotableTokens {
		for _, tokenID := range itemTokenIDs {
			if tokenID == destTokenID {
				originStr := strings.Split(keyTokenID, "-")[0]
				origin, err := strconv.Atoi(originStr)
				if err != nil {
					return nil, fmt.Errorf("error converting origin chainID: %w", err)
				}
				fee, err := m.feePricer.GetTotalFee(ctx, uint32(origin), uint32(chainID), destTokenID)
				if err != nil {
					return nil, fmt.Errorf("error getting total fee: %w", err)
				}
				quote := rfqAPIClient.APIQuotePutRequest{
					OriginChainID:   originStr,
					OriginTokenAddr: strings.Split(keyTokenID, "-")[1],
					DestChainID:     fmt.Sprint(chainID),
					DestTokenAddr:   address.Hex(),
					DestAmount:      balance.String(),
					MaxOriginAmount: balance.String(),
					FixedFee:        fee.String(),
				}
				quotes = append(quotes, quote)
			}
		}
	}
	return quotes, nil
}

// Submits a single quote.
func (m *Manager) submitQuote(quote rfqAPIClient.APIQuotePutRequest) error {
	err := m.rfqClient.PutQuote(&quote)
	if err != nil {
		return fmt.Errorf("error submitting quote: %w", err)
	}
	return nil
}
