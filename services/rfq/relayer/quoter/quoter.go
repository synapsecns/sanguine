// Package quoter submits quotes to the RFQ API for which assets the relayer is willing to relay.
package quoter

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	rfqAPIClient "github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
)

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
}

// NewQuoterManager creates a new QuoterManager.
func NewQuoterManager(metricsHandler metrics.Handler, quotableTokens map[string][]string, inventoryManager inventory.Manager, rfqAPIUrl string, relayerSigner signer.Signer) (*Manager, error) {
	rfqAPIClient, err := rfqAPIClient.NewAuthenticatedClient(metricsHandler, rfqAPIUrl, relayerSigner)
	if err != nil {
		return nil, fmt.Errorf("error creating RFQ API client: %w", err)
	}

	return &Manager{
		quotableTokens:   quotableTokens,
		inventoryManager: inventoryManager,
		rfqClient:        rfqAPIClient,
		relayerSigner:    relayerSigner,
	}, nil
}

// SetQuotableTokens sets the quotable tokens.
func (m *Manager) SetQuotableTokens(tokens map[string][]string) {
	m.quotableTokens = tokens
}

// ShouldProcess determines if a quote should be processed.
func (m *Manager) ShouldProcess(quote reldb.QuoteRequest) bool {
	// first check if token is valid
	// then check if we'll make money on it

	return true
}

// SubmitAllQuotes submits all quotes to the RFQ API.
func (m *Manager) SubmitAllQuotes() error {
	inv, err := m.inventoryManager.GetCommitableBalances(context.Background())
	if err != nil {
		return fmt.Errorf("error getting commitable balances: %w", err)
	}
	return m.prepareAndSubmitQuotes(inv)
}

// Prepares and submits quotes based on inventory.
func (m *Manager) prepareAndSubmitQuotes(inv map[int]map[common.Address]*big.Int) error {
	var allQuotes []rfqAPIClient.APIQuotePutRequest

	// First, generate all quotes
	for chainID, balances := range inv {
		for address, balance := range balances {
			quotes, err := m.GenerateQuotes(chainID, address, balance)
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
func (m *Manager) GenerateQuotes(chainID int, address common.Address, balance *big.Int) ([]rfqAPIClient.APIQuotePutRequest, error) {
	destTokenID := fmt.Sprintf("%d-%s", chainID, address.Hex())
	var quotes []rfqAPIClient.APIQuotePutRequest
	for keyTokenID, itemTokenIDs := range m.quotableTokens {
		for _, tokenID := range itemTokenIDs {
			if tokenID == destTokenID {
				quote := rfqAPIClient.APIQuotePutRequest{
					OriginChainID:   strings.Split(keyTokenID, "-")[0],
					OriginTokenAddr: strings.Split(keyTokenID, "-")[1],
					DestChainID:     fmt.Sprint(chainID),
					DestTokenAddr:   address.Hex(),
					DestAmount:      balance.String(),
					Price:           "1",
					MaxOriginAmount: balance.String(),
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
