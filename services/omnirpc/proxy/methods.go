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
)
