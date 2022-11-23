package gas_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/chain/gas"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"math/big"
)

func (s GasSuite) TestGasPriceEstimator() {
	testChain := simulated.NewSimulatedBackend(s.GetTestContext(), s.T())
	estimator := gas.NewGasPriceEstimator(s.GetTestContext(), testChain)
	height, err := testChain.BlockByNumber(s.GetTestContext(), nil)
	Nil(s.T(), err)

	// below ethconfig.FullNodeGPO.Blocks, price should be default
	price, err := estimator.EstimateGasPrice(s.GetTestContext(), height.NumberU64(), gas.GetConfig())
	Nil(s.T(), err)
	Equal(s.T(), gas.DefaultGasPriceEstimate, price)
	s.mockGasEvents(testChain)

	// price at new height should not be the same
	height, err = testChain.BlockByNumber(s.GetTestContext(), nil)
	Nil(s.T(), err)

	price, err = estimator.EstimateGasPrice(s.GetTestContext(), height.NumberU64(), gas.GetConfig())
	Nil(s.T(), err)
	NotEqual(s.T(), gas.DefaultGasPriceEstimate.Uint64(), price.Uint64())
	Less(s.T(), gas.GetConfig().IgnorePrice.Uint64(), price.Uint64())

	// finally, let's make sure the new gas price isn't ignored by our backend (base fee test)
	acct := testChain.GetFundedAccount(s.GetTestContext(), big.NewInt(params.Ether))
	to := common.BigToAddress(big.NewInt(0))

	signedTx, err := testChain.SignTx(types.NewTx(&types.LegacyTx{
		Gas:      21000,
		GasPrice: price,
		To:       &to,
		Value:    big.NewInt(params.GWei),
	}), types.LatestSignerForChainID(testChain.Chain.GetBigChainID()), acct.PrivateKey)
	Nil(s.T(), err)

	err = testChain.SendTransaction(s.GetTestContext(), signedTx)
	Nil(s.T(), err)

	testChain.WaitForConfirmation(s.GetTestContext(), signedTx)
}

// mockGasEvents mocks a bunch of gas events.
func (s GasSuite) mockGasEvents(chain *simulated.Backend) {
	// create a new transactor
	transactor := mocks.MockAccount(s.T())

	// fund the transactor w/ 50 eth
	chain.FundAccount(s.GetTestContext(), transactor.Address, *big.NewInt(0).Mul(big.NewInt(params.Ether), big.NewInt(50)))

	// get the current height, so we can discard any txes before the height where gas calculations start
	currentBlock, err := chain.BlockByNumber(s.GetTestContext(), nil)
	Nil(s.T(), err)
	height := int(currentBlock.NumberU64())
	// gasHeight is the height at which gas will be calculated back from
	gasHeight := gas.GetConfig().Blocks + height + 20

	var signedTx *types.Transaction

	destAddress := mocks.MockAddress()
	for i := height; i <= gasHeight; i++ {
		// generate a number between 1 and 256 and multiply it by gwei
		gasPrice := big.NewInt(0).Mul(big.NewInt(params.GWei), big.NewInt(int64(gofakeit.Uint8())+1))
		rawTx := types.NewTx(&types.LegacyTx{
			Gas:      simulated.BlockGasLimit,
			GasPrice: gasPrice,
			To:       &destAddress,
			Value:    big.NewInt(1),
		})
		signedTx, err = chain.SignTx(rawTx, types.LatestSignerForChainID(chain.ChainConfig().ChainID), transactor.PrivateKey)
		Nil(s.T(), err)

		err = chain.Client().SendTransaction(s.GetTestContext(), signedTx)
		Nil(s.T(), err)
		chain.Commit()
	}
	chain.WaitForConfirmation(s.GetTestContext(), signedTx)
}
