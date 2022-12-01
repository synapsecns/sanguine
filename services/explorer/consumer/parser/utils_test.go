package parser_test

import (
	"fmt"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	"math/big"
	"path/filepath"
	"testing"
)

func TestOpenYaml(t *testing.T) {
	path, err := filepath.Abs("../../static/tokenIDToCoinGeckoID.yaml")
	Nil(t, err)
	parsedYaml, err := parser.OpenYaml(path)
	Nil(t, err)
	NotNil(t, parsedYaml)
}

func TestGetAmountUSD(t *testing.T) {
	price := 0.44
	amount := parser.GetAmountUSD(big.NewInt(111100011), 2, &price)
	fmt.Println(amount)
	NotNil(t, amount)
}
