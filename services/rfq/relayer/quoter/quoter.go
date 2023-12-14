// Quoter submits quotes to the RFQ API for which assets the relayer is willing to relay.
package quoter

import (
	"context"

	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	rfqAPIClient "github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
)

type Quoter interface {
	// SubmitQuote submits a quote to the RFQ API.
	SubmitQuote() error
	// GetSelfQuotes gets relayer's live quote from the RFQ API.
	GetSelfQuotes() ([]*db.Quote, error)
}

type quoterManagerImpl struct {
	// inventoryManager is used to get the relayer's inventory.
	inventoryManager *inventory.InventoryManager
	// rfqClient is used to communicate with the RFQ API.
	rfqClient rfqAPIClient.Client
	// relayerAddress is the relayers address they interact on chain with
	relayerSigner signer.Signer
	// rfqAPIURL is the URL of the RFQ API.
	rfqAPIURL string
}

func NewQuoterManager(ctx context.Context, inventoryManager *inventory.InventoryManager, rfqAPIUrl string, relayerSigner signer.Signer) (interface{}, error) {
	rfqAPIClient, err := rfqAPIClient.NewClient(rfqAPIUrl, relayerSigner)
	if err != nil {
		return nil, err
	}

	return &quoterManagerImpl{
		inventoryManager: inventoryManager,
		rfqClient:        rfqAPIClient,
		relayerSigner:    relayerSigner,
	}, nil

}
