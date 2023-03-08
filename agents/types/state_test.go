package types_test

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"math/big"
	"testing"
	"time"
)

func TestHash(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	_, stateContract := deployManager.GetStateHarness(ctx, testBackend)

	root := common.BigToHash(big.NewInt(gofakeit.Int64()))

	var rootB32 [32]byte
	copy(rootB32[:], root[:])

	origin := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	blockNumber := randomUint40BigInt(t)
	timestamp := randomUint40BigInt(t)
	state := types.NewState(rootB32, origin, nonce, blockNumber, timestamp)
	stateHash, err := state.Hash()
	Nil(t, err)

	encodedState, err := types.EncodeState(state)
	Nil(t, err)

	stateContractHash, err := stateContract.Hash(&bind.CallOpts{Context: ctx}, encodedState)
	Nil(t, err)

	Equal(t, stateHash, stateContractHash)
}
