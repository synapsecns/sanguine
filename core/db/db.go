package db

import (
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/types"
)

// DB contains the synapse db.
type DB interface {
	// StoreCommittedMessage stores a committed message.
	StoreCommittedMessage(committedMessage types.CommittedMessage) error
}

// pebbleDB contains a rocksdb used to store merkle trees.
type pebbleDB struct {
	*pebble.DB
	entity string
}

// NewDB creates a new db.
func NewDB(dbPath, entity string) (DB, error) {
	db, err := pebble.Open(dbPath, &pebble.Options{})

	if err != nil {
		return nil, fmt.Errorf("could not create db: %w", err)
	}

	return &pebbleDB{DB: db, entity: entity}, nil
}

// StoreCommittedMessage stores a committed message.
func (d *pebbleDB) StoreCommittedMessage(committedMessage types.CommittedMessage) error {
	decodedMessage, err := types.DecodeMessage(committedMessage.Message())
	if err != nil {
		return fmt.Errorf("could not decode messages: %w", err)
	}

	err = d.StoreLeaf(committedMessage.LeafIndex(), decodedMessage.DestinationAndNonce(), committedMessage.Leaf())
	if err != nil {
		return fmt.Errorf("could not store leaf: %w", err)
	}

	err = d.StoreKeyedEncodable(MESSAGE, ToSlice(committedMessage.Leaf()), committedMessage.Message())
	if err != nil {
		return fmt.Errorf("could not store message: %w", err)
	}
	return nil
}

// StoreLeaf stores the leaf keyed by leaf_index.
func (d *pebbleDB) StoreLeaf(leafIndex uint32, destinationAndNonce uint64, leaf common.Hash) error {
	logger.Debugf("storing leaf hash keyed by index (%d) and dest+nonce (%d)", leafIndex, destinationAndNonce)

	err := d.StoreKeyedEncodable(LEAF, []byte(fmt.Sprintf("%d", destinationAndNonce)), leaf.Bytes())
	if err != nil {
		return fmt.Errorf("could not store destination encodable: %w", err)
	}

	err = d.StoreKeyedEncodable(LEAF, []byte(fmt.Sprintf("%d", leafIndex)), leaf.Bytes())
	if err != nil {
		return fmt.Errorf("could not store leaf index: %w", err)
	}

	return nil
}

// FullKey gets the full key.
func (d *pebbleDB) FullKey(prefix string, key []byte) []byte {
	return []byte(fmt.Sprintf("%s_%x_%s", d.entity, prefix, key))
}

// StoreKeyedEncodable stores a key + prefix.
func (d *pebbleDB) StoreKeyedEncodable(prefix string, key, value []byte) error {
	err := d.DB.Set(d.FullKey(prefix, key), value, &pebble.WriteOptions{Sync: true})
	if err != nil {
		return fmt.Errorf("could not store key encodable: %w", err)
	}
	return nil
}

// ToSlice converts a kappa value toa  byte slice.
func ToSlice(kappa [32]byte) []byte {
	rawKappa := make([]byte, len(kappa))
	copy(rawKappa, kappa[:])
	return rawKappa
}
