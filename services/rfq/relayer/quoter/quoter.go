// Quoter submits quotes to the RFQ API for which assets the relayer is willing to relay.
package quoter

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	rfqAPIClient "github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
)

type Quoter interface {
	// SubmitQuote submits a quote to the RFQ API.
	SubmitQuote() error
	// Submit All Quotes submits all quotes to the RFQ API.
	SubmitAllQuotes() error
	// GetSelfQuotes gets relayer's live quote from the RFQ API.
	GetSelfQuotes() ([]*db.Quote, error)
	// TOOO: Build this. ValidateQuote takes in a bridge event, and checks if the Relayer has a live quote, or should have a live quote out for this bridge event.
	// We do this by either saving all quotes in-memory, and refreshing via GetSelfQuotes() through the API
	// The first comparison is does bridge transaction OriginChainID+TokenAddr match with a quote + DestChainId+DestTokenAddr, then we look to see if we have enough amount to relay it + if the price fits our bounds (based on that the Relayer is relaying the destination token for the origin)
	// validateQuote(BridgeEvent)
}

type QuoterManager struct {
	// inventoryManager is used to get the relayer's inventory.
	inventoryManager inventory.InventoryManager
	// rfqClient is used to communicate with the RFQ API.
	rfqClient rfqAPIClient.Client
	// relayerSigner is the signer used by the relayer to interact on chain
	relayerSigner signer.Signer
	// quotableTokens are tokens that the relayer is willing to relay to & from
	quotableTokens map[string][]string
}

func (q *QuoterManager) SetQuotableTokens(tokens map[string][]string) {
	q.quotableTokens = tokens
}

func NewQuoterManager(ctx context.Context, quotableTokens map[string][]string, inventoryManager inventory.InventoryManager, rfqAPIUrl string, relayerSigner signer.Signer) (*QuoterManager, error) {
	rfqAPIClient, err := rfqAPIClient.NewClient(rfqAPIUrl, relayerSigner)
	if err != nil {
		return nil, err
	}

	return &QuoterManager{
		quotableTokens:   quotableTokens,
		inventoryManager: inventoryManager,
		rfqClient:        rfqAPIClient,
		relayerSigner:    relayerSigner,
	}, nil
}

// Gets the inventory and prepares quotes.
func (m *QuoterManager) SubmitAllQuotes() error {
	inv, err := m.inventoryManager.GetCommitableBalances(context.Background())
	if err != nil {
		return err
	}
	return m.prepareAndSubmitQuotes(inv)
}

// Prepares and submits quotes based on inventory.
func (m *QuoterManager) prepareAndSubmitQuotes(inv map[int]map[common.Address]*big.Int) error {
	var allQuotes []rfqAPIClient.APIQuotePutRequest

	// First, generate all quotes
	for chainId, balances := range inv {
		for address, balance := range balances {
			quotes, err := m.GenerateQuotes(chainId, address, balance)
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

// TODO: THIS LOOP IS BROKEN
// Essentially, if we know a destination chain token balance, then we just need to find which tokens are bridgeable to it.
// We can do this by looking at the quotableTokens map, and finding the key that matches the destination chain token.
// Generates quotes for a given chain ID, address, and balance.
func (m *QuoterManager) GenerateQuotes(chainId int, address common.Address, balance *big.Int) ([]rfqAPIClient.APIQuotePutRequest, error) {
	destTokenId := fmt.Sprintf("%d-%s", chainId, address.Hex())
	var quotes []rfqAPIClient.APIQuotePutRequest
	for keyTokenID, itemTokenIds := range m.quotableTokens {
		for _, tokenID := range itemTokenIds {
			if tokenID == destTokenId {
				quote := rfqAPIClient.APIQuotePutRequest{
					OriginChainID:   strings.Split(keyTokenID, "-")[0],
					OriginTokenAddr: strings.Split(keyTokenID, "-")[1],
					DestChainID:     fmt.Sprint(chainId),
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
func (m *QuoterManager) submitQuote(quote rfqAPIClient.APIQuotePutRequest) error {
	return m.rfqClient.PutQuote(&quote)
}
