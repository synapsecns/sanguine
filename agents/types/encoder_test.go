package types_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/types"
)

func TestRequestEncodeDecode(t *testing.T) {
	version := uint32(1)
	gasLimit := gofakeit.Uint64()
	gasDrop := randomUint64BigInt(t)

	ogRequest, err := types.EncodeRequest(types.NewRequest(version, gasLimit, gasDrop))
	Nil(t, err)

	decodedRequest := types.DecodeRequest(ogRequest)

	Equal(t, decodedRequest.Version(), version)
	Equal(t, decodedRequest.GasLimit(), gasLimit)
	Equal(t, decodedRequest.GasDrop(), gasDrop)
}

func TestEncodeDecodeTips(t *testing.T) {
	// we want to make sure we can deal w/ overflows
	summitTip := randomUint64BigInt(t)
	attestationTip := randomUint64BigInt(t)
	executionTip := randomUint64BigInt(t)
	deliveryTip := randomUint64BigInt(t)

	encodedTips, err := types.EncodeTips(types.NewTips(summitTip, attestationTip, executionTip, deliveryTip))
	Nil(t, err)

	decodedTips, err := types.DecodeTips(encodedTips)
	Nil(t, err)

	Equal(t, decodedTips.SummitTip(), summitTip)
	Equal(t, decodedTips.AttestationTip(), attestationTip)
	Equal(t, decodedTips.ExecutionTip(), executionTip)
	Equal(t, decodedTips.DeliveryTip(), deliveryTip)
}

func TestEncodeDecodeBaseMessage(t *testing.T) {
	senderBytes := randomUint64BigInt(t).Bytes()
	var sender [32]byte
	copy(sender[:], senderBytes)

	recipientBytes := randomUint64BigInt(t).Bytes()
	var recipient [32]byte
	copy(recipient[:], recipientBytes)

	summitTip := randomUint64BigInt(t)
	attestationTip := randomUint64BigInt(t)
	executionTip := randomUint64BigInt(t)
	deliveryTip := randomUint64BigInt(t)

	tips := types.NewTips(summitTip, attestationTip, executionTip, deliveryTip)

	version := uint32(1)
	gasLimit := gofakeit.Uint64()
	gasDrop := randomUint64BigInt(t)

	request := types.NewRequest(version, gasLimit, gasDrop)

	content := randomUint64BigInt(t).Bytes()

	encodedBaseMessage, err := types.EncodeBaseMessage(types.NewBaseMessage(sender, recipient, tips, request, content))
	Nil(t, err)

	decodedBaseMessage, err := types.DecodeBaseMessage(encodedBaseMessage)
	Nil(t, err)

	Equal(t, decodedBaseMessage.Sender(), sender)
	Equal(t, decodedBaseMessage.Recipient(), recipient)
	Equal(t, decodedBaseMessage.Tips().SummitTip().Uint64(), tips.SummitTip().Uint64())
	Equal(t, decodedBaseMessage.Tips().AttestationTip().Uint64(), tips.AttestationTip().Uint64())
	Equal(t, decodedBaseMessage.Tips().DeliveryTip().Uint64(), tips.DeliveryTip().Uint64())
	Equal(t, decodedBaseMessage.Tips().ExecutionTip().Uint64(), tips.ExecutionTip().Uint64())
	Equal(t, decodedBaseMessage.Request().Version(), request.Version())
	Equal(t, decodedBaseMessage.Request().GasLimit(), request.GasLimit())
	Equal(t, decodedBaseMessage.Request().GasDrop().Uint64(), request.GasDrop().Uint64())
	Equal(t, decodedBaseMessage.Content(), content)
}

func TestNewMessageEncodeDecode(t *testing.T) {
	origin := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	body := []byte(gofakeit.Sentence(gofakeit.Number(5, 15)))
	optimisticSeconds := gofakeit.Uint32()

	flag := types.MessageFlagManager

	header := types.NewHeader(
		flag, origin, nonce, destination, optimisticSeconds)

	newMessage := types.NewMessage(header, nil, body)

	Equal(t, newMessage.Header().Flag(), flag)
	Equal(t, newMessage.OriginDomain(), origin)
	Equal(t, newMessage.Nonce(), nonce)
	Equal(t, newMessage.DestinationDomain(), destination)
	Equal(t, newMessage.Body(), body)

	encodedMessage, err := types.EncodeMessage(newMessage)
	Nil(t, err)

	// make sure decode is same as encode
	decodedMessage, err := types.DecodeMessage(encodedMessage)
	Nil(t, err)

	Equal(t, newMessage.OriginDomain(), decodedMessage.OriginDomain())
	Equal(t, newMessage.Nonce(), decodedMessage.Nonce())
	Equal(t, newMessage.DestinationDomain(), decodedMessage.DestinationDomain())
	Equal(t, newMessage.Body(), decodedMessage.Body())
}

func TestHeaderEncodeDecode(t *testing.T) {
	origin := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	optimisticSeconds := gofakeit.Uint32()

	flag := types.MessageFlagManager
	ogHeader, err := types.EncodeHeader(types.NewHeader(flag, origin, nonce, destination, optimisticSeconds))
	Nil(t, err)

	decodedHeader, err := types.DecodeHeader(ogHeader)
	Nil(t, err)

	Equal(t, decodedHeader.Flag(), flag)
	Equal(t, decodedHeader.OriginDomain(), origin)
	Equal(t, decodedHeader.Nonce(), nonce)
	Equal(t, decodedHeader.DestinationDomain(), destination)
	Equal(t, decodedHeader.OptimisticSeconds(), optimisticSeconds)
}

func TestChainGasEncodeDecode(t *testing.T) {
	domain := gofakeit.Uint32()

	gasPrice := gofakeit.Uint16()
	dataPrice := gofakeit.Uint16()
	execBuffer := gofakeit.Uint16()
	amortAttCost := gofakeit.Uint16()
	etherPrice := gofakeit.Uint16()
	markup := gofakeit.Uint16()

	gasData := types.NewGasData(gasPrice, dataPrice, execBuffer, amortAttCost, etherPrice, markup)

	chainGas := types.NewChainGas(gasData, domain)

	encodedChainGas, err := types.EncodeChainGas(chainGas)
	Nil(t, err)

	decodedChainGas, err := types.DecodeChainGas(encodedChainGas)
	Nil(t, err)

	Equal(t, chainGas.Domain(), decodedChainGas.Domain())
	Equal(t, chainGas.GasData().GasPrice(), decodedChainGas.GasData().GasPrice())
	Equal(t, chainGas.GasData().DataPrice(), decodedChainGas.GasData().DataPrice())
	Equal(t, chainGas.GasData().DataPrice(), decodedChainGas.GasData().DataPrice())
	Equal(t, chainGas.GasData().ExecBuffer(), decodedChainGas.GasData().ExecBuffer())
	Equal(t, chainGas.GasData().AmortAttCost(), decodedChainGas.GasData().AmortAttCost())
	Equal(t, chainGas.GasData().EtherPrice(), decodedChainGas.GasData().EtherPrice())
	Equal(t, chainGas.GasData().Markup(), decodedChainGas.GasData().Markup())
}
