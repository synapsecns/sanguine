package base

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/db"
	synapseCommon "github.com/synapsecns/synapse-node/pkg/common"
	"github.com/synapsecns/synapse-node/pkg/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/big"
	"strings"
)

// UpdateEventStatus updates an events status. We need to do this here and not on, say, tx result because
// we don't marshall + store the logs for idempotent acces.
func (s Store) UpdateEventStatus(ctx context.Context, identifier string, status db.EventStatus) error {
	logger.Debugf("updating status for txid %s to %s", identifier, status.String())
	tx := s.DB().WithContext(ctx).Model(&EventModel{}).Where(fmt.Sprintf("%s = ?", TxHashFieldName), identifier).Update(EventStatusFieldName, status)
	if tx.Error != nil {
		return fmt.Errorf("could not update event status: %w", tx.Error)
	}
	return nil
}

// PutUserEvent idempotently stores a bridge event an event log.
func (s Store) PutUserEvent(ctx context.Context, originChainID big.Int, log types.CrossChainUserEventLog, status db.EventStatus) error {
	logger.Debugf("storing %s for tx %s with status %s", log.GetEventType().String(), synapseCommon.GetTxLogText(&originChainID, log.GetIdentifier()), status.String())

	tx := s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: TxHashFieldName}},
		DoUpdates: clause.AssignmentColumns([]string{EventStatusFieldName}),
	}).Create(&EventModel{
		Kappa:         bridge.KappaToSlice(bridge.GetKappa(log)),
		To:            log.GetContractAddress(),
		OriginChainID: originChainID.Uint64(),
		DestChainID:   log.GetDestinationChainID().Uint64(),
		Amount:        log.GetAmount().Uint64(),
		Token:         log.GetToken(),
		BlockNumber:   log.GetBlockNumber(),
		TXHash:        log.GetIdentifier(),
		EventStatus:   status,
		EventType:     log.GetEventType(),
	})

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// PutBridgeEvent idempotently stores an event triggered by the bridge.
func (s Store) PutBridgeEvent(ctx context.Context, chainID big.Int, log types.CrossChainBridgeEventLog) error {
	logger.Debugf("storing %s for tx %s", log.GetEventType().String(), synapseCommon.GetTxLogText(&chainID, log.GetIdentifier()))

	var originatingEvent EventModel
	tx := s.DB().WithContext(ctx).Where(&EventModel{Kappa: bridge.KappaToSlice(log.GetKappa())}).First(&originatingEvent)
	if tx.Error != nil {
		return fmt.Errorf("could not get event model. Needed for kappa: %w", tx.Error)
	}

	tx = s.DB().WithContext(ctx).Create(&BridgeEventModel{
		To:               log.GetContractAddress(),
		ChainID:          chainID.Uint64(),
		Amount:           log.GetAmount().Uint64(),
		Token:            log.GetToken(),
		BlockNumber:      log.GetBlockNumber(),
		TXHash:           log.GetIdentifier(),
		Kappa:            bridge.KappaToSlice(log.GetKappa()),
		EventType:        log.GetEventType(),
		OriginatingEvent: &originatingEvent,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// GetEVMEventsByStatus fetches events that match statuses passed into matchStatuses.
//nolint: dupl
func (s Store) GetEVMEventsByStatus(ctx context.Context, matchStatuses []db.EventStatus) (results []db.EVMUserEvent, err error) {
	var modelResults []EventModel

	inArgs := make([]int, len(matchStatuses))
	// used for logging
	statuses := make([]string, len(matchStatuses))
	for i := range matchStatuses {
		inArgs[i] = int(matchStatuses[i].Int())
		statuses[i] = matchStatuses[i].String()
	}

	tx := s.DB().WithContext(ctx).Model(&EventModel{}).Where(fmt.Sprintf("%s  IN ?", EventStatusFieldName), inArgs).Order(fmt.Sprintf("%s desc, %s desc", OriginChainIDFieldName, BlockNumberFieldName)).Preload(clause.Associations).Find(&modelResults)
	if tx.Error != nil {
		return results, fmt.Errorf("could not get results for matchStatuses: %s. Got error: %w", strings.Join(statuses, ","), tx.Error)
	}

	for _, dbResult := range modelResults {
		// override the originating event with the parent event, gorm won't automatically handle this case :(
		for i := range dbResult.ResultingTX {
			dbResult.ResultingTX[i].OriginatingEvent = dbResult
		}
		event, err := dbResult.ToEVMUserEvent()
		if err != nil {
			return results, fmt.Errorf("could not convert db result to event: %w", err)
		}
		results = append(results, *event)
	}

	return results, nil
}

// GetUserEvent gets a chain agnostic user event.
func (s Store) GetUserEvent(ctx context.Context, identifier string, chainID *big.Int) (hasEvent bool, event *db.CrossChainUserEvent, err error) {
	// check if the chain has the deposit
	var model *EventModel
	tx := s.DB().WithContext(ctx).Model(&EventModel{}).Where(EventModel{TXHash: identifier, OriginChainID: chainID.Uint64()}, TxHashFieldName, OriginChainIDFieldName).Preload(clause.Associations).First(&model)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}
		return false, nil, fmt.Errorf("could not query for deposit: %w", tx.Error)
	}

	res := model.ToCrossChainUserEvent()

	return true, &res, nil
}
