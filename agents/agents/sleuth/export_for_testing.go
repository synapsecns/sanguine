package sleuth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/types"
)

// CheckMessage gets the message to investigate.
func (s *Sleuth) CheckMessage(ctx context.Context, txHash common.Hash, chainID uint32) (*types.Message, error) {
	return s.checkMessage(ctx, txHash, chainID)
}

// CheckState sees if a message has been included in a state.
func (s *Sleuth) CheckState(ctx context.Context, nonce, chainID uint32) (bool, error) {
	return s.checkState(ctx, nonce, chainID)
}

// CheckSnapshot sees if a message has been included in a snapshot.
func (s *Sleuth) CheckSnapshot(ctx context.Context, nonce, chainID uint32) (*[][32]byte, error) {
	return s.checkSnapshot(ctx, nonce, chainID)
}

// CheckAttestation sees if a message has been involved in an attestation.
func (s *Sleuth) CheckAttestation(ctx context.Context, snapshotRoots [][32]byte, destinationDomain uint32) (bool, error) {
	return s.checkAttestation(ctx, snapshotRoots, destinationDomain)
}

// CheckExecuted sees if a message has been executed.
func (s *Sleuth) CheckExecuted(ctx context.Context, messageHash [32]byte, destinationDomain uint32) (bool, error) {
	return s.checkExecuted(ctx, messageHash, destinationDomain)
}
