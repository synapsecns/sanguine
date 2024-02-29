package db

import (
	"context"
	"gorm.io/gorm"
	"time"
)

// ChainListenerDB is the interface for the chain listener database.
type ChainListenerDB interface {
	// PutLatestBlock upsers the latest block on a given chain id to be new height.
	PutLatestBlock(ctx context.Context, chainID, height uint64) error
	// LatestBlockForChain gets the latest block for a given chain id.
	// will return ErrNoLatestBlockForChainID if no block exists for the chain.
	LatestBlockForChain(ctx context.Context, chainID uint64) (uint64, error)
}

// LastIndexed is used to make sure we haven't missed any events while offline.
// since we event source - rather than use a state machine this is needed to make sure we haven't missed any events
// by allowing us to go back and source any events we may have missed.
//
// this does not inherit from gorm.model to allow us to use ChainID as a primary key.
type LastIndexed struct {
	// CreatedAt is the creation time
	CreatedAt time.Time
	// UpdatedAt is the update time
	UpdatedAt time.Time
	// DeletedAt time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// ChainID is the chain id of the chain we're watching blocks on. This is our primary index.
	ChainID uint64 `gorm:"column:chain_id;primaryKey;autoIncrement:false"`
	// BlockHeight is the highest height we've seen on the chain
	BlockNumber int `gorm:"block_number"`
}

// GetAllModels gets all models to migrate
// see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func GetAllModels() (allModels []interface{}) {
	allModels = []interface{}{&LastIndexed{}}
	return allModels
}
