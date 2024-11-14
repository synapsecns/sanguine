package base

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/guard/guarddb"
)

func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	statusFieldName = namer.GetConsistentName("Status")
	transactionIDFieldName = namer.GetConsistentName("TransactionID")
}

var (
	// statusFieldName is the status field name.
	statusFieldName string
	// transactionIDFieldName is the transactions id field name.
	transactionIDFieldName string
)

// PendingProvenModel is the primary event model.
type PendingProvenModel struct {
	// CreatedAt is the creation time
	CreatedAt time.Time
	// UpdatedAt is the update time
	UpdatedAt time.Time
	// Origin is the origin chain id
	Origin uint32
	// RelayerAddress is the address of the relayer that called prove()
	RelayerAddress string
	// TransactionID is the transaction id of the event
	TransactionID string `gorm:"column:transaction_id;primaryKey"`
	// TxHash is the hash of the relay transaction on destination
	TxHash string
	// Status is the status of the event
	Status guarddb.PendingProvenStatus `gorm:"column:status;index:idx_guard_status_name"`
	// BlockNumber is the block number of the event
	BlockNumber uint64
}

// FromPendingProven converts a quote request to an object that can be stored in the db.
func FromPendingProven(proven guarddb.PendingProven) PendingProvenModel {
	return PendingProvenModel{
		Origin:         proven.Origin,
		RelayerAddress: proven.RelayerAddress.Hex(),
		TransactionID:  hexutil.Encode(proven.TransactionID[:]),
		TxHash:         proven.TxHash.Hex(),
		Status:         proven.Status,
		BlockNumber:    proven.BlockNumber,
	}
}

// ToPendingProven converts a db object to a pending proven.
func (p PendingProvenModel) ToPendingProven() (*guarddb.PendingProven, error) {
	txID, err := hexutil.Decode(p.TransactionID)
	if err != nil {
		return nil, fmt.Errorf("could not get transaction id: %w", err)
	}

	transactionID, err := sliceToArray(txID)
	if err != nil {
		return nil, fmt.Errorf("could not convert transaction id: %w", err)
	}

	return &guarddb.PendingProven{
		Origin:         p.Origin,
		RelayerAddress: common.HexToAddress(p.RelayerAddress),
		TransactionID:  transactionID,
		TxHash:         common.HexToHash(p.TxHash),
		Status:         p.Status,
		BlockNumber:    p.BlockNumber,
	}, nil
}

// BridgeRequestModel is the primary event model.
type BridgeRequestModel struct {
	// CreatedAt is the creation time
	CreatedAt time.Time
	// UpdatedAt is the update time
	UpdatedAt time.Time
	// TransactionID is the transaction id of the event
	TransactionID string `gorm:"column:transaction_id;primaryKey"`
	// OriginChainID is the origin chain for the transactions
	OriginChainID uint32
	// DestChainID is the destination chain for the tx
	DestChainID uint32
	// OriginSender is the original sender
	OriginSender string
	// DestRecipient is the recipient of the destination tx
	DestRecipient string
	// OriginToken is the origin token address
	OriginToken string
	// DestToken is the destination token address
	DestToken string
	// OriginAmount is the origin amount stored for sorting.
	// This is not the source of truth, but is approximate
	OriginAmount string
	// DestAmount is the destination amount stored for sorting.
	DestAmount string
	// Deadline is the deadline for the relay
	Deadline time.Time `gorm:"index"`
	// OriginNonce is the nonce on the origin chain in the app.
	// this is not effected by the message.sender nonce.
	OriginNonce int `gorm:"index"`
	// RawRequest is the raw request, hex encoded.
	RawRequest string
	// SendChainGas is true if the chain should send gas
	SendChainGas bool
}

// FromBridgeRequest converts a bridge request object to db model.
func FromBridgeRequest(request guarddb.BridgeRequest) BridgeRequestModel {
	return BridgeRequestModel{
		TransactionID: hexutil.Encode(request.TransactionID[:]),
		OriginChainID: request.Transaction.OriginChainId,
		DestChainID:   request.Transaction.DestChainId,
		OriginSender:  request.Transaction.OriginSender.String(),
		DestRecipient: request.Transaction.DestRecipient.String(),
		OriginToken:   request.Transaction.OriginToken.String(),
		RawRequest:    hexutil.Encode(request.RawRequest),
		SendChainGas:  request.Transaction.SendChainGas,
		DestToken:     request.Transaction.DestToken.String(),
		OriginAmount:  request.Transaction.OriginAmount.String(),
		DestAmount:    request.Transaction.DestAmount.String(),
		Deadline:      time.Unix(int64(request.Transaction.Deadline.Uint64()), 0),
		OriginNonce:   int(request.Transaction.Nonce.Uint64()),
	}
}

// ToBridgeRequest converts the bridge request db model to object.
func (b BridgeRequestModel) ToBridgeRequest() (*guarddb.BridgeRequest, error) {
	txID, err := hexutil.Decode(b.TransactionID)
	if err != nil {
		return nil, fmt.Errorf("could not get transaction id: %w", err)
	}

	req, err := hexutil.Decode(b.RawRequest)
	if err != nil {
		return nil, fmt.Errorf("could not get request: %w", err)
	}

	transactionID, err := sliceToArray(txID)
	if err != nil {
		return nil, fmt.Errorf("could not convert transaction id: %w", err)
	}

	originAmount, ok := new(big.Int).SetString(b.OriginAmount, 10)
	if !ok {
		return nil, errors.New("could not convert origin amount")
	}
	destAmount, ok := new(big.Int).SetString(b.DestAmount, 10)
	if !ok {
		return nil, errors.New("could not convert dest amount")
	}

	return &guarddb.BridgeRequest{
		TransactionID: transactionID,
		RawRequest:    req,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId: b.OriginChainID,
			DestChainId:   b.DestChainID,
			OriginSender:  common.HexToAddress(b.OriginSender),
			DestRecipient: common.HexToAddress(b.DestRecipient),
			OriginToken:   common.HexToAddress(b.OriginToken),
			SendChainGas:  b.SendChainGas,
			DestToken:     common.HexToAddress(b.DestToken),
			OriginAmount:  originAmount,
			DestAmount:    destAmount,
			Deadline:      big.NewInt(b.Deadline.Unix()),
			Nonce:         big.NewInt(int64(b.OriginNonce)),
		},
	}, nil
}

func sliceToArray(slice []byte) ([32]byte, error) {
	var arr [32]byte
	if len(slice) != 32 {
		return arr, errors.New("slice is not 32 bytes long")
	}
	copy(arr[:], slice)
	return arr, nil
}
