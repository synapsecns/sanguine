package anvil

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

const (
	anvilNamespace = "anvil"
	evmNamespace   = "evm"
)

// Client is a client for interacting with anvil.
// Note: this should not be used directly unless dealing with an external anvil instance.
// Please use the Backend instead.
type Client struct {
	*rpc.Client
}

// nolint: wrapcheck
func (c *Client) callAnvilContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	return c.CallContext(ctx, result, fmt.Sprintf("%s_%s", anvilNamespace, method), args...)
}

// nolint: wrapcheck
func (c *Client) callEvmContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	return c.CallContext(ctx, result, fmt.Sprintf("%s_%s", evmNamespace, method), args...)
}

// Dial dials an anvil instance.
func Dial(ctx context.Context, rpcURL string) (*Client, error) {
	c, err := rpc.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("could not create client: %w", err)
	}
	return &Client{c}, nil
}

// ImpersonateAccount impersonates an account.
// Send transactions impersonating an externally owned account or contract.
// While impersonating a contract, the contract functions can not be called. anvil_stopImpersonatingAccount must be used if the contract's functions are to be called again. See also EIP-3607.
func (c *Client) ImpersonateAccount(ctx context.Context, address common.Address) error {
	return c.callAnvilContext(ctx, nil, "impersonateAccount", address.String())
}

// StopImpersonatingAccount stops impersonating an account or contract if previously set with anvil_impersonateAccount.
func (c *Client) StopImpersonatingAccount(ctx context.Context, address common.Address) error {
	return c.callAnvilContext(ctx, nil, "stopImpersonatingAccount", address.String())
}

// GetAutomine returns whether automine is enabled.
func (c *Client) GetAutomine(ctx context.Context) (isAutomining bool, err error) {
	err = c.callAnvilContext(ctx, &isAutomining, "getAutomine")
	return
}

// Mine mines a number of blocks.
func (c *Client) Mine(ctx context.Context, blockCount uint) error {
	return c.callAnvilContext(ctx, nil, "mine", blockCount)
}

// DropTransaction drops a transactions from the memopool.
func (c *Client) DropTransaction(ctx context.Context, tx common.Hash) error {
	return c.callAnvilContext(ctx, nil, "dropTransaction", tx)
}

// Reset the fork to a fresh forked state, and optionally update the fork config.
func (c *Client) Reset(ctx context.Context) error {
	return c.callAnvilContext(ctx, nil, "reset")
}

// SetRPCURL sets the rpc url.
func (c *Client) SetRPCURL(ctx context.Context, rpcURL string) error {
	return c.callAnvilContext(ctx, nil, "setRpcUrl", rpcURL)
}

// SetBalance sets the balance of an account.
func (c *Client) SetBalance(ctx context.Context, address common.Address, balance uint64) error {
	return c.callAnvilContext(ctx, nil, "setBalance", address.String(), balance)
}

// SetCode sets the code of an account.
func (c *Client) SetCode(ctx context.Context, address common.Address, code []byte) error {
	return c.callAnvilContext(ctx, nil, "setCode", address.String(), hexutil.Encode(code))
}

// SetNonce sets the nonce of an account.
func (c *Client) SetNonce(ctx context.Context, address common.Address, nonce uint64) error {
	return c.callAnvilContext(ctx, nil, "setNonce", address.String(), nonce)
}

// SetStorageAt sets the storage of an account.
func (c *Client) SetStorageAt(ctx context.Context, address common.Address, key common.Hash, value common.Hash) error {
	return c.callAnvilContext(ctx, nil, "setStorageAt", address.String(), key, value)
}

// SetCoinbase sets the coinbase of the chain.
func (c *Client) SetCoinbase(ctx context.Context, coinbase common.Address) error {
	return c.callAnvilContext(ctx, nil, "setCoinbase", coinbase.String())
}

// SetLoggingEnabled Enable or disable logging.
func (c *Client) SetLoggingEnabled(ctx context.Context, enabled bool) error {
	return c.callAnvilContext(ctx, nil, "setLoggingEnabled", enabled)
}

// SetMinGasPrice sets the minimum gas price.
// this is disabled when eip-1559 is enabled.
func (c *Client) SetMinGasPrice(ctx context.Context, minGasPrice *big.Int) error {
	return c.callAnvilContext(ctx, nil, "setMinGasPrice", minGasPrice)
}

// SetNextBlockBaseFeePerGas sets the base fee per gas for the next block.
func (c *Client) SetNextBlockBaseFeePerGas(ctx context.Context, baseFee *big.Int) error {
	return c.callAnvilContext(ctx, nil, "setNextBlockBaseFeePerGas", baseFee)
}

// DumpState returns a hex string representing the complete state of the chain.
// Can be re-imported into a fresh/restarted instance of Anvil to reattain the same state.
func (c *Client) DumpState(ctx context.Context) (string, error) {
	var stateHex string
	err := c.callAnvilContext(ctx, &stateHex, "dumpState")
	return stateHex, err
}

// LoadState merges the contents of a hex string previously returned by DumpState
// into the current chain state. Will overwrite any colliding accounts/storage slots.
func (c *Client) LoadState(ctx context.Context, stateHex string) error {
	return c.callAnvilContext(ctx, nil, "loadState", stateHex)
}

// NodeInfo retrieves the configuration params for the currently running Anvil node.
func (c *Client) NodeInfo(ctx context.Context) (map[string]interface{}, error) {
	var nodeInfo map[string]interface{}
	err := c.callAnvilContext(ctx, &nodeInfo, "nodeInfo")
	return nodeInfo, err
}

