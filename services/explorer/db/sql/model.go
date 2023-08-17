package sql

import (
	"database/sql"
	"math/big"

	"github.com/synapsecns/sanguine/core/dbcommon"
)

// init define common field names. See package docs for an explanation of why we have to do this.
// Note: Some models share names. In cases where they do, we run the check against all names.
// This is cheap because it's only done at startup.
func init() {
	namer := dbcommon.NewNamer([]interface{}{&SwapEvent{}, &BridgeEvent{}})
	TxHashFieldName = namer.GetConsistentName("TxHash")
	ChainIDFieldName = namer.GetConsistentName("ChainID")
	BlockNumberFieldName = namer.GetConsistentName("BlockNumber")
	ContractAddressFieldName = namer.GetConsistentName("ContractAddress")
	InsertTimeFieldName = namer.GetConsistentName("InsertTime")
	EventTypeFieldName = namer.GetConsistentName("EventType")
	AmountFieldName = namer.GetConsistentName("Amount")
	EventIndexFieldName = namer.GetConsistentName("EventIndex")
	DestinationChainIDFieldName = namer.GetConsistentName("DestinationChainID")
	TokenFieldName = namer.GetConsistentName("Token")
	RecipientFieldName = namer.GetConsistentName("Recipient")
	KappaFieldName = namer.GetConsistentName("Kappa")
	DestinationKappaFieldName = namer.GetConsistentName("DestinationKappa")
	SenderFieldName = namer.GetConsistentName("Sender")
	TimeStampFieldName = namer.GetConsistentName("TimeStamp")
	AmountUSDFieldName = namer.GetConsistentName("AmountUSD")
	TokenDecimalFieldName = namer.GetConsistentName("TokenDecimal")
	AdminFeeUSDFieldName = namer.GetConsistentName("AdminFeeUSD")
	FeeUSDFieldName = namer.GetConsistentName("FeeUSD")
}

var (
	// TxHashFieldName is the field name of the tx hash.
	TxHashFieldName string
	// ChainIDFieldName gets the chain id field name.
	ChainIDFieldName string
	// BlockNumberFieldName is the name of the block number field.
	BlockNumberFieldName string
	// ContractAddressFieldName is the address of the contract.
	ContractAddressFieldName string
	// InsertTimeFieldName is the insert time field name.
	InsertTimeFieldName string
	// EventTypeFieldName is the event type field name.
	EventTypeFieldName string
	// AmountFieldName is the amount field name.
	AmountFieldName string
	// EventIndexFieldName is the event index field name.
	EventIndexFieldName string
	// DestinationChainIDFieldName is the destination chain id field name.
	DestinationChainIDFieldName string
	// TokenFieldName is the token field name.
	TokenFieldName string
	// RecipientFieldName is the recipient field name.
	RecipientFieldName string
	// KappaFieldName is the kappa field name.
	KappaFieldName string
	// DestinationKappaFieldName is the destination kappa field name.
	DestinationKappaFieldName string
	// SenderFieldName is the sender field name.
	SenderFieldName string
	// TimeStampFieldName is the timestamp field name.
	TimeStampFieldName string
	// AmountUSDFieldName is the amount in USD field name.
	AmountUSDFieldName string
	// TokenDecimalFieldName is the token decimal field name.
	TokenDecimalFieldName string
	// AdminFeeUSDFieldName is the usd admin fee field name.
	AdminFeeUSDFieldName string
	// FeeUSDFieldName is the fee in usd field name.
	FeeUSDFieldName string
)

// PageSize is the amount of entries per page of events.
var PageSize = 100

