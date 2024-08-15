package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// NewChainListenerStore creates a new transaction store.
func NewChainListenerStore(db *gorm.DB, metrics metrics.Handler, listenerName string) *Store {
	return &Store{
		db:           db,
		metrics:      metrics,
		listenerName: listenerName,
	}
}

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	db           *gorm.DB
	metrics      metrics.Handler
	listenerName string
}

// PutLatestBlock upserts the latest block into the database.
func (s *Store) PutLatestBlock(ctx context.Context, chainID, height uint64) error {
	tx := s.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: chainIDFieldName}, {Name: s.listenerName}},
		DoUpdates: clause.AssignmentColumns([]string{chainIDFieldName, blockNumberFieldName}),
	}).Create(&LastIndexed{
		ChainID:     chainID,
		BlockNumber: int(height),
	})

	if tx.Error != nil {
		return fmt.Errorf("could not block updated: %w", tx.Error)
	}
	return nil
}

// LatestBlockForChain gets the latest block for a chain.
func (s *Store) LatestBlockForChain(ctx context.Context, chainID uint64) (uint64, error) {
	blockWatchModel := LastIndexed{ChainID: chainID, ListenerName: s.listenerName}
	err := s.db.WithContext(ctx).First(&blockWatchModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, ErrNoLatestBlockForChainID
		}
		return 0, fmt.Errorf("could not fetch latest block: %w", err)
	}

	return uint64(blockWatchModel.BlockNumber), nil
}

// SetListenerName sets the listener name.
func (s *Store) SetListenerName(name string) {
	s.listenerName = name
}

func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	chainIDFieldName = namer.GetConsistentName("ChainID")
	blockNumberFieldName = namer.GetConsistentName("BlockNumber")
}

var (
	// chainIDFieldName gets the chain id field name.
	chainIDFieldName string
	// blockNumberFieldName is the name of the block number field.
	blockNumberFieldName string
)

// ErrNoLatestBlockForChainID is returned when no block exists for the chain.
var ErrNoLatestBlockForChainID = errors.New("no latest block for chainId")
