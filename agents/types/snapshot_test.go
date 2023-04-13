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

func TestSnapshotRootAndProofs(t *testing.T) {
	// TODO (Max Planck): Fix me
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	_, snapshotContract := deployManager.GetSnapshotHarness(ctx, testBackend)

	rootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
	rootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
	originA := gofakeit.Uint32()
	originB := gofakeit.Uint32()
	nonceA := gofakeit.Uint32()
	nonceB := gofakeit.Uint32()
	blockNumberA := randomUint40BigInt(t)
	blockNumberB := randomUint40BigInt(t)
	timestampA := randomUint40BigInt(t)
	timestampB := randomUint40BigInt(t)

	stateA := types.NewState(rootA, originA, nonceA, blockNumberA, timestampA)
	stateB := types.NewState(rootB, originB, nonceB, blockNumberB, timestampB)
	snapshot := types.NewSnapshot([]types.State{stateA, stateB})

	snapshotRoot, _, err := snapshot.SnapshotRootAndProofs()
	Nil(t, err)

	encodedSnapshot, err := types.EncodeSnapshot(snapshot)
	Nil(t, err)

	snapshotContractRoot, err := snapshotContract.Root(&bind.CallOpts{Context: ctx}, encodedSnapshot)
	Nil(t, err)

	Equal(t, snapshotRoot, snapshotContractRoot)
}