// CCTPEvent stores a cctp event.
type CCTPEvent struct {
	// InsertTime is the time the event was inserted into the database.
	InsertTime uint64 `gorm:"column:insert_time"`
	// ChainID is the chain ID of the chain in which the indexed event occurred.
	ChainID uint32 `gorm:"column:chain_id"`
	// TxHash is the transaction hash of the event.
	TxHash string `gorm:"column:tx_hash"`
	// ContractAddress is the address of the contract that generated the event.
	ContractAddress string `gorm:"column:contract_address"`
	// BlockNumber is the timestamp.
	BlockNumber uint64 `gorm:"column:block_number"`
	// EventType is the type of the event.
	EventType uint8 `gorm:"column:event_type"`
	// RequestID is the request ID of the CCTP transfer.
	RequestID string `gorm:"column:request_id"`

	// Token is either the address of the received token on destination or the address of the token burnt on origin.
	Token string `gorm:"column:token"`
	// Amount is the amount of the CCTP transfer.
	Amount *big.Int `gorm:"column:amount;type:UInt256"`
	// EventIndex is the index of the log.
	EventIndex uint64 `gorm:"column:event_index"`
	// AmountUSD is the amount of the CCTP transfer in USD.
	AmountUSD float64 `gorm:"column:amount_usd;type:Float64"`
	// OriginChainID is the chain ID of the CCTP transfer.
	OriginChainID *big.Int `gorm:"column:origin_chain_id;type:UInt256"`
	// DestinationChainID is the chain ID of the CCTP transfer.
	DestinationChainID *big.Int `gorm:"column:destination_chain_id;type:UInt256"`
	// Sender is the address of the sender.
	Sender sql.NullString `gorm:"column:sender"`
	// Nonce is the nonce of the CCTP transfer.
	Nonce sql.NullInt64 `gorm:"column:nonce"`
	// MintToken is the address of the minted token on destination
	MintToken sql.NullString `gorm:"column:mint_token"`
	// RequestVersion is the request version of the CCTP transfer.
	RequestVersion sql.NullInt32 `gorm:"column:request_version"`
	// FormattedRequest is the formatted request of the CCTP transfer.
	FormattedRequest sql.NullString `gorm:"column:formatted_request"`
	// Recipient is the recipient of the CCTP transfer.
	Recipient sql.NullString `gorm:"column:recipient"`
	// Fee is the fee of the CCTP transfer.
	Fee *big.Int `gorm:"column:fee;type:UInt256"`
	// FeeUSD is the fee of the CCTP transfer in USD terms.
	FeeUSD *float64 `gorm:"column:fee_usd;type:Float64"`
	// TokenDecimal is the token's decimal.
	TokenDecimal *uint8 `gorm:"column:token_decimal"`
	// TokenSymbol is the token's symbol from coin gecko.
	TokenSymbol string `gorm:"column:token_symbol"`
	// TimeStamp is the timestamp in which the record was inserted.
	TimeStamp *uint64 `gorm:"column:timestamp"`
}

