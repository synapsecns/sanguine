package gas_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/chain/gas"
	"gotest.tools/assert"
	"math/big"
	"testing"
	"time"
)

// cmpTestCase is a gas comparison test case.
type cmpTestCase struct {
	// x and y are the txes we're comparing
	x, y *types.Transaction
	// base fee is the base fee for the gas block
	baseFee *big.Int
	// expectedOutput is the expected output as an index
	expectedOutput int
	// description is hte description
	description string
	// nilBlock determines wether a nil block should be used for the test
	nilBlock bool
}

// getCmpTestCases gets the test cases we use to check against gas price test.
func getCmpTestCases() []*cmpTestCase {
	// this will serve as a fuzz test for legacy txes
	gofakeit.Seed(time.Now().UnixNano())

	// gasPriceTests test gas price v gas price
	// we use a random base fee for this test since it shouldn't matter.
	var gasPriceTests = []*cmpTestCase{
		{
			x: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(10),
			}),
			y: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(11),
			}),
			baseFee:        big.NewInt(int64(gofakeit.Uint64())),
			expectedOutput: -1,
			description:    "gas price x < gas price y",
		},
		{
			x: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(10),
			}),
			y: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(11),
			}),
			baseFee:        nil,
			expectedOutput: -1,
			// TODO: we should use ifacemaker to turn block into an interface application wide
			// https://github.com/vburenin/ifacemaker and make sure base fee is not called
			description: "nil base fee (will panic if EffectiveGasTipCmp() is used which is we want)",
		},
		{
			x: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(11),
			}),
			y: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(10),
			}),
			baseFee:        big.NewInt(int64(gofakeit.Uint64())),
			expectedOutput: 1,
			description:    "gas price y < gas price x",
		},
		{
			x: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(10),
			}),
			y: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(10),
			}),
			baseFee:        big.NewInt(int64(gofakeit.Uint64())),
			expectedOutput: 0,
			description:    "gas price y = gas price x",
		},
	}

	// combinedTests test legacy vs dynamic fee
	var combinedTests = []*cmpTestCase{
		{
			x: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(10),
			}),
			y: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(9),
				GasFeeCap: big.NewInt(10),
			}),
			baseFee:        big.NewInt(0),
			expectedOutput: 1,
			description:    "gas tip cap < gas price (legacy) w/ no base fee",
		},
		{
			x: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(10),
			}),
			y: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(9),
				GasFeeCap: big.NewInt(10),
			}),
			baseFee:        big.NewInt(1),
			expectedOutput: 0,
			description:    "gas tip cap < gas price (legacy) w/ base fee making it equal",
		},
		{
			x: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(10),
			}),
			y: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(9),
				GasFeeCap: big.NewInt(11),
			}),
			baseFee:        big.NewInt(2),
			expectedOutput: -1,
			description:    "gas tip cap < gas price (legacy) w/ base fee making less",
		},
	}

	// dynamicTests test dynamicFee vs dynamicFee
	var dynamicTests = []*cmpTestCase{
		{
			x: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(9),
				GasFeeCap: big.NewInt(10),
			}),
			y: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(9),
				GasFeeCap: big.NewInt(10),
			}),
			baseFee:        big.NewInt(1),
			expectedOutput: 0,
			description:    "gas tip cap = gas tip cap",
		},
		{
			x: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(2),
				GasFeeCap: big.NewInt(9),
			}),
			y: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(1),
				GasFeeCap: big.NewInt(10),
			}),
			baseFee:        big.NewInt(2),
			expectedOutput: 1,
			description:    "gas tip cap x > gas tip y cap and fee cap x < fee cap y w/ base fee pushing it over edge",
		},
		{
			x: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(9),
				GasFeeCap: big.NewInt(10),
			}),
			y: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(9),
				GasFeeCap: big.NewInt(10),
			}),
			nilBlock:       true,
			expectedOutput: 0,
			description:    "gas tip cap = gas tip cap and base fee is empty",
		},
		{
			x: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(10),
				GasFeeCap: big.NewInt(11),
			}),
			y: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(9),
				GasFeeCap: big.NewInt(12),
			}),
			nilBlock:       true,
			expectedOutput: 1,
			description:    "gas tip cap > gas tip cap and base fee is empty",
		},
	}

	return append(append(gasPriceTests, combinedTests...), dynamicTests...)
}

