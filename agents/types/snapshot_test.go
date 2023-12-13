package types_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
)

func TestSnapshotRootAndProofs(t *testing.T) {
	// TODO (Max Planck): Fix me
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	_, snapshotContract := deployManager.GetSnapshotHarness(ctx, testBackend)
	_, stateContract := deployManager.GetStateHarness(ctx, testBackend)

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

	gasPriceA := gofakeit.Uint16()
	dataPriceA := gofakeit.Uint16()
	execBufferA := gofakeit.Uint16()
	amortAttCostA := gofakeit.Uint16()
	etherPriceA := gofakeit.Uint16()
	markupA := gofakeit.Uint16()
	gasDataA := types.NewGasData(gasPriceA, dataPriceA, execBufferA, amortAttCostA, etherPriceA, markupA)

	stateA := types.NewState(rootA, originA, nonceA, blockNumberA, timestampA, gasDataA)

	gasPriceB := gofakeit.Uint16()
	dataPriceB := gofakeit.Uint16()
	execBufferB := gofakeit.Uint16()
	amortAttCostB := gofakeit.Uint16()
	etherPriceB := gofakeit.Uint16()
	markupB := gofakeit.Uint16()
	gasDataB := types.NewGasData(gasPriceB, dataPriceB, execBufferB, amortAttCostB, etherPriceB, markupB)

	stateB := types.NewState(rootB, originB, nonceB, blockNumberB, timestampB, gasDataB)
	snapshot := types.NewSnapshot([]types.State{stateA, stateB})

	snapshotRoot, _, err := snapshot.SnapshotRootAndProofs()
	Nil(t, err)

	encodedSnapshot, err := snapshot.Encode()
	Nil(t, err)

	snapshotContractStatesAmount, err := snapshotContract.StatesAmount(&bind.CallOpts{Context: ctx}, encodedSnapshot)
	Nil(t, err)

	Equal(t, big.NewInt(2), snapshotContractStatesAmount)

	snapshotContractStateA, err := snapshotContract.State(&bind.CallOpts{Context: ctx}, encodedSnapshot, big.NewInt(0))
	Nil(t, err)

	stateABytes, err := stateA.Encode()
	Nil(t, err)

	Equal(t, stateABytes, snapshotContractStateA)

	snapshotContractStateB, err := snapshotContract.State(&bind.CallOpts{Context: ctx}, encodedSnapshot, big.NewInt(1))
	Nil(t, err)

	stateBBytes, err := stateB.Encode()
	Nil(t, err)

	Equal(t, stateBBytes, snapshotContractStateB)

	stateALeaf, err := stateA.Hash()
	Nil(t, err)

	stateContractStateALeaf, err := stateContract.Leaf(&bind.CallOpts{Context: ctx}, stateABytes)
	Nil(t, err)

	Equal(t, stateALeaf, stateContractStateALeaf)

	stateBLeaf, err := stateB.Hash()
	Nil(t, err)

	stateContractStateBLeaf, err := stateContract.Leaf(&bind.CallOpts{Context: ctx}, stateBBytes)
	Nil(t, err)

	Equal(t, stateBLeaf, stateContractStateBLeaf)

	snapshotContractRoot, err := snapshotContract.CalculateRoot(&bind.CallOpts{Context: ctx}, encodedSnapshot)
	Nil(t, err)

	Equal(t, snapshotRoot, snapshotContractRoot)
}
