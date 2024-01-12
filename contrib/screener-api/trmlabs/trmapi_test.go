package trmlabs_test

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("skipping test in CI environment")
	}

	apiKey := os.Getenv("TRM_API_KEY")

	url := "https://api.trmlabs.com/"

	ctx := context.Background()

	client, err := trmlabs.NewClient(apiKey, url)
	if err != nil {
		t.Fatal(err)
	}

	results, err := client.ScreenAddress(ctx, "0x0000000000000000000000000000000000000000")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Print(results)
}
