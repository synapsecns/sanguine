package gas_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated/multibackend"
	"github.com/synapsecns/sanguine/ethergo/chain/gas"
	"github.com/synapsecns/sanguine/ethergo/chain/gas/backend/mocks"
	"math/big"
)

// TestGasPriceSetterFrom0 ensures gas price calculations are correct from a 0 block.
func (s GasSuite) TestGasPriceSetterFrom0London() {
	testChain := simulated.NewSimulatedBackend(s.GetTestContext(), s.T())
	NotNil(s.T(), testChain.ChainConfig().LondonBlock)

	gasSetter := gas.NewGasSetter(s.GetTestContext(), testChain)
	txContent := testChain.GetTxContext(s.GetTestContext(), nil)

	// txOpts content may auto set these, we don't want these to give false negatives on failing tests
	txContent.TransactOpts.GasFeeCap = nil
	txContent.TransactOpts.GasPrice = nil
	txContent.TransactOpts.GasPrice = nil

	gasBlock, err := testChain.BlockByNumber(s.GetTestContext(), nil)
	Nil(s.T(), err)

	err = gasSetter.SetGasFeeByBlock(s.GetTestContext(), txContent.TransactOpts, gasBlock.NumberU64(), gas.GetConfig().MaxPrice)
	Nil(s.T(), err)

	True(s.T(), txContent.TransactOpts.GasFeeCap.Cmp(txContent.TransactOpts.GasTipCap) > 0)
	Nil(s.T(), txContent.TransactOpts.GasPrice)
}

func (s GasSuite) TestGasPriceSetterFrom0PreLondon() {
	config := multibackend.NewConfigWithChainID(params.AllEthashProtocolChanges.ChainID)
	config.LondonBlock = nil
	config.ArrowGlacierBlock = nil
	config.GrayGlacierBlock = nil
	config.MergeNetsplitBlock = nil
	config.ShanghaiTime = nil
	config.CancunTime = nil

	testChain := simulated.NewSimulatedBackendWithConfig(s.GetTestContext(), s.T(), config)
	Nil(s.T(), testChain.ChainConfig().LondonBlock)

	gasSetter := gas.NewGasSetter(s.GetTestContext(), testChain)
	txContent := testChain.GetTxContext(s.GetTestContext(), nil)

	// txOpts content may auto set these, we don't want these to give false negatives on failing tests
	txContent.TransactOpts.GasFeeCap = nil
	txContent.TransactOpts.GasPrice = nil
	txContent.TransactOpts.GasPrice = nil

	gasBlock, err := testChain.BlockByNumber(s.GetTestContext(), nil)
	Nil(s.T(), err)

	err = gasSetter.SetGasFeeByBlock(s.GetTestContext(), txContent.TransactOpts, gasBlock.NumberU64(), gas.GetConfig().MaxPrice)
	Nil(s.T(), err)

	True(s.T(), txContent.TransactOpts.GasPrice.Cmp(big.NewInt(0)) > 0)
	Nil(s.T(), txContent.TransactOpts.GasFeeCap)
	Nil(s.T(), txContent.TransactOpts.GasTipCap)
}

// TestSetGasFeeMaxPrice makes sure set gas fee respects the max price.
func (s GasSuite) TestSetGasFeeMaxPriceLondon() {
	mockOracle := new(mocks.OracleBackendChain)
	gasSetter := gas.NewGasSetter(s.GetTestContext(), mockOracle)

	// enable london
	mockOracle.On("ChainConfig", mock.Anything).Return(&params.ChainConfig{
		LondonBlock: big.NewInt(0),
	})

	// suggest a gas price of 1000
	suggestedPrice := new(big.Int).Mul(big.NewInt(params.GWei), big.NewInt(1000))
	suggestedTip := new(big.Int).Mul(big.NewInt(params.GWei), big.NewInt(400))

	mockOracle.On("SuggestGasPrice", mock.Anything).Return(suggestedPrice, nil)
	mockOracle.On("SuggestGasTipCap", mock.Anything).Return(suggestedTip, nil)
	mockOracle.On("HeaderByNumber", mock.Anything, mock.Anything).Return(&types.Header{BaseFee: big.NewInt(1)}, nil)

	maxPrice := new(big.Int).Mul(big.NewInt(params.GWei), big.NewInt(750))

	testTransactor := &bind.TransactOpts{}
	err := gasSetter.SetGasFee(s.GetTestContext(), testTransactor, 1, maxPrice)
	Nil(s.T(), err)

	// make sure gas price is nil, this is a london tx
	Nil(s.T(), testTransactor.GasPrice)

	// if we go over, max price should equal fee cap
	Equal(s.T(), testTransactor.GasFeeCap.Cmp(maxPrice), 0)
}

// TestSetGasFeeMaxPrice makes sure set gas fee respects the max price.
func (s GasSuite) TestSetGasFeeMaxPricePreLondon() {
	mockOracle := new(mocks.OracleBackendChain)
	gasSetter := gas.NewGasSetter(s.GetTestContext(), mockOracle)

	// enable london
	mockOracle.On("ChainConfig", mock.Anything).Return(&params.ChainConfig{
		LondonBlock: nil,
	})

	// suggest a gas price of 1000
	suggestedPrice := new(big.Int).Mul(big.NewInt(params.GWei), big.NewInt(1000))
	mockOracle.On("SuggestGasPrice", mock.Anything).Return(suggestedPrice, nil)

	maxPrice := new(big.Int).Mul(big.NewInt(params.GWei), big.NewInt(750))

	testTransactor := &bind.TransactOpts{}
	err := gasSetter.SetGasFee(s.GetTestContext(), testTransactor, 1, maxPrice)
	Nil(s.T(), err)

	// make sure gas price is nil, this is a pre-london tx
	Nil(s.T(), testTransactor.GasFeeCap)
	Nil(s.T(), testTransactor.GasTipCap)

	// if we go over, max price should equal fee cap
	Equal(s.T(), testTransactor.GasPrice.Cmp(maxPrice), 0)
}
