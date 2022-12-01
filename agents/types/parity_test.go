package types_test

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
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

	sig := types.NewSignature(new(big.Int).SetUint64(uint64(gofakeit.Uint8())), new(big.Int).SetUint64(gofakeit.Uint64()), new(big.Int).SetUint64(gofakeit.Uint64()))

	attestKey := types.AttestationKey{
		Origin:      origin,
		Destination: destination,
		Nonce:       nonce,
	}
	signedAttestation := types.NewSignedAttestation(
		types.NewAttestation(attestKey.GetRawKey(), root),
		sig,
	)

	encodedSignature, err := types.EncodeSignature(sig)
	Nil(t, err)

	signedContractAttestation, err := attesationContract.FormatAttestation(&bind.CallOpts{Context: ctx}, origin, destination, nonce, root, encodedSignature)
	Nil(t, err)

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

	// create a new encoded message to test against
	header := types.NewHeader(origin, sender, nonce, destination, recipient, optimisticSeconds)
	headerBytes, err := types.EncodeHeader(header)
	Nil(t, err)
	tips := types.NewTips(notaryTip, broadcasterTip, proverTip, executorTip)
	tipsBytes, err := types.EncodeTips(tips)
	Nil(t, err)
	testMessage := types.NewMessage(header, tips, body)
	testMessageLeaf, err := testMessage.ToLeaf()
	Nil(t, err)

	messageLeaf, err := messageContract.MessageHash(&bind.CallOpts{Context: ctx}, headerBytes, tipsBytes, body)
	Nil(t, err)

	decodedMessageLeaf, err := decodedMessage.ToLeaf()
	Nil(t, err)

	Equal(t, decodedMessage.OriginDomain(), origin)
	Equal(t, decodedMessage.Sender(), sender)
	Equal(t, decodedMessage.Nonce(), nonce)
	Equal(t, decodedMessage.DestinationDomain(), destination)
	Equal(t, decodedMessage.Body(), body)
	Equal(t, messageLeaf, decodedMessageLeaf)
	Equal(t, messageLeaf, testMessageLeaf)
}

func TestDispatchMessageParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()
	chainID := gofakeit.Uint32()
	deployManager := testutil.NewDeployManager(t)
	simulatedChain := geth.NewEmbeddedBackendForChainID(ctx, t, big.NewInt(int64(chainID)))

	originContract, originRef := deployManager.GetOrigin(ctx, simulatedChain)
	transactOpts := simulatedChain.GetTxContext(ctx, nil)

	dispatchSink := make(chan *origin.OriginDispatch)
	sub, err := originRef.WatchDispatch(&bind.WatchOpts{Context: ctx}, dispatchSink, [][32]byte{}, []uint32{}, []uint32{})
	Nil(t, err)

	destination := chainID + 1
	recipient := [32]byte{byte(gofakeit.Uint32())}
	optimisticSeconds := gofakeit.Uint32()
	notaryTip := big.NewInt(int64(int(gofakeit.Uint32())))
	broadcasterTip := big.NewInt(int64(int(gofakeit.Uint32())))
	proverTip := big.NewInt(int64(int(gofakeit.Uint32())))
	executorTip := big.NewInt(int64(int(gofakeit.Uint32())))
	tips := types.NewTips(notaryTip, broadcasterTip, proverTip, executorTip)
	encodedTips, err := types.EncodeTips(tips)
	Nil(t, err)
	messageBytes := []byte{byte(gofakeit.Uint32())}

	transactOpts.Value = types.TotalTips(tips)

	tx, err := originRef.Dispatch(transactOpts.TransactOpts, destination, recipient, optimisticSeconds, encodedTips, messageBytes)
	Nil(t, err)
	simulatedChain.WaitForConfirmation(ctx, tx)

	sender, err := simulatedChain.Signer().Sender(tx)
	Nil(t, err)

	// create the agents type message
	testHeader := types.NewHeader(chainID, sender.Hash(), uint32(tx.Nonce()+1), destination, recipient, optimisticSeconds)
	testMessage := types.NewMessage(testHeader, tips, messageBytes)
	testMessageLeaf, err := testMessage.ToLeaf()
	Nil(t, err)

	testMessageVersion := testMessage.Version()
	testMessageHeader := testMessage.Header()
	testMessageHeaderVersion := testMessageHeader.Version()
	testMessageHeaderOrigin := testMessageHeader.OriginDomain()
	testMessageHeaderSender := testMessageHeader.Sender()
	testMessageHeaderNonce := testMessageHeader.Nonce()
	testMessageHeaderDestination := testMessageHeader.DestinationDomain()
	testMessageHeaderRecipient := testMessageHeader.Recipient()
	testMessageHeaderOptimisticSeconds := testMessageHeader.OptimisticSeconds()
	testMessageTips := testMessage.Tips()
	testMessageBody := testMessage.Body()
	testMessageOriginDomain := testMessage.OriginDomain()
	testMessageSender := testMessage.Sender()
	testMessageNonce := testMessage.Nonce()
	testMessageDestinationDomain := testMessage.DestinationDomain()
	testMessageRecipient := testMessage.Recipient()
	testMessageOptimisticSeconds := testMessage.OptimisticSeconds()

	watchCtx, cancelWatch := context.WithTimeout(ctx, time.Second*30)
	defer cancelWatch()

	select {
	// check for errors and fail
	case <-watchCtx.Done():
		t.Error(t, fmt.Errorf("test context completed %w", ctx.Err()))
	case <-sub.Err():
		t.Error(t, sub.Err())
	// get dispatch event
	case item := <-dispatchSink:
		parser, err := origin.NewParser(originContract.Address())
		Nil(t, err)

		committedMessage, ok := parser.ParseDispatch(item.Raw)
		True(t, ok)
		message, err := types.DecodeMessage(committedMessage.Message())
		Nil(t, err)

		messageLeaf, err := message.ToLeaf()
		Nil(t, err)

		messageVersion := message.Version()
		messageHeader := message.Header()
		messageHeaderVersion := messageHeader.Version()
		messageHeaderOrigin := messageHeader.OriginDomain()
		messageHeaderSender := messageHeader.Sender()
		messageHeaderNonce := messageHeader.Nonce()
		messageHeaderDestination := messageHeader.DestinationDomain()
		messageHeaderRecipient := messageHeader.Recipient()
		messageHeaderOptimisticSeconds := messageHeader.OptimisticSeconds()
		messageTips := message.Tips()
		messageBody := message.Body()
		messageOriginDomain := message.OriginDomain()
		messageSender := message.Sender()
		messageNonce := message.Nonce()
		messageDestinationDomain := message.DestinationDomain()
		messageRecipient := message.Recipient()
		messageOptimisticSeconds := message.OptimisticSeconds()

		Equal(t, messageLeaf, testMessageLeaf)
		Equal(t, messageVersion, testMessageVersion)
		Equal(t, messageHeader, testMessageHeader)
		Equal(t, messageHeaderVersion, testMessageHeaderVersion)
		Equal(t, messageHeaderOrigin, testMessageHeaderOrigin)
		Equal(t, messageHeaderSender, testMessageHeaderSender)
		Equal(t, messageHeaderNonce, testMessageHeaderNonce)
		Equal(t, messageHeaderDestination, testMessageHeaderDestination)
		Equal(t, messageHeaderRecipient, testMessageHeaderRecipient)
		Equal(t, messageHeaderOptimisticSeconds, testMessageHeaderOptimisticSeconds)
		Equal(t, messageTips, testMessageTips)
		Equal(t, messageBody, testMessageBody)
		Equal(t, messageOriginDomain, testMessageOriginDomain)
		Equal(t, messageSender, testMessageSender)
		Equal(t, messageNonce, testMessageNonce)
		Equal(t, messageDestinationDomain, testMessageDestinationDomain)
		Equal(t, messageRecipient, testMessageRecipient)
		Equal(t, messageOptimisticSeconds, testMessageOptimisticSeconds)

		break
	}
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
	attestKeyFromRaw := types.NewAttestionKey(rawKey)
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
