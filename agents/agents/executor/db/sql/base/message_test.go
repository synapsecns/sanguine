package base_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/sql/base"
	"math/big"
	"testing"
)

func TestMessageDBMessageParity(t *testing.T) {
	chainID := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	message := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
	blockNumber := gofakeit.Uint64()
	executed := gofakeit.Bool()
	minimumTimeSet := gofakeit.Bool()
	minimumTime := gofakeit.Uint64()
	initialDBMessage := db.DBMessage{
		ChainID:        &chainID,
		Destination:    &destination,
		Nonce:          &nonce,
		Message:        &message,
		BlockNumber:    &blockNumber,
		Executed:       &executed,
		MinimumTimeSet: &minimumTimeSet,
		MinimumTime:    &minimumTime,
	}

	initialMessage := base.DBMessageToMessage(initialDBMessage)

	finalDBMessage := base.MessageToDBMessage(initialMessage)

	finalMessage := base.DBMessageToMessage(finalDBMessage)

	Equal(t, initialDBMessage, finalDBMessage)
	Equal(t, initialMessage, finalMessage)
}
