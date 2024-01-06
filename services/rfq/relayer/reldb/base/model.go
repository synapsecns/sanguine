package base

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"gorm.io/gorm"
)

func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	chainIDFieldName = namer.GetConsistentName("ChainID")
	blockNumberFieldName = namer.GetConsistentName("BlockNumber")
	statusFieldName = namer.GetConsistentName("Status")
	transactionIDFieldName = namer.GetConsistentName("TransactionID")
	destTxHashFieldName = namer.GetConsistentName("DestTxHash")
}

var (
	// chainIDFieldName gets the chain id field name.
	chainIDFieldName string
	// blockNumberFieldName is the name of the block number field.
	blockNumberFieldName string

	statusFieldName string
	// transactionIDFieldName is the transactions id field name.
	transactionIDFieldName string
	// destTxHashFieldName is the dest tx hash field name.
	destTxHashFieldName string
)

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

// RequestForQuote is the primary event model.
type RequestForQuote struct {
	// CreatedAt is the creation time
	CreatedAt time.Time
	// UpdatedAt is the update time
	UpdatedAt     time.Time
	TransactionID string `gorm:"column:chain_id;primaryKey;autoIncrement:false"` // TODO: change to transaction_id
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
	// OriginAmountOriginal is the origin amount used for preicison
	OriginAmountOriginal string
	// OriginTokenDecimals is the origin token decimals
	OriginTokenDecimals uint8
	// DestToken is the destination token address
	DestToken string
	// DestTokenDecimals is the destination token decimal count
	DestTokenDecimals uint8
	// OriginAmount is the origin amount stored for sorting.
	// This is not the source of truth, but is approximate
	OriginAmount decimal.Decimal `gorm:"index"`
	// DestAmountOriginal is the original amount used for precision
	DestAmountOriginal string
	// DestAmountOriginal is the original destination amount
	DestAmount decimal.Decimal `gorm:"index"`
	DestTxHash string
	Deadline   time.Time `gorm:"index"`
	// OriginNonce is the nonce on the origin chain in the app.
	// this is not effected by the message.sender nonce.
	OriginNonce int `gorm:"index"`
	// Status is the current status of the event
	Status reldb.QuoteRequestStatus
	// BlockNumber is the block number of the event
	BlockNumber uint64
	// RawRequest is the raw request, hex encoded.
	RawRequest string
	// SendChainGas is true if the chain should send gas
	SendChainGas bool
}

// FromQuoteRequest converts a quote request to an object that can be stored in the db.
// TODO: add validation for deadline > uint64
// TODO: roundtripper test.
func FromQuoteRequest(request reldb.QuoteRequest) RequestForQuote {
	return RequestForQuote{
		TransactionID:        hexutil.Encode(request.TransactionID[:]),
		OriginChainID:        request.Transaction.OriginChainId,
		DestChainID:          request.Transaction.DestChainId,
		OriginSender:         request.Transaction.OriginSender.String(),
		DestRecipient:        request.Transaction.DestRecipient.String(),
		OriginToken:          request.Transaction.OriginToken.String(),
		OriginTokenDecimals:  request.OriginTokenDecimals,
		RawRequest:           hexutil.Encode(request.RawRequest),
		SendChainGas:         request.Transaction.SendChainGas,
		DestTokenDecimals:    request.DestTokenDecimals,
		DestToken:            request.Transaction.DestToken.String(),
		DestTxHash:           request.DestTxHash.String(),
		OriginAmountOriginal: request.Transaction.OriginAmount.String(),
		OriginAmount:         decimal.NewFromBigInt(request.Transaction.OriginAmount, int32(request.OriginTokenDecimals)),
		DestAmountOriginal:   request.Transaction.DestAmount.String(),
		DestAmount:           decimal.NewFromBigInt(request.Transaction.DestAmount, int32(request.DestTokenDecimals)),
		Deadline:             time.Unix(int64(request.Transaction.Deadline.Uint64()), 0),
		OriginNonce:          int(request.Transaction.Nonce.Uint64()),
		Status:               request.Status,
		BlockNumber:          request.BlockNumber,
	}
}

// ToQuoteRequest converts a db object to a quote request.
func (r RequestForQuote) ToQuoteRequest() (*reldb.QuoteRequest, error) {
	txID, err := hexutil.Decode(r.TransactionID)
	if err != nil {
		return nil, fmt.Errorf("could not get transaction id: %w", err)
	}

	req, err := hexutil.Decode(r.RawRequest)
	if err != nil {
		return nil, fmt.Errorf("could not get request: %w", err)
	}

	transactionID, err := sliceToArray(txID)
	if err != nil {
		return nil, fmt.Errorf("could not convert transaction id: %w", err)
	}

	return &reldb.QuoteRequest{
		OriginTokenDecimals: r.OriginTokenDecimals,
		DestTokenDecimals:   r.DestTokenDecimals,
		TransactionID:       transactionID,
		RawRequest:          req,
		Sender:              common.HexToAddress(r.OriginSender),
		BlockNumber:         r.BlockNumber,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId: r.OriginChainID,
			DestChainId:   r.DestChainID,
			OriginSender:  common.HexToAddress(r.OriginSender),
			DestRecipient: common.HexToAddress(r.DestRecipient),
			OriginToken:   common.HexToAddress(r.OriginToken),
			SendChainGas:  r.SendChainGas,
			DestToken:     common.HexToAddress(r.DestToken),
			OriginAmount:  new(big.Int).Div(r.OriginAmount.BigInt(), big.NewInt(int64(math.Pow10(int(r.OriginTokenDecimals))))),
			// OriginAmount: new(big.Int).Div(r.OriginAmount.BigInt(), big.NewInt(int64(r.OriginTokenDecimals))),
			DestAmount: new(big.Int).Div(r.DestAmount.BigInt(), big.NewInt(int64(math.Pow10(int(r.DestTokenDecimals))))),
			Deadline:   big.NewInt(r.Deadline.Unix()),
			Nonce:      big.NewInt(int64(r.OriginNonce)),
		},
		Status:     r.Status,
		DestTxHash: common.HexToHash(r.DestTxHash),
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
