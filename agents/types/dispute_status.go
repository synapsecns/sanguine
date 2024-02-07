package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core"
)

// DisputeStatus is the dispute status interface.
type DisputeStatus interface {
	// Flag is the current status flag of the dispute.
	Flag() DisputeFlagType
	// Rival is the address of the rival.
	Rival() common.Address
	// FraudProver is the address of the fraud prover for this dispute.
	FraudProver() common.Address
	// DisputePtr is the index of the dispute.
	DisputePtr() *big.Int
}

type disputeStatus struct {
	flag        DisputeFlagType
	rival       common.Address
	fraudProver common.Address
	disputePtr  *big.Int
}

// NewDisputeStatus creates a new dispute status.
func NewDisputeStatus(flag DisputeFlagType, rival, fraudProver common.Address, disputePtr *big.Int) DisputeStatus {
	return &disputeStatus{
		flag:        flag,
		rival:       rival,
		fraudProver: fraudProver,
		disputePtr:  disputePtr,
	}
}

func (s disputeStatus) Flag() DisputeFlagType {
	return s.flag
}

func (s disputeStatus) Rival() common.Address {
	return s.rival
}

func (s disputeStatus) FraudProver() common.Address {
	return s.fraudProver
}

func (s disputeStatus) DisputePtr() *big.Int {
	return core.CopyBigInt(s.disputePtr)
}

var _ DisputeStatus = disputeStatus{}
