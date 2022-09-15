package swap_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/bridgeconfig"
)

func (c *ConfigSuite) TestParser() {
	_, bridgeConfigContract := c.deployManager.GetBridgeConfigV3(c.GetTestContext(), c.testBackend)

	fetcher, err := bridgeconfig.NewFetcher(bridgeConfigContract.Address(), c.testBackend)
	Nil(c.T(), err)

	curentBlockNumber, err := c.testBackend.BlockNumber(c.GetTestContext())
	Nil(c.T(), err)

	for _, testToken := range testTokens {
		tokenID, err := fetcher.GetTokenID(c.GetTestContext(), uint32(testToken.ChainId.Uint64()), uint32(curentBlockNumber), common.HexToAddress(testToken.TokenAddress))

		Nil(c.T(), err)
		Equal(c.T(), tokenID, testToken.tokenID)
		token, err := fetcher.GetToken(c.GetTestContext(), uint32(testToken.ChainId.Uint64()), uint32(curentBlockNumber), tokenID)
		Nil(c.T(), err)
		Equal(c.T(), common.HexToAddress(testToken.TokenAddress).String(), common.HexToAddress(token.TokenAddress).String())
		Equal(c.T(), testToken.SwapFee, token.SwapFee)
		Equal(c.T(), testToken.IsUnderlying, token.IsUnderlying)
	}
}
