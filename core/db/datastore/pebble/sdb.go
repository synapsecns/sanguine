package pebble

import (
	"errors"
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/core/types"
	"strconv"
)

// pebbleDB contains a rocksdb used to store merkle trees.
// TODO: deprecate in favor of mysql.
type pebbleDB struct {
	*pebble.DB
	entity string
}

// NewMessageDB creates a new db.
func NewMessageDB(dbPath, entity string) (db.MessageDB, error) {
	messageDB, err := pebble.Open(dbPath, &pebble.Options{})

	if err != nil {
		return nil, fmt.Errorf("could not create messageDB: %w", err)
	}

	return &pebbleDB{DB: messageDB, entity: entity}, nil
}

func (d *pebbleDB) RetrieveProducedUpdate(root common.Hash) (types.SignedUpdate, error) {
	rawUpdate, _, err := d.Get(d.FullKey(UpdaterProducedUpdate, root.Bytes()))
	if err != nil {
		return nil, fmt.Errorf("could not get latest height: %w", err)
	}

	decodedUpdate, err := types.DecodeSignedUpdate(rawUpdate)
	if err != nil {
		return nil, fmt.Errorf("could not decode signed update: %w", err)
	}

	return decodedUpdate, nil
}

func (d *pebbleDB) StoreProducedUpdate(previousRoot common.Hash, update types.SignedUpdate) error {
	signedUpdate, err := types.EncodeSignedUpdate(update)
	if err != nil {
		return fmt.Errorf("could not encode signed update: %w", err)
	}

	err = d.StoreKeyedEncodable(UpdaterProducedUpdate, previousRoot.Bytes(), signedUpdate)
	if err != nil {
		return fmt.Errorf("could not store produced update: %w", err)
	}

	return nil
}

// RetrieveLatestRoot gets the latest root.
func (d *pebbleDB) RetrieveLatestRoot() (common.Hash, error) {
	rawRoot, _, err := d.Get(d.FullKey(LatestRoot, []byte("")))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not get latest height: %w", err)
	}

	return common.BytesToHash(rawRoot), nil
}

// StoreLatestRoot stores the latest root.
func (d *pebbleDB) StoreLatestRoot(latestRoot common.Hash) error {
	err := d.StoreKeyedEncodable(LatestRoot, []byte(""), latestRoot.Bytes())
	if err != nil {
		return fmt.Errorf("could not store latest root: %w", err)
	}
	return nil
}

// GetIndexedHeight gets the indexed height.
func (d *pebbleDB) GetIndexedHeight(domain string) (uint32, error) {
	rawHeight, _, err := d.Get(d.FullKey(HEIGHT, []byte(domain)))
	if err != nil {
		return 0, fmt.Errorf("could not get height: %w", err)
	}

	height, err := strconv.Atoi(string(rawHeight))
	if err != nil {
		return 0, fmt.Errorf("could not get indexed height: %w", err)
	}

	return uint32(height), nil
}

// StoreIndexedHeight stores the most recent indexed height.
func (d *pebbleDB) StoreIndexedHeight(domain string, height uint32) error {
	err := d.StoreKeyedEncodable(HEIGHT, []byte(domain), []byte(fmt.Sprintf("%d", height)))
	if err != nil {
		return fmt.Errorf("could not store height: %w", err)
	}
	return nil
}

