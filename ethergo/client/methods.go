package client

import (
	"github.com/benbjohnson/immutable"
	"github.com/ipfs/go-log"
	"strings"
)

var logger = log.Logger("ethergo-client-logger")

// RPCMethod is an enum type for an rpc method.
type RPCMethod string

// NOTE: any changes here must be added to allMethods list below.

// ETH METHODS:.
const (
	// ChainIDMethod is used to retrieve the current chain ID for transaction replay protection.
	ChainIDMethod RPCMethod = "eth_chainId"
	// BlockByHashMethod gets a block by hash.
	BlockByHashMethod RPCMethod = "eth_getBlockByHash"
	// BlockByNumberMethod gets a block by number.
	BlockByNumberMethod RPCMethod = "eth_getBlockByNumber"
	// BlockNumberMethod gets the latest block number.
	BlockNumberMethod RPCMethod = "eth_blockNumber"
	// TransactionByHashMethod returns the transaction with the given hash.
	TransactionByHashMethod RPCMethod = "eth_getTransactionByHash"
	// TransactionByBlockHashAndIndexMethod returns a transaction by its hash and index.
	TransactionByBlockHashAndIndexMethod RPCMethod = "eth_getTransactionByBlockHashAndIndex"
	// TransactionCountByHashMethod get sthe transaction count by hash.
	TransactionCountByHashMethod RPCMethod = "eth_getBlockTransactionCountByHash"
	// TransactionReceiptByHashMethod gets the transaction receipt from a hash.
	TransactionReceiptByHashMethod RPCMethod = "eth_getTransactionReceipt"
	// SyncProgressMethod gets the sync progress.
	SyncProgressMethod RPCMethod = "eth_syncing"
	// GetBalanceMethod gets the balance for a given address.
	GetBalanceMethod RPCMethod = "eth_getBalance"
	// StorageAtMethod returns the value of key in the contract storage of the given account.
	StorageAtMethod RPCMethod = "eth_getStorageAt"
	// GetCodeMethod returns the contract code of the given account.
	GetCodeMethod RPCMethod = "eth_getCode"
	// TransactionCountMethod returns the account nonce of the given account.
	TransactionCountMethod RPCMethod = "eth_getTransactionCount"
	// GetLogsMethod filters logs.
	GetLogsMethod RPCMethod = "eth_getLogs"
	// CallMethod calls a contract.
	CallMethod RPCMethod = "eth_call"
	// GasPriceMethod gets the gas price.
	GasPriceMethod RPCMethod = "eth_gasPrice"
	// MaxPriorityMethod gets the max priority fee.
	MaxPriorityMethod RPCMethod = "eth_maxPriorityFeePerGas"
	// FeeHistoryMethod gets the fee history.
	FeeHistoryMethod RPCMethod = "eth_feeHistory"
	// EstimateGasMethod tries to estimate the gas needed to execute a specific transaction.
	EstimateGasMethod RPCMethod = "eth_estimateGas"
	// PendingTransactionCountMethod gets the pending transaction count.
	PendingTransactionCountMethod RPCMethod = "eth_getBlockTransactionCountByNumber"
	// SendRawTransactionMethod sends a raw tx.
	SendRawTransactionMethod RPCMethod = "eth_sendRawTransaction"
	// SubscribeMethod subscribes to an event.
	SubscribeMethod RPCMethod = "eth_subscribe"
	// HarmonyGetLogsMethod filters logs on harmony.
	HarmonyGetLogsMethod RPCMethod = "hmy_getLogs"
	// HarmonyGetReceiptMethod gets a receipt on harmony
	// this should always return Hash rather than Ethereum Hash
	HarmonyGetReceiptMethod RPCMethod = "hmy_getTransactionReceipt"
)

// NetMethods:.
const (
	// NetVersionMethod gets the network version.
	NetVersionMethod  RPCMethod = "net_version"
	Web3VersionMethod RPCMethod = "web3_clientVersion"
)

// allMethods gets all available rpc methods.
var allMethods = []RPCMethod{ChainIDMethod, BlockByHashMethod, BlockByNumberMethod, BlockNumberMethod,
	BlockNumberMethod, TransactionByHashMethod, TransactionByBlockHashAndIndexMethod, TransactionCountByHashMethod,
	TransactionReceiptByHashMethod, SyncProgressMethod, GetBalanceMethod, SubscribeMethod, NetVersionMethod, Web3VersionMethod, HarmonyGetReceiptMethod, HarmonyGetLogsMethod}

var methodMap *immutable.Map[RPCMethod, string]

func init() {
	methodLowerMap := immutable.NewMapBuilder[RPCMethod, string](nil)
	for _, method := range allMethods {
		methodLowerMap.Set(method, strings.ToLower(string(method)))
	}
	methodMap = methodLowerMap.Map()
}

// Comparable converts an rpc method to lowercase based on a preset map
// can be used for case sensitive comparison.
func (r RPCMethod) Comparable() string {
	res, ok := methodMap.Get(r)
	if !ok {
		logger.Warnf("rpc method not found for %s", r)
	}

	return res
}

// String returns the string representation of an rpc method.
func (r RPCMethod) String() string {
	return string(r)
}
