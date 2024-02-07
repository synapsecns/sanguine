package types_test

import (
	"context"
	"crypto/rand"
	"math/big"
	"testing"
	"time"

	"github.com/synapsecns/sanguine/core/testsuite"

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
	summitTip := big.NewInt(int64(gofakeit.Uint32()))
	attestationTip := big.NewInt(int64(gofakeit.Uint32()))
	executionTip := big.NewInt(int64(gofakeit.Uint32()))
	deliveryTip := big.NewInt(int64(gofakeit.Uint32()))

	solidityFormattedTips, err := handle.EncodeTips(&bind.CallOpts{Context: ctx},
		summitTip.Uint64(), attestationTip.Uint64(), executionTip.Uint64(), deliveryTip.Uint64())
	Nil(t, err)

	goTips, err := types.EncodeTips(types.NewTips(summitTip, attestationTip, executionTip, deliveryTip))
	Nil(t, err)

	Equal(t, solidityFormattedTips, new(big.Int).SetBytes(goTips), testsuite.BigIntComparer())

	decodedTips, err := types.DecodeTips(goTips)
	Nil(t, err)

	Equal(t, decodedTips.SummitTip(), summitTip)
	Equal(t, decodedTips.AttestationTip(), attestationTip)
	Equal(t, decodedTips.ExecutionTip(), executionTip)
	Equal(t, decodedTips.DeliveryTip(), deliveryTip)

	// Check the conversion into a big.Int
	goTipsBigInt, err := types.EncodeTipsBigInt(types.NewTips(summitTip, attestationTip, executionTip, deliveryTip))
	Nil(t, err)

	solidityTipsBigInt, err := handle.EncodeTips(&bind.CallOpts{Context: ctx}, summitTip.Uint64(), attestationTip.Uint64(), executionTip.Uint64(), deliveryTip.Uint64())
	Nil(t, err)

	Equal(t, goTipsBigInt.Bytes(), solidityTipsBigInt.Bytes())
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

func TestEncodeGasDataParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	gasPrice := gofakeit.Uint16()
	dataPrice := gofakeit.Uint16()
	execBuffer := gofakeit.Uint16()
	amortAttCost := gofakeit.Uint16()
	etherPrice := gofakeit.Uint16()
	markup := gofakeit.Uint16()

	_, gasDataContract := deployManager.GetGasDataHarness(ctx, testBackend)

	contractData, err := gasDataContract.EncodeGasData(&bind.CallOpts{Context: ctx}, gasPrice, dataPrice, execBuffer, amortAttCost, etherPrice, markup)
	Nil(t, err)

	goFormattedData, err := types.EncodeGasData(types.NewGasData(gasPrice, dataPrice, execBuffer, amortAttCost, etherPrice, markup))
	Nil(t, err)
	Equal(t, contractData.Bytes(), goFormattedData)

	gasDataFromBytes, err := types.DecodeGasData(goFormattedData)
	Nil(t, err)
	Equal(t, gasPrice, gasDataFromBytes.GasPrice())
	Equal(t, dataPrice, gasDataFromBytes.DataPrice())
	Equal(t, execBuffer, gasDataFromBytes.ExecBuffer())
	Equal(t, amortAttCost, gasDataFromBytes.AmortAttCost())
	Equal(t, etherPrice, gasDataFromBytes.EtherPrice())
	Equal(t, markup, gasDataFromBytes.Markup())
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
	_, gasDataContract := deployManager.GetGasDataHarness(ctx, testBackend)

	gasPrice := gofakeit.Uint16()
	dataPrice := gofakeit.Uint16()
	execBuffer := gofakeit.Uint16()
	amortAttCost := gofakeit.Uint16()
	etherPrice := gofakeit.Uint16()
	markup := gofakeit.Uint16()
	gasContractData, err := gasDataContract.EncodeGasData(&bind.CallOpts{Context: ctx}, gasPrice, dataPrice, execBuffer, amortAttCost, etherPrice, markup)
	Nil(t, err)

	contractData, err := stateContract.FormatState(&bind.CallOpts{Context: ctx}, rootB32, origin, nonce, blockNumber, timestamp, gasContractData)
	Nil(t, err)

	gasData := types.NewGasData(gasPrice, dataPrice, execBuffer, amortAttCost, etherPrice, markup)

	goFormattedData, err := types.NewState(rootB32, origin, nonce, blockNumber, timestamp, gasData).Encode()
	Nil(t, err)
	Equal(t, contractData, goFormattedData)

	stateFromBytes, err := types.DecodeState(goFormattedData)
	Nil(t, err)
	Equal(t, rootB32, stateFromBytes.Root())
	Equal(t, origin, stateFromBytes.Origin())
	Equal(t, nonce, stateFromBytes.Nonce())
	Equal(t, blockNumber, stateFromBytes.BlockNumber())
	Equal(t, timestamp, stateFromBytes.Timestamp())
	Equal(t, gasPrice, stateFromBytes.GasData().GasPrice())
	Equal(t, dataPrice, stateFromBytes.GasData().DataPrice())
	Equal(t, execBuffer, stateFromBytes.GasData().ExecBuffer())
	Equal(t, amortAttCost, stateFromBytes.GasData().AmortAttCost())
	Equal(t, etherPrice, stateFromBytes.GasData().EtherPrice())
	Equal(t, markup, stateFromBytes.GasData().Markup())
}

func TestEncodeReceiptParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	_, receiptHarness := deployManager.GetReceiptHarness(ctx, testBackend)

	origin := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	messageHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
	snapshotRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))
	stateIndex := gofakeit.Uint8()
	attNotary := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	firstExecutor := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	finalExecutor := common.BigToAddress(big.NewInt(gofakeit.Int64()))

	receipt := types.NewReceipt(origin, destination, messageHash, snapshotRoot, stateIndex, attNotary, firstExecutor, finalExecutor)

	encodedReceipt, err := receipt.Encode()
	Nil(t, err)

	solEncodedReceipt, err := receiptHarness.FormatReceipt(&bind.CallOpts{Context: ctx}, origin, destination, messageHash, snapshotRoot, stateIndex, attNotary, firstExecutor, finalExecutor)
	Nil(t, err)

	Equal(t, encodedReceipt, solEncodedReceipt)

	decodedReceipt, err := types.DecodeReceipt(encodedReceipt)
	Nil(t, err)

	Equal(t, receipt.Origin(), decodedReceipt.Origin())
	Equal(t, receipt.Destination(), decodedReceipt.Destination())
	Equal(t, receipt.MessageHash(), decodedReceipt.MessageHash())
	Equal(t, receipt.SnapshotRoot(), decodedReceipt.SnapshotRoot())
	Equal(t, receipt.StateIndex(), decodedReceipt.StateIndex())
	Equal(t, receipt.AttestationNotary(), decodedReceipt.AttestationNotary())
	Equal(t, receipt.FirstExecutor(), decodedReceipt.FirstExecutor())
	Equal(t, receipt.FinalExecutor(), decodedReceipt.FinalExecutor())
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

	gasPriceA := gofakeit.Uint16()
	dataPriceA := gofakeit.Uint16()
	execBufferA := gofakeit.Uint16()
	amortAttCostA := gofakeit.Uint16()
	etherPriceA := gofakeit.Uint16()
	markupA := gofakeit.Uint16()
	gasDataA := types.NewGasData(gasPriceA, dataPriceA, execBufferA, amortAttCostA, etherPriceA, markupA)

	gasPriceB := gofakeit.Uint16()
	dataPriceB := gofakeit.Uint16()
	execBufferB := gofakeit.Uint16()
	amortAttCostB := gofakeit.Uint16()
	etherPriceB := gofakeit.Uint16()
	markupB := gofakeit.Uint16()
	gasDataB := types.NewGasData(gasPriceB, dataPriceB, execBufferB, amortAttCostB, etherPriceB, markupB)

	stateA := types.NewState(rootA, originA, nonceA, blockNumberA, timestampA, gasDataA)
	stateB := types.NewState(rootB, originB, nonceB, blockNumberB, timestampB, gasDataB)

	var statesAB [][]byte
	stateABytes, err := stateA.Encode()
	Nil(t, err)
	statesAB = append(statesAB, stateABytes)
	stateBBytes, err := stateB.Encode()
	Nil(t, err)
	statesAB = append(statesAB, stateBBytes)

	contractData, err := snapshotContract.FormatSnapshot(&bind.CallOpts{Context: ctx}, statesAB)
	Nil(t, err)

	goFormattedData, err := types.NewSnapshot([]types.State{stateA, stateB}).Encode()
	Nil(t, err)

	Equal(t, contractData, goFormattedData)

	snapshotFromBytes, err := types.DecodeSnapshot(goFormattedData)
	Nil(t, err)
	Equal(t, stateA.Root(), snapshotFromBytes.States()[0].Root())
	Equal(t, stateA.Origin(), snapshotFromBytes.States()[0].Origin())
	Equal(t, stateA.Nonce(), snapshotFromBytes.States()[0].Nonce())
	Equal(t, stateA.BlockNumber(), snapshotFromBytes.States()[0].BlockNumber())
	Equal(t, stateA.Timestamp(), snapshotFromBytes.States()[0].Timestamp())
	Equal(t, stateA.GasData().GasPrice(), snapshotFromBytes.States()[0].GasData().GasPrice())
	Equal(t, stateA.GasData().DataPrice(), snapshotFromBytes.States()[0].GasData().DataPrice())
	Equal(t, stateA.GasData().ExecBuffer(), snapshotFromBytes.States()[0].GasData().ExecBuffer())
	Equal(t, stateA.GasData().AmortAttCost(), snapshotFromBytes.States()[0].GasData().AmortAttCost())
	Equal(t, stateA.GasData().EtherPrice(), snapshotFromBytes.States()[0].GasData().EtherPrice())
	Equal(t, stateA.GasData().Markup(), snapshotFromBytes.States()[0].GasData().Markup())

	Equal(t, stateB.Root(), snapshotFromBytes.States()[1].Root())
	Equal(t, stateB.Origin(), snapshotFromBytes.States()[1].Origin())
	Equal(t, stateB.Nonce(), snapshotFromBytes.States()[1].Nonce())
	Equal(t, stateB.BlockNumber(), snapshotFromBytes.States()[1].BlockNumber())
	Equal(t, stateB.Timestamp(), snapshotFromBytes.States()[1].Timestamp())
	Equal(t, stateB.GasData().GasPrice(), snapshotFromBytes.States()[1].GasData().GasPrice())
	Equal(t, stateB.GasData().DataPrice(), snapshotFromBytes.States()[1].GasData().DataPrice())
	Equal(t, stateB.GasData().ExecBuffer(), snapshotFromBytes.States()[1].GasData().ExecBuffer())
	Equal(t, stateB.GasData().AmortAttCost(), snapshotFromBytes.States()[1].GasData().AmortAttCost())
	Equal(t, stateB.GasData().EtherPrice(), snapshotFromBytes.States()[1].GasData().EtherPrice())
	Equal(t, stateB.GasData().Markup(), snapshotFromBytes.States()[1].GasData().Markup())

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

	var rootB32, dataHashB32 [32]byte
	copy(rootB32[:], snapRoot[:])
	copy(dataHashB32[:], agentRoot[:])

	nonce := gofakeit.Uint32()
	blockNumber := randomUint40BigInt(t)
	timestamp := randomUint40BigInt(t)

	contractData, err := attestationContract.FormatAttestation(&bind.CallOpts{Context: ctx}, rootB32, dataHashB32, nonce, blockNumber, timestamp)
	Nil(t, err)

	goFormattedData, err := types.NewAttestation(rootB32, dataHashB32, nonce, blockNumber, timestamp).Encode()
	Nil(t, err)

	Equal(t, contractData, goFormattedData)

	attestationFromBytes, err := types.DecodeAttestation(goFormattedData)
	Nil(t, err)
	Equal(t, rootB32, attestationFromBytes.SnapshotRoot())
	Equal(t, dataHashB32, attestationFromBytes.DataHash())
	Equal(t, nonce, attestationFromBytes.Nonce())
	Equal(t, blockNumber, attestationFromBytes.BlockNumber())
	Equal(t, timestamp, attestationFromBytes.Timestamp())

	// Testing data hash.
	gasData := types.NewGasData(gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16())

	gasDataBytes, err := types.EncodeGasData(gasData)
	Nil(t, err)

	gasDataHash := crypto.Keccak256Hash(gasDataBytes)

	var agentRootB32, gasDataHashB32 [32]byte
	copy(agentRootB32[:], agentRoot[:])
	copy(gasDataHashB32[:], gasDataHash[:])

	attestationDataHash := types.GetAttestationDataHash(agentRootB32, gasDataHashB32)
	attestation := types.NewAttestation([32]byte{1}, attestationDataHash, 1, big.NewInt(1), big.NewInt(1))

	contractDataHashFromVals, err := attestationContract.DataHash(&bind.CallOpts{Context: ctx}, agentRootB32, gasDataHashB32)
	Nil(t, err)

	Equal(t, contractDataHashFromVals, attestationDataHash)

	encodedDataHashAttestation, err := attestation.Encode()
	Nil(t, err)

	contractDataHashFromAtt, err := attestationContract.DataHash0(&bind.CallOpts{Context: ctx}, encodedDataHashAttestation)
	Nil(t, err)

	Equal(t, contractDataHashFromAtt, attestationDataHash)
}

func TestBaseMessageEncodeParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)
	_, baseMessageContract := deployManager.GetBaseMessageHarness(ctx, testBackend)
	_, tipsContract := deployManager.GetTipsHarness(ctx, testBackend)
	_, requestContract := deployManager.GetRequestHarness(ctx, testBackend)
	_, messageContract := deployManager.GetMessageHarness(ctx, testBackend)
	_, headerContract := deployManager.GetHeaderHarness(ctx, testBackend)

	// Generate some fake data.
	flag := types.MessageFlagBase
	origin := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	optimisticSeconds := gofakeit.Uint32()
	summitTip := gofakeit.Uint64()
	attestationTip := gofakeit.Uint64()
	executionTip := gofakeit.Uint64()
	deliveryTip := gofakeit.Uint64()
	sender := common.BigToHash(big.NewInt(gofakeit.Int64()))
	recipient := common.BigToHash(big.NewInt(gofakeit.Int64()))
	gasDrop := randomUint96BigInt(t)
	gasLimit := gofakeit.Uint64()
	version := gofakeit.Uint32()
	content := []byte{byte(gofakeit.Int64())}

	formattedHeader, err := headerContract.EncodeHeader(&bind.CallOpts{Context: ctx}, uint8(flag), origin, nonce, destination, optimisticSeconds)
	Nil(t, err)

	formattedTips, err := tipsContract.EncodeTips(&bind.CallOpts{Context: ctx}, summitTip, attestationTip, executionTip, deliveryTip)
	Nil(t, err)

	formattedRequest, err := requestContract.EncodeRequest(&bind.CallOpts{Context: ctx}, gasDrop, gasLimit, version)
	Nil(t, err)

	formattedBaseMessage, err := baseMessageContract.FormatBaseMessage(&bind.CallOpts{Context: ctx}, formattedTips, sender, recipient, formattedRequest, content)
	Nil(t, err)

	formattedMessage, err := messageContract.FormatMessage(&bind.CallOpts{Context: ctx}, formattedHeader, formattedBaseMessage)
	Nil(t, err)

	decodedMessage, err := types.DecodeMessage(formattedMessage)
	Nil(t, err)

	// Header parity.
	Equal(t, decodedMessage.Header().Flag(), flag)
	Equal(t, decodedMessage.Header().OriginDomain(), origin)
	Equal(t, decodedMessage.Header().Nonce(), nonce)
	Equal(t, decodedMessage.Header().DestinationDomain(), destination)
	Equal(t, decodedMessage.Header().OptimisticSeconds(), optimisticSeconds)

	// BaseMessage parity.
	senderBytes32 := [32]byte{}
	copy(senderBytes32[:], sender.Bytes()[:32])
	recipientBytes32 := [32]byte{}
	copy(recipientBytes32[:], recipient.Bytes()[:32])
	Equal(t, decodedMessage.BaseMessage().Sender(), senderBytes32)
	Equal(t, decodedMessage.BaseMessage().Recipient(), recipientBytes32)
	Equal(t, decodedMessage.BaseMessage().Content(), content)

	// Leaf parity.
	bodyLeaf, err := baseMessageContract.BodyLeaf(&bind.CallOpts{Context: ctx}, formattedBaseMessage)
	Nil(t, err)
	goBodyLeaf, err := decodedMessage.BaseMessage().BodyLeaf()
	Nil(t, err)
	Equal(t, bodyLeaf[:], goBodyLeaf)

	baseMessageLeaf, err := baseMessageContract.Leaf(&bind.CallOpts{Context: ctx}, formattedBaseMessage)
	Nil(t, err)
	goBaseMessageLeaf, err := decodedMessage.BaseMessage().Leaf()
	Nil(t, err)
	Equal(t, baseMessageLeaf, goBaseMessageLeaf)

	headerLeaf, err := headerContract.Leaf(&bind.CallOpts{Context: ctx}, formattedHeader)
	Nil(t, err)
	goHeaderLeaf, err := decodedMessage.Header().Leaf()
	Nil(t, err)
	Equal(t, headerLeaf, goHeaderLeaf)

	leaf, err := messageContract.Leaf(&bind.CallOpts{Context: ctx}, formattedMessage)
	Nil(t, err)
	goLeaf, err := decodedMessage.ToLeaf()
	Nil(t, err)
	Equal(t, leaf, goLeaf)
}