// StoreLatestMessage stores the latest committed message.
func (d *pebbleDB) StoreLatestMessage(committedMessage types.CommittedMessage) error {
	// If there is no latest root, or if this update is on the latest root
	// update latest root
	latestLeaf, err := d.RetrieveLatestLeafIndex()
	if err == nil {
		if latestLeaf == committedMessage.LeafIndex()-1 {
			err = d.UpdateLatestLeafIndex(committedMessage.LeafIndex())
			if err != nil {
				return fmt.Errorf("could not store latest leaf index: %w", err)
			}
		} else {
			logger.Debugf("Attempted to store message not building off latest leaf index. Latest leaf index: %d. Attempted leaf index: %d.", latestLeaf, committedMessage.LeafIndex())
		}
	}
	if errors.Is(err, pebble.ErrNotFound) {
		err = d.UpdateLatestLeafIndex(committedMessage.LeafIndex())
		if err != nil {
			return fmt.Errorf("could not store latest leaf index: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("could not store message: %w", err)
	}

	return d.StoreCommittedMessage(committedMessage)
}

func (d *pebbleDB) UpdateLatestLeafIndex(leafIndex uint32) error {
	err := d.StoreKeyedEncodable("", []byte(LatestLeafIndex), []byte(strconv.Itoa(int(leafIndex))))
	if err != nil {
		return fmt.Errorf("could not store latest leaf: %w", err)
	}
	return nil
}

// RetrieveLatestLeafIndex gets the latest leaf index.
func (d *pebbleDB) RetrieveLatestLeafIndex() (uint32, error) {
	value, _, err := d.DB.Get(d.FullKey("", []byte(LatestLeafIndex)))
	if err != nil {
		return 0, fmt.Errorf("could not retrieve key: %w", err)
	}

	index, err := strconv.Atoi(string(value))
	if err != nil {
		return 0, fmt.Errorf("could not retrieve leaf index: %w", err)
	}

	return uint32(index), nil
}

// StoreCommittedMessage stores a committed message.
func (d *pebbleDB) StoreCommittedMessage(committedMessage types.CommittedMessage) error {
	// first we decode the underlying message to get the destination and nonce. These are used as a key for leaf storage
	decodedMessage, err := types.DecodeMessage(committedMessage.Message())
	if err != nil {
		return fmt.Errorf("could not decode messages: %w", err)
	}

	// now we store the leaf by both the leaf index and the destination and nonce
	err = d.StoreLeaf(committedMessage.LeafIndex(), decodedMessage.DestinationAndNonce(), committedMessage.Leaf())
	if err != nil {
		return fmt.Errorf("could not store leaf: %w", err)
	}

	// next up we encode the message so we can key it to a leaf.
	encodedMessage, err := committedMessage.Encode()
	if err != nil {
		return fmt.Errorf("could not encode committed message: %w", err)
	}

	// finally we store the encoded message
	err = d.StoreKeyedEncodable(MESSAGE, ToSlice(committedMessage.Leaf()), encodedMessage)
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

// MessageByNonce retrieves a raw committed message by its leaf hash.
func (d *pebbleDB) MessageByNonce(destination, nonce uint32) (types.CommittedMessage, error) {
	leaf, err := d.LeafByNonce(destination, nonce)
	if err != nil {
		return nil, fmt.Errorf("could not get leaf: %w", err)
	}

	message, err := d.MessageByLeaf(leaf)
	if err != nil {
		return nil, fmt.Errorf("could not get message by leaf: %w", err)
	}
	return message, nil
}

// MessageByLeaf fetches a message by leaf.
func (d *pebbleDB) MessageByLeaf(leaf common.Hash) (types.CommittedMessage, error) {
	value, _, err := d.DB.Get(d.FullKey(MESSAGE, leaf.Bytes()))
	if err != nil {
		return nil, fmt.Errorf("could not get message by leaf: %w", err)
	}

	committedMessage, err := types.DecodeCommittedMessage(value)
	if err != nil {
		return nil, fmt.Errorf("could not decode message: %w", err)
	}

	return committedMessage, nil
}

// LeafByNonce retrieves the leaf hash keyed by destination and nonce.
func (d *pebbleDB) LeafByNonce(destination, nonce uint32) (common.Hash, error) {
	destAndNonce := types.CombineDestinationAndNonce(destination, nonce)
	value, _, err := d.DB.Get(d.FullKey(LEAF, []byte(fmt.Sprintf("%d", destAndNonce))))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not get leaf by nonce: %w", err)
	}

	if len(value) != common.HashLength {
		return common.Hash{}, fmt.Errorf("leaves must have length of exactly %d, got %d (value: %s)", common.HashLength, len(value), value)
	}

	return common.BytesToHash(value), nil
}

// LeafByIndex gets a leaf by the index.
func (d *pebbleDB) LeafByIndex(leafIndex uint32) (common.Hash, error) {
	value, _, err := d.DB.Get(d.FullKey(LEAF, []byte(fmt.Sprintf("%d", leafIndex))))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not get leaf by nonce: %w", err)
	}

	if len(value) != common.HashLength {
		return common.Hash{}, fmt.Errorf("leaves must have length of exactly %d, got %d (value: %s)", common.HashLength, len(value), value)
	}

	return common.BytesToHash(value), nil
}

// MessageByLeafIndex retrieves a message by its leaf index.
func (d *pebbleDB) MessageByLeafIndex(leafIndex uint32) (types.CommittedMessage, error) {
	leaf, err := d.LeafByIndex(leafIndex)
	if err != nil {
		return nil, fmt.Errorf("could not get leaf at index %d: %w", leafIndex, err)
	}

	return d.MessageByLeaf(leaf)
}

func (d *pebbleDB) StoreProof(leafIndex uint32, proof types.Proof) error {
	encodedProof, err := proof.Encode()
	if err != nil {
		return fmt.Errorf("could not encode proof: %w", err)
	}

	err = d.StoreKeyedEncodable(PROOF, []byte(fmt.Sprintf("%d", leafIndex)), encodedProof)
	if err != nil {
		return fmt.Errorf("could not store proof: %w", err)
	}
	return nil
}

func (d *pebbleDB) ProofByLeafIndex(leafIndex uint32) (types.Proof, error) {
	dbProof, _, err := d.DB.Get(d.FullKey(PROOF, []byte(fmt.Sprintf("%d", leafIndex))))
	if err != nil {
		return nil, fmt.Errorf("could not get db proof: %w", err)
	}

	decodedProof, err := types.DecodeProof(dbProof)
	if err != nil {
		return nil, fmt.Errorf("could not decode proof: %w", err)
	}

	return decodedProof, nil
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

func (d *pebbleDB) GetMessageLatestBlockEnd() (height uint32, err error) {
	rawHeight, _, err := d.Get(d.FullKey(MessagesLastBlockEnd, []byte("")))
	if err != nil {
		return 0, fmt.Errorf("could not get height: %w", err)
	}

	uncastHeight, err := strconv.Atoi(string(rawHeight))
	if err != nil {
		return 0, fmt.Errorf("could not get indexed height: %w", err)
	}

	return uint32(uncastHeight), nil
}

// StoreMessageLatestBlockEnd stores the latest message block.
func (d *pebbleDB) StoreMessageLatestBlockEnd(height uint32) error {
	err := d.StoreKeyedEncodable(MessagesLastBlockEnd, []byte(""), []byte(fmt.Sprintf("%d", height)))
	if err != nil {
		return fmt.Errorf("could not store height: %w", err)
	}
	return nil
}

// ToSlice converts a kappa value toa  byte slice.
func ToSlice(kappa [32]byte) []byte {
	rawKappa := make([]byte, len(kappa))
	copy(rawKappa, kappa[:])
	return rawKappa
}