// SetAutomine enables or disables automatic mining of new blocks with each new transaction.
func (c *Client) SetAutomine(ctx context.Context, enabled bool) error {
	return c.callEvmContext(ctx, nil, "setAutomine", enabled)
}

// SetIntervalMining sets the mining behavior to interval with the given interval (seconds).
func (c *Client) SetIntervalMining(ctx context.Context, interval int) error {
	return c.callEvmContext(ctx, nil, "setIntervalMining", interval)
}

// Snapshot captures the state of the blockchain at the current block.
func (c *Client) Snapshot(ctx context.Context) (snapshotID string, err error) {
	err = c.callEvmContext(ctx, &snapshotID, "snapshot")
	return
}

// Revert reverts the state of the blockchain to a previous snapshot.
func (c *Client) Revert(ctx context.Context, snapshotID string) error {
	return c.callEvmContext(ctx, nil, "revert", snapshotID)
}

// IncreaseTime jumps forward in time by the given amount of time, in seconds.
func (c *Client) IncreaseTime(ctx context.Context, seconds int64) error {
	return c.callEvmContext(ctx, nil, "increaseTime", seconds)
}

// SetNextBlockTimestamp sets the exact timestamp for the next block.
func (c *Client) SetNextBlockTimestamp(ctx context.Context, timestamp int64) error {
	return c.callEvmContext(ctx, nil, "setNextBlockTimestamp", timestamp)
}

// SetBlockTimestampInterval sets a block timestamp interval.
func (c *Client) SetBlockTimestampInterval(ctx context.Context, interval int64) error {
	return c.callAnvilContext(ctx, nil, "setBlockTimestampInterval", interval)
}

// RemoveBlockTimestampInterval removes an anvil_setBlockTimestampInterval if it exists.
func (c *Client) RemoveBlockTimestampInterval(ctx context.Context) error {
	return c.callAnvilContext(ctx, nil, "removeBlockTimestampInterval")
}

// SetBlockGasLimit sets the block gas limit for the following blocks.
func (c *Client) SetBlockGasLimit(ctx context.Context, gasLimit uint64) error {
	return c.callEvmContext(ctx, nil, "setBlockGasLimit", gasLimit)
}

// EvmMine mines a single block.
func (c *Client) EvmMine(ctx context.Context) error {
	return c.callEvmContext(ctx, nil, "mine")
}

// EnableTraces enables call traces for transactions that are returned to the user when they execute a transaction.
func (c *Client) EnableTraces(ctx context.Context) error {
	return c.callAnvilContext(ctx, nil, "enableTraces")
}

// anvilTransactionLegacy represents a transaction that will serialize to the correct JSON.
type anvilTransactionLegacy struct {
	From            string `json:"from"`
	To              string `json:"to"`
	GasPrice        string `json:"gasPrice"`
	Gas             string `json:"gas"`
	Value           string `json:"value,omitempty"`
	Data            string `json:"data"`
	Nonce           string `json:"nonce"`
	TransactionType string `json:"type"`
}

// anvilTransactionDynamic represents a transaction that will serialize to the correct JSON.
type anvilTransactionDynamic struct {
	From                 string `json:"from"`
	To                   string `json:"to"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	Gas                  string `json:"gas"`
	Value                string `json:"value,omitempty"`
	Data                 string `json:"data"`
	Nonce                string `json:"nonce"`
	TransactionType      string `json:"type"`
}

// SendUnsignedTransaction sends a transaction to the anvil node.
// It is the responsibility of the caller to call impersonateAccount and revertImpersonatedAccount.
func (c *Client) SendUnsignedTransaction(ctx context.Context, from common.Address, tx *types.Transaction) error {
	var anTx interface{}
	if tx.Type() == types.LegacyTxType {
		anTx = anvilTransactionLegacy{
			From:            from.Hex(),
			To:              tx.To().Hex(),
			GasPrice:        bigIntToString(tx.GasPrice().Int64()),
			Gas:             bigIntToString(int64(tx.Gas())),
			Data:            hex.EncodeToString(tx.Data()),
			Nonce:           bigIntToString(int64(tx.Nonce())),
			Value:           fmt.Sprintf("%x", tx.Value()),
			TransactionType: bigIntToString(int64(tx.Type())),
		}
	} else {
		anTx = anvilTransactionDynamic{
			From:                 from.Hex(),
			To:                   tx.To().Hex(),
			MaxFeePerGas:         bigIntToString(tx.GasPrice().Int64()),
			MaxPriorityFeePerGas: bigIntToString(tx.GasPrice().Int64()),
			Gas:                  bigIntToString(int64(tx.Gas())),
			Data:                 hex.EncodeToString(tx.Data()),
			Nonce:                bigIntToString(int64(tx.Nonce())),
			Value:                fmt.Sprintf("%x", tx.Value()),
			TransactionType:      bigIntToString(int64(tx.Type())),
		}
	}

	// nolint: wrapcheck
	return c.CallContext(ctx, nil, "eth_sendTransaction", anTx)
}

func bigIntToString(n int64) string {
	b := big.NewInt(n).Bytes()
	if len(b) == 0 {
		return "0x00"
	}
	return hex.EncodeToString(b)
}

// Close closes the conn.
func (c *Client) Close() {
	c.Client.Close()
}
