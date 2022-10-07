package consumer_test

import (
	"context"
	"fmt"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer"
	"testing"
)

func TestGetDefiLlamaData(t *testing.T) {
	knownPrice := float64(19830.33068615768)
	knownSymbol := "BTC"
	timestamp := 1664980469
	coinGeckoID := "bitcoin"
	ctx := context.TODO()
	price, symbol := consumer.GetDefiLlamaData(ctx, timestamp, &coinGeckoID)
	NotNil(t, price)
	NotNil(t, symbol)
	Equal(t, *price, knownPrice)
	Equal(t, *symbol, knownSymbol)
}
func TestOpenYaml(t *testing.T) {
	parsedYaml, err := consumer.OpenYaml()
	fmt.Println(parsedYaml)
	Nil(t, err)
	NotNil(t, parsedYaml)
}

func TestGetTokenMetadataWithTokenID(t *testing.T) {
	timestamp := 1664980469
	tokenID := "synFRAX"
	ctx := context.TODO()
	price, symbol := consumer.GetTokenMetadataWithTokenID(ctx, timestamp, &tokenID)
	NotNil(t, price)
	NotNil(t, symbol)
}
