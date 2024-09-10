package botmd_test

import (
	"context"
	"testing"

	"github.com/synapsecns/sanguine/contrib/opbot/botmd"
	"github.com/synapsecns/sanguine/core/metrics"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
)

func TestStripLinks(t *testing.T) {
	testLink := "<https://example.com|example>"
	expected := "example"

	if got := botmd.StripLinks(testLink); got != expected {
		t.Errorf("StripLinks(%s) = %s; want %s", testLink, got, expected)
	}
}

func TestTxAge(t *testing.T) {
	notExpected := "unknown time ago" // should be a definite time

	status := &relapi.GetQuoteRequestResponse{
		OriginTxHash:  "0x954264d120f5f3cf50edc39ebaf88ea9dc647d9d6843b7a120ed3677e23d7890",
		OriginChainID: 421611,
	}

	ctx := context.Background()

	client := omnirpcClient.NewOmnirpcClient("https://arb1.arbitrum.io/rpc", metrics.Get())
	cc, err := client.GetChainClient(ctx, int(status.OriginChainID))
	if err != nil {
		t.Fatalf("GetChainClient() failed: %v", err)
	}

	if got := botmd.GetTxAge(context.Background(), cc, status); got == notExpected {
		t.Errorf("TxAge(%s) = %s; want not %s", status.OriginTxHash, got, notExpected)
	}
}
