package db

// LogFilter is a filter to use when querying the database for logs.
type LogFilter struct {
	ContractAddress string
	ChainID         uint32
	BlockNumber     uint64
	TxHash          string
	TxIndex         uint64
	BlockHash       string
	Index           uint64
	Confirmed       bool
}

// ReceiptFilter is a filter to use when querying the database for receipts.
type ReceiptFilter struct {
	ChainID          uint32
	TxHash           string
	ContractAddress  string
	BlockHash        string
	BlockNumber      uint64
	TransactionIndex uint64
	Confirmed        bool
}

// EthTxFilter is a filter to use when querying the database for eth transactions.
type EthTxFilter struct {
	TxHash      string
	ChainID     uint32
	BlockHash   string
	BlockNumber uint64
	Confirmed   bool
}

// FailedLogFilter is a filter to use when querying the database for failed blocks.
type FailedLogFilter struct {
	ChainID         uint32
	ContractAddress string
	TxHash          string
	BlockIndex      uint64
	BlockNumber     uint64
}
