package anvil_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func (a *AnvilSuite) TestGetTxContext() {
	res := a.backend.GetTxContext(a.GetTestContext(), nil)

	tx, err := a.backend.SignTx(types.NewTx(&types.LegacyTx{
		To:       &common.Address{},
		Value:    big.NewInt(0),
		Gas:      res.GasLimit,
		GasPrice: res.GasPrice,
	}), a.backend.Signer(), res.PrivateKey)
	Nil(a.T(), err)

	err = a.backend.SendTransaction(a.GetTestContext(), tx)
	Nil(a.T(), err)

	a.backend.WaitForConfirmation(a.GetTestContext(), tx)
}
