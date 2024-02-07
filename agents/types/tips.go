package types

import (
	"github.com/synapsecns/sanguine/core"
	"math/big"
)

const (
	// TipsSize is the size of the tips in bytes.
	TipsSize = 8 * 4
	// ShiftSummitTip is the shift for the summit tip.
	ShiftSummitTip = 24 * 8
	// ShiftAttestationTip is the shift for the attestation tip.
	ShiftAttestationTip = 16 * 8
	// ShiftExecutionTip is the shift for the execution tip.
	ShiftExecutionTip = 8 * 8
)

// Tips contain tips used for scientizing different agents.
type Tips interface {
	// SummitTip gets the tips for the agent work on summit
	SummitTip() *big.Int
	// AttestationTip gets the tips for the doing the attestation
	AttestationTip() *big.Int
	// ExecutionTip gets the tips for executing the message
	ExecutionTip() *big.Int
	// DeliveryTip gets the tips for delivering the message receipt to summit
	DeliveryTip() *big.Int
}

// NewTips creates a new tips type.
func NewTips(summitTip, attestationTip, executionTip, deliveryTip *big.Int) Tips {
	return tips{
		summitTip:      summitTip,
		attestationTip: attestationTip,
		executionTip:   executionTip,
		deliveryTip:    deliveryTip,
	}
}

// tips implements Tips.
type tips struct {
	summitTip, attestationTip, executionTip, deliveryTip *big.Int
}

func (t tips) SummitTip() *big.Int {
	return core.CopyBigInt(t.summitTip)
}

func (t tips) AttestationTip() *big.Int {
	return core.CopyBigInt(t.attestationTip)
}

func (t tips) ExecutionTip() *big.Int {
	return core.CopyBigInt(t.executionTip)
}

func (t tips) DeliveryTip() *big.Int {
	return core.CopyBigInt(t.deliveryTip)
}

var _ Tips = tips{}

// TotalTips gets the combined value of the tips.
func TotalTips(tips Tips) *big.Int {
	vals := []*big.Int{tips.SummitTip(), tips.AttestationTip(), tips.ExecutionTip(), tips.DeliveryTip()}
	total := new(big.Int)

	for _, val := range vals {
		total = new(big.Int).Add(total, val)
	}

	return total
}
