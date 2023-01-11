package executor_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	execTypes "github.com/synapsecns/sanguine/agents/agents/executor/types"
	"testing"
)

func TestBinarySearchAttestationForNonce(t *testing.T) {
	one := uint32(1)
	two := uint32(2)
	four := uint32(4)
	five := uint32(5)
	attestations := []execTypes.DBAttestation{
		{
			Nonce: &one,
		},
		{
			Nonce: &two,
		},
		{
			Nonce: &four,
		},
		{
			Nonce: &five,
		},
	}

	assert.Equal(t, &one, executor.BinarySearchAttestationsForNonce(attestations, 1).Nonce)
	assert.Equal(t, &two, executor.BinarySearchAttestationsForNonce(attestations, 2).Nonce)
	assert.Equal(t, &four, executor.BinarySearchAttestationsForNonce(attestations, 3).Nonce)
	assert.Equal(t, &four, executor.BinarySearchAttestationsForNonce(attestations, 4).Nonce)
	assert.Equal(t, &five, executor.BinarySearchAttestationsForNonce(attestations, 5).Nonce)
	assert.Nil(t, executor.BinarySearchAttestationsForNonce(attestations, 6))
}
