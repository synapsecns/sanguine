package signer_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"math/big"
	"testing"
)

func (s *SignerSuite) TestSignAndSubmit() {
	s.RunOnAllSigners(func(testSigner signer.Signer) {
		testBackend := simulated.NewSimulatedBackend(s.GetTestContext(), s.T())

		testBackend.FundAccount(s.GetTestContext(), testSigner.Address(), *big.NewInt(params.Ether))

		transactor, err := testSigner.GetTransactor(s.GetTestContext(), testBackend.GetBigChainID())
		s.NoError(err)

		gasPrice, err := testBackend.SuggestGasPrice(s.GetTestContext())
		s.NoError(err)

		signedTx, err := transactor.Signer(transactor.From, types.NewTx(&types.LegacyTx{
			To:       &common.Address{},
			Value:    big.NewInt(params.GWei),
			GasPrice: gasPrice,
			Gas:      21000,
		}))
		assert.Nil(s.T(), err)

		err = testBackend.SendTransaction(s.GetTestContext(), signedTx)
		assert.Nil(s.T(), err)

		testBackend.WaitForConfirmation(s.GetTestContext(), signedTx)
	})
}

func TestEncodeSigner(t *testing.T) {
	v := new(big.Int).SetUint64(gofakeit.Uint64())
	r := new(big.Int).SetUint64(gofakeit.Uint64())
	s := new(big.Int).SetUint64(gofakeit.Uint64())

	sig := signer.NewSignature(v, r, s)

	rawSig := signer.Encode(sig)

	// test decoding
	decoded := signer.DecodeSignature(rawSig)
	assert.True(t, signer.IsEqual(sig, decoded))

	exampleSigner := ethTypes.FrontierSigner{}

	newR, newS, _, err := exampleSigner.SignatureValues(ethTypes.NewTx(&ethTypes.LegacyTx{
		V: v,
		R: r,
		S: s,
	}), rawSig)

	assert.Nil(t, err)

	// skip v, this gets modified by eth
	assert.Equal(t, s, newS)
	assert.Equal(t, r, newR)
}
