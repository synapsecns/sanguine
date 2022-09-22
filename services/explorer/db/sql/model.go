package sql

import (
	"math/big"
)

// SwapEvent stores data for emitted events from the Swap contract.
type SwapEvent struct {
	// ContractAddress is the address of the contract that generated the event
	ContractAddress string `gorm:"column:contract_address;primaryKey"`
	// ChainID is the chain id of the contract that generated the event
	ChainID uint32 `gorm:"column:chain_id;primaryKey;auto_increment:false"`
	// BlockNumber is the block number of the event
	BlockNumber uint64 `gorm:"column:block_number;primaryKey;auto_increment:false"`
	// TxHash is the transaction hash of the event
	TxHash string `gorm:"column:tx_hash;primaryKey"`
	// EventType is the type of the event
	EventType uint8 `gorm:"column:event_type;primaryKey;auto_increment:false"`

	// TokenIndex is the index of the token in the pool
	TokenIndex *uint8 `gorm:"column:token_index"`
	// Amount is the amount of tokens
	Amount *big.Int `gorm:"column:amount;type:UInt256"`
	// AmountFee is the amount of fees
	AmountFee *big.Int `gorm:"column:amount_fee;type:UInt256"`
	// ProtocolFee is the protocol fee
	ProtocolFee *big.Int `gorm:"column:protocol_fee;type:UInt256"`
	// Buyer is the address of the buyer
	Buyer *string `gorm:"column:buyer"`
	// TokensSold is the amount of tokens sold
	TokensSold *big.Int `gorm:"column:tokens_sold;type:UInt256"`
	// SoldID is the id of the token sold
	SoldID *big.Int `gorm:"column:sold_id;type:UInt256"`
	// TokensBought is the amount of tokens bought
	TokensBought *big.Int `gorm:"column:tokens_bought;type:UInt256"`
	// BoughtID is the id of the token bought
	BoughtID *big.Int `gorm:"column:bought_id;type:UInt256"`
	// Provider is the address of the provider
	Provider *string `gorm:"column:provider"`
	// TokenAmounts is the amounts of each token to transact
	TokenAmounts []*big.Int `gorm:"column:token_amounts;type:Array(UInt256)"`
	// Fees is the fees for each token
	Fees []*big.Int `gorm:"column:fees;type:Array(UInt256)"`
	// Invariant is the invariant of the pool
	Invariant *big.Int `gorm:"column:invariant;type:UInt256"`
	// LPTokenAmount is the amount of LP tokens
	LPTokenAmount *big.Int `gorm:"column:lp_token_amount;type:UInt256"`
	// LPTokenSupply is the supply of the LP token
	LPTokenSupply *big.Int `gorm:"column:lp_token_supply;type:UInt256"`
	// NewAdminFee is the new admin fee
	NewAdminFee *big.Int `gorm:"column:new_admin_fee;type:UInt256"`
	// NewSwapFee is the new swap fee
	NewSwapFee *big.Int `gorm:"column:new_swap_fee;type:UInt256"`
	// OldA is the old A value
	OldA *big.Int `gorm:"column:old_a;type:UInt256"`
	// NewA is the new A value
	NewA *big.Int `gorm:"column:new_a;type:UInt256"`
	// InitialTime is the initial time
	InitialTime *big.Int `gorm:"column:initial_time;type:UInt256"`
	// FutureTime is the future time
	FutureTime *big.Int `gorm:"column:future_time;type:UInt256"`
	// CurrentA is the current A value
	CurrentA *big.Int `gorm:"column:current_a;type:UInt256"`
	// Time is the time
	Time *big.Int `gorm:"column:time;type:UInt256"`
	// TokenID is the token's ID
	TokenID *string `gorm:"column:token_id"`
}

// BridgeEvent stores data for emitted events from the Bridge contract.
type BridgeEvent struct {
	// ContractAddress is the address of the contract that generated the event
	ContractAddress string `gorm:"column:contract_address;primaryKey"`
	// ChainID is the chain id of the contract that generated the event
	ChainID uint32 `gorm:"column:chain_id;primaryKey;auto_increment:false"`
	// EventType is the type of the event
	EventType uint8 `gorm:"column:event_type;primaryKey;auto_increment:false"`
	// BlockNumber is the block number of the event
	BlockNumber uint64 `gorm:"column:block_number;primaryKey;auto_increment:false"`
	// TxHash is the transaction hash of the event
	TxHash string `gorm:"column:tx_hash;primaryKey"`
	// Token is the address of the token
	Token string `gorm:"column:token"`
	// Amount is the amount of tokens
	Amount big.Int `gorm:"column:amount;type:UInt256"`

	// Recipient is the address to send the tokens to
	Recipient *string `gorm:"column:recipient"`
	// DestinationChainID is the chain id of the chain to send the tokens to
	DestinationChainID *uint32 `gorm:"column:destination_chain_id"`
	// Fee is the fee
	Fee *big.Int `gorm:"column:fee;type:UInt256"`
	// Kappa is theFee keccak256 hash of the transaction
	Kappa *string `gorm:"column:kappa"`
	// TokenIndexFrom is the index of the from token in the pool
	TokenIndexFrom *uint8 `gorm:"column:token_index_from"`
	// TokenIndexTo is the index of the to token in the pool
	TokenIndexTo *uint8 `gorm:"column:token_index_to"`
	// MinDy is the minimum amount of tokens to receive
	MinDy *big.Int `gorm:"column:min_dy;type:UInt256"`
	// Deadline is the deadline of the transaction
	Deadline *big.Int `gorm:"column:deadline;type:UInt256"`
	// SwapSuccess is whether the swap was successful
	SwapSuccess *uint8 `gorm:"column:swap_success"`
	// SwapTokenIndex is the index of the token in the pool
	SwapTokenIndex *uint8 `gorm:"column:swap_token_index"`
	// SwapMinAmount is the minimum amount of tokens to receive
	SwapMinAmount *big.Int `gorm:"column:swap_min_amount;type:UInt256"`
	// SwapDeadline is the deadline of the swap transaction
	SwapDeadline *big.Int `gorm:"column:swap_deadline;type:UInt256"`
	// TokenID is the token's ID
	TokenID *string `gorm:"column:token_id"`
}
