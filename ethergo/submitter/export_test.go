package submitter

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/submitter/config"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"go.opentelemetry.io/otel/attribute"
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
func TxToAttributes(transaction *types.Transaction, UUID string) []attribute.KeyValue {
	return txToAttributes(transaction, UUID)
}

// SortTxes exports sortTxesByChainID for testing.
func SortTxes(txs []db.TX) map[uint64][]db.TX {
	return sortTxesByChainID(txs)
}

// GroupTxesByNonce exports groupTxesByNonce for testing.
func GroupTxesByNonce(txs []db.TX) map[uint64][]db.TX {
	return groupTxesByNonce(txs)
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

// NewTestTransactionSubmitter wraps TestTransactionSubmitter in a TransactionSubmitter interface.
func NewTestTransactionSubmitter(metrics metrics.Handler, signer signer.Signer, fetcher ClientFetcher, db db.Service, config *config.Config) TestTransactionSubmitter {
	txSubmitter := NewTransactionSubmitter(metrics, signer, fetcher, db, config)
	//nolint: forcetypeassert
	return txSubmitter.(TestTransactionSubmitter)
}

// TestTransactionSubmitter is a TransactionSubmitter interface for testing.
type TestTransactionSubmitter interface {
	TransactionSubmitter
	// SetGasPrice exports setGasPrice for testing.
	SetGasPrice(ctx context.Context, client client.EVM,
		transactor *bind.TransactOpts, bigChainID *big.Int, prevTx *types.Transaction) (err error)
	// GetGasBlock exports getGasBlock for testing.
	GetGasBlock(ctx context.Context, client client.EVM, chainID int) (gasBlock *types.Header, err error)
	// GetNonce exports getNonce for testing.
	GetNonce(parentCtx context.Context, chainID *big.Int, address common.Address) (_ uint64, err error)
	// CheckAndSetConfirmation exports checkAndSetConfirmation for testing.
	CheckAndSetConfirmation(ctx context.Context, chainClient client.EVM, txes []db.TX) error
}

// SetGasPrice exports setGasPrice for testing.
func (t *txSubmitterImpl) SetGasPrice(ctx context.Context, client client.EVM,
	transactor *bind.TransactOpts, bigChainID *big.Int, prevTx *types.Transaction) (err error) {
	return t.setGasPrice(ctx, client, transactor, bigChainID, prevTx)
}

// GetGasBlock exports getGasBlock for testing.
func (t *txSubmitterImpl) GetGasBlock(ctx context.Context, client client.EVM, chainID int) (gasBlock *types.Header, err error) {
	return t.getGasBlock(ctx, client, chainID)
}

// GetNonce exports getNonce for testing.
func (t *txSubmitterImpl) GetNonce(parentCtx context.Context, chainID *big.Int, address common.Address) (_ uint64, err error) {
	return t.getNonce(parentCtx, chainID, address)
}

// CheckAndSetConfirmation exports checkAndSetConfirmation for testing.
func (t *txSubmitterImpl) CheckAndSetConfirmation(ctx context.Context, chainClient client.EVM, txes []db.TX) error {
	return t.checkAndSetConfirmation(ctx, chainClient, txes)
}

// SortTxes exports sortTxesByChainID for testing.
func SortTxes(txs []db.TX) map[uint64][]db.TX {
	return sortTxesByChainID(txs)
}

// GroupTxesByNonce exports groupTxesByNonce for testing.
func GroupTxesByNonce(txs []db.TX) map[uint64][]db.TX {
	return groupTxesByNonce(txs)
}

// SetForceNoFallback exports setForceNoFallback for testing.
func SetForceNoFallback(val bool) {
	forceNoFallbackIfZero = val
}

// export for testing
// Warning! Changing this value will not change the underlying value!
// this is exported for testing to allow a return to underlying value after the test
// TODO: consider doing a SetForTest.
var ForceNoFallbackIfZero = forceNoFallbackIfZero
