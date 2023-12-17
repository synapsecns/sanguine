package base_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb/base"
)

func TestRoundtripBetweenFromQuoteRequestAndToQuoteRequest(t *testing.T) {
	// Step 1: Setup
	originalRequest := reldb.QuoteRequest{
		OriginTokenDecimals: 18,
		DestTokenDecimals:   6,
		BlockNumber:         9,
		TransactionId:       [32]byte{},
		RawRequest:          []byte(gofakeit.Paragraph(1, 2, 3, " ")),
		Sender:              common.Address{},
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId: 1,
			DestChainId:   2,
			OriginSender:  common.Address{},
			DestRecipient: common.Address{},
			OriginToken:   common.Address{},
			DestToken:     common.Address{},
			OriginAmount:  big.NewInt(1000),
			DestAmount:    big.NewInt(2000),
			Deadline:      big.NewInt(time.Now().Unix()),
			Nonce:         big.NewInt(1),
		},
		Status: reldb.QuoteRequestStatus(1),
	}

	// Step 2: Test FromQuoteRequest
	requestForQuote := base.FromQuoteRequest(originalRequest)

	// Step 3: Test ToQuoteRequest
	roundTrippedRequest, err := requestForQuote.ToQuoteRequest()
	if err != nil {
		t.Errorf("ToQuoteRequest returned an error: %v", err)
	}

	// Step 4: Assertions
	// Compare all the fields of originalRequest and roundTrippedRequest
	if !reflect.DeepEqual(originalRequest, *roundTrippedRequest) {
		t.Errorf("Round tripped request did not match original request. Original: %+v, RoundTripped: %+v", originalRequest, *roundTrippedRequest)
	}
}
