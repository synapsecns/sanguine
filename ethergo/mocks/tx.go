package mocks

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"math/big"
	"testing"
)

// MockBlocksOnBackend mocks count blocks on a backend.
func MockBlocksOnBackend(ctx context.Context, t *testing.T, simulatedBackend backends.SimulatedTestBackend, count int) (blocks []*types.Block) {
	t.Helper()

	acct := simulatedBackend.GetFundedAccount(ctx, big.NewInt(params.Ether))
	lastBlock, err := simulatedBackend.BlockByNumber(ctx, nil)
	assert.Nil(t, err)

	// there should be done tx per block on simulated backend but we want to make sure
	defer func() {
		assert.Equal(t, count, len(blocks))
	}()

	for len(blocks) < count {
		MockTx(ctx, t, simulatedBackend, acct, types.LegacyTxType)
		newBlock, err := simulatedBackend.BlockByNumber(ctx, nil)
		assert.Nil(t, err)

		if newBlock.Number().Uint64() > lastBlock.Number().Uint64() {
			blocks = append(blocks, newBlock)
			lastBlock = newBlock
		}
	}
	return blocks
}

// GetMockTxes gets count txes from a new mock account. These are real txes ona  simulated chain
// useful for testing db interactions/adding noise to listeners.
func GetMockTxes(ctx context.Context, t *testing.T, count int, txType uint8) (txes []*types.Transaction) {
	t.Helper()
	simulatedBackend := simulated.NewSimulatedBackend(ctx, t)

	acct := simulatedBackend.GetFundedAccount(ctx, big.NewInt(params.Ether))
	for i := 0; i < count; i++ {
		txes = append(txes, MockTx(ctx, t, simulatedBackend, acct, txType))
	}
	return txes
}

// MockTx mocks a transaction on a simulated backend.
func MockTx(ctx context.Context, t *testing.T, simulatedBackend backends.SimulatedTestBackend, acct *keystore.Key, txType uint8) *types.Transaction {
	t.Helper()
	nonce, err := simulatedBackend.PendingNonceAt(ctx, acct.Address)
	assert.Nil(t, err)

	to := common.BigToAddress(big.NewInt(0))
	value := big.NewInt(params.GWei)

	var rawTx *types.Transaction
	switch txType {
	case types.LegacyTxType:
		gasPrice, err := simulatedBackend.SuggestGasPrice(ctx)
		assert.Nil(t, err)

		rawTx = types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      21000,
			To:       &to,
			Value:    value,
		})
	case types.DynamicFeeTxType:
		gasTipCap, err := simulatedBackend.SuggestGasTipCap(ctx)
		assert.Nil(t, err)

		latestBlock, err := simulatedBackend.BlockByNumber(ctx, nil)
		assert.Nil(t, err)

		rawTx = types.NewTx(&types.DynamicFeeTx{
			ChainID:   simulatedBackend.GetBigChainID(),
			Nonce:     nonce,
			GasTipCap: gasTipCap,
			GasFeeCap: big.NewInt(0).Add(latestBlock.BaseFee(), gasTipCap),
			Gas:       21000,
			To:        &to,
			Value:     value,
		})
	default:
		t.Errorf("tx type %d unspoorted", txType)
	}

	signedTx, err := simulatedBackend.SignTx(rawTx, simulatedBackend.Signer(), acct.PrivateKey)
	assert.Nil(t, err)

	err = simulatedBackend.SendTransaction(ctx, signedTx)
	assert.Nil(t, err)

	return signedTx
}