// BridgeEvent stores data for emitted events from the Bridge contract.
type BridgeEvent struct {
	// InsertTime is the time the event was inserted into the database.
	InsertTime uint64 `gorm:"column:insert_time"`
	// ContractAddress is the address of the contract that generated the event.
	ContractAddress string `gorm:"column:contract_address"`
	// ChainID is the chain id of the contract that generated the event.
	ChainID uint32 `gorm:"column:chain_id"`
	// EventType is the type of the event.
	EventType uint8 `gorm:"column:event_type"`
	// BlockNumber is the block number of the event.
	BlockNumber uint64 `gorm:"column:block_number"`
	// TxHash is the transaction hash of the event.
	TxHash string `gorm:"column:tx_hash"`
	// Token is the address of the token.
	Token string `gorm:"column:token"`
	// Amount is the amount of tokens.
	Amount *big.Int `gorm:"column:amount;type:UInt256"`
	// EventIndex is the index of the log.
	EventIndex uint64 `gorm:"column:event_index"`
	// DestinationKappa is the destination kappa.
	DestinationKappa string `gorm:"column:destination_kappa"`
	// Sender is the address of the sender.
	Sender string `gorm:"column:sender"`

	// Recipient is the address to send the tokens to.
	Recipient sql.NullString `gorm:"column:recipient"`
	// RecipientBytes is the recipient address in bytes.
	RecipientBytes sql.NullString `gorm:"column:recipient_bytes"`
	// DestinationChainID is the chain id of the chain to send the tokens to.
	DestinationChainID *big.Int `gorm:"column:destination_chain_id;type:UInt256"`
	// Fee is the fee.
	Fee *big.Int `gorm:"column:fee;type:UInt256"`
	// Kappa is the keccak256 hash of the transaction.
	Kappa sql.NullString `gorm:"column:kappa"`
	// TokenIndexFrom is the index of the from token in the pool.
	TokenIndexFrom *big.Int `gorm:"column:token_index_from;type:UInt256"`
	// TokenIndexTo is the index of the to token in the pool.
	TokenIndexTo *big.Int `gorm:"column:token_index_to;type:UInt256"`
	// MinDy is the minimum amount of tokens to receive.
	MinDy *big.Int `gorm:"column:min_dy;type:UInt256"`
	// Deadline is the deadline of the transaction.
	Deadline *big.Int `gorm:"column:deadline;type:UInt256"`
	// SwapSuccess is whether the swap was successful.
	SwapSuccess *big.Int `gorm:"column:swap_success;type:UInt256"`
	// SwapTokenIndex is the index of the token in the pool.
	SwapTokenIndex *big.Int `gorm:"column:swap_token_index;type:UInt256"`
	// SwapMinAmount is the minimum amount of tokens to receive.
	SwapMinAmount *big.Int `gorm:"column:swap_min_amount;type:UInt256"`
	// SwapDeadline is the deadline of the swap transaction.
	SwapDeadline *big.Int `gorm:"column:swap_deadline;type:UInt256"`
	// AmountUSD is the amount in USD.
	AmountUSD *float64 `gorm:"column:amount_usd;type:Float64"`
	// FeeUSD is the fee amount in USD.
	FeeUSD *float64 `gorm:"column:fee_usd;type:Float64"`
	// TokenDecimal is the token's decimal.
	TokenDecimal *uint8 `gorm:"column:token_decimal"`
	// TokenSymbol is the token's symbol from coin gecko.
	TokenSymbol sql.NullString `gorm:"column:token_symbol"`
	// TimeStamp is the timestamp of the block in which the event occurred.
	TimeStamp *uint64 `gorm:"column:timestamp"`
}

