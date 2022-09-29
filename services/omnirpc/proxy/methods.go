package proxy

const (
	// ChainIDMethod is used to retrieve the current chain ID for transaction replay protection.
	ChainIDMethod = "eth_chainId"
	// BlockByHashMethod gets a block by hash
	BlockByHashMethod = "eth_getBlockByHash"
	// BlockByNumberMethod gets a block by number
	BlockByNumberMethod = "eth_getBlockByNumber"
	// BlockNumberMethod gets the latest block number
	BlockNumberMethod = "eth_blockNumber"
	// TransactionByHashMethod returns the transaction with the given hash
	TransactionByHashMethod = "eth_getTransactionByHash"
	// TransactionByBlockHashAndIndexMethod returns a transaction by its hash and index
	TransactionByBlockHashAndIndexMethod = "eth_getTransactionByBlockHashAndIndex"
	// TransactionCountByHashMethod get sthe transaction count by hash
	TransactionCountByHashMethod = "eth_getBlockTransactionCountByHash"
	// TransactionReceiptByHashMethod gets the transaction receipt from a hash
	TransactionReceiptByHashMethod = "eth_getTransactionReceipt"
	// SyncProgressMethod gets the sync progress
	SyncProgressMethod = "eth_syncing"
	// GetBalanceMethod gets the balance for a given address
	GetBalanceMethod = "eth_getBalance"
	// StorageAtMethod returns the value of key in the contract storage of the given account.
	StorageAtMethod = "eth_getStorageAt"
	// CodeAtMethod returns the contract code of the given account.
	CodeAtMethod = "eth_getCode"
	// NonceAtMethod returns the account nonce of the given account.
	NonceAtMethod = "eth_getTransactionCount"
)
