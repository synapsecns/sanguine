package base

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
)

func (s Store) StoreTokenSwap(ctx context.Context, rawLog types.Log, swap swap.SwapUtilsTokenSwap, chainID uint32) {

	//dbTx := s.DB().
	//	WithContext(ctx).
	//	Clauses(clause.OnConflict{
	//		Columns:   []clause.Column{{Name: TxHashFieldName}, {Name: ChainIDFieldName}},
	//		DoNothing: true,
	//	}).Create()
}
