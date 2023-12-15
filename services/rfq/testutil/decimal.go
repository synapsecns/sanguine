package testutil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"math/big"
	"reflect"
	"testing"
)

// MustAdjustAmount multiplies a token by the number of decimals in the amount.
// see AdjustAmount.
func MustAdjustAmount(ctx context.Context, tb testing.TB, amount *big.Int, handler interface{}) (res *big.Int) {
	tb.Helper()
	res, err := AdjustAmount(ctx, amount, handler)
	assert.Nil(tb, err)
	return res
}

// AdjustAmount multiplies a token by the number of decimals in the amount.
// this does not use an interface with a Decimals() methods since usdt returns a (non-erc 20 compliant)
// *big.Int rather than a uint8 for decimals. If the handler is determined to be usdt (via a type switch)
// that decimals method is called, otherwise the standard decimals method is called.
func AdjustAmount(ctx context.Context, amount *big.Int, handler interface{}) (res *big.Int, err error) {
	decimalCount, err := GetDecimals(ctx, handler)
	if err != nil {
		return nil, fmt.Errorf("could not get decimals: %w", err)
	}
	decimalMultiplier := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimalCount)), nil)
	return big.NewInt(0).Mul(decimalMultiplier, amount), nil
}

// ERC20DecimalHandler gets the decimal count from a standard erc20 token.
type ERC20DecimalHandler interface {
	// Decimals gets the erc-20 decimals
	Decimals(opts *bind.CallOpts) (uint8, error)
}

// BigIntDecimalHandler is the decimal handler for tokens which return a *big.Int
// this is non-standard: but it's done by usdt token (see: https://tether.to/).
type BigIntDecimalHandler interface {
	// Decimals gets the decimals from the big int
	Decimals(opts *bind.CallOpts) (*big.Int, error)
}

// GetDecimals gets decimals from token that adheres to either the tether or the erc-20 standard.
func GetDecimals(ctx context.Context, handler interface{}) (res uint8, err error) {
	var decimalCount *big.Int
	switch decimalHandler := handler.(type) {
	case BigIntDecimalHandler:
		decimalCount, err = decimalHandler.Decimals(&bind.CallOpts{Context: ctx})
	case ERC20DecimalHandler:
		var decimals uint8
		decimals, err = decimalHandler.Decimals(&bind.CallOpts{Context: ctx})
		decimalCount = big.NewInt(int64(decimals))
	default:
		err = fmt.Errorf("no handler available for type %s", reflect.TypeOf(decimalHandler))
	}
	if err != nil {
		return res, fmt.Errorf("could not get decimals: %w", err)
	}
	return uint8(decimalCount.Uint64()), nil
}
