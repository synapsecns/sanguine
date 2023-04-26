package types

import (
	"github.com/synapsecns/sanguine/core"
	"math/big"
)

// Tips contain tips used for scientizing different agents.
type Tips interface {
	// NotaryTip gets the tips for the notary
	NotaryTip() *big.Int
	// BroadcasterTip gets the tips for the broadcaster
	BroadcasterTip() *big.Int
	// ProverTip gets the tips for the prover
	ProverTip() *big.Int
	// ExecutorTip gets the tips for the executor
	ExecutorTip() *big.Int
}

// NewTips creates a new tips type.
func NewTips(notaryTip, broadcasterTip, proverTip, executorTip *big.Int) Tips {
	return tips{
		notaryTip:      notaryTip,
		broadcasterTip: broadcasterTip,
		proverTip:      proverTip,
		executorTip:    executorTip,
	}
}

// tips implements Tips.
type tips struct {
	notaryTip, broadcasterTip, proverTip, executorTip *big.Int
}

func (t tips) NotaryTip() *big.Int {
	return core.CopyBigInt(t.notaryTip)
}

func (t tips) BroadcasterTip() *big.Int {
	return core.CopyBigInt(t.broadcasterTip)
}

func (t tips) ProverTip() *big.Int {
	return core.CopyBigInt(t.proverTip)
}

func (t tips) ExecutorTip() *big.Int {
	return core.CopyBigInt(t.executorTip)
}

var _ Tips = tips{}

// TotalTips gets the combined value of the tips.
func TotalTips(tips Tips) *big.Int {
	vals := []*big.Int{tips.NotaryTip(), tips.ExecutorTip(), tips.BroadcasterTip(), tips.ProverTip()}
	total := new(big.Int)

	for _, val := range vals {
		total = new(big.Int).Add(total, val)
	}

	return total
}
