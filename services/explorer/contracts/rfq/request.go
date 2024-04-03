package rfq

import "C"
import (
	"math/big"

	"github.com/synapsecns/sanguine/services/explorer/types/rfq"

	"github.com/ethereum/go-ethereum/common"
)

func (e SynapseRFQBridgeRequestedEvent) GetTxHash() common.Hash {
	return e.Raw.TxHash
}

func (e SynapseRFQBridgeRequestedEvent) GetContractAddress() common.Address {
	return e.Raw.Address
}
func (e SynapseRFQBridgeRequestedEvent) GetBlockNumber() uint64 {
	return e.Raw.BlockNumber
}

func (e SynapseRFQBridgeRequestedEvent) GetEventType() rfq.EventType {

	return rfq.BridgeRequestedEvent
}

func (e SynapseRFQBridgeRequestedEvent) GetEventIndex() uint64 {
	return uint64(e.Raw.TxIndex)
}

func (e SynapseRFQBridgeRequestedEvent) GetTransactionID() [32]byte {
	return e.TransactionID
}

func (e SynapseRFQBridgeRequestedEvent) GetSender() string {
	str := e.Sender.String()
	return &str
}

func (e SynapseRFQBridgeRequestedEvent) GetRequest() *[]byte {
	return e.Request
}

func (e SynapseRFQBridgeRequestedEvent) GetDestinationChainID() *big.Int {
	return e.DestinationChainID
}

func (e SynapseRFQBridgeRequestedEvent) GetOriginToken() *string {
	return &e.OriginToken.string()
}

func (e SynapseRFQBridgeRequestedEvent) GetDestinationToken() *string {
	return e.DestinationToken.string()
}

func (e SynapseRFQBridgeRequestedEvent) GetOriginAmount() *big.Int {
	return e.OriginAmount
}

func (e SynapseRFQBridgeRequestedEvent) GetDestinationAmount() *big.Int {
	return e.DestinationAmount
}

func (e SynapseRFQBridgeRequestedEvent) SendChainGas() bool {
	return e.SendChainGas
}

var _ rfq.EventLog = &SynapseRFQBridgeRequestedEvent{}

func (e SynapseRFQBridgeRelayedEvent) GetTxHash() common.Hash {
	return e.Raw.TxHash
}

func (e SynapseRFQBridgeRelayedEvent) GetContractAddress() common.Address {
	return e.Raw.Address
}
func (e SynapseRFQBridgeRelayedEvent) GetBlockNumber() uint64 {
	return e.Raw.BlockNumber
}

func (e SynapseRFQBridgeRelayedEvent) GetEventType() rfq.EventType {

	return rfq.BridgeRelayedEvent
}

func (e SynapseRFQBridgeRelayedEvent) GetEventIndex() uint64 {
	return uint64(e.Raw.TxIndex)
}

func (e SynapseRFQBridgeRelayedEvent) GetTransactionID() [32]byte {
	return e.TransactionID
}

func (e SynapseRFQBridgeRequestedEvent) GetRelayer() string {
	return e.Relayer
}

func (e SynapseRFQBridgeRelayedEvent) GetRecipient() string {
	str := e.Recipient.String()
	return &str
}

func (e SynapseRFQBridgeRelayedEvent) GetOriginChainID() *big.Int {
	return e.OriginChainID
}

func (e SynapseRFQBridgeRelayedEvent) GetOriginToken() *string {
	return &e.OriginToken.string()
}

func (e SynapseRFQBridgeRelayedEvent) GetDestinationToken() *string {
	return e.DestinationToken.string()
}

func (e SynapseRFQBridgeRelayedEvent) GetOriginAmount() *big.Int {
	return e.OriginAmount
}

func (e SynapseRFQBridgeRelayedEvent) GetDestinationAmount() *big.Int {
	return e.DestinationAmount
}

func (e SynapseRFQBridgeRelayedEvent) GetChainGasAmount() *big.Int {
	return e.ChainGasAmount
}

var _ rfq.EventLog = &SynapseRFQBridgeRelayedEvent{}
