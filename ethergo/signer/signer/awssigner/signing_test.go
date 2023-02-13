package awssigner_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/awssigner"
	"math/big"
)

// TestSigning tests the signing process.
func (k *KMSSuite) TestSigning() {
	testSigner := awssigner.NewSignerFromMockKMS(k.GetTestContext(), k.T())

	testBackend := simulated.NewSimulatedBackend(k.GetTestContext(), k.T())

	testBackend.FundAccount(k.GetTestContext(), testSigner.Address(), *big.NewInt(params.Ether))

	transactor, err := testSigner.GetTransactor(k.GetTestContext(), testBackend.GetBigChainID())
	Nil(k.T(), err)

	gasPrice, err := testBackend.SuggestGasPrice(k.GetTestContext())
	Nil(k.T(), err)

	signedTx, err := transactor.Signer(transactor.From, types.NewTx(&types.LegacyTx{
		To:       &common.Address{},
		Value:    big.NewInt(params.GWei),
		GasPrice: gasPrice,
		Gas:      21000,
	}))
	Nil(k.T(), err)

	err = testBackend.SendTransaction(k.GetTestContext(), signedTx)
	Nil(k.T(), err)

	testBackend.WaitForConfirmation(k.GetTestContext(), signedTx)
}
