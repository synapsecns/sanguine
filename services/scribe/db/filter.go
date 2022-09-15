package db

// BuildLogFilter builds a log filter from nullable parameters.
func BuildLogFilter(contractAddress *string, blockNumber *int, txHash *string, txIndex *int, blockHash *string, index *int, confirmed *bool) LogFilter {
	logFilter := LogFilter{}
	if contractAddress != nil {
		logFilter.ContractAddress = *contractAddress
	}
	if blockNumber != nil {
		logFilter.BlockNumber = uint64(*blockNumber)
	}
	if txHash != nil {
		logFilter.TxHash = *txHash
	}
	if txIndex != nil {
		logFilter.TxIndex = uint64(*txIndex)
	}
	if blockHash != nil {
		logFilter.BlockHash = *blockHash
	}
	if index != nil {
		logFilter.Index = uint64(*index)
	}
	if confirmed != nil {
		logFilter.Confirmed = *confirmed
	}
	return logFilter
}

// BuildReceiptFilter builds a receipt filter.
func BuildReceiptFilter(txHash *string, contractAddress *string, blockHash *string, blockNumber *int, transactionIndex *int, confirmed *bool) ReceiptFilter {
	receiptFilter := ReceiptFilter{}
	if txHash != nil {
		receiptFilter.TxHash = *txHash
	}
	if contractAddress != nil {
		receiptFilter.ContractAddress = *contractAddress
	}
	if blockHash != nil {
		receiptFilter.BlockHash = *blockHash
	}
	if blockNumber != nil {
		receiptFilter.BlockNumber = uint64(*blockNumber)
	}
	if transactionIndex != nil {
		receiptFilter.TransactionIndex = uint64(*transactionIndex)
	}
	if confirmed != nil {
		receiptFilter.Confirmed = *confirmed
	}
	return receiptFilter
}

// BuildEthTxFilter cannot build eth tx filter.
func BuildEthTxFilter(txHash *string, blockNumber *int, blockHash *string, confirmed *bool) EthTxFilter {
	ethTxFilter := EthTxFilter{}
	if txHash != nil {
		ethTxFilter.TxHash = *txHash
	}
	if blockNumber != nil {
		ethTxFilter.BlockNumber = uint64(*blockNumber)
	}
	if blockHash != nil {
		ethTxFilter.BlockHash = *blockHash
	}
	if confirmed != nil {
		ethTxFilter.Confirmed = *confirmed
	}
	return ethTxFilter
}
