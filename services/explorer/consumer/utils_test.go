package consumer

import (
	"fmt"
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDefiLlamaData(t *testing.T) {
	knownPrice := float64(19830.33068615768)
	knownSymbol := "BTC"
	timestamp := 1664980469
	coinGeckoID := "bitcoin"
	price, symbol := GetDefiLlamaData(timestamp, &coinGeckoID)
	NotNil(t, price)
	NotNil(t, symbol)
	Equal(t, *price, knownPrice)
	Equal(t, *symbol, knownSymbol)
}
func TestOpenYaml(t *testing.T) {
	parsedYaml, err := OpenYaml("tokenIDToCoinGeckoID.yaml")
	fmt.Println(parsedYaml)
	Nil(t, err)
	NotNil(t, parsedYaml)
}

func TestGetTokenMetadataWithTokenID(t *testing.T) {
	timestamp := 1664980469
	tokenID := "synFRAX"
	price, symbol := GetTokenMetadataWithTokenID(timestamp, &tokenID, "tokenIDtoCoinGeckoID.yaml")
	NotNil(t, price)
	NotNil(t, symbol)
}
