package types

import "math/big"

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

// ChainGassesToSnapGas converts a slice of ChainGas to a slice of big.Int.
func ChainGassesToSnapGas(chainGasses []ChainGas) (snapGasses []*big.Int, err error) {
	snapGasses = make([]*big.Int, len(chainGasses))
	for i, cg := range chainGasses {
		snapGas, err := EncodeChainGas(cg)
		if err != nil {
			return nil, err
		}

		snapGasses[i] = new(big.Int).SetBytes(snapGas)
	}

	return snapGasses, nil
}

var _ ChainGas = chainGas{}
