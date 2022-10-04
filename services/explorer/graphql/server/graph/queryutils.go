package graph

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"time"
)

func (r *queryResolver) getChainIDs(ctx context.Context, chainID *int) ([]uint32, error) {
	var chainIDs []uint32
	// if the chain ID is not specified, get all chain IDs
	if chainID == nil {
		chainIDsInt, err := r.DB.GetAllChainIDs(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get all chain IDs: %w", err)
		}
		chainIDs = append(chainIDs, chainIDsInt...)
	} else {
		chainIDs = append(chainIDs, uint32(*chainID))
	}
	return chainIDs, nil
}

func (r *queryResolver) getDirectionIn(direction *model.Direction) bool {
	var directionIn bool
	if direction != nil {
		directionIn = *direction == model.DirectionIn
	} else {
		directionIn = true
	}
	return directionIn
}

func (r *queryResolver) getTargetTime(hours *int) uint64 {
	var targetTime uint64
	if hours == nil {
		targetTime = uint64(time.Now().Add(-time.Hour * 24).Unix())
	} else {
		targetTime = uint64(time.Now().Add(-time.Hour * time.Duration(*hours)).Unix())
	}
	return targetTime
}

func (r *queryResolver) getTokenAddressesByChainID(ctx context.Context, chainID uint32) ([]string, error) {
	tokenAddresses, err := r.DB.GetTokenAddressesByChainID(ctx, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to get token addresses by chain ID: %w", err)
	}
	return tokenAddresses, nil
}
