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
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

func TestEncodeTipsParity(t *testing.T) {
	// TODO (joe): re-enabled this
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	_, handle := deployManager.GetTipsHarness(ctx, testBackend)

	// we want to make sure we can deal w/ overflows
	summitTip := randomUint64BigInt(t)
	attestationTip := randomUint64BigInt(t)
	executionTip := randomUint64BigInt(t)
	deliveryTip := randomUint64BigInt(t)

	solidityFormattedTips, err := handle.EncodeTips(&bind.CallOpts{Context: ctx},
		summitTip.Uint64(), attestationTip.Uint64(), executionTip.Uint64(), deliveryTip.Uint64())
	Nil(t, err)

	goTips, err := types.EncodeTips(types.NewTips(summitTip, attestationTip, executionTip, deliveryTip))
	Nil(t, err)

	Equal(t, solidityFormattedTips.Bytes(), goTips)

	decodedTips, err := types.DecodeTips(goTips)
	Nil(t, err)

	Equal(t, decodedTips.SummitTip(), summitTip)
	Equal(t, decodedTips.AttestationTip(), attestationTip)
	Equal(t, decodedTips.ExecutionTip(), executionTip)
	Equal(t, decodedTips.DeliveryTip(), deliveryTip)
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

func randomUint64BigInt(tb testing.TB) *big.Int {
	tb.Helper()

	// Max random value, a 130-bits integer, i.e 2^96 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(64), nil).Sub(max, big.NewInt(1))

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

	testWallet, err := wallet.FromRandom()
	Nil(t, err)
	testSigner := localsigner.NewSigner(testWallet.PrivateKey())

	basicHashToSign := crypto.Keccak256Hash([]byte{0x1})
	firstSignature, err := crypto.Sign(basicHashToSign[:], testWallet.PrivateKey())
	Nil(t, err)
	encPubKey := crypto.FromECDSAPub(testWallet.PublicKey())
	Equal(t, 65, len(encPubKey))
	Equal(t, 65, len(firstSignature))
	True(t, crypto.VerifySignature(encPubKey, basicHashToSign[:], firstSignature[:crypto.RecoveryIDOffset]))

	testSignature, testSignedEncodedSnapshot, testSnapshotHash, err := snapshotFromBytes.SignSnapshot(ctx, testSigner)
	Nil(t, err)
	Equal(t, goFormattedData, testSignedEncodedSnapshot)
	Greater(t, len(testSnapshotHash), 0)
	encodedSignature, err := types.EncodeSignature(testSignature)
	Nil(t, err)
	True(t, crypto.VerifySignature(crypto.FromECDSAPub(testWallet.PublicKey()), core.BytesToSlice(testSnapshotHash), encodedSignature[:crypto.RecoveryIDOffset]))
}

/*
	func VerifySignature(pubkey, digestHash, signature []byte) bool {
		return secp256k1.VerifySignature(pubkey, digestHash, signature)
	}
*/

func TestEncodeAttestationParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	_, attestationContract := deployManager.GetAttestationHarness(ctx, testBackend)

	snapRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))
	agentRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))

	var rootB32, agentRootB32 [32]byte
	copy(rootB32[:], snapRoot[:])
	copy(agentRootB32[:], agentRoot[:])

	nonce := gofakeit.Uint32()
	blockNumber := randomUint40BigInt(t)
	timestamp := randomUint40BigInt(t)

	contractData, err := attestationContract.FormatAttestation(&bind.CallOpts{Context: ctx}, rootB32, agentRootB32, nonce, blockNumber, timestamp)
	Nil(t, err)

	goFormattedData, err := types.EncodeAttestation(types.NewAttestation(rootB32, agentRootB32, nonce, blockNumber, timestamp))
	Nil(t, err)

	Equal(t, contractData, goFormattedData)

	attestationFromBytes, err := types.DecodeAttestation(goFormattedData)
	Nil(t, err)
	Equal(t, rootB32, attestationFromBytes.SnapshotRoot())
	Equal(t, agentRootB32, attestationFromBytes.AgentRoot())
	Equal(t, nonce, attestationFromBytes.Nonce())
	Equal(t, blockNumber, attestationFromBytes.BlockNumber())
	Equal(t, timestamp, attestationFromBytes.Timestamp())
}

func TestMessageEncodeParity(t *testing.T) {
	// TODO (joeallen): FIX ME
	// t.Skip()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)
	_, messageContract := deployManager.GetMessageHarness(ctx, testBackend)
	_, headerContract := deployManager.GetHeaderHarness(ctx, testBackend)

	// generate some fake data
	origin := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	body := []byte(gofakeit.Sentence(gofakeit.Number(5, 15)))
	optimisticSeconds := gofakeit.Uint32()

	flag := uint8(1)

	formattedHeader, err := headerContract.EncodeHeader(&bind.CallOpts{Context: ctx}, origin, nonce, destination, optimisticSeconds)
	Nil(t, err)

	goHeader, err := types.EncodeHeader(types.NewHeader(origin, nonce, destination, optimisticSeconds))
	Nil(t, err)
	formattedHeaderFromGo := new(big.Int).SetBytes(goHeader)

	Equal(t, formattedHeader, formattedHeaderFromGo)

	formattedMessage, err := messageContract.FormatMessage(&bind.CallOpts{Context: ctx}, flag, formattedHeader, body)
	Nil(t, err)

	decodedMessage, err := types.DecodeMessage(formattedMessage)
	Nil(t, err)

	Equal(t, decodedMessage.OriginDomain(), origin)
	Equal(t, decodedMessage.Nonce(), nonce)
	Equal(t, decodedMessage.DestinationDomain(), destination)
	Equal(t, decodedMessage.OptimisticSeconds(), optimisticSeconds)
	Equal(t, decodedMessage.Body(), body)
}

func TestHeaderEncodeParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)
	_, headerHarnessContract := deployManager.GetHeaderHarness(ctx, testBackend)

	origin := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	optimisticSeconds := gofakeit.Uint32()

	solHeader, err := headerHarnessContract.EncodeHeader(&bind.CallOpts{Context: ctx},
		origin,
		nonce,
		destination,
		optimisticSeconds,
	)
	Nil(t, err)

	goHeader, err := types.EncodeHeader(types.NewHeader(origin, nonce, destination, optimisticSeconds))
	Nil(t, err)

	Equal(t, goHeader, solHeader.Bytes())
}
