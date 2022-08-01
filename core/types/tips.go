package types

import (
	"github.com/synapsecns/synapse-node/pkg/common"
	"math/big"
)

// Tips contain tips used for scientizing different agents.
type Tips interface {
	// UpdaterTip gets the tips for the updater
	UpdaterTip() *big.Int
	// RelayerTip gets the tips for the updater
	RelayerTip() *big.Int
	// ProverTip gets the tips for the prover
	ProverTip() *big.Int
	// ProcessorTip gets the tips for the processor
	ProcessorTip() *big.Int
}

func NewTips(updaterTip, relayerTip, proverTip, processorTip *big.Int) Tips {
	return tips{
		updaterTip:   updaterTip,
		relayerTip:   relayerTip,
		proverTip:    proverTip,
		processorTip: processorTip,
	}
}

// tips implements Tips.
type tips struct {
	updaterTip, relayerTip, proverTip, processorTip *big.Int
}

func (t tips) UpdaterTip() *big.Int {
	return common.CopyBigInt(t.updaterTip)
}

func (t tips) RelayerTip() *big.Int {
	return common.CopyBigInt(t.relayerTip)
}

func (t tips) ProverTip() *big.Int {
	return common.CopyBigInt(t.proverTip)
}

func (t tips) ProcessorTip() *big.Int {
	return common.CopyBigInt(t.processorTip)
}

var _ Tips = tips{}

// TotalTips gets the combined value of the tips
func TotalTips(tips Tips) *big.Int {
	vals := []*big.Int{tips.UpdaterTip(), tips.ProcessorTip(), tips.RelayerTip(), tips.ProverTip()}
	total := new(big.Int)

	for _, val := range vals {
		total = new(big.Int).Add(total, val)
	}

	return total
}
