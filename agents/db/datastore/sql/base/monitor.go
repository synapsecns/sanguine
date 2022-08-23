package base

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/types"
	"gorm.io/gorm"
)

// StoreDispatchMessage takes a message and stores the information.
func (s Store) StoreDispatchMessage(ctx context.Context, message types.Message) error {
	dxTx := s.DB().WithContext(ctx).Create(&DispatchMessage{
		DMOrigin:            message.OriginDomain(),
		DMSender:            message.Sender().String(),
		DMNonce:             message.Nonce(),
		DMDestination:       message.DestinationDomain(),
		DMRecipient:         message.Recipient().String(),
		DMOptimisticSeconds: message.OptimisticSeconds(),
		DMNotaryTip:         message.Tips().NotaryTip().Bytes(),
		DMBroadcasterTip:    message.Tips().BroadcasterTip().Bytes(),
		DMProverTip:         message.Tips().ProverTip().Bytes(),
		DMExecutorTip:       message.Tips().ExecutorTip().Bytes(),
		DMBody:              message.Body(),
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
func (s Store) GetDelinquentMessages(ctx context.Context, destinationDomain uint32) ([]types.Message, error) {
	var delinquentMessages []types.Message
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

	// Format SQL query strings.
	joinStatement := fmt.Sprintf(
		"LEFT OUTER JOIN %s ON %s.%s = %s.%s",
		aaTable,
		aaTable,
		NonceFieldName,
		dmTable,
		NonceFieldName,
	)
	whereStatement := fmt.Sprintf(
		"%s.%s IS NULL AND %s.destination = %d",
		aaTable,
		NonceFieldName,
		dmTable,
		destinationDomain,
	)

	// Run an inverse join on the nonces between dispatched messages and accepted attestations on a given destination domain.
	rows, err := s.DB().WithContext(ctx).
		Model(&DispatchMessage{}).
		Select(dmTable + ".*").
		Joins(joinStatement).
		Where(whereStatement).
		Rows()
	if err != nil {
		return []types.Message{}, fmt.Errorf("could not get rows: %w", err)
	}
	if rows.Err() != nil {
		return []types.Message{}, fmt.Errorf("could not get rows: %w", rows.Err())
	}
	for rows.Next() {
		err = s.DB().ScanRows(rows, &res)
		if err != nil {
			return []types.Message{}, fmt.Errorf("could not scan rows: %w", err)
		}
		// Create a new Message based on the data, and append to the returned list.
		delinquentMessage := dispatchMessageToMessage(res)
		delinquentMessages = append(delinquentMessages, delinquentMessage)
	}
	return delinquentMessages, nil
}

// Take a DispatchMessage and convert it into a Message.
func dispatchMessageToMessage(d DispatchMessage) types.Message {
	header := types.NewHeader(
		d.DMOrigin,
		common.HexToHash(d.DMSender),
		d.DMNonce,
		d.DMDestination,
		common.HexToHash(d.DMRecipient),
		d.DMOptimisticSeconds,
	)
	tips := types.NewTips(
		new(big.Int).SetBytes(d.DMNotaryTip),
		new(big.Int).SetBytes(d.DMBroadcasterTip),
		new(big.Int).SetBytes(d.DMProverTip),
		new(big.Int).SetBytes(d.DMExecutorTip),
	)
	return types.NewMessage(
		header,
		tips,
		d.DMBody,
	)
}
