package trmlabs_test

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/contrib/screener-api/screener"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("skipping test in CI environment")
	}

	apiKey := os.Getenv("TRM_API_KEY")
	apiKey = "0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9"

	ctx := context.Background()

	test := gofakeit.Name()

	client, err := screener.NewSimpleScreener(apiKey, fmt.Sprintf("/Users/%s/Downloads/risk_engine_2024-01-12_001124.csv", test))
	if err != nil {
		t.Fatal(err)
	}

	fat := mocks.MockAddress()
	results, err := client.ScreenAddress(ctx, common.HexToAddress(fat.String()))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Print(results)
}
