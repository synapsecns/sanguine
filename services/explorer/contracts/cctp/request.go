package cctp

import (
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

// GetOriginChainID gets the origin chain ID for the event.
func (s SynapseCCTPCircleRequestSent) GetOriginChainID() *big.Int {
	return nil
}

// GetDestinationChainID gets the destination chain ID for the event.
func (s SynapseCCTPCircleRequestSent) GetDestinationChainID() *big.Int {
	return s.ChainId
}

// GetSender gets the sender of the CCTP tokens on the origin chain.
func (s SynapseCCTPCircleRequestSent) GetSender() *common.Address {
	return &s.Sender
}

// GetNonce gets the nonce of the CCTP message on the origin chain.
func (s SynapseCCTPCircleRequestSent) GetNonce() *uint64 {
	return &s.Nonce
}

// GetBurnToken gets the address of the Circle token that was burnt.
func (s SynapseCCTPCircleRequestSent) GetBurnToken() *common.Address {
	return &s.Token
}

// GetMintToken gets the address of the minted Circle token.
func (s SynapseCCTPCircleRequestSent) GetMintToken() *common.Address {
	return nil
}

// GetSentAmount gets the amount of Circle tokens burnt.
func (s SynapseCCTPCircleRequestSent) GetSentAmount() *big.Int {
	return s.Amount
}

// GetReceivedAmount gets the received amount by the recipient.
func (s SynapseCCTPCircleRequestSent) GetReceivedAmount() *big.Int {
	return nil
}

// GetRequestVersion gets the version of the request format.
func (s SynapseCCTPCircleRequestSent) GetRequestVersion() *uint32 {
	return &s.RequestVersion
}

// GetFormattedRequest gets the formatted request for the action to take on the destination chain.
func (s SynapseCCTPCircleRequestSent) GetFormattedRequest() []byte {
	return s.FormattedRequest
}

// GetRequestID gets the unique identifier of the request.
func (s SynapseCCTPCircleRequestSent) GetRequestID() [32]byte {
	return s.RequestID
}

// GetRecipient gets the end recipient of the tokens on this chain.
func (s SynapseCCTPCircleRequestSent) GetRecipient() *common.Address {
	return nil
}

// GetFee gets the fee paid for fulfilling the request, in minted tokens.
func (s SynapseCCTPCircleRequestSent) GetFee() *big.Int {
	return nil
}

// GetToken gets the address of the token that the recipient received.
func (s SynapseCCTPCircleRequestSent) GetToken() *common.Address {
	return nil
}

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

// GetOriginChainID gets the origin chain ID for the event.
func (s SynapseCCTPCircleRequestFulfilled) GetOriginChainID() *big.Int {
	return big.NewInt(int64(s.OriginDomain))
}

// GetDestinationChainID gets the destination chain ID for the event.
func (s SynapseCCTPCircleRequestFulfilled) GetDestinationChainID() *big.Int {
	return nil
}

// GetSender gets the sender of the CCTP tokens on the origin chain.
func (s SynapseCCTPCircleRequestFulfilled) GetSender() *common.Address {
	return nil
}

// GetNonce gets the nonce of the CCTP message on the origin chain.
func (s SynapseCCTPCircleRequestFulfilled) GetNonce() *uint64 {
	return nil
}

// GetBurnToken gets the address of the Circle token that was burnt.
func (s SynapseCCTPCircleRequestFulfilled) GetBurnToken() *common.Address {
	return nil
}

// GetMintToken gets the address of the minted Circle token.
func (s SynapseCCTPCircleRequestFulfilled) GetMintToken() *common.Address {
	return &s.MintToken
}

// GetReceivedAmount gets the received amount by the recipient.
func (s SynapseCCTPCircleRequestFulfilled) GetReceivedAmount() *big.Int {
	return s.Amount
}

// GetRequestVersion gets the version of the request format.
func (s SynapseCCTPCircleRequestFulfilled) GetRequestVersion() *uint32 {
	return nil
}

// GetFormattedRequest gets the formatted request for the action to take on the destination chain.
func (s SynapseCCTPCircleRequestFulfilled) GetFormattedRequest() []byte {
	return []byte{}
}

// GetRequestID gets the unique identifier of the request.
func (s SynapseCCTPCircleRequestFulfilled) GetRequestID() [32]byte {
	return [32]byte{}
}

// GetRecipient gets the end recipient of the tokens on this chain.
func (s SynapseCCTPCircleRequestFulfilled) GetRecipient() *common.Address {
	return &s.Recipient
}

// GetFee gets the fee paid for fulfilling the request, in minted tokens.
func (s SynapseCCTPCircleRequestFulfilled) GetFee() *big.Int {
	return s.Fee
}

// GetToken gets the address of the token that the recipient received.
func (s SynapseCCTPCircleRequestFulfilled) GetToken() *common.Address {
	return &s.Token
}
