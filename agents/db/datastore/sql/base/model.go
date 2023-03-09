package base

import (
	"math/big"
	"time"

	"github.com/synapsecns/sanguine/core/dbcommon"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/types"
	"gorm.io/gorm"
)

// define common field names. See package docs  for an explanation of why we have to do this.
// note: some models share names. In cases where they do, we run the check against all names.
// This is cheap because it's only done at startup.
func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	NonceFieldName = namer.GetConsistentName("Nonce")
	DomainIDFieldName = namer.GetConsistentName("DomainID")
	BlockNumberFieldName = namer.GetConsistentName("BlockNumber")
	OriginFieldName = namer.GetConsistentName("IPOrigin")
	DestinationFieldName = namer.GetConsistentName("IPDestination")
	AttestationStateFieldName = namer.GetConsistentName("IPAttestationState")
}

var (
	// NonceFieldName is the field name of the nonce.
	NonceFieldName string
	// DomainIDFieldName gets the chain id field name.
	DomainIDFieldName string
	// BlockNumberFieldName is the name of the block number field.
	BlockNumberFieldName string
	// LeafIndexFieldName is the field name of the leaf index.
	LeafIndexFieldName string
	// OriginFieldName is the name of the origin field.
	OriginFieldName string
	// DestinationFieldName is the name of the destination field.
	DestinationFieldName string
	// AttestationStateFieldName is the name of the attestation state field.
	AttestationStateFieldName string
)

// RawEthTX contains a raw evm transaction that is unsigned
// note: idx_id contains a composite index of (chain_id,nonce)
type RawEthTX struct {
	gorm.Model
	// From is the sender of the transaction
	From string `gorm:"from"`
	// To is the contract address the transaction was sent to.
	To string `gorm:"index"`
	// ChainID is the chain id the transaction hash will be sent on
	ChainID uint64 `gorm:"column:chain_id;uniqueIndex:idx_id"`
	// Nonce is the nonce of the raw evm tx
	Nonce uint64 `gorm:"column:nonce;uniqueIndex:idx_id"`
	// RawTx is the raw serialized transaction
	RawTx []byte `gorm:"column:raw_tx"`
}

// ProcessedEthTx contains a processed ethereum transaction.
type ProcessedEthTx struct {
	TxHash string `gorm:"txhash;uniqueIndex:idx_txhash;size:66"`
	// RawTx is the raw serialized transaction
	RawTx []byte `gorm:"column:raw_tx"`
	// RawEthTx is the txid that caused the event
	RawEthTx uint
	// OriginatingEvent is the event that originated the tx
	EthTx RawEthTX `gorm:"foreignkey:RawEthTx"`
	// GasFeeCap contains the gas fee cap stored in wei
	GasFeeCap uint64
	// GasTipCap contains the gas tip cap stored in wei
	GasTipCap uint64
}

// BlockEndModel is used to make sure we haven't missed any events while offline.
// since we event source - rather than use a state machine this is needed to make sure we haven't missed any events
// by allowing us to go back and source any events we may have missed.
//
// this does not inherit from gorm.model to allow us to use ChainID as a primary key.
type BlockEndModel struct {
	// CreatedAt is the creation time
	CreatedAt time.Time
	// UpdatedAt is the update time
	UpdatedAt time.Time
	// DeletedAt time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// DomainID is the chain id of the chain we're watching blocks on. This is our primary index.
	DomainID uint32 `gorm:"column:domain_id;primaryKey;autoIncrement:false"`
	// BlockHeight is the highest height we've seen on the chain
	BlockNumber uint32 `gorm:"block_number"`
}

// DispatchMessage is used to store information about dispatched messages from the Origin contract.
// Monitoring uses these messages' nonces to check for missing messages on destination chains.
type DispatchMessage struct {
	// DMOrigin is the origin chainID of the message.
	DMOrigin uint32 `gorm:"column:origin"`
	// DMSender is the sender of the message.
	DMSender string `gorm:"column:sender"`
	// DMNonce is the nonce of the message.
	DMNonce uint32 `gorm:"column:nonce"`
	// DMDestination is the destination chainID of the message.
	DMDestination uint32 `gorm:"column:destination"`
	// DMRecipient is the recipient of the message.
	DMRecipient string `gorm:"column:recipient"`
	// DMOptimisticSeconds is the optimistic seconds of the message.
	DMOptimisticSeconds uint32 `gorm:"column:optimistic_seconds"`
	// DMNotaryTip is the notary tip of the message.
	DMNotaryTip []byte `gorm:"column:notary_tip"`
	// DMBroadcasterTip is the broadcaster tip of the message.
	DMBroadcasterTip []byte `gorm:"column:broadcaster_tip"`
	// DMProverTip is the prover tip of the message.
	DMProverTip []byte `gorm:"column:prover_tip"`
	// DMExecutorTip is the executor tip of the message.
	DMExecutorTip []byte `gorm:"column:executor_tip"`
	// DMBody is the body of the message.
	DMBody []byte `gorm:"column:body"`
}

