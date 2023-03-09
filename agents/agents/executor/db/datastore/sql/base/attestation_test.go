package base_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/base"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	"math/big"
	"testing"
)

func TestAttestationDBAttestationParity(t *testing.T) {
	chainID := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	root := common.BigToHash(big.NewInt(gofakeit.Int64()))
	blockNumberDestination := gofakeit.Uint64()
	destinationBlockTime := gofakeit.Uint64()

	initialDBAttestation := types.DBAttestation{
		ChainID:                &chainID,
		Destination:            &destination,
		Nonce:                  &nonce,
		Root:                   &root,
		DestinationBlockNumber: &blockNumberDestination,
		DestinationTimestamp:   &destinationBlockTime,
	}

	initialAttestation := base.DBAttestationToAttestation(initialDBAttestation)

	finalDBAttestation := base.AttestationToDBAttestation(initialAttestation)

	finalAttestation := base.DBAttestationToAttestation(finalDBAttestation)

	Equal(t, initialDBAttestation, finalDBAttestation)
	Equal(t, initialAttestation, finalAttestation)
}