// SwapEvent stores data for emitted events from the Swap contract.
type SwapEvent struct {
	// InsertTime is the time the event was inserted into the database.
	InsertTime uint64 `gorm:"column:insert_time"`
	// ContractAddress is the address of the contract that generated the event.
	ContractAddress string `gorm:"column:contract_address"`
	// ChainID is the chain id of the contract that generated the event.
	ChainID uint32 `gorm:"column:chain_id"`
	// BlockNumber is the block number of the event.
	BlockNumber uint64 `gorm:"column:block_number"`
	// TxHash is the transaction hash of the event.
	TxHash string `gorm:"column:tx_hash"`
	// EventType is the type of the event.
	EventType uint8 `gorm:"column:event_type"`
	// EventIndex is the index of the log.
	EventIndex uint64 `gorm:"column:event_index"`
	// Sender is the address of the sender.
	Sender string `gorm:"column:sender"`

	// Amount is the amount of tokens.
	Amount map[uint8]string `gorm:"column:amount;type:Map(UInt8, String)"`
	// Fee is the amount of fees.
	Fee map[uint8]string `gorm:"column:fee;type:Map(UInt8, String)"`
	// AdminFee is the amount of admin fees.
	AdminFee map[uint8]string `gorm:"column:admin_fee;type:Map(UInt8, String)"`
	// AmountUSD is the amount in USD.
	AmountUSD map[uint8]float64 `gorm:"column:amount_usd;type:Map(UInt8, Float64)"`
	// FeeAmountUSD is the fee amount in USD.
	FeeUSD map[uint8]float64 `gorm:"column:fee_usd;type:Map(UInt8, Float64)"`
	// AdminFeeAmountUSD is the admin fee amount in USD.
	AdminFeeUSD map[uint8]float64 `gorm:"column:admin_fee_usd;type:Map(UInt8, Float64)"`
	// ProtocolFee is the protocol fee.
	ProtocolFee *big.Int `gorm:"column:protocol_fee;type:UInt256"`
	// Buyer is the address of the buyer.
	Buyer sql.NullString `gorm:"column:buyer"`
	// TokensSold is the amount of tokens sold.
	TokensSold *big.Int `gorm:"column:tokens_sold;type:UInt256"`
	// SoldID is the id of the token sold.
	SoldID *big.Int `gorm:"column:sold_id;type:UInt256"`
	// TokensBought is the amount of tokens bought.
	TokensBought *big.Int `gorm:"column:tokens_bought;type:UInt256"`
	// BoughtID is the id of the token bought.
	BoughtID *big.Int `gorm:"column:bought_id;type:UInt256"`
	// Provider is the address of the provider.
	Provider sql.NullString `gorm:"column:provider"`
	// Invariant is the invariant of the pool.
	Invariant *big.Int `gorm:"column:invariant;type:UInt256"`
	// LPTokenAmount is the amount of LP tokens.
	LPTokenAmount *big.Int `gorm:"column:lp_token_amount;type:UInt256"`
	// LPTokenSupply is the supply of the LP token.
	LPTokenSupply *big.Int `gorm:"column:lp_token_supply;type:UInt256"`
	// NewAdminFee is the new admin fee.
	NewAdminFee *big.Int `gorm:"column:new_admin_fee;type:UInt256"`
	// NewSwapFee is the new swap fee.
	NewSwapFee *big.Int `gorm:"column:new_swap_fee;type:UInt256"`
	// OldA is the old A value.
	OldA *big.Int `gorm:"column:old_a;type:UInt256"`
	// NewA is the new A value.
	NewA *big.Int `gorm:"column:new_a;type:UInt256"`
	// InitialTime is the initial time.
	InitialTime *big.Int `gorm:"column:initial_time;type:UInt256"`
	// FutureTime is the future time.
	FutureTime *big.Int `gorm:"column:future_time;type:UInt256"`
	// CurrentA is the current A value.
	CurrentA *big.Int `gorm:"column:current_a;type:UInt256"`
	// Time is the time.
	Time *big.Int `gorm:"column:time;type:UInt256"`
	// Receiver is the address of the receiver.
	Receiver sql.NullString `gorm:"column:receiver"`
	// TokenPrices are the prices of each token at the given time.
	TokenPrice map[uint8]float64 `gorm:"column:token_price;type:Map(UInt8, Float64)"`
	// TokenDecimal is the token's decimal.
	TokenDecimal map[uint8]uint8 `gorm:"column:token_decimal;type:Map(UInt8, UInt8)"`
	// TokenSymbol is the token's symbol from coingecko.
	TokenSymbol map[uint8]string `gorm:"column:token_symbol;type:Map(UInt8, String)"`
	// TokenSymbol is the token's symbol from coingecko.
	TokenCoinGeckoID map[uint8]string `gorm:"column:token_coingecko_id;type:Map(UInt8, String)"`
	// TimeStamp is the timestamp of the block in which the event occurred.
	TimeStamp *uint64 `gorm:"column:timestamp"`
}

// LastBlock stores the last block number that the explorer has backfilled to on each chain.
type LastBlock struct {
	// ChainID is the chain id of the chain.
	ChainID uint32 `gorm:"column:chain_id"`
	// BlockNumber is the last block number that the explorer has backfilled to.
	BlockNumber uint64 `gorm:"column:block_number"`
	// ContractAddress is the address of the contract that generated the event.
	ContractAddress string `gorm:"column:contract_address"`
}

// TokenIndex stores the data for each token index on each chain.
type TokenIndex struct {
	// ChainID is the chain id of the chain.
	ChainID uint32 `gorm:"column:chain_id"`
	// TokenIndex is the token index in the pool.
	TokenIndex uint8 `gorm:"column:token_index"`
	// TokenAddress is the address of the token.
	TokenAddress string `gorm:"column:token_address"`
	// ContractAddress is the address of the contract that generated the event.
	ContractAddress string `gorm:"column:contract_address"`
}