// TODO: Add separate tests for BaseMessage and ManagerMessage.
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
	flag := types.MessageFlagManager
	origin := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	body := []byte(gofakeit.Sentence(gofakeit.Number(5, 15)))
	optimisticSeconds := gofakeit.Uint32()

	formattedHeader, err := headerContract.EncodeHeader(&bind.CallOpts{Context: ctx}, uint8(flag), origin, nonce, destination, optimisticSeconds)
	Nil(t, err)

	formattedMessage, err := messageContract.FormatMessage(&bind.CallOpts{Context: ctx}, formattedHeader, body)
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

	flag := types.MessageFlagManager
	origin := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	optimisticSeconds := gofakeit.Uint32()

	solHeader, err := headerHarnessContract.EncodeHeader(&bind.CallOpts{Context: ctx},
		uint8(flag),
		origin,
		nonce,
		destination,
		optimisticSeconds,
	)
	Nil(t, err)

	goHeader, err := types.EncodeHeader(types.NewHeader(flag, origin, nonce, destination, optimisticSeconds))
	Nil(t, err)

	Equal(t, goHeader, solHeader.Bytes())

	goHeaderHash, err := types.NewHeader(flag, origin, nonce, destination, optimisticSeconds).Leaf()
	Nil(t, err)

	solHeaderHash, err := headerHarnessContract.Leaf(&bind.CallOpts{Context: ctx}, solHeader)
	Nil(t, err)

	Equal(t, goHeaderHash, solHeaderHash)
}
