package testutil_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/services/rfq/testutil"

	"math/big"
)

func (s *TestUtilSuite) TestDecimals() {
	simulatedBackend := simulated.NewSimulatedBackend(s.GetTestContext(), s.T())
	deployHelper := testutil.NewDeployManager(s.T())

	// test out the usdt token (non-standard, 6 decimal)
	_, usdtRef := deployHelper.GetUSDT(s.GetTestContext(), simulatedBackend)
	s.adjustAmountTestDecimals(usdtRef, 1e6)

	// test out usdc token (standard, 6 decimal)
	_, usdcRef := deployHelper.GetUSDC(s.GetTestContext(), simulatedBackend)
	s.adjustAmountTestDecimals(usdcRef, 1e6)

	_, daiRef := deployHelper.GetDAI(s.GetTestContext(), simulatedBackend)
	s.adjustAmountTestDecimals(daiRef, 1e18)
}

func (s *TestUtilSuite) adjustAmountTestDecimals(handler interface{}, multiplier uint64) {
	amount, err := testutil.AdjustAmount(s.GetTestContext(), big.NewInt(1), handler)
	Nil(s.T(), err)
	Equal(s.T(), amount.Uint64(), 1*multiplier)

	// make sure amount is accounted for
	amount, err = testutil.AdjustAmount(s.GetTestContext(), big.NewInt(2), handler)
	Nil(s.T(), err)
	Equal(s.T(), amount.Uint64(), 2*multiplier)

	amount, err = testutil.AdjustAmount(s.GetTestContext(), big.NewInt(0), handler)
	Nil(s.T(), err)
	Equal(s.T(), amount.Uint64(), uint64(0))
}
