// Package balance keeps track of the balance
package balance

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core"
	"math/big"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/bindings"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/config"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/utils"

	EVMClient "github.com/synapsecns/sanguine/ethergo/client"
)

const precision = 5

// IBalanceManager is the interface for a BalanceManager.
type IBalanceManager interface {
	IncrementBalance(tokenID string, amount *big.Int)
	DecrementBalance(tokenID string, amount *big.Int)
	GetAllOnChainBalances(ctx context.Context) error
	GetOnChainBalance(ctx context.Context, tokenID string) error
	GetBalance(tokenID string) *Balance
	GetNormalizedBalance(tokenID string) (float64, error)
}

type balanceManageImpl struct {
	clients        map[uint32]EVMClient.EVM
	erc20s         map[string]*bindings.MockERC20
	balances       map[string]*Balance
	relayerAddress common.Address
	// TODO: use mapMutex
	mux sync.RWMutex // Mutex to ensure thread safety
}

// NewBalanceManager creates a new balance manager.
func NewBalanceManager(clients map[uint32]EVMClient.EVM, assets []config.AssetConfig, relayerAddress common.Address) (IBalanceManager, error) {
	erc20s := make(map[string]*bindings.MockERC20)
	balances := make(map[string]*Balance)
	for _, asset := range assets {
		tokenAddress := common.HexToAddress(asset.Address)
		tokenID := utils.GenerateTokenID(asset.ChainID, tokenAddress)
		erc20Instance, err := bindings.NewMockERC20(tokenAddress, clients[asset.ChainID])
		if err != nil {
			return nil, fmt.Errorf("could not create ERC20 instance for token %s on chain %d: %w", tokenAddress.String(), asset.ChainID, err)
		}
		erc20s[tokenID] = erc20Instance
		balances[tokenID] = &Balance{}
	}
	return &balanceManageImpl{
		erc20s:         erc20s,
		clients:        clients,
		relayerAddress: relayerAddress,
		balances:       balances,
	}, nil
}

// IncrementBalance increments the balance of an asset by a given amount.
func (i *balanceManageImpl) IncrementBalance(tokenID string, amount *big.Int) {
	i.mux.Lock()
	defer i.mux.Unlock()

	currentBalance, ok := i.balances[tokenID]
	if !ok {
		currentBalance = &Balance{}
	}
	newBalance := new(big.Int).Add(currentBalance.Amount, amount)
	i.balances[tokenID].Amount = core.CopyBigInt(newBalance)
}

// DecrementBalance decrements the balance of an asset by a given amount.
func (i *balanceManageImpl) DecrementBalance(tokenID string, amount *big.Int) {
	i.mux.Lock()
	defer i.mux.Unlock()

	currentBalance, ok := i.balances[tokenID]
	if !ok {
		currentBalance = &Balance{}
	}
	newBalance := new(big.Int).Sub(currentBalance.Amount, amount)
	i.balances[tokenID].Amount = core.CopyBigInt(newBalance)
}

func (i *balanceManageImpl) GetBalance(tokenID string) *Balance {
	i.mux.RLock()
	defer i.mux.RUnlock()

	balance, ok := i.balances[tokenID]
	if !ok {
		return nil
	}
	return balance
}

func (i *balanceManageImpl) GetNormalizedBalance(tokenID string) (float64, error) {
	i.mux.RLock()
	defer i.mux.RUnlock()

	balance, ok := i.balances[tokenID]
	if !ok {
		return 0, nil
	}
	// Get normalized balance
	return balance.ToFloat64()
}

// GetAllOnChainBalances fetches the latest balance for all assets in the balance handler.
func (i *balanceManageImpl) GetAllOnChainBalances(ctx context.Context) error {
	for tokenID := range i.balances {
		err := i.GetOnChainBalance(ctx, tokenID)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetOnChainBalance fetches the latest balance of an asset from the blockchain for a given chain ID and token address (tokenID).
func (i *balanceManageImpl) GetOnChainBalance(ctx context.Context, tokenID string) error {
	ops := &bind.CallOpts{
		Context: ctx,
	}
	balance, err := i.erc20s[tokenID].BalanceOf(ops, i.relayerAddress)
	if err != nil {
		return fmt.Errorf("could not retrieve balance of token ID %s: %w", tokenID, err)
	}

	decimals, err := i.erc20s[tokenID].Decimals(ops)
	if err != nil {
		return fmt.Errorf("could not retrieve decimals of token ID %s: %w", tokenID, err)
	}
	newBalance := &Balance{
		Amount:   core.CopyBigInt(balance),
		Decimals: decimals,
	}

	i.mux.Lock()
	defer i.mux.Unlock()

	// Set balance
	i.balances[tokenID] = newBalance
	return nil
}

// Balance represents a token balance.
type Balance struct {
	Amount   *big.Int
	Decimals uint8
}

// ToFloat64 converts a balance to a float64.
func (b *Balance) ToFloat64() (float64, error) {
	// Get normalized balance
	decimalMultiplier := new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(b.Decimals)), nil))
	normalizedBalance := new(big.Float).Quo(new(big.Float).SetInt(b.Amount), decimalMultiplier)
	trueAmountStr := normalizedBalance.SetMode(big.AwayFromZero).Text('f', precision)
	amount, err := strconv.ParseFloat(trueAmountStr, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse float64 from string %s: %w", trueAmountStr, err)
	}
	return amount, nil
}
