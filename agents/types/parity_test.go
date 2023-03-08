package types_test

import (
	"context"
	"crypto/rand"
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

func TestEncodeTipsParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	_, handle := deployManager.GetTipsHarness(ctx, testBackend)

	// make sure constants match
	tipsVersion, err := handle.TipsVersion(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, tipsVersion, types.TipsVersion)

	notaryOffset, err := handle.OffsetNotary(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, notaryOffset, big.NewInt(types.OffsetNotary))

	relayerOffset, err := handle.OffsetBroadcaster(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, relayerOffset, big.NewInt(types.OffsetBroadcaster))

	proverOffset, err := handle.OffsetProver(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, proverOffset, big.NewInt(types.OffsetProver))

	processorOffset, err := handle.OffsetExecutor(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, processorOffset, big.NewInt(types.OffsetExecutor))

	// we want to make sure we can deal w/ overflows
	notaryTip := randomUint96BigInt(t)
	broadcasterTip := randomUint96BigInt(t)
	proverTip := randomUint96BigInt(t)
	executorTip := randomUint96BigInt(t)

	solidityFormattedTips, err := handle.FormatTips(&bind.CallOpts{Context: ctx}, notaryTip, broadcasterTip, proverTip, executorTip)
	Nil(t, err)

	goTips, err := types.EncodeTips(types.NewTips(notaryTip, broadcasterTip, proverTip, executorTip))
	Nil(t, err)

	Equal(t, goTips, solidityFormattedTips)
}

// randomUint96BigInt is a helper method for generating random uint96 values
// see:  https://stackoverflow.com/a/45428754
func randomUint96BigInt(tb testing.TB) *big.Int {
	tb.Helper()

	// Max random value, a 130-bits integer, i.e 2^96 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(96), nil).Sub(max, big.NewInt(1))

	// Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	Nil(tb, err)

	return n
}

func randomUint40BigInt(tb testing.TB) *big.Int {
	tb.Helper()

	// Max random value, a 130-bits integer, i.e 2^96 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(40), nil).Sub(max, big.NewInt(1))

	// Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	Nil(tb, err)

	return n
}

func TestEncodeStateParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	root := common.BigToHash(big.NewInt(gofakeit.Int64()))

	var rootB32 [32]byte
	copy(rootB32[:], root[:])

	origin := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	blockNumber := randomUint40BigInt(t)
	timestamp := randomUint40BigInt(t)

	_, stateContract := deployManager.GetStateHarness(ctx, testBackend)

	contractData, err := stateContract.FormatState(&bind.CallOpts{Context: ctx}, rootB32, origin, nonce, blockNumber, timestamp)
	Nil(t, err)

	goFormattedData, err := types.EncodeState(types.NewState(rootB32, origin, nonce, blockNumber, timestamp))
	Nil(t, err)
	Equal(t, contractData, goFormattedData)

	stateFromBytes, err := types.DecodeState(goFormattedData)
	Nil(t, err)
	Equal(t, rootB32, stateFromBytes.Root())
	Equal(t, origin, stateFromBytes.Origin())
	Equal(t, nonce, stateFromBytes.Nonce())
	Equal(t, blockNumber, stateFromBytes.BlockNumber())
	Equal(t, timestamp, stateFromBytes.Timestamp())
}

func TestEncodeSnapshotParity(t *testing.T) {
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

	var statesAB [][]byte
	stateABytes, err := types.EncodeState(stateA)
	Nil(t, err)
	statesAB = append(statesAB, stateABytes)
	stateBBytes, err := types.EncodeState(stateB)
	Nil(t, err)
	statesAB = append(statesAB, stateBBytes)

	contractData, err := snapshotContract.FormatSnapshot(&bind.CallOpts{Context: ctx}, statesAB)
	Nil(t, err)

	goFormattedData, err := types.EncodeSnapshot(types.NewSnapshot([]types.State{stateA, stateB}))
	Nil(t, err)

	Equal(t, contractData, goFormattedData)

	snapshotFromBytes, err := types.DecodeSnapshot(goFormattedData)
	Nil(t, err)
	Equal(t, stateA.Root(), snapshotFromBytes.States()[0].Root())
	Equal(t, stateA.Origin(), snapshotFromBytes.States()[0].Origin())
	Equal(t, stateA.Nonce(), snapshotFromBytes.States()[0].Nonce())
	Equal(t, stateA.BlockNumber(), snapshotFromBytes.States()[0].BlockNumber())
	Equal(t, stateA.Timestamp(), snapshotFromBytes.States()[0].Timestamp())

	Equal(t, stateB.Root(), snapshotFromBytes.States()[1].Root())
	Equal(t, stateB.Origin(), snapshotFromBytes.States()[1].Origin())
	Equal(t, stateB.Nonce(), snapshotFromBytes.States()[1].Nonce())
	Equal(t, stateB.BlockNumber(), snapshotFromBytes.States()[1].BlockNumber())
	Equal(t, stateB.Timestamp(), snapshotFromBytes.States()[1].Timestamp())
}

