package types

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

const (
	receiptOffsetOrigin        = 0
	receiptOffsetDestination   = 4
	receiptOffsetMessageHash   = 8
	receiptOffsetSnapshotRoot  = 40
	receiptOffsetStateIndex    = 72
	receiptOffsetAttNotary     = 73
	receiptOffsetFirstExecutor = 93
	receiptOffsetFinalExecutor = 113
	receiptSize                = 133
)

// Receipt is the receipt interface.
type Receipt interface {
	Encoder
	// Origin is the origin of the receipt.
	Origin() uint32
	// Destination is the destination of the receipt.
	Destination() uint32
	// MessageHash is the hash of the message.
	MessageHash() [32]byte
	// SnapshotRoot is the root of the Snapshot Merkle Tree.
	SnapshotRoot() [32]byte
	// StateIndex is the state index of the receipt.
	StateIndex() uint8
	// AttestationNotary is the notary of the attestation.
	AttestationNotary() common.Address
	// FirstExecutor is the first executor of the receipt.
	FirstExecutor() common.Address
	// FinalExecutor is the final executor of the receipt.
	FinalExecutor() common.Address
	// SignReceipt returns the signature of the receipt payload signed by the signer.
	SignReceipt(ctx context.Context, signer signer.Signer, valid bool) (signer.Signature, []byte, common.Hash, error)
}

type receipt struct {
	origin            uint32
	destination       uint32
	messageHash       [32]byte
	snapshotRoot      [32]byte
	stateIndex        uint8
	attestationNotary common.Address
	firstExecutor     common.Address
	finalExecutor     common.Address
}

// NewReceipt creates a new receipt.
func NewReceipt(origin uint32, destination uint32, messageHash [32]byte, snapshotRoot [32]byte, stateIndex uint8, attestationNotary common.Address, firstExecutor common.Address, finalExecutor common.Address) Receipt {
	return &receipt{
		origin:            origin,
		destination:       destination,
		messageHash:       messageHash,
		snapshotRoot:      snapshotRoot,
		stateIndex:        stateIndex,
		attestationNotary: attestationNotary,
		firstExecutor:     firstExecutor,
		finalExecutor:     finalExecutor,
	}
}

func (r receipt) Origin() uint32 {
	return r.origin
}

func (r receipt) Destination() uint32 {
	return r.destination
}

func (r receipt) MessageHash() [32]byte {
	return r.messageHash
}

func (r receipt) SnapshotRoot() [32]byte {
	return r.snapshotRoot
}

func (r receipt) StateIndex() uint8 {
	return r.stateIndex
}

func (r receipt) AttestationNotary() common.Address {
	return r.attestationNotary
}

func (r receipt) FirstExecutor() common.Address {
	return r.firstExecutor
}

func (r receipt) FinalExecutor() common.Address {
	return r.finalExecutor
}

func (r receipt) SignReceipt(ctx context.Context, signer signer.Signer, valid bool) (signer.Signature, []byte, common.Hash, error) {
	var receiptSalt string
	if valid {
		receiptSalt = ReceiptValidSalt
	} else {
		receiptSalt = ReceiptInvalidSalt
	}
	return signEncoder(ctx, signer, r, receiptSalt)
}
