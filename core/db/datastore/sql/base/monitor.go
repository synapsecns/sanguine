package base

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/types"
	"gorm.io/gorm"
)

// StoreDispatchMessage takes a message and stores the information.
func (s Store) StoreDispatchMessage(ctx context.Context, message types.Message) error {
	dxTx := s.DB().WithContext(ctx).Create(&DispatchMessage{
		DMOrigin:      message.OriginDomain(),
		DMSender:      message.Sender().String(),
		DMNonce:       message.Nonce(),
		DMDestination: message.DestinationDomain(),
		DMRecipient:   message.Recipient().String(),
	})
	if dxTx.Error != nil {
		return fmt.Errorf("could not insert dispatch message: %w", dxTx.Error)
	}
	return nil
}

// StoreAcceptedAttestation stores an accepted attestation from a destination.
func (s Store) StoreAcceptedAttestation(ctx context.Context, destinationDomain uint32, attestation types.Attestation) error {
	root := attestation.Root()
	dxTx := s.DB().WithContext(ctx).Create(&AcceptedAttestation{
		AAOriginDomain:      attestation.Domain(),
		AANonce:             attestation.Nonce(),
		AARoot:              "0x" + hex.EncodeToString(root[:]),
		AADestinationDomain: destinationDomain,
	})
	if dxTx.Error != nil {
		return fmt.Errorf("could not insert accepted attestation: %w", dxTx.Error)
	}
	return nil
}

// GetDelinquentMessages gets messages that were sent on the origin chain,
// but never received on the destination chain.
func (s Store) GetDelinquentMessages(ctx context.Context, destinationDomain uint32) ([]types.DelinquentMessage, error) {
	var delinquentMessages []types.DelinquentMessage
	var res DispatchMessage
	stmt := &gorm.Statement{DB: s.DB().WithContext(ctx)}
	// Get the SQL table name of the DispatchMessage table.
	err := stmt.Parse(&DispatchMessage{})
	if err != nil {
		return nil, fmt.Errorf("could not parse dispatch message table name: %w", err)
	}
	dmTable := stmt.Schema.Table
	// Get the SQL table name of the AcceptedAttestation table.
	err = stmt.Parse(&AcceptedAttestation{})
	if err != nil {
		return nil, fmt.Errorf("could not parse accepted attestation table name: %w", err)
	}
	aaTable := stmt.Schema.Table

	// Run an inverse join on the nonces between dispatched messages and accepted attestations on a given destination domain.
	rows, err := s.DB().
		WithContext(ctx).
		Model(&DispatchMessage{}).
		Select(dmTable+".*").
		Joins("LEFT OUTER JOIN "+aaTable+" ON "+aaTable+".nonce = "+dmTable+".nonce").
		Where(aaTable+".nonce IS NULL AND "+dmTable+".destination = ?", destinationDomain).
		Rows()
	if err != nil {
		return []types.DelinquentMessage{}, fmt.Errorf("could not get rows: %w", err)
	}
	if rows.Err() != nil {
		return []types.DelinquentMessage{}, fmt.Errorf("could not get rows: %w", rows.Err())
	}
	for rows.Next() {
		err = s.DB().ScanRows(rows, &res)
		if err != nil {
			return []types.DelinquentMessage{}, fmt.Errorf("could not scan rows: %w", err)
		}
		// Create a new DelinquentMessage based on the data, and append to the returned list.
		delinquentMessage := types.NewDelinquentMessage(
			res.DMOrigin,
			common.HexToHash(res.DMSender),
			res.DMNonce,
			res.DMDestination,
			common.HexToHash(res.DMRecipient),
		)
		delinquentMessages = append(delinquentMessages, delinquentMessage)
	}
	return delinquentMessages, nil
}
