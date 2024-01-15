package client_test

import (
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"gotest.tools/assert"
	"math"
	"math/big"
	"time"
)

// TestConfigFromID makes sure configs are properly returned by id.
func (c ClientSuite) TestConfigFromID() {
	// make sure all current configs are properly returned
	for _, chainConfig := range client.ChainConfigs {
		res := client.ConfigFromID(chainConfig.ChainID)
		assert.DeepEqual(c.T(), res, chainConfig, testsuite.BigIntComparer())
	}

	// make sure we don't panic on nil
	NotPanics(c.T(), func() {
		client.ConfigFromID(nil)
	})
}

// TestConfig tests the config.
func (c ClientSuite) TestConfig() {
	testConfig := client.Config{}

	mockAddress := mocks.MockAddress()
	testConfig.SetEthBridgeAddress(mockAddress)
	Equal(c.T(), testConfig.GetEthBridgeAddress(), mockAddress)
}

// TestChainSigner is a sanity check that our configs correctly use the signer.
func (c ClientSuite) TestChainSigner() {
	for _, config := range client.ChainConfigs {
		signer := types.MakeSigner(config, big.NewInt(0).SetUint64(math.MaxUint64), uint64(time.Now().Unix()))
		False(c.T(), signer.Equal(types.FrontierSigner{}))
		False(c.T(), signer.Equal(types.HomesteadSigner{}))
	}
}
