package anvil_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/util"
	"math/big"
)

// Make sure a hardforked rpc will have a balance.
func (a *AnvilSuite) TestFundAccount() {
	fundedAccount := base.MockAccount(a.T())
	ether := big.NewInt(params.Ether)
	a.backend.FundAccount(a.GetTestContext(), fundedAccount.Address, *ether)

	realBalance, err := a.backend.BalanceAt(a.GetTestContext(), fundedAccount.Address, nil)
	a.Require().NoError(err)

	Equal(a.T(), ether, realBalance)
}

func (a *AnvilSuite) TestGetTxContext() {
	res := a.backend.GetTxContext(a.GetTestContext(), nil)

	prevCount, err := a.counter.GetCount(&bind.CallOpts{Context: a.GetTestContext()})
	a.Require().NoError(err)

	res.TransactOpts.NoSend = true
	tx, err := a.counter.IncrementCounter(res.TransactOpts)
	a.Require().NoError(err)

	sender, err := util.TxToCall(tx)
	a.Require().NoError(err)

	Equal(a.T(), res.TransactOpts.From, sender.From)

	a.backend.WaitForConfirmation(a.GetTestContext(), tx)

	newCount, err := a.counter.GetCount(&bind.CallOpts{Context: a.GetTestContext()})
	a.Require().NoError(err)

	Equal(a.T(), prevCount.Uint64()+1, newCount.Uint64())
}

func (a *AnvilSuite) TestImpersonateAccount() {
	ogCount, err := a.counter.GetVitalikCount(&bind.CallOpts{Context: a.GetTestContext()})
	a.Require().NoError(err)

	// impersonate vitalik, and send the fren some eth
	err = a.backend.ImpersonateAccount(a.GetTestContext(), vitalik, func(transactOpts *bind.TransactOpts) *types.Transaction {
		tx, err := a.counter.VitalikIncrement(transactOpts)
		a.Require().NoError(err)

		return tx
	})
	a.Require().NoError(err)

	vitalikCount, err := a.counter.GetVitalikCount(&bind.CallOpts{Context: a.GetTestContext()})
	a.Require().NoError(err)

	Equal(a.T(), ogCount.Uint64()+10, vitalikCount.Uint64())
}
