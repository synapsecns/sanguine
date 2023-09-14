package domains

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// DisputeStatus is the return type from DisputeStatus.
type DisputeStatus struct {
	// DisputeFlag is the 0: None; 1: Pending; 2: Slashed.
	DisputeFlag uint8
	// Rival is the address of the rival.
	Rival common.Address
	// FraudProver is the address of the fraud prover.
	FraudProver common.Address
	// DisputePtr is the value of the dispute pointer.
	DisputePtr *big.Int
}
