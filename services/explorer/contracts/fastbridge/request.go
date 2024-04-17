package fastbridge

import "C"
import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/fastbridge"
)

// GetTxHash gets the tx hash for the event.
func (e FastBridgeBridgeRequested) GetTxHash() common.Hash {
	return e.Raw.TxHash
}

// GetContractAddress gets the contract address the event occurred on.
func (e FastBridgeBridgeRequested) GetContractAddress() common.Address {
	return e.Raw.Address
}

// GetBlockNumber gets the block number for the event.
func (e FastBridgeBridgeRequested) GetBlockNumber() uint64 {
	return e.Raw.BlockNumber
}

// GetEventType gets the event type for the event.
func (e FastBridgeBridgeRequested) GetEventType() fastbridge.EventType {
	return fastbridge.BridgeRequestedEvent
}

// GetEventIndex gets the event index for the event.
func (e FastBridgeBridgeRequested) GetEventIndex() uint64 {
	return uint64(e.Raw.TxIndex)
}

// GetTransactionID gets the transaction id for the event.
func (e FastBridgeBridgeRequested) GetTransactionID() [32]byte {
	return e.TransactionId
}

// GetSender gets the sender for the event.
func (e FastBridgeBridgeRequested) GetSender() *string {
	str := e.Sender.String()
	return &str
}

// GetRequest gets the request for the event.
func (e FastBridgeBridgeRequested) GetRequest() *[]byte {
	return &e.Request
}

// GetOriginChainID gets the origin chain id for the event.
func (e FastBridgeBridgeRequested) GetOriginChainID() *big.Int {
	return nil
}

// GetDestChainID gets the destination chain id for the event.
func (e FastBridgeBridgeRequested) GetDestChainID() *big.Int {
	return big.NewInt(int64(e.DestChainId))
}

// GetOriginToken gets the origin token for the event.
func (e FastBridgeBridgeRequested) GetOriginToken() common.Address {
	return e.OriginToken
}

// GetDestToken gets the destination token for the event.
func (e FastBridgeBridgeRequested) GetDestToken() common.Address {
	return e.DestToken
}

// GetOriginAmount gets the origin amount for the event.
func (e FastBridgeBridgeRequested) GetOriginAmount() *big.Int {
	return e.OriginAmount
}

// GetDestAmount gets the destination amount for the event.
func (e FastBridgeBridgeRequested) GetDestAmount() *big.Int {
	return e.DestAmount
}

// GetSendChainGas gets the send chain gas for the event.
func (e FastBridgeBridgeRequested) GetSendChainGas() *bool {
	return &e.SendChainGas
}

// GetChainGasAmount gets the chain gas amount for the event.
func (e FastBridgeBridgeRequested) GetChainGasAmount() *big.Int {
	return nil
}

// GetRelayer gets the relayer address for the event.
func (e FastBridgeBridgeRequested) GetRelayer() *string {
	return nil
}

// GetTo gets the to address for the event.
func (e FastBridgeBridgeRequested) GetTo() *string {
	return nil
}

// GetTxHash gets the tx hash for the event.
func (e FastBridgeBridgeRelayed) GetTxHash() common.Hash {
	return e.Raw.TxHash
}

// GetContractAddress gets the contract address the event occurred on.
func (e FastBridgeBridgeRelayed) GetContractAddress() common.Address {
	return e.Raw.Address
}

// GetBlockNumber gets the block number for the event.
func (e FastBridgeBridgeRelayed) GetBlockNumber() uint64 {
	return e.Raw.BlockNumber
}

// GetEventType gets the event type for the event.
func (e FastBridgeBridgeRelayed) GetEventType() fastbridge.EventType {

	return fastbridge.BridgeRelayedEvent
}

// GetEventIndex gets the event index for the event.
func (e FastBridgeBridgeRelayed) GetEventIndex() uint64 {
	return uint64(e.Raw.TxIndex)
}

// GetTransactionID gets the transaction id for the event.
func (e FastBridgeBridgeRelayed) GetTransactionID() [32]byte {
	return e.TransactionId
}

// GetRelayer gets the relayer address for the event.
func (e FastBridgeBridgeRelayed) GetRelayer() *string {
	str := e.Relayer.String()
	return &str
}

// GetTo gets the to address for the event.
func (e FastBridgeBridgeRelayed) GetTo() *string {
	str := e.To.String()
	return &str
}

// GetOriginChainID gets the origin chain id for the event.
func (e FastBridgeBridgeRelayed) GetOriginChainID() *big.Int {
	return big.NewInt(int64(e.OriginChainId))
}

// GetDestChainID gets the destination chain id for the event.
func (e FastBridgeBridgeRelayed) GetDestChainID() *big.Int {
	return nil
}

// GetOriginToken gets the origin token for the event.
func (e FastBridgeBridgeRelayed) GetOriginToken() common.Address {
	return e.OriginToken
}

// GetDestToken gets the destination token for the event.
func (e FastBridgeBridgeRelayed) GetDestToken() common.Address {
	return e.DestToken
}

// GetOriginAmount gets the origin amount for the event.
func (e FastBridgeBridgeRelayed) GetOriginAmount() *big.Int {
	return e.OriginAmount
}

// GetDestAmount gets the destination amount for the event.
func (e FastBridgeBridgeRelayed) GetDestAmount() *big.Int {
	return e.DestAmount
}

// GetChainGasAmount gets the chain gas amount for the event.
func (e FastBridgeBridgeRelayed) GetChainGasAmount() *big.Int {
	return e.ChainGasAmount
}

// GetRequest gets the request for the event.
func (e FastBridgeBridgeRelayed) GetRequest() *[]byte {
	return nil
}

// GetSendChainGas gets the send chain gas for the event.
func (e FastBridgeBridgeRelayed) GetSendChainGas() *bool {
	return nil
}

// GetSender gets the sender for the event.
func (e FastBridgeBridgeRelayed) GetSender() *string {
	return nil
}