func TestMessageEncodeParity(t *testing.T) {
	// TODO (joeallen): FIX ME
	t.Skip()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)
	_, messageContract := deployManager.GetMessageHarness(ctx, testBackend)

	// TODO (joeallen): FIX ME
	// check constant parity
	// version, err := messageContract.MessageVersion0(&bind.CallOpts{Context: ctx})
	// Nil(t, err)
	// Equal(t, version, types.MessageVersion)

	headerOffset, err := messageContract.OffsetHeader(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, headerOffset, big.NewInt(int64(types.HeaderOffset)))

	// generate some fake data
	origin := gofakeit.Uint32()
	sender := common.BigToHash(big.NewInt(gofakeit.Int64()))
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	recipient := common.BigToHash(big.NewInt(gofakeit.Int64()))
	body := []byte(gofakeit.Sentence(gofakeit.Number(5, 15)))
	optimisticSeconds := gofakeit.Uint32()

	notaryTip := randomUint96BigInt(t)
	broadcasterTip := randomUint96BigInt(t)
	proverTip := randomUint96BigInt(t)
	executorTip := randomUint96BigInt(t)

	formattedMessage, err := messageContract.FormatMessage1(&bind.CallOpts{Context: ctx}, origin, sender, nonce, destination, recipient, optimisticSeconds, notaryTip, broadcasterTip, proverTip, executorTip, body)
	Nil(t, err)

	decodedMessage, err := types.DecodeMessage(formattedMessage)
	Nil(t, err)

	Equal(t, decodedMessage.OriginDomain(), origin)
	Equal(t, decodedMessage.Sender(), sender)
	Equal(t, decodedMessage.Nonce(), nonce)
	Equal(t, decodedMessage.DestinationDomain(), destination)
	Equal(t, decodedMessage.Body(), body)
}

func TestHeaderEncodeParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)
	_, headerHarnessContract := deployManager.GetHeaderHarness(ctx, testBackend)

	origin := gofakeit.Uint32()
	sender := common.BigToHash(big.NewInt(gofakeit.Int64()))
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	recipient := common.BigToHash(big.NewInt(gofakeit.Int64()))
	optimisticSeconds := gofakeit.Uint32()

	solHeader, err := headerHarnessContract.FormatHeader(&bind.CallOpts{Context: ctx},
		origin,
		sender,
		nonce,
		destination,
		recipient,
		optimisticSeconds,
	)
	Nil(t, err)

	goHeader, err := types.EncodeHeader(types.NewHeader(origin, sender, nonce, destination, recipient, optimisticSeconds))
	Nil(t, err)

	Equal(t, goHeader, solHeader)

	headerVersion, err := headerHarnessContract.HeaderVersion(&bind.CallOpts{Context: ctx})
	Nil(t, err)

	Equal(t, headerVersion, types.HeaderVersion)
}

func TestAttestationKey(t *testing.T) {
	origin := uint32(1)
	destination := uint32(2)
	nonce := uint32(3)
	attestKey := types.AttestationKey{
		Origin:      origin,
		Destination: destination,
		Nonce:       nonce,
	}
	rawKey := attestKey.GetRawKey()
	attestKeyFromRaw := types.NewAttestationKey(rawKey)
	Equal(t, attestKey.Origin, attestKeyFromRaw.Origin)
	Equal(t, attestKey.Destination, attestKeyFromRaw.Destination)
	Equal(t, attestKey.Nonce, attestKeyFromRaw.Nonce)
}

func TestAttestedDomains(t *testing.T) {
	origin := uint32(1)
	destination := uint32(2)
	attestDomains := types.AttestedDomains{
		Origin:      origin,
		Destination: destination,
	}
	rawDomains := attestDomains.GetRawDomains()
	attestDomainsFromRaw := types.NewAttestedDomains(rawDomains)
	Equal(t, attestDomains.Origin, attestDomainsFromRaw.Origin)
	Equal(t, attestDomains.Destination, attestDomainsFromRaw.Destination)
}

func TestAttestedAgentCounts(t *testing.T) {
	guardCount := uint32(1)
	notaryCount := uint32(2)
	attestationAgentCounts := types.AttestationAgentCounts{
		GuardCount:  guardCount,
		NotaryCount: notaryCount,
	}
	rawDomains := attestationAgentCounts.GetRawAgentCounts()
	attestationAgentCountsFromRaw := types.NewAttestationAgentCounts(rawDomains)
	Equal(t, attestationAgentCounts.GuardCount, attestationAgentCountsFromRaw.GuardCount)
	Equal(t, attestationAgentCounts.NotaryCount, attestationAgentCountsFromRaw.NotaryCount)
}
