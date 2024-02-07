package base_test

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/sql/base"
	"math/big"
	"testing"
)

func TestDBStateToState(t *testing.T) {
	snapshotRoot := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
	root := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
	chainID := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	originBlockNumber := gofakeit.Uint64()
	originTimestamp := gofakeit.Uint64()
	proof := []string{common.BigToHash(big.NewInt(gofakeit.Int64())).String(), common.BigToHash(big.NewInt(gofakeit.Int64())).String()}
	stateIndex := gofakeit.Uint32()
	blockNumber := gofakeit.Uint64()

	gasPrice := gofakeit.Uint16()
	dataPrice := gofakeit.Uint16()
	execBuffer := gofakeit.Uint16()
	amortAttCost := gofakeit.Uint16()
	etherPrice := gofakeit.Uint16()
	markup := gofakeit.Uint16()

	proofJSON, err := json.Marshal(proof)
	if err != nil {
		panic(err)
	}

	initialDBState := db.DBState{
		SnapshotRoot:      &snapshotRoot,
		Root:              &root,
		ChainID:           &chainID,
		Nonce:             &nonce,
		OriginBlockNumber: &originBlockNumber,
		OriginTimestamp:   &originTimestamp,
		Proof:             (*json.RawMessage)(&proofJSON),
		StateIndex:        &stateIndex,
		BlockNumber:       &blockNumber,
		GDGasPrice:        &gasPrice,
		GDDataPrice:       &dataPrice,
		GDExecBuffer:      &execBuffer,
		GDAmortAttCost:    &amortAttCost,
		GDEtherPrice:      &etherPrice,
		GDMarkup:          &markup,
	}

	initialState := base.DBStateToState(initialDBState)

	finalDBState := base.StateToDBState(initialState)

	finalState := base.DBStateToState(finalDBState)

	assert.Equal(t, initialDBState, finalDBState)
	assert.Equal(t, initialState, finalState)
}
