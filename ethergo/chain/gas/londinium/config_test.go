package londinium_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/chain/gas/londinium"
	"math/big"
)

func (l LondoniumSuite) TestLondoniumConfig() {
	defaultPrice := big.NewInt(gofakeit.Int64())

	newConfig := gasprice.Config{
		Blocks:           int(gofakeit.Int64()),
		Percentile:       gofakeit.Number(1, 100),
		MaxHeaderHistory: int(gofakeit.Int64()),
		MaxBlockHistory:  int(gofakeit.Int64()),
		Default:          defaultPrice,
		MaxPrice:         big.NewInt(0).Mul(defaultPrice, big.NewInt(2)),
		IgnorePrice:      big.NewInt(0).Div(defaultPrice, big.NewInt(2)),
	}

	londiniumConfig := londinium.ToLondiniumConfig(newConfig)

	Equal(l.T(), newConfig.Blocks, londiniumConfig.Blocks)
	Equal(l.T(), newConfig.Percentile, londiniumConfig.Percentile)
	Equal(l.T(), newConfig.Default, londiniumConfig.Default)
	Equal(l.T(), newConfig.MaxPrice, londiniumConfig.MaxPrice)
}
