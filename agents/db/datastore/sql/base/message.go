package base

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/agents/db"

	"github.com/Thor-x86/nullable"
	"github.com/synapsecns/sanguine/agents/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/clause"
)

// StoreCommittedMessage stores a raw committed message building off the leaf index
// this method is idempotent.
func (s Store) StoreCommittedMessage(ctx context.Context, domainID uint32, message types.CommittedMessage) error {
	decodedMessage, err := types.DecodeMessage(message.Message())
	if err != nil {
		return fmt.Errorf("could not decode message for insertion")
	}

	// workaround for sqlite issue: https://github.com/hashicorp/go-dbw/issues/2
	// currently, we can't process on conflcicts with composite keys, so we emulate the behavior here for testing
	// can be fixed after: https://github.com/go-gorm/gorm/issues/4879
	baseQuery := s.DB().WithContext(ctx)
	if s.db.Dialector.Name() != (sqlite.Dialector{}).Name() {
		baseQuery = baseQuery.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: DomainIDFieldName}, {Name: LeafIndexFieldName}},
			DoNothing: true,
		})
	} else {
		var count int64
		tx := s.DB().WithContext(ctx).Model(&CommittedMessage{}).Where(&CommittedMessage{
			CMDomainID: domainID,
			CMNonce:    decodedMessage.Nonce(),
		}).Count(&count)

		if tx.Error != nil {
			return fmt.Errorf("could not check keys: %w", tx.Error)
		}

		// emulate already exists behavior
		if count > 0 {
			return nil
		}
	}

	tx := baseQuery.Create(&CommittedMessage{
		CMDomainID:          domainID,
		CMMessage:           message.Message(),
		CMLeaf:              hashToSlice(message.Leaf()),
		CMOrigin:            decodedMessage.OriginDomain(),
		CMNonce:             decodedMessage.Nonce(),
		CMDestination:       decodedMessage.DestinationDomain(),
		CMBody:              decodedMessage.Body(),
		CMOptimisticSeconds: decodedMessage.OptimisticSeconds(),
	})

	if tx.Error != nil {
		return fmt.Errorf("could not store committed message updated: %w", tx.Error)
	}
	return nil
}

// RetrieveLatestCommittedMessageNonce gets the latest commitedd message by nonce.
func (s Store) RetrieveLatestCommittedMessageNonce(ctx context.Context, domainID uint32) (_ uint32, err error) {
	var nonce nullable.Uint32

	selectMaxNonce := fmt.Sprintf("max(`%s`)", NonceFieldName)

	tx := s.DB().WithContext(ctx).Model(&CommittedMessage{}).Select(selectMaxNonce).Where(CommittedMessage{CMDomainID: domainID}).Scan(&nonce)

	if tx.Error != nil {
		return 0, fmt.Errorf("could not get nonce for chain id: %w", tx.Error)
	}

	// if no nonces, return the corresponding error.
	if nonce.Get() == nil {
		return 0, db.ErrNoNonceForDomain
	}
	return *nonce.Get(), nil
}

// hashToSlice converts a kappa value toa  byte slice.
func hashToSlice(hash [32]byte) []byte {
	rawKappa := make([]byte, len(hash))
	copy(rawKappa, hash[:])
	return rawKappa
}
