package types

import (
	"github.com/synapsecns/synapse-node/pkg/common"
	"math/big"
)

// Tips contain tips used for scientizing different agents.
type Tips interface {
	// Version gets the version of the tips header
	Version() uint16
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
		version:        tipsVersion,
		notaryTip:      notaryTip,
		broadcasterTip: broadcasterTip,
		proverTip:      proverTip,
		executorTip:    executorTip,
	}
}

// tips implements Tips.
type tips struct {
	version                                           uint16
	notaryTip, broadcasterTip, proverTip, executorTip *big.Int
}

func (t tips) Version() uint16 {
	return t.version
}

func (t tips) NotaryTip() *big.Int {
	return common.CopyBigInt(t.notaryTip)
}

func (t tips) BroadcasterTip() *big.Int {
	return common.CopyBigInt(t.broadcasterTip)
}

func (t tips) ProverTip() *big.Int {
	return common.CopyBigInt(t.proverTip)
}

func (t tips) ExecutorTip() *big.Int {
	return common.CopyBigInt(t.executorTip)
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
