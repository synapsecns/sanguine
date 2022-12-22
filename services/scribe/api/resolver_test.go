package api_test

import (
	"github.com/synapsecns/sanguine/services/scribe/graphql"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
)

func (g APISuite) TestLogResolvers() {
	chainID := uint32(1)
	blockNumber := uint64(16131419)
	// store a transaction
	tx := g.buildEthTx()
	err := g.db.StoreEthTx(g.GetTestContext(), tx, chainID, common.BigToHash(big.NewInt(gofakeit.Int64())), blockNumber, gofakeit.Uint64())
	Nil(g.T(), err)
	// store a log
	log := g.buildLog(common.BigToAddress(big.NewInt(gofakeit.Int64())), blockNumber)
	log.TxHash = tx.Hash()
	err = g.db.StoreLogs(g.GetTestContext(), chainID, log)
	Nil(g.T(), err)
	// store a receipt
	receipt := g.buildReceipt(common.BigToAddress(big.NewInt(gofakeit.Int64())), blockNumber)
	receipt.TxHash = tx.Hash()
	receipt.Logs = []*types.Log{&log}
	err = g.db.StoreReceipt(g.GetTestContext(), chainID, receipt)
	Nil(g.T(), err)

	// test the log's resolver for the transaction and receipt
	logResolver, err := g.gqlClient.GetLogsResolvers(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	retrievedTx := logResolver.Response[0].Transaction
	Equal(g.T(), retrievedTx.ChainID, int(chainID))
	Equal(g.T(), retrievedTx.TxHash, tx.Hash().String())
	Equal(g.T(), retrievedTx.Protected, tx.Protected())
	Equal(g.T(), retrievedTx.Type, int(tx.Type()))
	Equal(g.T(), retrievedTx.Data, common.Bytes2Hex(tx.Data()))
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
	Equal(g.T(), retrievedReceipt.Bloom, common.Bytes2Hex(receipt.Bloom.Bytes()))
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
	Equal(g.T(), retrievedTx.Data, common.Bytes2Hex(tx.Data()))
	Equal(g.T(), retrievedTx.Gas, int(tx.Gas()))
	Equal(g.T(), retrievedTx.GasPrice, int(tx.GasPrice().Uint64()))
	Equal(g.T(), retrievedTx.GasTipCap, tx.GasTipCap().String())
	Equal(g.T(), retrievedTx.GasFeeCap, tx.GasFeeCap().String())
	Equal(g.T(), retrievedTx.Value, tx.Value().String())
	Equal(g.T(), retrievedTx.Nonce, int(tx.Nonce()))
	Equal(g.T(), retrievedTx.To, tx.To().String())
	parsedLog, err := graphql.ParseLog(*receiptResolver.Response[0].Logs[0])
	Nil(g.T(), err)
	Equal(g.T(), *parsedLog, log)

	// test the transaction's resolver for the receipt and logs
	txResolver, err := g.gqlClient.GetTransactionsResolvers(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	retrievedReceipt = txResolver.Response[0].Receipt
	Equal(g.T(), retrievedReceipt.ChainID, int(chainID))
	Equal(g.T(), retrievedReceipt.Type, int(receipt.Type))
	Equal(g.T(), retrievedReceipt.PostState, string(receipt.PostState))
	Equal(g.T(), retrievedReceipt.Status, int(receipt.Status))
	Equal(g.T(), retrievedReceipt.CumulativeGasUsed, int(receipt.CumulativeGasUsed))
	Equal(g.T(), retrievedReceipt.Bloom, common.Bytes2Hex(receipt.Bloom.Bytes()))
	Equal(g.T(), retrievedReceipt.TxHash, receipt.TxHash.String())
	Equal(g.T(), retrievedReceipt.ContractAddress, receipt.ContractAddress.String())
	Equal(g.T(), retrievedReceipt.GasUsed, int(receipt.GasUsed))
	Equal(g.T(), retrievedReceipt.BlockNumber, int(receipt.BlockNumber.Int64()))
	Equal(g.T(), retrievedReceipt.TransactionIndex, int(receipt.TransactionIndex))
	parsedLog, err = graphql.ParseLog(*txResolver.Response[0].Logs[0])
	Nil(g.T(), err)
	Equal(g.T(), *parsedLog, log)
}
