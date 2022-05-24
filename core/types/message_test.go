package types_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/types"
	"math/big"
	"testing"
)

func TestNewMessage(t *testing.T) {
	origin := gofakeit.Uint32()
	sender := common.BigToHash(big.NewInt(gofakeit.Int64()))
	nonce := gofakeit.Uint32()
	destination := common.BigToHash(big.NewInt(gofakeit.Int64()))
	body := []byte(gofakeit.Sentence(gofakeit.Number(5, 15)))

	newMessage := types.NewMessage(origin, sender, nonce, destination, body)

	Equal(t, newMessage.Origin(), origin)
	Equal(t, newMessage.Sender(), sender)
	Equal(t, newMessage.Nonce(), nonce)
	Equal(t, newMessage.Destination(), destination)
	Equal(t, newMessage.Body(), body)
}

func TestNewCommittedMessage(t *testing.T) {
	leafIndex := gofakeit.Uint32()
	committedRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))
	message := []byte(gofakeit.Sentence(gofakeit.Number(5, 15)))

	committedMessage := types.NewCommittedMessage(leafIndex, committedRoot, message)

	Equal(t, leafIndex, committedMessage.LeafIndex())
	Equal(t, committedRoot, committedMessage.CommitedRoot())
	Equal(t, message, committedMessage.Message())
}
