package types

import (
	"github.com/ethereum/go-ethereum/common"
	common2 "github.com/synapsecns/synapse-node/pkg/common"
	"math/big"
)

// SignedUpdateWithMeta gets the signed update with meta data.
type SignedUpdateWithMeta interface {
	// SignedUpdate gets the signed update
	SignedUpdate() SignedUpdate
	// Metadata gets the metadata
	Metadata() UpdateMeta
}

// NewSignedUpdateWithMeta returns a struct that conforms to SignedUpdateWithMeta.
func NewSignedUpdateWithMeta(signedUpdate SignedUpdate, metadata UpdateMeta) SignedUpdateWithMeta {
	return signedUpdateWithMeta{
		signedUpdate: signedUpdate,
		metadata:     metadata,
	}
}

type signedUpdateWithMeta struct {
	signedUpdate SignedUpdate
	metadata     UpdateMeta
}

func (s signedUpdateWithMeta) SignedUpdate() SignedUpdate {
	return s.signedUpdate
}

func (s signedUpdateWithMeta) Metadata() UpdateMeta {
	return s.metadata
}

var _ SignedUpdateWithMeta = signedUpdateWithMeta{}

// SignedUpdate is an update.
type SignedUpdate interface {
	// Update is the update
	Update() Update
	// Signature is the signature of the update
	Signature() Signature
}

// NewSignedUpdate gets a signed update.
func NewSignedUpdate(update Update, signature Signature) SignedUpdate {
	return signedUpdate{
		update:    update,
		signature: signature,
	}
}

// signedUpdate is a struct that conforms to SignedUpdate.
type signedUpdate struct {
	update    Update
	signature Signature
}

func (s signedUpdate) Update() Update {
	return s.update
}

func (s signedUpdate) Signature() Signature {
	return s.signature
}

var _ SignedUpdate = signedUpdate{}

// Update gets the update.
type Update interface {
	// HomeDomain gets the home chain
	HomeDomain() uint32
	// PreviousRoot gets the hash of the old merkle root
	PreviousRoot() common.Hash
	// NewRoot gets the new root of the tree
	NewRoot() common.Hash
}

type update struct {
	// homeDomain is the domain of the home contract
	homeDomain uint32
	// previousRoot is the last merkle root
	previousRoot common.Hash
	// newRoot is the next merkle root
	newRoot common.Hash
}

// NewUpdate creates a new update type.
func NewUpdate(homeDomain uint32, previousRoot, newRoot common.Hash) Update {
	return update{
		homeDomain:   homeDomain,
		previousRoot: previousRoot,
		newRoot:      newRoot,
	}
}

// HomeDomain gets the home domain.
func (u update) HomeDomain() uint32 {
	return u.homeDomain
}

// PreviousRoot gets the previous root.
func (u update) PreviousRoot() common.Hash {
	return u.previousRoot
}

// NewRoot gets the root of the update.
func (u update) NewRoot() common.Hash {
	return u.newRoot
}

var _ Update = update{}

// UpdateMeta gets metadata stored about the update.
type UpdateMeta interface {
	// BlockNumber gets the block number
	BlockNumber() uint64
	// Timestamp of the block
	Timestamp() *uint64
}

// Signature creates a new signature.
type Signature interface {
	// V gets the v value of the signature
	V() *big.Int
	// R is the r value of the signature
	R() *big.Int
	// S is the s value of the signature
	S() *big.Int
}

// signature contains an ecdsa signature
// one of the reasons we use interfaces here is to ensure the underlying data structures
// are not accidentally mutated. To ensure this, we copy big ints before returning.
type signature struct {
	v, r, s *big.Int
}

// NewSignature creates a new signature.
func NewSignature(v, r, s *big.Int) Signature {
	return signature{
		v: common2.CopyBigInt(v),
		r: common2.CopyBigInt(r),
		s: common2.CopyBigInt(s),
	}
}

func (s signature) V() *big.Int {
	return common2.CopyBigInt(s.v)
}

func (s signature) R() *big.Int {
	return common2.CopyBigInt(s.r)
}

func (s signature) S() *big.Int {
	return common2.CopyBigInt(s.s)
}

var _ Signature = signature{}

type updateMeta struct {
	blockNumber uint64
	timestamp   *uint64
}

// NewUpdateMeta gets an update meta object.
func NewUpdateMeta(blockNumber uint64, timestamp *uint64) UpdateMeta {
	return &updateMeta{
		blockNumber: blockNumber,
		timestamp:   timestamp,
	}
}

// BlockNumber gets the block number of the update.
func (u updateMeta) BlockNumber() uint64 {
	return u.blockNumber
}

// Timestamp gets the timestamp of the update. This can be null.
func (u updateMeta) Timestamp() *uint64 {
	return u.timestamp
}

var _ UpdateMeta = updateMeta{}