// SwapFees stores the admin and swap fees.
type SwapFees struct {
	// ChainID is the chain id of the chain.
	ChainID uint32 `gorm:"column:chain_id"`
	// ContractAddress is the address of the contract that generated the event.
	ContractAddress string `gorm:"column:contract_address"`
	// BlockNumber is the timestamp.
	BlockNumber uint64 `gorm:"column:block_number"`
	// FeeType is the type of fee.
	FeeType string `gorm:"column:fee_type"`
	// Fee the fee.
	Fee uint64 `gorm:"column:fee"`
}

// MessageBusEvent stores data for emitted events from the message bus contract.
type MessageBusEvent struct {
	// InsertTime is the time the event was inserted into the database
	InsertTime uint64 `gorm:"column:insert_time"`
	// ContractAddress is the address of the contract that generated the event
	ContractAddress string `gorm:"column:contract_address"`
	// ChainID is the chain id of the contract that generated the event
	ChainID uint32 `gorm:"column:chain_id"`
	// BlockNumber is the block number of the event
	BlockNumber uint64 `gorm:"column:block_number"`
	// TxHash is the transaction hash of the event
	TxHash string `gorm:"column:tx_hash"`
	// EventType is the type of the event
	EventType uint8 `gorm:"column:event_type"`
	// EventIndex is the index of the log
	EventIndex uint64 `gorm:"column:event_index"`
	// Sender is the address of the sender
	Sender string `gorm:"column:sender"`

	// MessageId is the message id of the event.
	MessageID sql.NullString `gorm:"column:message_id"`
	// SourceChainID is the chain id of the message's source chain.
	SourceChainID *big.Int `gorm:"column:source_chain_id;type:UInt256"`
	// Status is the status of the event.
	Status sql.NullString `gorm:"column:status"`
	// SourceAddress is the address that the message will be passed from.
	SourceAddress sql.NullString `gorm:"column:source_address"`
	// DestinationAddress is the address that the message will be passed to.
	DestinationAddress sql.NullString `gorm:"column:destination_address"`
	// DestinationChainID is the chain id of the message's destination chain.
	DestinationChainID *big.Int `gorm:"column:destination_chain_id;type:UInt256"`
	// Nonce is the nonce of the message. It is equivalent to the nonce on the origin chain.
	Nonce *big.Int `gorm:"column:nonce;type:UInt256"`
	// Message is the message.
	Message sql.NullString `gorm:"column:message"`
	// Receiver is the receiver of the event.
	Receiver sql.NullString `gorm:"column:receiver"`
	// Options are the options chosen for the message.
	Options sql.NullString `gorm:"column:options"`
	// Fee is the fee of the message.
	Fee *big.Int `gorm:"column:fee;type:UInt256"`
	// FeeUSD is the fee of the message.
	FeeUSD *float64 `gorm:"column:fee_usd;type:Float64"`
	// RevertedReason is the reason a call was reverted.
	RevertedReason sql.NullString `gorm:"column:reverted_reason"`
	// TimeStamp is the timestamp in which the record was inserted.
	TimeStamp *uint64 `gorm:"column:timestamp"`
}

// TODO clean up the comments in this

