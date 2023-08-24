package cctp

import "C"
import (
	"github.com/synapsecns/sanguine/services/explorer/types/cctp"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseCCTPCircleRequestSent) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetBlockNumber gets the block number for the event.
func (s SynapseCCTPCircleRequestSent) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetTxHash gets the unique identifier (txhash) for the event.
func (s SynapseCCTPCircleRequestSent) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetEventType gets the event type for the event.
func (s SynapseCCTPCircleRequestSent) GetEventType() cctp.EventType {
	return cctp.CircleRequestSentEvent
}

// GetRequestID gets the unique identifier of the request.
func (s SynapseCCTPCircleRequestSent) GetRequestID() [32]byte {
	return s.RequestID
}

// GetOriginChainID gets the origin chain ID for the event.
func (s SynapseCCTPCircleRequestSent) GetOriginChainID() *big.Int {
	return nil
}

// GetDestinationChainID gets the destination chain ID for the event.
func (s SynapseCCTPCircleRequestSent) GetDestinationChainID() *big.Int {
	return s.ChainId
}

// GetSender gets the sender of the CCTP tokens on the origin chain.
func (s SynapseCCTPCircleRequestSent) GetSender() *string {
	str := s.Sender.String()
	return &str
}

// GetNonce gets the nonce of the CCTP message on the origin chain.
func (s SynapseCCTPCircleRequestSent) GetNonce() *uint64 {
	return &s.Nonce
}

// GetMintToken gets the address of the minted Circle token.
func (s SynapseCCTPCircleRequestSent) GetMintToken() *string {
	return nil
}

// GetSentAmount gets the amount of Circle tokens burnt.
func (s SynapseCCTPCircleRequestSent) GetSentAmount() *big.Int {
	return s.Amount
}

// GetAmount gets the amount from the transfer.
func (s SynapseCCTPCircleRequestSent) GetAmount() *big.Int {
	return s.Amount
}

// GetRequestVersion gets the version of the request format.
func (s SynapseCCTPCircleRequestSent) GetRequestVersion() *uint32 {
	return &s.RequestVersion
}

// GetFormattedRequest gets the formatted request for the action to take on the destination chain.
func (s SynapseCCTPCircleRequestSent) GetFormattedRequest() *[]byte {
	return &s.FormattedRequest
}

// GetRecipient gets the end recipient of the tokens on this chain.
func (s SynapseCCTPCircleRequestSent) GetRecipient() *string {
	return nil
}

// GetFee gets the fee paid for fulfilling the request, in minted tokens.
func (s SynapseCCTPCircleRequestSent) GetFee() *big.Int {
	return nil
}

// GetToken gets the address of the token that the recipient received.
func (s SynapseCCTPCircleRequestSent) GetToken() string {
	return s.Token.String()
}

// GetEventIndex gets the tx index of the event in the block it was executed in.
func (s SynapseCCTPCircleRequestSent) GetEventIndex() uint64 {
	return uint64(s.Raw.TxIndex)
}

var _ cctp.EventLog = &SynapseCCTPCircleRequestSent{}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseCCTPCircleRequestFulfilled) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetBlockNumber gets the block number for the event.
func (s SynapseCCTPCircleRequestFulfilled) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetTxHash gets the unique identifier (txhash) for the event.
func (s SynapseCCTPCircleRequestFulfilled) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetEventType gets the event type for the event.
func (s SynapseCCTPCircleRequestFulfilled) GetEventType() cctp.EventType {
	return cctp.CircleRequestFulfilledEvent
}

// GetRequestID gets the unique identifier of the request.
func (s SynapseCCTPCircleRequestFulfilled) GetRequestID() [32]byte {
	return s.RequestID
}

// GetOriginChainID gets the origin chain ID for the event.
func (s SynapseCCTPCircleRequestFulfilled) GetOriginChainID() *big.Int {
	// domain to chain mapping TODO move to static mapping
	domainToChain := []int64{1, 43114, 10, 42161}
	if s.OriginDomain >= uint32(len(domainToChain)) { // Catch if the domain is not in the mapping (explorer lagging behind addition of new chains)
		return big.NewInt(int64(s.OriginDomain))
	}
	return big.NewInt(domainToChain[s.OriginDomain])
}

// GetDestinationChainID gets the destination chain ID for the event.
func (s SynapseCCTPCircleRequestFulfilled) GetDestinationChainID() *big.Int {
	return nil
}

// GetSender gets the sender of the CCTP tokens on the origin chain.
func (s SynapseCCTPCircleRequestFulfilled) GetSender() *string {
	return nil
}

// GetNonce gets the nonce of the CCTP message on the origin chain.
func (s SynapseCCTPCircleRequestFulfilled) GetNonce() *uint64 {
	return nil
}

// GetMintToken gets the address of the minted Circle token.
func (s SynapseCCTPCircleRequestFulfilled) GetMintToken() *string {
	str := s.MintToken.String()
	return &str
}

// GetAmount is the amount from the transfer.
func (s SynapseCCTPCircleRequestFulfilled) GetAmount() *big.Int {
	return s.Amount
}

// GetRequestVersion gets the version of the request format.
func (s SynapseCCTPCircleRequestFulfilled) GetRequestVersion() *uint32 {
	return nil
}

// GetFormattedRequest gets the formatted request for the action to take on the destination chain.
func (s SynapseCCTPCircleRequestFulfilled) GetFormattedRequest() *[]byte {
	return nil
}

// GetRecipient gets the end recipient of the tokens on this chain.
func (s SynapseCCTPCircleRequestFulfilled) GetRecipient() *string {
	str := s.Recipient.String()
	return &str
}

// GetFee gets the fee paid for fulfilling the request, in minted tokens.
func (s SynapseCCTPCircleRequestFulfilled) GetFee() *big.Int {
	return s.Fee
}

// GetToken gets the address of the token that the recipient received.
func (s SynapseCCTPCircleRequestFulfilled) GetToken() string {
	return s.Token.String()
}

// GetEventIndex gets the tx index of the event in the block it was executed in.
func (s SynapseCCTPCircleRequestFulfilled) GetEventIndex() uint64 {
	return uint64(s.Raw.TxIndex)
}

var _ cctp.EventLog = &SynapseCCTPCircleRequestFulfilled{}
