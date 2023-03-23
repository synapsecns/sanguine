package testhelper_test

import (
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"math/big"
)

func (s *TestHelperSuite) TestOminrpcServer() {
	omnirpcServer := testhelper.NewOmnirpcServer(s.GetTestContext(), s.T(), s.testBackends...)

	for _, testBackend := range s.testBackends {
		ogBlock, err := testBackend.BlockByNumber(s.GetTestContext(), big.NewInt(1))
		Nil(s.T(), err)

		omnirpcURL := testhelper.GetURL(omnirpcServer, testBackend)
		omnirpcEthClient, err := ethclient.DialContext(s.GetTestContext(), omnirpcURL)
		Nil(s.T(), err)

		omniBlock, err := omnirpcEthClient.BlockByNumber(s.GetTestContext(), big.NewInt(1))
		Nil(s.T(), err)

		Equal(s.T(), ogBlock.Hash(), omniBlock.Hash())
	}
}
