package types_test

import (
	"testing"
)

func TestHash(t *testing.T) {
	// TODO (joe): FIX ME
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// defer cancel()

	// testBackend := simulated.NewSimulatedBackend(ctx, t)
	// deployManager := testutil.NewDeployManager(t)

	// _, stateContract := deployManager.GetStateHarness(ctx, testBackend)

	// root := common.BigToHash(big.NewInt(gofakeit.Int64()))

	// var rootB32 [32]byte
	// copy(rootB32[:], root[:])

	// origin := gofakeit.Uint32()
	// nonce := gofakeit.Uint32()
	// blockNumber := randomUint40BigInt(t)
	// timestamp := randomUint40BigInt(t)
	// state := types.NewState(rootB32, origin, nonce, blockNumber, timestamp)
	// stateHash, err := state.Hash()
	// Nil(t, err)

	// encodedState, err := types.EncodeState(state)
	// Nil(t, err)

	// stateContractHash, err := stateContract.Hash(&bind.CallOpts{Context: ctx}, encodedState)
	// Nil(t, err)

	// Equal(t, stateHash, stateContractHash)
}