// HybridBridgeEvent is the datatype for a returned bridge event.
type HybridBridgeEvent struct {
	// FInsertTime is the time the event was inserted into the database.
	FInsertTime uint64 `gorm:"column:finsert_time"`
	// FContractAddress is the address of the contract that generated the event.
	FContractAddress string `gorm:"column:fcontract_address"`
	// FChainID is the chain id of the contract that generated the event.
	FChainID uint32 `gorm:"column:fchain_id"`
	// FEventType is the type of the event.
	FEventType uint8 `gorm:"column:fevent_type"`
	// FBlockNumber is the block number of the event.
	FBlockNumber uint64 `gorm:"column:fblock_number"`
	// FTxHash is the transaction hash of the event.
	FTxHash string `gorm:"column:ftx_hash"`
	// FToken is the address of the token.
	FToken string `gorm:"column:ftoken"`
	// FAmount is the amount of tokens.
	FAmount *big.Int `gorm:"column:famount;type:UInt256"`
	// FEventIndex is the index of the log.
	FEventIndex uint64 `gorm:"column:fevent_index"`
	// FDestinationKappa is the destination kappa.
	FDestinationKappa string `gorm:"column:fdestination_kappa"`
	// FSender is the address of the sender.
	FSender string `gorm:"column:fsender"`

	// FRecipient is the address to send the tokens to.
	FRecipient sql.NullString `gorm:"column:frecipient"`
	// FRecipientBytes is the recipient address in bytes.
	FRecipientBytes sql.NullString `gorm:"column:frecipient_bytes"`
	// FDestinationChainID is the chain id of the chain to send the tokens to.
	FDestinationChainID *big.Int `gorm:"column:fdestination_chain_id;type:UInt256"`
	// FFee is the fee.
	FFee *big.Int `gorm:"column:ffee;type:UInt256"`
	// FKappa is theFee keccak256 hash of the transaction.
	FKappa sql.NullString `gorm:"column:fkappa"`
	// FTokenIndexFrom is the index of the from token in the pool.
	FTokenIndexFrom *big.Int `gorm:"column:ftoken_index_from;type:UInt256"`
	// FTokenIndexTo is the index of the to token in the pool.
	FTokenIndexTo *big.Int `gorm:"column:ftoken_index_to;type:UInt256"`
	// FMinDy is the minimum amount of tokens to receive.
	FMinDy *big.Int `gorm:"column:fmin_dy;type:UInt256"`
	// FDeadline is the deadline of the transaction.
	FDeadline *big.Int `gorm:"column:fdeadline;type:UInt256"`
	// FSwapSuccess is whether the swap was successful.
	FSwapSuccess *big.Int `gorm:"column:fswap_success;type:UInt256"`
	// FSwapTokenIndex is the index of the token in the pool.
	FSwapTokenIndex *big.Int `gorm:"column:fswap_token_index;type:UInt256"`
	// FSwapMinAmount is the minimum amount of tokens to receive.
	FSwapMinAmount *big.Int `gorm:"column:fswap_min_amount;type:UInt256"`
	// FSwapDeadline is the deadline of the swap transaction.
	FSwapDeadline *big.Int `gorm:"column:fswap_deadline;type:UInt256"`
	// FTokenID is the token's ID.
	FTokenID sql.NullString `gorm:"column:ftoken_id"`
	// FAmountUSD is the amount in USD.
	FAmountUSD *float64 `gorm:"column:famount_usd;type:Float64"`
	// FFeeAmountUSD is the fee amount in USD.
	FFeeAmountUSD *float64 `gorm:"column:ffee_amount_usd;type:Float64"`
	// FTokenDecimal is the token's decimal.
	FTokenDecimal *uint8 `gorm:"column:ftoken_decimal"`
	// FTokenSymbol is the token's symbol from coin gecko.
	FTokenSymbol sql.NullString `gorm:"column:ftoken_symbol"`
	// FTimeStamp is the timestamp of the block in which the event occurred.
	FTimeStamp *uint64 `gorm:"column:ftimestamp"`
	// TInsertTime is the time the event was inserted into the database.
	TInsertTime uint64 `gorm:"column:finsert_time"`
	// TContractAddress is the address of the contract that generated the event.
	TContractAddress string `gorm:"column:tcontract_address"`
	// TChainID is the chain id of the contract that generated the event.
	TChainID uint32 `gorm:"column:tchain_id"`
	// TEventType is the type of the event.
	TEventType uint8 `gorm:"column:tevent_type"`
	// TBlockNumber is the block number of the event.
	TBlockNumber uint64 `gorm:"column:tblock_number"`
	// TTxHash is the transaction hash of the event.
	TTxHash string `gorm:"column:ttx_hash"`
	// TToken is the address of the token.
	TToken string `gorm:"column:ttoken"`
	// TAmount is the amount of tokens.
	TAmount *big.Int `gorm:"column:tamount;type:UInt256"`
	// TEventIndex is the index of the log.
	TEventIndex uint64 `gorm:"column:tevent_index"`
	// TDestinationKappa is the destination kappa.
	TDestinationKappa string `gorm:"column:tdestination_kappa"`
	// TSender is the address of the sender.
	TSender string `gorm:"column:tsender"`

	// TRecipient is the address to send the tokens to.
	TRecipient sql.NullString `gorm:"column:trecipient"`
	// TRecipientBytes is the recipient address in bytes.
	TRecipientBytes sql.NullString `gorm:"column:trecipient_bytes"`
	// TDestinationChainID is the chain id of the chain to send the tokens to.
	TDestinationChainID *big.Int `gorm:"column:tdestination_chain_id;type:UInt256"`
	// TFee is the fee.
	TFee *big.Int `gorm:"column:tfee;type:UInt256"`
	// TKappa is theFee keccak256 hash of the transaction.
	TKappa sql.NullString `gorm:"column:tkappa"`
	// TTokenIndexFrom is the index of the from token in the pool.
	TTokenIndexFrom *big.Int `gorm:"column:ttoken_index_from;type:UInt256"`
	// TTokenIndexTo is the index of the to token in the pool.
	TTokenIndexTo *big.Int `gorm:"column:ttoken_index_to;type:UInt256"`
	// TMinDy is the minimum amount of tokens to receive.
	TMinDy *big.Int `gorm:"column:tmin_dy;type:UInt256"`
	// TDeadline is the deadline of the transaction.
	TDeadline *big.Int `gorm:"column:tdeadline;type:UInt256"`
	// TSwapSuccess is whether the swap was successful.
	TSwapSuccess *big.Int `gorm:"column:tswap_success;type:UInt256"`
	// TSwapTokenIndex is the index of the token in the pool.
	TSwapTokenIndex *big.Int `gorm:"column:tswap_token_index;type:UInt256"`
	// TSwapMinAmount is the minimum amount of tokens to receive.
	TSwapMinAmount *big.Int `gorm:"column:tswap_min_amount;type:UInt256"`
	// TSwapDeadline is the deadline of the swap transaction.
	TSwapDeadline *big.Int `gorm:"column:tswap_deadline;type:UInt256"`
	// TTokenID is the token's ID.
	TTokenID sql.NullString `gorm:"column:ttoken_id"`
	// TAmountUSD is the amount in USD.
	TAmountUSD *float64 `gorm:"column:tamount_usd;type:Float64"`
	// TFeeAmountUSD is the fee amount in USD.
	TFeeAmountUSD *float64 `gorm:"column:tfee_amount_usd;type:Float64"`
	// TTokenDecimal is the token's decimal.
	TTokenDecimal *uint8 `gorm:"column:ttoken_decimal"`
	// TTokenSymbol is the token's symbol from coin gecko.
	TTokenSymbol sql.NullString `gorm:"column:ttoken_symbol"`
	// TTimeStamp is the timestamp of the block in which the event occurred.
	TTimeStamp *uint64 `gorm:"column:ttimestamp"`
}

