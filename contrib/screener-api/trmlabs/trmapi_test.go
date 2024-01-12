package trmlabs_test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/contrib/screener-api/screener"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("skipping test in CI environment")
	}

	apiKey := os.Getenv("TRM_API_KEY")

	ctx := context.Background()

	client, err := screener.NewSimpleScreener(apiKey, "/Users/jake/Downloads/risk_engine_2024-01-12_001124.csv")
	if err != nil {
		t.Fatal(err)
	}

	results, err := client.ScreenAddress(ctx, common.HexToAddress("0x0000000000000000000000000000000000000000"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Print(results)
}
