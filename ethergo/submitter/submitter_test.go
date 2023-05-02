package submitter_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/client/mocks"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/ethergo/submitter/config"
	"math/big"
)

func (s *SubmitterSuite) TestSetGasPrice() {
	wall, err := wallet.FromRandom()
	s.Require().NoError(err)

	signer := localsigner.NewSigner(wall.PrivateKey())

	chainID := s.testBackends[0].GetBigChainID()

	client := new(mocks.EVM)

	transactor, err := signer.GetTransactor(s.GetTestContext(), chainID)
	s.Require().NoError(err)

	cfg := &config.Config{}
	ts := submitter.NewTestTransactionSubmitter(s.metrics, signer, s, cfg)

	// 1. Test with gas price set, but not one that exceeds max (not eip-1559)
	gasPrice := new(big.Int).SetUint64(gofakeit.Uint64())
	maxPrice := new(big.Int).Add(gasPrice, new(big.Int).SetUint64(1))
	cfg.SetGlobalMaxGasPrice(maxPrice)

	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Return(gasPrice, nil)
	err = ts.SetGasPrice(s.GetTestContext(), client, transactor, chainID, nil)
	s.Require().NoError(err)

	s.Equal(gasPrice, transactor.GasPrice, testsuite.BigIntComparer())

	// 2. Test with gas price set, but one that exceeds max, should return max (not eip-1559)
	maxPrice = new(big.Int).Sub(gasPrice, new(big.Int).SetUint64(1))
	cfg.SetGlobalMaxGasPrice(maxPrice)

	err = ts.SetGasPrice(s.GetTestContext(), client, transactor, chainID, nil)
	s.Require().NoError(err)
	s.Equal(maxPrice, transactor.GasPrice, testsuite.BigIntComparer())

	// 3. Test with gas price set, but one that exceeds max, should return max (legacy tx)
	cfg.SetGlobalEIP1559Support(true)
	tipCap := new(big.Int).SetUint64(uint64(gofakeit.Uint32()))
	client.On(testsuite.GetFunctionName(client.SuggestGasTipCap), mock.Anything).Return(tipCap, nil)

	err = ts.SetGasPrice(s.GetTestContext(), client, transactor, chainID, nil)
	s.Require().NoError(err)

	s.Equal(tipCap, transactor.GasTipCap, testsuite.BigIntComparer())
	s.Equal(maxPrice, transactor.GasFeeCap, testsuite.BigIntComparer())

	// 4. Test with bump (TODO)
	// 5. Test with bump and max (TODO)
}
