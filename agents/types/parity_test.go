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

func TestEncodeAttestationParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	origin := gofakeit.Uint32()
	destination := origin + 1
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))

	_, attesationContract := deployManager.GetAttestationHarness(ctx, testBackend)

	contractData, err := attesationContract.FormatAttestationData(&bind.CallOpts{Context: ctx}, origin, destination, nonce, root)
	Nil(t, err)

	attestKey := types.AttestationKey{
		Origin:      origin,
		Destination: destination,
		Nonce:       nonce,
	}
	goFormattedData, err := types.EncodeAttestation(types.NewAttestation(attestKey.GetRawKey(), root))
	Nil(t, err)
	Equal(t, contractData, goFormattedData)

	attestationFromBytes, err := types.DecodeAttestation(goFormattedData)
	Nil(t, err)
	Equal(t, origin, attestationFromBytes.Origin())
	Equal(t, destination, attestationFromBytes.Destination())
	Equal(t, nonce, attestationFromBytes.Nonce())
	Equal(t, root, common.Hash(attestationFromBytes.Root()))
}

func TestEncodeSignedAttestationParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	_, attesationContract := deployManager.GetAttestationHarness(ctx, testBackend)

	origin := gofakeit.Uint32()
	destination := origin + 1
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))

	sigGuard1 := types.NewSignature(new(big.Int).SetUint64(uint64(gofakeit.Uint8())), new(big.Int).SetUint64(gofakeit.Uint64()), new(big.Int).SetUint64(gofakeit.Uint64()))
	encodedGuardSignature1, err := types.EncodeSignature(sigGuard1)
	Nil(t, err)
	sigGuard2 := types.NewSignature(new(big.Int).SetUint64(uint64(gofakeit.Uint8())), new(big.Int).SetUint64(gofakeit.Uint64()), new(big.Int).SetUint64(gofakeit.Uint64()))
	encodedGuardSignature2, err := types.EncodeSignature(sigGuard2)
	Nil(t, err)
	sigGuard3 := types.NewSignature(new(big.Int).SetUint64(uint64(gofakeit.Uint8())), new(big.Int).SetUint64(gofakeit.Uint64()), new(big.Int).SetUint64(gofakeit.Uint64()))
	encodedGuardSignature3, err := types.EncodeSignature(sigGuard3)
	Nil(t, err)

	sigNotary1 := types.NewSignature(new(big.Int).SetUint64(uint64(gofakeit.Uint8())), new(big.Int).SetUint64(gofakeit.Uint64()), new(big.Int).SetUint64(gofakeit.Uint64()))
	encodedNotarySignature1, err := types.EncodeSignature(sigNotary1)
	Nil(t, err)
	sigNotary2 := types.NewSignature(new(big.Int).SetUint64(uint64(gofakeit.Uint8())), new(big.Int).SetUint64(gofakeit.Uint64()), new(big.Int).SetUint64(gofakeit.Uint64()))
	encodedNotarySignature2, err := types.EncodeSignature(sigNotary2)
	Nil(t, err)

	attestKey := types.AttestationKey{
		Origin:      origin,
		Destination: destination,
		Nonce:       nonce,
	}

	attestation := types.NewAttestation(attestKey.GetRawKey(), root)

	encodedAttestation, err := types.EncodeAttestation(attestation)
	Nil(t, err)

	encodedGuardSignatures := []byte{}
	encodedNotarySignatures := []byte{}
	encodedGuardSignatures = append(encodedGuardSignatures, encodedGuardSignature1...)
	encodedGuardSignatures = append(encodedGuardSignatures, encodedGuardSignature2...)
	encodedGuardSignatures = append(encodedGuardSignatures, encodedGuardSignature3...)

	encodedNotarySignatures = append(encodedNotarySignatures, encodedNotarySignature1...)
	encodedNotarySignatures = append(encodedNotarySignatures, encodedNotarySignature2...)
	signedContractAttestation, err := attesationContract.FormatAttestation(
		&bind.CallOpts{Context: ctx},
		encodedAttestation,
		encodedGuardSignatures,
		encodedNotarySignatures,
	)
	Nil(t, err)

	signedAttestation := types.NewSignedAttestation(
		types.NewAttestation(attestKey.GetRawKey(), root),
		[]types.Signature{sigGuard1, sigGuard2, sigGuard3},
		[]types.Signature{sigNotary1, sigNotary2},
	)

	goData, err := types.EncodeSignedAttestation(signedAttestation)
	Nil(t, err)

	Equal(t, signedContractAttestation, goData)
}

func TestMessageEncodeParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)
	_, messageContract := deployManager.GetMessageHarness(ctx, testBackend)

	// check constant parity
	version, err := messageContract.MessageVersion0(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, version, types.MessageVersion)

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
