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
)

// CopyTransactOpts exports copyTransactOpts for testing.
func CopyTransactOpts(opts *bind.TransactOpts) *bind.TransactOpts {
	return copyTransactOpts(opts)
}

// SortTxes exports sortTxesByChainID for testing.
func SortTxes(txs []db.TX, maxPerChain int) map[uint64][]db.TX {
	return sortTxesByChainID(txs, maxPerChain)
}

// GroupTxesByNonce exports groupTxesByNonce for testing.
func GroupTxesByNonce(txs []db.TX) map[uint64][]db.TX {
	return groupTxesByNonce(txs)
}

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

// Outersection exports outersection for testing.
func Outersection(set, superset []*big.Int) []*big.Int {
	return outersection(set, superset)
}

// MapToBigIntSlice exports mapToBigIntSlice for testing.
func MapToBigIntSlice[T any](m map[uint64]T) []*big.Int {
	return mapToBigIntSlice(m)
}
