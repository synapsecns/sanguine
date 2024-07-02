package base

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/services/rfq/guard/guarddb"
)

func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	statusFieldName = namer.GetConsistentName("Status")
	transactionIDFieldName = namer.GetConsistentName("TransactionID")
	originTxHashFieldName = namer.GetConsistentName("OriginTxHash")
	destTxHashFieldName = namer.GetConsistentName("DestTxHash")
	rebalanceIDFieldName = namer.GetConsistentName("RebalanceID")
	relayNonceFieldName = namer.GetConsistentName("RelayNonce")
}

var (
	// statusFieldName is the status field name.
	statusFieldName string
	// transactionIDFieldName is the transactions id field name.
	transactionIDFieldName string
	// originTxHashFieldName is the origin tx hash field name.
	originTxHashFieldName string
	// destTxHashFieldName is the dest tx hash field name.
	destTxHashFieldName string
	// rebalanceIDFieldName is the rebalances id field name.
	rebalanceIDFieldName string
	// relayNonceFieldName is the relay nonce field name.
	relayNonceFieldName string
)

// PendingProvenModel is the primary event model.
type PendingProvenModel struct {
	// CreatedAt is the creation time
	CreatedAt time.Time
	// UpdatedAt is the update time
	UpdatedAt time.Time
	// TransactionID is the transaction id of the event
	TransactionID string `gorm:"column:transaction_id;primaryKey"`
	// TxHash is the hash of the relay transaction on destination
	TxHash string
	// Status is the status of the event
	Status guarddb.PendingProvenStatus
}

// FromPendingProven converts a quote request to an object that can be stored in the db.
func FromPendingProven(proven guarddb.PendingProven) PendingProvenModel {
	return PendingProvenModel{
		TransactionID: hexutil.Encode(proven.TransactionID[:]),
		TxHash:        proven.TxHash.Hex(),
		Status:        proven.Status,
	}
}

var emptyHash = common.HexToHash("").Hex()

func hashToNullString(h common.Hash) sql.NullString {
	if h.Hex() == emptyHash {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{
		String: h.Hex(),
		Valid:  true,
	}
}

func stringToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
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
		TransactionID: transactionID,
		TxHash:        common.HexToHash(p.TxHash),
		Status:        p.Status,
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
