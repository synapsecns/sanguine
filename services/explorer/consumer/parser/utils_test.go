package parser_test

import (
	"math/big"
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
)

func TestGetAmountUSD(t *testing.T) {
	price := 0.44
	amount := parser.GetAmountUSD(big.NewInt(111100011), 2, &price)
	NotNil(t, amount)
}
