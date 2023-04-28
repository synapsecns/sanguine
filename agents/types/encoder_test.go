package types_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/types"
)

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

func TestNewMessageEncodeDecode(t *testing.T) {
	origin := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	body := []byte(gofakeit.Sentence(gofakeit.Number(5, 15)))
	optimisticSeconds := gofakeit.Uint32()

	header := types.NewHeader(
		origin, nonce, destination, optimisticSeconds)

	flag := types.MessageFlagBase

	newMessage := types.NewMessage(flag, header, body)

	Equal(t, newMessage.Flag(), flag)
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

	ogHeader, err := types.EncodeHeader(types.NewHeader(origin, nonce, destination, optimisticSeconds))
	Nil(t, err)

	decodedHeader, err := types.DecodeHeader(ogHeader)
	Nil(t, err)

	Equal(t, decodedHeader.OriginDomain(), origin)
	Equal(t, decodedHeader.Nonce(), nonce)
	Equal(t, decodedHeader.DestinationDomain(), destination)
	Equal(t, decodedHeader.OptimisticSeconds(), optimisticSeconds)
}
