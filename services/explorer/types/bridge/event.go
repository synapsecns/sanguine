package bridge

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// EventLog is the interface for all bridge events.
type EventLog interface {
	// GetContractAddress returns the contract address of the log.
	GetContractAddress() common.Address
	// GetChainID returns the chain id of the log.
	GetChainID() uint32
	// GetBlockNumber returns the block number of the log.
	GetBlockNumber() uint64
	// GetEventType returns the event type of the log.
	GetEventType() EventType
}

// TokenDepositLog is the interface for the token deposit event.
type TokenDepositLog interface {
	EventLog
	// GetTo returns the recipient of the token deposit.
	GetTo() *common.Address
	// GetChainId returns the chain id of the token deposit.
	GetChainId() *big.Int
	// GetToken returns the token of the token deposit.
	// TODO
	// GetAmount returns the amount of the token deposit.
	GetAmount() *big.Int
}

// TokenRedeemLog is the interface for the token redeem event.
type TokenRedeemLog interface {
	EventLog
	// GetTo returns the recipient of the token redeem.
	GetTo() *common.Address
	// GetChainId returns the chain id of the token redeem.
	GetChainId() *big.Int
	// GetToken returns the token of the token redeem.
	// TODO
	// GetAmount returns the amount of the token redeem.
	GetAmount() *big.Int
}

// TokenWithdrawLog is the interface for the token withdraw event.
type TokenWithdrawLog interface {
	EventLog
	// GetTo returns the recipient of the token withdraw.
	GetTo() *common.Address
	// GetToken returns the token of the token withdraw.
	// TODO
	// GetAmount returns the amount of the token withdraw.
	GetAmount() *big.Int
	// GetFee returns the fee of the token withdraw.
	GetFee() *big.Int
	// GetKappa returns the kappa of the token withdraw.
	GetKappa() [32]byte
}

// TokenMintLog is the interface for the token mint event.
type TokenMintLog interface {
	EventLog
	// GetTo returns the recipient of the token mint.
	GetTo() *common.Address
	// GetToken returns the token of the token mint.
	// TODO
	// GetAmount returns the amount of the token mint.
	GetAmount() *big.Int
	// GetFee returns the fee of the token mint.
	GetFee() *big.Int
	// GetKappa returns the kappa of the token mint.
	GetKappa() [32]byte
}

// TokenDepositAndSwapLog is the interface for the token deposit and swap event.
type TokenDepositAndSwapLog interface {
	EventLog
	// GetTo returns the recipient of the token deposit and swap.
	GetTo() *common.Address
	// GetChainId returns the chain id of the token deposit and swap.
	GetChainId() *big.Int
	// GetToken returns the token of the token deposit and swap.
	// TODO
	// GetAmount returns the amount of the token deposit and swap.
	GetAmount() *big.Int
	// GetTokenIndexFrom returns the token index of the token deposit and swap.
	GetTokenIndexFrom() uint8
	// GetTokenIndexTo returns the token index of the token deposit and swap.
	GetTokenIndexTo() uint8
	// GetMinDy returns the min dy of the token deposit and swap.
	GetMinDy() *big.Int
	// GetDeadline returns the deadline of the token deposit and swap.
	GetDeadline() *big.Int
}

// TokenMintAndSwapLog is the interface for the token mint and swap event.
type TokenMintAndSwapLog interface {
	EventLog
	// GetTo returns the recipient of the token mint and swap.
	GetTo() *common.Address
	// GetToken returns the token of the token mint and swap.
	// TODO
	// GetAmount returns the amount of the token mint and swap.
	GetAmount() *big.Int
	// GetFee returns the fee of the token mint and swap.
	GetFee() *big.Int
	// GetTokenIndexFrom returns the token index of the token mint and swap.
	GetTokenIndexFrom() uint8
	// GetTokenIndexTo returns the token index of the token mint and swap.
	GetTokenIndexTo() uint8
	// GetMinDy returns the min dy of the token mint and swap.
	GetMinDy() *big.Int
	// GetDeadline returns the deadline of the token mint and swap.
	GetDeadline() *big.Int
	// GetSwapSuccess returns the swap success of the token mint and swap.
	GetSwapSuccess() bool
	// GetKappa returns the kappa of the token mint and swap.
	GetKappa() [32]byte
}

// TokenRedeemAndSwapLog is the interface for the token redeem and swap event.
type TokenRedeemAndSwapLog interface {
	EventLog
	// GetTo returns the recipient of the token redeem and swap.
	GetTo() *common.Address
	// GetChainId returns the chain id of the token redeem and swap.
	GetChainId() *big.Int
	// GetToken returns the token of the token redeem and swap.
	// TODO
	// GetAmount returns the amount of the token redeem and swap.
	GetAmount() *big.Int
	// GetTokenIndexFrom returns the token index of the token redeem and swap.
	GetTokenIndexFrom() uint8
	// GetTokenIndexTo returns the token index of the token redeem and swap.
	GetTokenIndexTo() uint8
	// GetMinDy returns the min dy of the token redeem and swap.
	GetMinDy() *big.Int
	// GetDeadline returns the deadline of the token redeem and swap.
	GetDeadline() *big.Int
}

// TokenRedeemAndRemoveLog is the interface for the token redeem and remove event.
type TokenRedeemAndRemoveLog interface {
	EventLog
	// GetTo returns the recipient of the token redeem and remove.
	GetTo() *common.Address
	// GetChainId returns the chain id of the token redeem and remove.
	GetChainId() *big.Int
	// GetToken returns the token of the token redeem and remove.
	// TODO
	// GetAmount returns the amount of the token redeem and remove.
	GetAmount() *big.Int
	// GetSwapTokenIndex returns the swap token index of the token redeem and remove.
	GetSwapTokenIndex() uint8
	// GetSwapMinAmount returns the swap min amount of the token redeem and remove.
	GetSwapMinAmount() *big.Int
	// GetSwapDeadline returns the swap deadline of the token redeem and remove.
	GetSwapDeadline() *big.Int
}

// TokenWithdrawAndRemoveLog is the interface for the token withdraw and remove event.
type TokenWithdrawAndRemoveLog interface {
	EventLog
	// GetTo returns the recipient of the token withdraw and remove.
	GetTo() *common.Address
	// GetToken returns the token of the token withdraw and remove.
	// TODO
	// GetAmount returns the amount of the token withdraw and remove.
	GetAmount() *big.Int
	// GetFee returns the fee of the token withdraw and remove.
	GetFee() *big.Int
	// GetSwapTokenIndex returns the swap token index of the token withdraw and remove.
	GetSwapTokenIndex() uint8
	// GetSwapMinAmount returns the swap min amount of the token withdraw and remove.
	GetSwapMinAmount() *big.Int
	// GetSwapDeadline returns the swap deadline of the token withdraw and remove.
	GetSwapDeadline() *big.Int
	// GetSwapSuccess returns the swap success of the token withdraw and remove.
	GetSwapSuccess() bool
	// GetKappa returns the kappa of the token withdraw and remove.
	GetKappa() [32]byte
}

// TokenRedeemV2Log is the interface for the token redeem v2 event.
type TokenRedeemV2Log interface {
	EventLog
	// GetTo returns the recipient of the token redeem v2.
	GetTo() *common.Address
	// GetChainId returns the chain id of the token redeem v2.
	GetChainId() *big.Int
	// GetToken returns the token of the token redeem v2.
	// TODO
	// GetAmount returns the amount of the token redeem v2.
	GetAmount() *big.Int
}