// AcceptedAttestation is used to track every received accepted attestation over all mirrors.
// Monitoring uses these accepted attestations' nonces to check for missing messages on destination chains.
type AcceptedAttestation struct {
	// AAOrigin is the chainID of the Origin contract.
	AAOrigin uint32 `gorm:"column:origin"`
	// AADestination is the chainID of the Destination contract.
	AADestination uint32 `gorm:"column:destination"`
	// AANonce is the nonce of the attestation.
	AANonce uint32 `gorm:"column:nonce"`
	// AARoot is the root of the attestation.
	AARoot string `gorm:"column:root"`
}

// CommittedMessage is a committed message
// it allows for querying on both the committed message and the underlying fields.
type CommittedMessage struct {
	gorm.Model
	CMVersion uint16 `gorm:"column:cm_version"`
	// CMDomainID is the id of the domain we're renaming
	CMDomainID uint32 `gorm:"column:domain_id;uniqueIndex:cm_idx_id"`
	// CMMessage is the fully detailed message that was created
	CMMessage []byte `gorm:"column:message"`
	// CMLeaf is the leaf
	CMLeaf []byte `gorm:"column:leaf"`
	// CMOrigin returns the slip-44 of the message
	CMOrigin uint32 `gorm:"column:origin"`
	// CMSender is the sender of the message
	CMSender []byte `gorm:"column:sender"`
	// CMNonce is the nonce of the message
	CMNonce uint32 `gorm:"column:nonce;uniqueIndex:cm_idx_id"`
	// CMDestination is the sip-44 destination of the message
	CMDestination uint32 `gorm:"column:destination"`
	// CMRecipient is the recipient of the message
	CMRecipient []byte `gorm:"column:recipient"`
	// CMBody is the body of the message
	CMBody []byte `gorm:"column:body"`
	// CMOptimisticSeconds is the optimistic seconds of the message
	CMOptimisticSeconds uint32 `gorm:"column:optimistic_seconds"`
	// CMNotaryTip is the notarytip
	CMNotaryTip []byte `gorm:"column:notary_tip"`
	// CMBroadcasterTip is the relayer tip
	CMBroadcasterTip []byte `gorm:"column:broadcaster_tip"`
	// CMProverTip is the prover tip
	CMProverTip []byte `gorm:"column:prover_tip"`
	// CMExecutorTip is the processor tip
	CMExecutorTip []byte `gorm:"column:executor_tip"`
}

// Version gets the message version.
func (c CommittedMessage) Version() uint16 {
	return c.CMVersion
}

// Header gets the header.
func (c CommittedMessage) Header() types.Header {
	return types.NewHeader(c.OriginDomain(), c.Sender(), c.Nonce(), c.DestinationDomain(), c.Recipient(), c.OptimisticSeconds())
}

// Tips gets the tips.
func (c CommittedMessage) Tips() types.Tips {
	return types.NewTips(new(big.Int).SetBytes(c.CMNotaryTip), new(big.Int).SetBytes(c.CMBroadcasterTip), new(big.Int).SetBytes(c.CMProverTip), new(big.Int).SetBytes(c.CMExecutorTip))
}

// OriginDomain returns the Slip-44 ID.
func (c CommittedMessage) OriginDomain() uint32 {
	return c.CMOrigin
}

// Sender is the address of the sender.
func (c CommittedMessage) Sender() common.Hash {
	return common.BytesToHash(c.CMSender)
}

// Nonce is the count of all previous messages to the destination.
func (c CommittedMessage) Nonce() uint32 {
	return c.CMNonce
}

// DestinationDomain is the slip-44 id of the destination.
func (c CommittedMessage) DestinationDomain() uint32 {
	return c.CMDestination
}

// Recipient is the address of the recipient.
func (c CommittedMessage) Recipient() common.Hash {
	return common.BytesToHash(c.CMRecipient)
}

// Body is the message contents.
func (c CommittedMessage) Body() []byte {
	return c.CMBody
}

// ToLeaf converts a leaf to a keccac256.
func (c CommittedMessage) ToLeaf() (leaf [32]byte, err error) {
	return common.BytesToHash(c.CMLeaf), nil
}

// OptimisticSeconds gets the optimistic seconds count.
func (c CommittedMessage) OptimisticSeconds() uint32 {
	return c.CMOptimisticSeconds
}

// Message gets the message.
func (c CommittedMessage) Message() []byte {
	return c.CMMessage
}

// Leaf gets the leaf.
func (c CommittedMessage) Leaf() [32]byte {
	return common.BytesToHash(c.CMLeaf)
}

// Encode encodes the message
// Deprecated: will be removed.
func (c CommittedMessage) Encode() ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

var _ types.CommittedMessage = CommittedMessage{}

var _ types.Message = CommittedMessage{}
