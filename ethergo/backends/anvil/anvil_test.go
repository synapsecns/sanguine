package anvil_test

import (
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"math/big"
)

// Make sure a hardforked rpc will have a balance.
func (a *AnvilSuite) TestFundAccount() {
	fundedAccount := base.MockAccount(a.T())
	ether := big.NewInt(params.Ether)
	a.backend.FundAccount(a.GetTestContext(), fundedAccount.Address, *ether)

	realBalance, err := a.backend.BalanceAt(a.GetTestContext(), fundedAccount.Address, nil)
	Nil(a.T(), err)

	Equal(a.T(), ether, realBalance)
}
