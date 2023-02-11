package gcpsigner_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/gcpsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/gcpsigner/gcpmock"
)

func (g *GCPSignerSuite) TestSign() {
	mc, err := gcpmock.NewMockClient()
	Nil(g.T(), err, "should make mock client")

	mk, err := gcpsigner.NewManagedKey(g.GetTestContext(), mc, "test")
	Nil(g.T(), err, "should create managed key")

	Equal(g.T(), mk.Address(), mc.Address())

	testTransactor, err := mk.GetTransactor(g.GetTestContext(), params.MainnetChainConfig.ChainID)
	Nil(g.T(), err, "should get transactor")

	testTx := types.NewTransaction(0, common.Address{}, testTransactor.Value, testTransactor.GasLimit, testTransactor.GasPrice, []byte{})
	signedTx, err := testTransactor.Signer(testTransactor.From, testTx)
	Nil(g.T(), err, "should sign signature")
	NotNil(g.T(), signedTx)

	signedTx.RawSignatureValues()

	recoveredSender, err := types.Sender(types.LatestSignerForChainID(params.MainnetChainConfig.ChainID), signedTx)
	Nil(g.T(), err, "should recover signature")

	Equal(g.T(), recoveredSender, mk.Address())
}
