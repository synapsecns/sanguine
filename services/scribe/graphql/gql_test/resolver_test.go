package gql_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
)

func (g GQLSuite) TestLogResolvers() {
	chainID := gofakeit.Uint32()
	// store a transaction
	tx := g.buildEthTx()
	err := g.db.StoreEthTx(g.GetTestContext(), tx, chainID, gofakeit.Uint64())
	Nil(g.T(), err)
	// store a log
	log := g.buildLog(common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint64())
	log.TxHash = tx.Hash()
	err = g.db.StoreLog(g.GetTestContext(), log, chainID)
	Nil(g.T(), err)
	// store a receipt
	receipt := g.buildReceipt(common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint64())
	receipt.TxHash = tx.Hash()
	receipt.Logs = []*types.Log{&log}
	err = g.db.StoreReceipt(g.GetTestContext(), receipt, chainID)
	Nil(g.T(), err)

	// test the log's resolver for the transaction and receipt
	logResolver, err := g.gqlClient.GetLogsResolvers(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	retrievedTx := logResolver.Response[0].Transaction
	Equal(g.T(), retrievedTx.ChainID, int(chainID))
	Equal(g.T(), retrievedTx.TxHash, tx.Hash().String())
	Equal(g.T(), retrievedTx.Protected, tx.Protected())
	Equal(g.T(), retrievedTx.Type, int(tx.Type()))
	Equal(g.T(), retrievedTx.Data, common.BytesToHash(tx.Data()).String())
	Equal(g.T(), retrievedTx.Gas, int(tx.Gas()))
	Equal(g.T(), retrievedTx.GasPrice, int(tx.GasPrice().Uint64()))
	Equal(g.T(), retrievedTx.GasTipCap, tx.GasTipCap().String())
	Equal(g.T(), retrievedTx.GasFeeCap, tx.GasFeeCap().String())
	Equal(g.T(), retrievedTx.Value, tx.Value().String())
	Equal(g.T(), retrievedTx.Nonce, int(tx.Nonce()))
	Equal(g.T(), retrievedTx.To, tx.To().String())
	retrievedReceipt := logResolver.Response[0].Receipt
	Equal(g.T(), retrievedReceipt.ChainID, int(chainID))
	Equal(g.T(), retrievedReceipt.Type, int(receipt.Type))
	Equal(g.T(), retrievedReceipt.PostState, string(receipt.PostState))
	Equal(g.T(), retrievedReceipt.Status, int(receipt.Status))
	Equal(g.T(), retrievedReceipt.CumulativeGasUsed, int(receipt.CumulativeGasUsed))
	Equal(g.T(), retrievedReceipt.Bloom, common.BytesToHash(receipt.Bloom.Bytes()).String())
	Equal(g.T(), retrievedReceipt.TxHash, receipt.TxHash.String())
	Equal(g.T(), retrievedReceipt.ContractAddress, receipt.ContractAddress.String())
	Equal(g.T(), retrievedReceipt.GasUsed, int(receipt.GasUsed))
	Equal(g.T(), retrievedReceipt.BlockNumber, int(receipt.BlockNumber.Int64()))
	Equal(g.T(), retrievedReceipt.TransactionIndex, int(receipt.TransactionIndex))

	// test the receipt's resolver for the transaction and logs
	receiptResolver, err := g.gqlClient.GetReceiptsResolvers(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	retrievedTx = receiptResolver.Response[0].Transaction
	Equal(g.T(), retrievedTx.ChainID, int(chainID))
	Equal(g.T(), retrievedTx.TxHash, tx.Hash().String())
	Equal(g.T(), retrievedTx.Protected, tx.Protected())
	Equal(g.T(), retrievedTx.Type, int(tx.Type()))
	Equal(g.T(), retrievedTx.Data, common.BytesToHash(tx.Data()).String())
	Equal(g.T(), retrievedTx.Gas, int(tx.Gas()))
	Equal(g.T(), retrievedTx.GasPrice, int(tx.GasPrice().Uint64()))
	Equal(g.T(), retrievedTx.GasTipCap, tx.GasTipCap().String())
	Equal(g.T(), retrievedTx.GasFeeCap, tx.GasFeeCap().String())
	Equal(g.T(), retrievedTx.Value, tx.Value().String())
	Equal(g.T(), retrievedTx.Nonce, int(tx.Nonce()))
	Equal(g.T(), retrievedTx.To, tx.To().String())
	retrievedLog := receiptResolver.Response[0].Logs[0]
	// convert receiptTopics
	var receiptTopics []string
	for _, topic := range log.Topics {
		receiptTopics = append(receiptTopics, topic.String())
	}
	Equal(g.T(), retrievedLog.ContractAddress, log.Address.String())
	Equal(g.T(), retrievedLog.ChainID, int(chainID))
	Equal(g.T(), retrievedLog.Topics, receiptTopics)
	Equal(g.T(), retrievedLog.Data, common.BytesToHash(log.Data).String())
	Equal(g.T(), retrievedLog.BlockNumber, int(log.BlockNumber))
	Equal(g.T(), retrievedLog.TxHash, log.TxHash.String())
	Equal(g.T(), retrievedLog.TxIndex, int(log.TxIndex))
	Equal(g.T(), retrievedLog.BlockHash, log.BlockHash.String())
	Equal(g.T(), retrievedLog.Index, int(log.Index))
	Equal(g.T(), retrievedLog.Removed, log.Removed)

	// test the transaction's resolver for the receipt and logs
	txResolver, err := g.gqlClient.GetTransactionsResolvers(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	retrievedReceipt = txResolver.Response[0].Receipt
	Equal(g.T(), retrievedReceipt.ChainID, int(chainID))
	Equal(g.T(), retrievedReceipt.Type, int(receipt.Type))
	Equal(g.T(), retrievedReceipt.PostState, string(receipt.PostState))
	Equal(g.T(), retrievedReceipt.Status, int(receipt.Status))
	Equal(g.T(), retrievedReceipt.CumulativeGasUsed, int(receipt.CumulativeGasUsed))
	Equal(g.T(), retrievedReceipt.Bloom, common.BytesToHash(receipt.Bloom.Bytes()).String())
	Equal(g.T(), retrievedReceipt.TxHash, receipt.TxHash.String())
	Equal(g.T(), retrievedReceipt.ContractAddress, receipt.ContractAddress.String())
	Equal(g.T(), retrievedReceipt.GasUsed, int(receipt.GasUsed))
	Equal(g.T(), retrievedReceipt.BlockNumber, int(receipt.BlockNumber.Int64()))
	Equal(g.T(), retrievedReceipt.TransactionIndex, int(receipt.TransactionIndex))
	retrievedLog = txResolver.Response[0].Logs[0]
	// convert txTopics
	var txTopics []string
	for _, topic := range log.Topics {
		txTopics = append(txTopics, topic.String())
	}
	Equal(g.T(), retrievedLog.ContractAddress, log.Address.String())
	Equal(g.T(), retrievedLog.ChainID, int(chainID))
	Equal(g.T(), retrievedLog.Topics, txTopics)
	Equal(g.T(), retrievedLog.Data, common.BytesToHash(log.Data).String())
	Equal(g.T(), retrievedLog.BlockNumber, int(log.BlockNumber))
	Equal(g.T(), retrievedLog.TxHash, log.TxHash.String())
	Equal(g.T(), retrievedLog.TxIndex, int(log.TxIndex))
	Equal(g.T(), retrievedLog.BlockHash, log.BlockHash.String())
	Equal(g.T(), retrievedLog.Index, int(log.Index))
	Equal(g.T(), retrievedLog.Removed, log.Removed)
}