// TestCompareGas tests the gas comparison method.
func (s GasSuite) TestCompareGas() {
	s.T().Parallel()

	testCases := getCmpTestCases()
	for _, tc := range testCases {
		tc := tc // capture range variable
		s.T().Run(tc.description, func(t *testing.T) {
			t.Parallel()

			var gasBlock *types.Block
			if !tc.nilBlock {
				// test the straight forward txOpts
				gasBlock = types.NewBlock(&types.Header{
					BaseFee: tc.baseFee,
				}, nil, nil, nil, nil)
			}

			var out int
			if gasBlock != nil {
				out = compareGas(tc.x, tc.y, gasBlock)
			} else {
				out = compareGas(tc.x, tc.y, nil)
			}
			Equal(t, tc.expectedOutput, out)

			// test the reverse as a sanity check. Multiplying by engative 1 should be the opposite
			// except in cases where
			reverseOut := compareGas(tc.y, tc.x, gasBlock)
			Equal(t, tc.expectedOutput*-1, reverseOut)

			optsOut := compareGas(gas.OptsToComparableTx(s.toTransactOpts(tc.x)), gas.OptsToComparableTx(s.toTransactOpts(tc.y)), gasBlock)
			Equal(t, tc.expectedOutput, optsOut)
		})
	}
}

func compareGas(x *types.Transaction, y *types.Transaction, gasBlock *types.Block) (out int) {
	if gasBlock != nil {
		out = gas.CompareGas(x, y, gasBlock.BaseFee())
	} else {
		out = gas.CompareGas(x, y, nil)
	}
	return out
}

func (s GasSuite) toTransactOpts(rawTx *types.Transaction) *bind.TransactOpts {
	switch rawTx.Type() {
	case types.LegacyTxType:
		return &bind.TransactOpts{
			GasPrice:  rawTx.GasPrice(),
			GasFeeCap: nil,
			GasTipCap: nil,
		}
	case types.DynamicFeeTxType:
		return &bind.TransactOpts{
			GasFeeCap: rawTx.GasFeeCap(),
			GasTipCap: rawTx.GasTipCap(),
		}
	}
	s.T().Errorf("unknown txOpts type %d", rawTx.Type())
	return nil
}

// gtTestCase is a test case for testing if a txes gas fee is greater than threshold.
type gtTestCase struct {
	// txOpts is the transaction we're testing against
	tx *types.Transaction
	// threshold is the threshold to pass into the test
	threshold *big.Int
	// expectedOutput is the expected output of the test case
	expectedOutput bool
	// description is a description of the test
	description string
}

func getGasGtTestCases() []*gtTestCase {
	var legacyTxCases = []*gtTestCase{
		{
			tx: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(500),
			}),
			threshold:      big.NewInt(499),
			expectedOutput: true,
			description:    "legacy txOpts greater then",
		},
		{
			tx: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(499),
			}),
			threshold:      big.NewInt(500),
			expectedOutput: false,
			description:    "legacy txOpts less then",
		},
		{
			tx: types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(500),
			}),
			threshold:      big.NewInt(500),
			expectedOutput: false,
			description:    "legacy txOpts equal",
		},
	}

	var dynamicTxCases = []*gtTestCase{
		{
			tx: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(499),
				GasFeeCap: big.NewInt(500),
			}),
			threshold:      big.NewInt(499),
			expectedOutput: true,
			description:    "legacy txOpts greater then",
		},
		{
			tx: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(498),
				GasFeeCap: big.NewInt(499),
			}),
			threshold:      big.NewInt(500),
			expectedOutput: false,
			description:    "legacy txOpts less then",
		},
		{
			tx: types.NewTx(&types.DynamicFeeTx{
				GasTipCap: big.NewInt(499),
				GasFeeCap: big.NewInt(500),
			}),
			threshold:      big.NewInt(500),
			expectedOutput: false,
			description:    "legacy txOpts equal",
		},
	}

	return append(legacyTxCases, dynamicTxCases...)
}

// TestGasGT tests the gas greater then method.
func (s GasSuite) TestGasGt() {
	s.T().Parallel()

	testCases := getGasGtTestCases()
	for _, tc := range testCases {
		tc := tc // capture range variable
		s.T().Run(tc.description, func(t *testing.T) {
			t.Parallel()

			out := gas.FeeGreaterThan(tc.tx, tc.threshold)
			Equal(t, tc.expectedOutput, out)

			optOut := gas.FeeGreaterThan(gas.OptsToComparableTx(s.toTransactOpts(tc.tx)), tc.threshold)
			Equal(t, tc.expectedOutput, optOut)
		})
	}
}

// bumpTestCase tests gas price bumping.
type bumpTestCase struct {
	// base fee is the base fee for the gas block
	baseFee *big.Int
	// txOpts to try to bump
	txOpts *bind.TransactOpts
	// precentIncrease is the percent we should try bumping by
	percentIncrease int
	// expectedGasPrice. Only checked on a legacy txOpts
	expectedGasPrice *big.Int
	// expectedGasTipCap is the expected gas tip cap. Only checked on dynamic txes
	expectedGasTipCap *big.Int
	// expectedGasFeeCap is the expected fee cap
	expectedGasFeeCap *big.Int
	// description is the description of the test case
	description string
}

