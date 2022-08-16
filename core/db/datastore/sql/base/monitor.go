package base

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/synapsecns/sanguine/core/types"
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
