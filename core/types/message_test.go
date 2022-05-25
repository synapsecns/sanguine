package types_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/types"
	"math/big"
	"testing"
)

func TestNewMessageEncodeDecode(t *testing.T) {
	origin := gofakeit.Uint32()
	sender := common.BigToHash(big.NewInt(gofakeit.Int64()))
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	body := []byte(gofakeit.Sentence(gofakeit.Number(5, 15)))
	recipient := common.BigToHash(big.NewInt(gofakeit.Int64()))

	newMessage := types.NewMessage(origin, sender, nonce, destination, body, recipient)

	Equal(t, newMessage.Origin(), origin)
	Equal(t, newMessage.Sender(), sender)
	Equal(t, newMessage.Nonce(), nonce)
	Equal(t, newMessage.Destination(), destination)
	Equal(t, newMessage.Body(), body)

	encodedMessage, err := newMessage.Encode()
	Nil(t, err)

	// make sure decode is same as encode
	decodedMessage, err := types.DecodeMessage(encodedMessage)
	Nil(t, err)

	Equal(t, newMessage.Origin(), decodedMessage.Origin())
	Equal(t, newMessage.Sender(), decodedMessage.Sender())
	Equal(t, newMessage.Nonce(), decodedMessage.Nonce())
	Equal(t, newMessage.Destination(), decodedMessage.Destination())
	Equal(t, newMessage.Body(), decodedMessage.Body())
}

func TestNewCommittedMessageEncodeDecode(t *testing.T) {
	leafIndex := gofakeit.Uint32()
	committedRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))
	message := []byte(gofakeit.Sentence(gofakeit.Number(5, 15)))

	committedMessage := types.NewCommittedMessage(leafIndex, committedRoot, message)

	Equal(t, leafIndex, committedMessage.LeafIndex())
	Equal(t, committedRoot, committedMessage.CommitedRoot())
	Equal(t, message, committedMessage.Message())

	encodedMessage, err := committedMessage.Encode()
	Nil(t, err)

	decodedMessage, err := types.DecodeCommittedMessage(encodedMessage)
	Nil(t, err)

	Equal(t, decodedMessage.Message(), committedMessage.Message())
	Equal(t, decodedMessage.CommitedRoot(), committedMessage.CommitedRoot())
	Equal(t, decodedMessage.Leaf(), committedMessage.Leaf())
	Equal(t, decodedMessage.LeafIndex(), committedMessage.LeafIndex())
}
