package london

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/synapsecns/sanguine/ethergo/chain/gas/backend"
	"math/big"
	"sync"
)

// FeeOracle is the fee oracle to use for v2 gas price calculations.
type FeeOracle struct {
	// height is the height used to determine the basefee
	height int
	// oracleBackend is the oracle backend oracleBackend
	oracleBackend HeightOracleBackend
	// oracle is the oracke
	oracle *gasprice.Oracle
	// structMux is the struct mutex
	structMux sync.Mutex
	// tipCap is the tip cap for the height. Since we have two methods
	// and both require a tip cap calulation, we use these as a kind of cache
	tipCap *big.Int
	// feeCap is the fee cap for the height. See above for details
	feeCap *big.Int
}

// NewFeeOracle creates a new fee oracle.
func NewFeeOracle(chain backend.OracleBackendChain, height uint64, config gasprice.Config) FeeOracle {
	oracleBackend := NewOracleBackendFromHeight(chain, height)
	return FeeOracle{
		height:        int(height),
		oracleBackend: oracleBackend,
		oracle:        gasprice.NewOracle(NewOracleBackendFromHeight(chain, height), config),
	}
}

// calculate calculates the tipCap and fee cap and sets them if not already set.
func (f *FeeOracle) calculate(ctx context.Context) (err error) {
	f.structMux.Lock()
	defer f.structMux.Unlock()
	// these are only both set on successful calculation
	if f.feeCap == nil || f.tipCap == nil {
		f.tipCap, err = f.oracle.SuggestTipCap(ctx)
		if err != nil {
			return fmt.Errorf("could not calculate tip cap: %w", err)
		}

		block, err := f.oracleBackend.BlockByNumber(ctx, rpc.BlockNumber(f.height))
		if err != nil {
			return fmt.Errorf("could not get block %d: %w", f.height, err)
		}
		f.feeCap = big.NewInt(0).Add(block.BaseFee(), f.tipCap)
	}
	return nil
}

// SuggestTipCap gets the suggested tip cap in a deterministic manner.
func (f *FeeOracle) SuggestTipCap(ctx context.Context) (*big.Int, error) {
	err := f.calculate(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get tip cap: %w", err)
	}
	return f.tipCap, nil
}

// SuggestFeeCap calculates the suggested fee cap in a deterministic manner.
// because this is done from gas block, it is advisable to bump this number by a percentage.
func (f *FeeOracle) SuggestFeeCap(ctx context.Context) (*big.Int, error) {
	err := f.calculate(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get tip cap: %w", err)
	}
	return f.feeCap, nil
}