// HybridMessageBusEvent stores data for emitted events from the message bus contract after joining origin and destination events.
type HybridMessageBusEvent struct {
	// InsertTime is the time the event was inserted into the database
	FInsertTime uint64 `gorm:"column:insert_time"`
	// ContractAddress is the address of the contract that generated the event
	FContractAddress string `gorm:"column:contract_address"`
	// ChainID is the chain id of the contract that generated the event
	FChainID uint32 `gorm:"column:chain_id"`
	// BlockNumber is the block number of the event
	FBlockNumber uint64 `gorm:"column:block_number"`
	// TxHash is the transaction hash of the event
	FTxHash string `gorm:"column:tx_hash"`
	// EventType is the type of the event
	FEventType uint8 `gorm:"column:event_type"`
	// EventIndex is the index of the log
	FEventIndex uint64 `gorm:"column:event_index"`
	// Sender is the address of the sender
	FSender string `gorm:"column:sender"`

	// MessageId is the message id of the event.
	FMessageID sql.NullString `gorm:"column:message_id"`
	// SourceChainID is the chain id of the message's source chain.
	FSourceChainID *big.Int `gorm:"column:source_chain_id;type:UInt256"`
	// Status is the status of the event.
	FStatus sql.NullString `gorm:"column:status"`
	// SourceAddress is the address that the message will be passed from.
	FSourceAddress sql.NullString `gorm:"column:source_address"`
	// DestinationAddress is the address that the message will be passed to.
	FDestinationAddress sql.NullString `gorm:"column:destination_address"`
	// DestinationChainID is the chain id of the message's destination chain.
	FDestinationChainID *big.Int `gorm:"column:destination_chain_id;type:UInt256"`
	// Nonce is the nonce of the message. It is equivalent to the nonce on the origin chain.
	FNonce *big.Int `gorm:"column:nonce;type:UInt256"`
	// Message is the message.
	FMessage sql.NullString `gorm:"column:message"`
	// Receiver is the receiver of the event.
	FReceiver sql.NullString `gorm:"column:receiver"`
	// Options are the options chosen for the message.
	FOptions sql.NullString `gorm:"column:options"`
	// Fee is the fee of the message.
	FFee *big.Int `gorm:"column:fee;type:UInt256"`
	// FeeUSD is the fee of the message.
	FFeeUSD *float64 `gorm:"column:fee_usd;type:Float64"`
	// RevertedReason is the reason a call was reverted.
	FRevertedReason sql.NullString `gorm:"column:reverted_reason"`
	// TimeStamp is the timestamp in which the record was inserted.
	FTimeStamp *uint64 `gorm:"column:timestamp"`

	// InsertTime is the time the event was inserted into the database
	TInsertTime uint64 `gorm:"column:t.insert_time"`
	// ContractAddress is the address of the contract that generated the event
	TContractAddress string `gorm:"column:t.contract_address"`
	// ChainID is the chain id of the contract that generated the event
	TChainID uint32 `gorm:"column:t.chain_id"`
	// BlockNumber is the block number of the event
	TBlockNumber uint64 `gorm:"column:t.block_number"`
	// TxHash is the transaction hash of the event
	TTxHash string `gorm:"column:t.tx_hash"`
	// EventType is the type of the event
	TEventType uint8 `gorm:"column:t.event_type"`
	// EventIndex is the index of the log
	TEventIndex uint64 `gorm:"column:t.event_index"`
	// Sender is the address of the sender
	TSender string `gorm:"column:t.sender"`

	// MessageId is the message id of the event.
	TMessageID sql.NullString `gorm:"column:t.message_id"`
	// SourceChainID is the chain id of the message's source chain.
	TSourceChainID *big.Int `gorm:"column:t.source_chain_id;type:UInt256"`
	// Status is the status of the event.
	TStatus sql.NullString `gorm:"column:t.status"`
	// SourceAddress is the address that the message will be passed from.
	TSourceAddress sql.NullString `gorm:"column:t.source_address"`
	// DestinationAddress is the address that the message will be passed to.
	TDestinationAddress sql.NullString `gorm:"column:t.destination_address"`
	// DestinationChainID is the chain id of the message's destination chain.
	TDestinationChainID *big.Int `gorm:"column:t.destination_chain_id;type:UInt256"`
	// Nonce is the nonce of the message. It is equivalent to the nonce on the origin chain.
	TNonce *big.Int `gorm:"column:t.nonce;type:UInt256"`
	// Message is the message.
	TMessage sql.NullString `gorm:"column:t.message"`
	// Receiver is the receiver of the event.
	TReceiver sql.NullString `gorm:"column:t.receiver"`
	// Options are the options chosen for the message.
	TOptions sql.NullString `gorm:"column:t.options"`
	// Fee is the fee of the message.
	TFee *big.Int `gorm:"column:t.fee;type:UInt256"`
	// FeeUSD is the fee of the message.
	TFeeUSD *float64 `gorm:"column:t.fee_usd;type:Float64"`
	// RevertedReason is the reason a call was reverted.
	TRevertedReason sql.NullString `gorm:"column:t.reverted_reason"`
	// TimeStamp is the timestamp in which the record was inserted.
	TTimeStamp *uint64 `gorm:"column:t.timestamp"`
}
