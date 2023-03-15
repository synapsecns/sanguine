package anvil_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
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

func (a *AnvilSuite) TestImpersonateAccount() {
	vitalik := common.HexToAddress("0xd8da6bf26964af9d7eed9e03e53415d37aa96045")
	vitalikFren := base.MockAccount(a.T())

	// impersonate vitalik, and send the fren some eth
	a.backend.ImpersonateAccount(a.GetTestContext(), vitalik, func(auth backends.AuthType) {
		// have vitalikFren sign since the signer shouldn't matter and he's broke rn
		tx, err := a.backend.SignTx(types.NewTx(&types.LegacyTx{
			To:       &vitalikFren.Address,
			Value:    big.NewInt(params.Ether),
			Gas:      auth.GasLimit,
			GasPrice: auth.GasPrice,
		}), a.backend.Signer(), vitalikFren.PrivateKey)
		Nil(a.T(), err)

		Nil(a.T(), a.backend.SendTransaction(a.GetTestContext(), tx))

		a.backend.WaitForConfirmation(a.GetTestContext(), tx)
	})

	// make sure the fren got the eth
	realBalance, err := a.backend.BalanceAt(a.GetTestContext(), vitalikFren.Address, nil)
	Nil(a.T(), err)
	NotEqual(a.T(), big.NewInt(0).Cmp(realBalance), 0)

	// now sure the fren is no longer vitalik
	txOpts := a.backend.GetTxContext(a.GetTestContext(), nil)

	// have vitalikFren sign since the signer shouldn't matter and he's broke rn
	tx, err := a.backend.SignTx(types.NewTx(&types.LegacyTx{
		To: &vitalikFren.Address,
		// oh no that's more than you were given
		Value:    big.NewInt(params.Ether * 5),
		Gas:      txOpts.GasLimit,
		GasPrice: txOpts.GasPrice,
	}), a.backend.Signer(), vitalikFren.PrivateKey)
	Nil(a.T(), err)

	// should error
	NotNil(a.T(), a.backend.SendTransaction(a.GetTestContext(), tx))
}
