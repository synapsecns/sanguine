package consumer_test

import (
	"context"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer"
	"os"
	"path/filepath"
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
	pwd, _ := os.Getwd()
	path := pwd + filepath.Clean("/tokenIDToCoinGeckoID.yaml")
	parsedYaml, err := consumer.OpenYaml(path)
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

func TestGetTokenMetadataFailure(t *testing.T) {
	timestamp := 1664980469
	tokenID := "this is not a token"
	ctx := context.TODO()
	price, symbol := consumer.GetTokenMetadataWithTokenID(ctx, timestamp, &tokenID)
	Nil(t, price)
	Nil(t, symbol)
}
