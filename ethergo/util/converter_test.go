package util_test

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/util"
	"math/big"
	"testing"
)

// TestTxToCall tests the conversion of a transaction to call.
func TestTxToCall(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	simulatedBackend := simulated.NewSimulatedBackend(ctx, t)
	// accountCount is the number of accounts to test
	const accountCount = 10
	// txCount is the number of txes to test per account
	const txCount = 20
	for i := 0; i < accountCount; i++ {
		acct := simulatedBackend.GetFundedAccount(ctx, big.NewInt(params.Ether))
		for i := 0; i < txCount; i++ {
			// create the mock signed tx
			mockedSignedTx := mocks.MockTx(ctx, t, simulatedBackend, acct, types.LegacyTxType)
			// convert the tx to a call
			mockTxAsCall, err := util.TxToCall(mockedSignedTx)
			Nil(t, err)

			// check equality
			Equal(t, mockTxAsCall.Value, mockedSignedTx.Value())
			Equal(t, mockTxAsCall.From, acct.Address)
			Equal(t, mockTxAsCall.To, mockedSignedTx.To())
			Equal(t, mockTxAsCall.Gas, mockedSignedTx.Gas())
			Equal(t, mockTxAsCall.GasPrice, mockedSignedTx.GasPrice())
			Equal(t, mockTxAsCall.Data, mockedSignedTx.Data())
			Equal(t, mockTxAsCall.AccessList, mockedSignedTx.AccessList())
			Equal(t, mockTxAsCall.GasFeeCap, mockedSignedTx.GasFeeCap())
			Equal(t, mockTxAsCall.GasTipCap, mockedSignedTx.GasTipCap())
		}
	}
}
