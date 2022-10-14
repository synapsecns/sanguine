//nolint:revive,golint
package message

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/types/message"
	"math/big"
)

// GetEventType gets the execute event type.
func (m MessageBusUpgradeableExecuted) GetEventType() message.EventType {
	return message.ExecutedEvent
}

// GetRaw gets the raw logs.
func (m MessageBusUpgradeableExecuted) GetRaw() ethTypes.Log {
	return m.Raw
}

// GetContractAddress gets the contract address the event occurred on.
func (m MessageBusUpgradeableExecuted) GetContractAddress() common.Address {
	return m.Raw.Address
}

// GetBlockNumber gets the block number for the event.
func (m MessageBusUpgradeableExecuted) GetBlockNumber() uint64 {
	return m.Raw.BlockNumber
}

// GetTxHash gets the unique identifier (txhash) for the Execute event.
func (m MessageBusUpgradeableExecuted) GetTxHash() common.Hash {
	return m.Raw.TxHash
}

// GetEventIndex gets the block index of the event.
func (m MessageBusUpgradeableExecuted) GetEventIndex() uint64 {
	return uint64(m.Raw.Index)
}

// GetMessageId gets the message id for the event.
func (m MessageBusUpgradeableExecuted) GetMessageId() string {
	return common.Bytes2Hex(m.MessageId[:])
}

// GetSourceChainID gets the source chain id for the event.
func (m MessageBusUpgradeableExecuted) GetSourceChainID() *big.Int {
	return big.NewInt(int64(m.SrcChainId))
}

// GetSourceAddress gets the source address for the event.
func (m MessageBusUpgradeableExecuted) GetSourceAddress() *string {
	return nil
}

// GetStatus gets the status for the event.
func (m MessageBusUpgradeableExecuted) GetStatus() *string {
	txStatus := []string{"Null", "Success", "Fail"}
	return &txStatus[m.Status]
}

// GetDestinationAddress gets the destination address for the event.
func (m MessageBusUpgradeableExecuted) GetDestinationAddress() *string {
	destinationAddress := m.DstAddress.String()
	return &destinationAddress
}

// GetDestinationChainID gets the destination chain id for the event.
func (m MessageBusUpgradeableExecuted) GetDestinationChainID() *big.Int {
	return nil
}

// GetNonce gets the source nonce for the event.
func (m MessageBusUpgradeableExecuted) GetNonce() *big.Int {
	return big.NewInt(int64(m.SrcNonce))
}

// GetMessage gets the message for the event.
func (m MessageBusUpgradeableExecuted) GetMessage() *string {
	return nil
}

// GetReceiver gets the receiver for the event.
func (m MessageBusUpgradeableExecuted) GetReceiver() *string {
	return nil
}

// GetOptions gets the options for the event.
func (m MessageBusUpgradeableExecuted) GetOptions() *string {
	return nil
}

// GetFee gets the fee for the event.
func (m MessageBusUpgradeableExecuted) GetFee() *big.Int {
	return nil
}

var _ message.EventLog = &MessageBusUpgradeableExecuted{}

// GetEventType gets the execute event type.
func (m MessageBusUpgradeableMessageSent) GetEventType() message.EventType {
	return message.ExecutedEvent
}

// GetRaw gets the raw logs.
func (m MessageBusUpgradeableMessageSent) GetRaw() ethTypes.Log {
	return m.Raw
}

// GetContractAddress gets the contract address the event occurred on.
func (m MessageBusUpgradeableMessageSent) GetContractAddress() common.Address {
	return m.Raw.Address
}

// GetBlockNumber gets the block number for the event.
func (m MessageBusUpgradeableMessageSent) GetBlockNumber() uint64 {
	return m.Raw.BlockNumber
}

// GetTxHash gets the unique identifier (txhash) for the Execute event.
func (m MessageBusUpgradeableMessageSent) GetTxHash() common.Hash {
	return m.Raw.TxHash
}

// GetEventIndex gets the block index of the event.
func (m MessageBusUpgradeableMessageSent) GetEventIndex() uint64 {
	return uint64(m.Raw.Index)
}

// GetMessageId gets the message id for the event.
func (m MessageBusUpgradeableMessageSent) GetMessageId() string {
	return common.Bytes2Hex(m.MessageId[:])
}

// GetSourceChainID gets the source chain id for the event.
func (m MessageBusUpgradeableMessageSent) GetSourceChainID() *big.Int {
	return m.SrcChainID
}

// GetSourceAddress gets the source address for the event.
func (m MessageBusUpgradeableMessageSent) GetSourceAddress() *string {
	sourceAddress := m.Sender.String()
	return &sourceAddress
}

// GetStatus gets the status for the event.
func (m MessageBusUpgradeableMessageSent) GetStatus() *string {
	return nil
}

// GetDestinationAddress gets the destination address for the event.
func (m MessageBusUpgradeableMessageSent) GetDestinationAddress() *string {
	return nil
}

// GetDestinationChainID gets the destination chain id for the event.
func (m MessageBusUpgradeableMessageSent) GetDestinationChainID() *big.Int {
	return m.DstChainId
}

// GetNonce gets the source nonce for the event.
func (m MessageBusUpgradeableMessageSent) GetNonce() *big.Int {
	return big.NewInt(int64(m.Nonce))
}

// GetMessage gets the message for the event.
func (m MessageBusUpgradeableMessageSent) GetMessage() *string {
	message := common.Bytes2Hex(m.Message[:])
	return &message
}

// GetReceiver gets the receiver for the event.
func (m MessageBusUpgradeableMessageSent) GetReceiver() *string {
	receiver := common.Bytes2Hex(m.Receiver[:])
	return &receiver
}

// GetOptions gets the options for the event.
func (m MessageBusUpgradeableMessageSent) GetOptions() *string {
	options := common.Bytes2Hex(m.Options[:])
	return &options
}

// GetFee gets the fee for the event.
func (m MessageBusUpgradeableMessageSent) GetFee() *big.Int {
	return m.Fee
}

var _ message.EventLog = &MessageBusUpgradeableMessageSent{}
