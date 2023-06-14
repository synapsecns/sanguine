package types

const (
	chainGasOffsetDomain  = 0
	chainGasOffsetGasData = 4
	chainGasSize          = 16
)

// ChainGas is the ChainGas interface.
type ChainGas interface {
	// GasData is the gas data for the chain.
	GasData() GasData
	// Domain is domain id of the chain.
	Domain() uint32
}

type chainGas struct {
	gasData GasData
	domain  uint32
}

// NewChainGas creates a new chaingas.
func NewChainGas(gasData GasData, domain uint32) ChainGas {
	return &chainGas{
		gasData: gasData,
		domain:  domain,
	}
}

func (g chainGas) GasData() GasData {
	return g.gasData
}

func (g chainGas) Domain() uint32 {
	return g.domain
}

var _ ChainGas = chainGas{}