// checkExpected checks the expected output against the transaction opts.
// this works because *bind.TransactOpts is passed by points.
func (b *bumpTestCase) checkExpected(tb testing.TB) {
	tb.Helper()

	if gas.IsDynamicTx(b.txOpts) {
		Equal(tb, b.txOpts.GasFeeCap.Cmp(b.expectedGasFeeCap), 0)
		Equal(tb, b.txOpts.GasTipCap.Cmp(b.expectedGasTipCap), 0)
	} else {
		Equal(tb, b.txOpts.GasPrice.Cmp(b.expectedGasPrice), 0)
	}
}

// getBumpTestCases gets test cases for bumping.
func getBumpTestCases() []*bumpTestCase {
	// this will serve as a fuzz test for legacy txes
	gofakeit.Seed(time.Now().UnixNano())
	testCfg := gas.GetConfig()

	var legacyTxTestCases = []*bumpTestCase{
		{
			txOpts: &bind.TransactOpts{
				GasPrice: big.NewInt(100),
			},
			percentIncrease:  10,
			baseFee:          big.NewInt(int64(gofakeit.Uint64())),
			expectedGasPrice: big.NewInt(110),
			description:      "bumps gas price by 10%",
		},
		{
			txOpts: &bind.TransactOpts{
				GasPrice: core.CopyBigInt(testCfg.MaxPrice),
			},
			baseFee:          big.NewInt(int64(gofakeit.Uint64())),
			percentIncrease:  10,
			expectedGasPrice: core.CopyBigInt(testCfg.MaxPrice),
			description:      "gas price does not exceed max price",
		},
	}

	var dynamicTxTestCases = []*bumpTestCase{
		{
			txOpts: &bind.TransactOpts{
				GasTipCap: big.NewInt(10),
				GasFeeCap: big.NewInt(100),
			},
			baseFee:           big.NewInt(1),
			percentIncrease:   10,
			expectedGasTipCap: big.NewInt(11),
			expectedGasFeeCap: testCfg.MaxPrice,
			description:       "bump tip by 10%",
		},
		{
			txOpts: &bind.TransactOpts{
				GasTipCap: big.NewInt(10),
				GasFeeCap: core.CopyBigInt(testCfg.MaxPrice),
			},
			baseFee:           big.NewInt(1),
			percentIncrease:   10,
			expectedGasTipCap: big.NewInt(11),
			expectedGasFeeCap: core.CopyBigInt(testCfg.MaxPrice),
			description:       "bump only tip cap (fee cap exceeds max price)",
		},
		{
			txOpts: &bind.TransactOpts{
				GasTipCap: big.NewInt(2),
				GasFeeCap: core.CopyBigInt(testCfg.MaxPrice),
			},
			baseFee:           big.NewInt(0).Sub(core.CopyBigInt(testCfg.MaxPrice), big.NewInt(2)),
			percentIncrease:   10,
			expectedGasTipCap: big.NewInt(2),
			expectedGasFeeCap: core.CopyBigInt(testCfg.MaxPrice),
			description:       "bump nothing (fee cap exceeds max price + base fee, but gas tip alone does not exceed fee cap)",
		},
		{
			txOpts: &bind.TransactOpts{
				GasTipCap: core.CopyBigInt(testCfg.MaxPrice),
				GasFeeCap: core.CopyBigInt(testCfg.MaxPrice),
			},
			baseFee:           big.NewInt(1),
			percentIncrease:   10,
			expectedGasTipCap: core.CopyBigInt(testCfg.MaxPrice),
			expectedGasFeeCap: core.CopyBigInt(testCfg.MaxPrice),
			description:       "bump nothing (tip cap exceeds max price)",
		},
	}

	return append(legacyTxTestCases, dynamicTxTestCases...)
}

func (s GasSuite) TestBumpGasPrice() {
	s.T().Parallel()

	testCases := getBumpTestCases()
	for _, tc := range testCases {
		tc := tc // capture range variable
		s.T().Run(tc.description, func(t *testing.T) {
			t.Parallel()

			gasBlock := types.NewBlock(&types.Header{
				BaseFee: tc.baseFee,
			}, nil, nil, nil, nil)

			gas.BumpGasFees(tc.txOpts, tc.percentIncrease, gasBlock.BaseFee(), gas.GetConfig().MaxPrice)
			tc.checkExpected(s.T())
		})
	}
}

// TestMin makes sure the correct value is returned from the min comparer.
func (s GasSuite) TestMin() {
	a := big.NewInt(0)
	b := big.NewInt(1)

	assert.DeepEqual(s.T(), a, gas.Min(a, b), testsuite.BigIntComparer())
	assert.DeepEqual(s.T(), a, gas.Min(b, a), testsuite.BigIntComparer())
}
