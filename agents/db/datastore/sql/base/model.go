package base

import (
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

// SentMessage is used to store information about sent messages from the Origin contract.
// Monitoring uses these messages' nonces to check for missing messages on destination chains.
type SentMessage struct {
	// SMOrigin is the origin chainID of the message.
	SMOrigin uint32 `gorm:"column:origin"`
	// SMSender is the sender of the message.
	SMSender string `gorm:"column:sender"`
	// SMNonce is the nonce of the message.
	SMNonce uint32 `gorm:"column:nonce"`
	// SMDestination is the destination chainID of the message.
	SMDestination uint32 `gorm:"column:destination"`
	// SMRecipient is the recipient of the message.
	SMRecipient string `gorm:"column:recipient"`
	// SMOptimisticSeconds is the optimistic seconds of the message.
	SMOptimisticSeconds uint32 `gorm:"column:optimistic_seconds"`
	// SMNotaryTip is the notary tip of the message.
	SMNotaryTip []byte `gorm:"column:notary_tip"`
	// SMBroadcasterTip is the broadcaster tip of the message.
	SMBroadcasterTip []byte `gorm:"column:broadcaster_tip"`
	// SMProverTip is the prover tip of the message.
	SMProverTip []byte `gorm:"column:prover_tip"`
	// SMExecutorTip is the executor tip of the message.
	SMExecutorTip []byte `gorm:"column:executor_tip"`
	// SMBody is the body of the message.
	SMBody []byte `gorm:"column:body"`
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
	CMFlag uint8 `gorm:"column:cm_flag"`
	// CMDomainID is the id of the domain we're renaming
	CMDomainID uint32 `gorm:"column:domain_id;uniqueIndex:cm_idx_id"`
	// CMMessage is the fully detailed message that was created
	CMMessage []byte `gorm:"column:message"`
	// CMLeaf is the leaf
	CMLeaf []byte `gorm:"column:leaf"`
	// CMOrigin returns the slip-44 of the message
	CMOrigin uint32 `gorm:"column:origin"`
	// CMNonce is the nonce of the message
	CMNonce uint32 `gorm:"column:nonce;uniqueIndex:cm_idx_id"`
	// CMDestination is the sip-44 destination of the message
	CMDestination uint32 `gorm:"column:destination"`
	// CMBody is the body of the message
	CMBody []byte `gorm:"column:body"`
	// CMOptimisticSeconds is the optimistic seconds of the message
	CMOptimisticSeconds uint32 `gorm:"column:optimistic_seconds"`
}

// Flag gets the message flag.
func (c CommittedMessage) Flag() types.MessageFlag {
	return types.MessageFlag(c.CMFlag)
}

// Header gets the header.
func (c CommittedMessage) Header() types.Header {
	return types.NewHeader(c.OriginDomain(), c.Nonce(), c.DestinationDomain(), c.OptimisticSeconds())
}

// OriginDomain returns the Slip-44 ID.
func (c CommittedMessage) OriginDomain() uint32 {
	return c.CMOrigin
}

// Nonce is the count of all previous messages to the destination.
func (c CommittedMessage) Nonce() uint32 {
	return c.CMNonce
}

// DestinationDomain is the slip-44 id of the destination.
func (c CommittedMessage) DestinationDomain() uint32 {
	return c.CMDestination
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

// BaseMessage gets the base message if it exists.
func (c CommittedMessage) BaseMessage() types.BaseMessage {
	if types.MessageFlag(c.CMFlag) == types.MessageFlagManager {
		return nil
	}
	baseMessage, err := types.DecodeBaseMessage(c.CMMessage)
	if err != nil {
		return nil
	}

	return baseMessage
}

// Message gets the message.
func (c CommittedMessage) Message() []byte {
	if types.MessageFlag(c.CMFlag) == types.MessageFlagBase {
		return []byte{}
	}
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
