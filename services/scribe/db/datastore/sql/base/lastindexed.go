package base

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm/clause"
)

const lastIndexedLivefillKey = "0x0000000000000000000000000000000000000000"

// StoreLastIndexed stores the last indexed block number for a contract.
// It updates the value if there is a previous last indexed value, and creates a new
// entry if there's no previous value.
func (s Store) StoreLastIndexed(parentCtx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64, livefillAtHead bool) (err error) {
	ctx, span := s.metrics.Tracer().Start(parentCtx, "StoreLastIndexed", trace.WithAttributes(
		attribute.String("contractAddress", contractAddress.String()),
		attribute.Int("chainID", int(chainID)),
		attribute.Int("blockNumber", int(blockNumber)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	address := contractAddress.String()
	if livefillAtHead {
		address = lastIndexedLivefillKey
	}

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: ContractAddressFieldName}, {Name: ChainIDFieldName}},
			DoUpdates: clause.AssignmentColumns([]string{BlockNumberFieldName}),
			Where: clause.Where{
				Exprs: []clause.Expression{
					clause.And(
						clause.Where{
							Exprs: []clause.Expression{
								clause.Eq{
									Column: clause.Column{Name: ContractAddressFieldName},
									Value:  address,
								},
								clause.Eq{
									Column: clause.Column{Name: ChainIDFieldName},
									Value:  chainID,
								},
							},
						},
						clause.Lt{
							Column: clause.Column{Name: BlockNumberFieldName},
							Value:  blockNumber,
						},
					),
				},
			},
		}).
		Create(&LastIndexedInfo{
			ContractAddress: address,
			ChainID:         chainID,
			BlockNumber:     blockNumber,
		})
	if dbTx.Error != nil {
		return fmt.Errorf("could not update last indexed info: %w", dbTx.Error)
	}
	return nil
}

// RetrieveLastIndexed retrieves the last indexed block number for a contract.
func (s Store) RetrieveLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32, livefill bool) (uint64, error) {
	entry := LastIndexedInfo{}
	address := contractAddress.String()
	if livefill {
		address = lastIndexedLivefillKey
	}

	dbTx := s.DB().WithContext(ctx).
		Model(&LastIndexedInfo{}).
		Where(&LastIndexedInfo{
			ContractAddress: address,
			ChainID:         chainID,
		}).
		First(&entry)
	if dbTx.RowsAffected == 0 {
		return 0, nil
	}
	if dbTx.Error != nil {
		return 0, fmt.Errorf("could not retrieve last indexed info: %w", dbTx.Error)
	}
	return entry.BlockNumber, nil
}

// StoreLastIndexedMultiple stores the last indexed block numbers for numerous contracts.
func (s Store) StoreLastIndexedMultiple(parentCtx context.Context, contractAddresses []common.Address, chainID uint32, blockNumber uint64) error {
	g, groupCtx := errgroup.WithContext(parentCtx)

	for i := range contractAddresses {
		index := i
		g.Go(func() error {
			err := s.StoreLastIndexed(groupCtx, contractAddresses[index], chainID, blockNumber, false)
			if err != nil {
				return fmt.Errorf("could not backfill: %w", err)
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not store last indexed: %w", err)
	}

	return nil
}

// RetrieveLastIndexedMultiple retrieves the last indexed block numbers for numerous contracts.
func (s Store) RetrieveLastIndexedMultiple(ctx context.Context, contractAddresses []common.Address, chainID uint32) (map[common.Address]uint64, error) {
	var entries []LastIndexedInfo
	addrStrings := make([]string, len(contractAddresses))
	for i, addr := range contractAddresses {
		addrStrings[i] = addr.String()
	}

	dbTx := s.DB().WithContext(ctx).
		Model(&LastIndexedInfo{}).
		Where("contract_address in ? AND chain_id = ?", addrStrings, chainID).
		Find(&entries)

	if dbTx.Error != nil {
		return nil, fmt.Errorf("could not retrieve last indexed info: %w", dbTx.Error)
	}

	result := make(map[common.Address]uint64)
	for _, addr := range contractAddresses {
		result[addr] = 0
	}

	for _, entry := range entries {
		addr := common.HexToAddress(entry.ContractAddress)
		result[addr] = entry.BlockNumber
	}

	return result, nil
}
