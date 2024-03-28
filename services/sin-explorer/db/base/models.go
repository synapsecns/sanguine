package base

// InterchainTransactionSent is the event emitted when a transaction is sent to another chain.
type InterchainTransactionSent struct {
	TransactionId   string `gorm:"column:transaction_id;primaryKey"`
	DbNonce         uint64 `gorm:"index"`
	EntryIndex      uint64 `gorm:"index"`
	DstChainId      uint64 `gorm:"index"`
	SrcSender       string `gorm:"index"`
	DstReceiver     string `gorm:"index"`
	VerificationFee string
	ExecutionFee    string
	Options         string
	Message         string
	TransactionHash string `gorm:"index"`
}

type InterchainTransactionReceived struct {
	TransactionId   string `gorm:"column:transaction_id;primaryKey"`
	DbNonce         uint64 `gorm:"index"`
	EntryIndex      uint64 `gorm:"index"`
	SrcChainId      uint64 `gorm:"index"`
	SrcSender       string `gorm:"index"`
	DstReceiver     string `gorm:"index"`
	TransactionHash string `gorm:"index"`
}
