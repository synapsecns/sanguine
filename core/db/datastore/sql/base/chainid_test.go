package base_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/db/datastore/sql/base"
	"math/big"
	"testing"
)

func TestGetChainID(t *testing.T) {
	// legacy tx type is always false since we can't derive a chain id from v without signing
	testLegacyTx := types.NewTx(&types.LegacyTx{})
	hasType, _ := base.GetChainID(testLegacyTx)

	False(t, hasType)

	// dynamic tx with a chain id
	testDynamicTx := types.NewTx(&types.DynamicFeeTx{
		ChainID: new(big.Int).SetUint64(gofakeit.Uint64()),
	})
	hasType, chainID := base.GetChainID(testDynamicTx)

	True(t, hasType)
	Equal(t, chainID.Uint64(), testDynamicTx.ChainId().Uint64())

	// dynamic tx with no chain id
	testDynamicTxNoID := types.NewTx(&types.DynamicFeeTx{})
	hasType, _ = base.GetChainID(testDynamicTxNoID)
	False(t, hasType)
}
