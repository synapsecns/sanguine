//nolint:revive,golint
package messagebus

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/types/messagebus"
)

// GetEventType gets the execute event type.
func (m MessageBusUpgradeableExecuted) GetEventType() messagebus.EventType {
	return messagebus.ExecutedEvent
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

// GetMessageID gets the messagebus id for the event.
func (m MessageBusUpgradeableExecuted) GetMessageID() *string {
	messagebusID := common.Bytes2Hex(m.MessageId[:])
	return &messagebusID
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

// GetMessage gets the messagebus for the event.
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

// GetRevertReason gets the fee for the event.
func (m MessageBusUpgradeableExecuted) GetRevertReason() *string {
	return nil
}

var _ messagebus.EventLog = &MessageBusUpgradeableExecuted{}

// GetEventType gets the execute event type.
func (m MessageBusUpgradeableMessageSent) GetEventType() messagebus.EventType {
	return messagebus.MessageSentEvent
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

// GetMessageID gets the messagebus id for the event.
func (m MessageBusUpgradeableMessageSent) GetMessageID() *string {
	messagebusID := common.Bytes2Hex(m.MessageId[:])
	return &messagebusID
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

// GetMessage gets the messagebus for the event.
func (m MessageBusUpgradeableMessageSent) GetMessage() *string {
	messagebus := common.Bytes2Hex(m.Message)
	return &messagebus
}

// GetReceiver gets the receiver for the event.
func (m MessageBusUpgradeableMessageSent) GetReceiver() *string {
	receiver := common.Bytes2Hex(m.Receiver[:])
	return &receiver
}

// GetOptions gets the options for the event.
func (m MessageBusUpgradeableMessageSent) GetOptions() *string {
	options := common.Bytes2Hex(m.Options)
	return &options
}

// GetFee gets the fee for the event.
func (m MessageBusUpgradeableMessageSent) GetFee() *big.Int {
	return m.Fee
}

// GetRevertReason gets the fee for the event.
func (m MessageBusUpgradeableMessageSent) GetRevertReason() *string {
	return nil
}

var _ messagebus.EventLog = &MessageBusUpgradeableMessageSent{}

// GetEventType gets the execute event type.
func (m MessageBusUpgradeableCallReverted) GetEventType() messagebus.EventType {
	return messagebus.CallRevertedEvent
}

// GetRaw gets the raw logs.
func (m MessageBusUpgradeableCallReverted) GetRaw() ethTypes.Log {
	return m.Raw
}

// GetContractAddress gets the contract address the event occurred on.
func (m MessageBusUpgradeableCallReverted) GetContractAddress() common.Address {
	return m.Raw.Address
}

// GetBlockNumber gets the block number for the event.
func (m MessageBusUpgradeableCallReverted) GetBlockNumber() uint64 {
	return m.Raw.BlockNumber
}

// GetTxHash gets the unique identifier (txhash) for the Execute event.
func (m MessageBusUpgradeableCallReverted) GetTxHash() common.Hash {
	return m.Raw.TxHash
}

// GetEventIndex gets the block index of the event.
func (m MessageBusUpgradeableCallReverted) GetEventIndex() uint64 {
	return uint64(m.Raw.Index)
}

// GetMessageID gets the messagebus id for the event.
func (m MessageBusUpgradeableCallReverted) GetMessageID() *string {
	return nil
}

// GetSourceChainID gets the source chain id for the event.
func (m MessageBusUpgradeableCallReverted) GetSourceChainID() *big.Int {
	return nil
}

// GetSourceAddress gets the source address for the event.
func (m MessageBusUpgradeableCallReverted) GetSourceAddress() *string {
	return nil
}

// GetStatus gets the status for the event.
func (m MessageBusUpgradeableCallReverted) GetStatus() *string {
	return nil
}

// GetDestinationAddress gets the destination address for the event.
func (m MessageBusUpgradeableCallReverted) GetDestinationAddress() *string {
	return nil
}

// GetDestinationChainID gets the destination chain id for the event.
func (m MessageBusUpgradeableCallReverted) GetDestinationChainID() *big.Int {
	return nil
}

// GetNonce gets the source nonce for the event.
func (m MessageBusUpgradeableCallReverted) GetNonce() *big.Int {
	return nil
}

// GetMessage gets the messagebus for the event.
func (m MessageBusUpgradeableCallReverted) GetMessage() *string {
	return nil
}

// GetReceiver gets the receiver for the event.
func (m MessageBusUpgradeableCallReverted) GetReceiver() *string {
	return nil
}

// GetOptions gets the options for the event.
func (m MessageBusUpgradeableCallReverted) GetOptions() *string {
	return nil
}

// GetFee gets the fee for the event.
func (m MessageBusUpgradeableCallReverted) GetFee() *big.Int {
	return nil
}

// GetRevertReason gets the fee for the event.
func (m MessageBusUpgradeableCallReverted) GetRevertReason() *string {
	reason := m.Reason
	return &reason
}

var _ messagebus.EventLog = &MessageBusUpgradeableCallReverted{}
