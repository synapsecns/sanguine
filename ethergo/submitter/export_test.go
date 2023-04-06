package submitter

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.opentelemetry.io/otel/attribute"
	"math/big"
)

// CopyTransactOpts exports copyTransactOpts for testing.
func CopyTransactOpts(opts *bind.TransactOpts) *bind.TransactOpts {
	return copyTransactOpts(opts)
}

// NullFieldAttribute is a constant used to test the null field attribute.
// it exports the underlying constant for testing.
const NullFieldAttribute = nullFieldAttribute

func AddressPtrToString(address *common.Address) string {
	return addressPtrToString(address)
}

// BigPtrToString converts a big.Int pointer to a string.
func BigPtrToString(num *big.Int) string {
	return bigPtrToString(num)
}

// TxToAttributes exports txToAttributes for testing.
func TxToAttributes(transaction *types.Transaction) []attribute.KeyValue {
	return txToAttributes(transaction)
}

const (
	// HashAttr exports hashAttr for testing.
	HashAttr = hashAttr
	// FromAttr exports fromAttr for testing.
	FromAttr = fromAttr
	// ToAttr exports toAttr for testing.
	ToAttr = toAttr
	// DataAttr exports dataAttr for testing.
	DataAttr = dataAttr
	// ValueAttr exports valueAttr for testing.
	ValueAttr = valueAttr
	// NonceAttr exports nonceAttr for testing.
	NonceAttr = nonceAttr
	// GasLimitAttr exports gasLimitAttr for testing.
	GasLimitAttr = gasLimitAttr
	// ChainIDAttr exports chainIDAttr for testing.
	ChainIDAttr = chainIDAttr
	// GasPriceAttr exports gasPriceAttr for testing.
	GasPriceAttr = gasPriceAttr
	// GasFeeCapAttr exports gasFeeCapAttr for testing.
	GasFeeCapAttr = gasFeeCapAttr
	// GasTipCapAttr exports gasTipCapAttr for testing.
	GasTipCapAttr = gasTipCapAttr
)
