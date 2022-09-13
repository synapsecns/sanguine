package gql_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
)

func (g GQLSuite) TestRetrieveData() {
	contractAddressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	contractAddressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	chainID := gofakeit.Uint32()

	// create and store logs, receipts, and txs
	var log types.Log
	var receipt types.Receipt
	var tx *types.Transaction
	var err error
	for blockNumber := 0; blockNumber < 10; blockNumber++ {
		// create and store logs
		log = g.buildLog(contractAddressA, uint64(blockNumber))
		err = g.db.StoreLog(g.GetTestContext(), log, chainID)
		Nil(g.T(), err)
		log = g.buildLog(contractAddressB, uint64(blockNumber))
		err = g.db.StoreLog(g.GetTestContext(), log, chainID)
		Nil(g.T(), err)
		// create and store receipts
		receipt = g.buildReceipt(contractAddressA, uint64(blockNumber))
		err = g.db.StoreReceipt(g.GetTestContext(), receipt, chainID)
		Nil(g.T(), err)
		receipt = g.buildReceipt(contractAddressB, uint64(blockNumber))
		err = g.db.StoreReceipt(g.GetTestContext(), receipt, chainID)
		Nil(g.T(), err)
		// create and store txs
		tx = g.buildEthTx()
		err = g.db.StoreEthTx(g.GetTestContext(), tx, chainID, common.BigToHash(big.NewInt(gofakeit.Int64())), uint64(blockNumber))
		Nil(g.T(), err)
		tx = g.buildEthTx()
		err = g.db.StoreEthTx(g.GetTestContext(), tx, chainID, common.BigToHash(big.NewInt(gofakeit.Int64())), uint64(blockNumber))
		Nil(g.T(), err)
	}

	// test get logs and get logs in a range
	logs, err := g.gqlClient.GetLogs(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	// there were 20 logs created (2 per loop, in a loop of 10)
	Equal(g.T(), len(logs.Response), 20)
	logsRange, err := g.gqlClient.GetLogsRange(g.GetTestContext(), int(chainID), 2, 5, 1)
	Nil(g.T(), err)
	// from 2-5, there were 8 logs created (2 per loop, in a range of 4)
	Equal(g.T(), len(logsRange.Response), 8)

	// test get receipts and get receipts in a range
	receipts, err := g.gqlClient.GetReceipts(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	// there were 20 receipts created (2 per loop, in a loop of 10)
	Equal(g.T(), len(receipts.Response), 20)
	receiptsRange, err := g.gqlClient.GetReceiptsRange(g.GetTestContext(), int(chainID), 1, 7, 1)
	Nil(g.T(), err)
	// from 1-7, there were 14 receipts created (2 per loop, in a range of 7)
	Equal(g.T(), len(receiptsRange.Response), 14)

	// test get transactions and get transactions in a range
	txs, err := g.gqlClient.GetTransactions(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	// there were 20 txs created (2 per loop, in a loop of 10)
	Equal(g.T(), len(txs.Response), 20)
	txsRange, err := g.gqlClient.GetTransactionsRange(g.GetTestContext(), int(chainID), 3, 8, 1)
	Nil(g.T(), err)
	// from 3-8, there were 12 txs created (2 per loop, in a range of 6)
	Equal(g.T(), len(txsRange.Response), 12)
}

func (g GQLSuite) TestLogDataEquality() {
	// create a log
	chainID := gofakeit.Uint32()
	log := g.buildLog(common.BigToAddress(big.NewInt(gofakeit.Int64())), uint64(gofakeit.Uint32()))

	// store it
	err := g.db.StoreLog(g.GetTestContext(), log, chainID)
	Nil(g.T(), err)

	// retrieve it
	logs, err := g.gqlClient.GetLogs(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	retrievedLog := logs.Response[0]

	// convert topics
	var topics []string
	for _, topic := range log.Topics {
		topics = append(topics, topic.String())
	}

	// check that the data is equal
	Equal(g.T(), retrievedLog.ContractAddress, log.Address.String())
	Equal(g.T(), retrievedLog.ChainID, int(chainID))
	Equal(g.T(), retrievedLog.Topics, topics)
	Equal(g.T(), retrievedLog.Data, common.BytesToHash(log.Data).String())
	Equal(g.T(), retrievedLog.BlockNumber, int(log.BlockNumber))
	Equal(g.T(), retrievedLog.TxHash, log.TxHash.String())
	Equal(g.T(), retrievedLog.TxIndex, int(log.TxIndex))
	Equal(g.T(), retrievedLog.BlockHash, log.BlockHash.String())
	Equal(g.T(), retrievedLog.Index, int(log.Index))
	Equal(g.T(), retrievedLog.Removed, log.Removed)
}

func (g GQLSuite) TestReceiptDataEquality() {
	// create a receipt
	chainID := gofakeit.Uint32()
	receipt := g.buildReceipt(common.BigToAddress(big.NewInt(gofakeit.Int64())), uint64(gofakeit.Uint32()))

	// store it
	err := g.db.StoreReceipt(g.GetTestContext(), receipt, chainID)
	Nil(g.T(), err)

	// retrieve it
	receipts, err := g.gqlClient.GetReceipts(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	retrievedReceipt := receipts.Response[0]

	// check that the data is equal
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
}

func (g GQLSuite) TestTransactionDataEquality() {
	// create a transaction
	chainID := gofakeit.Uint32()
	blockNumber := uint64(gofakeit.Uint32())
	tx := g.buildEthTx()

	// store it
	err := g.db.StoreEthTx(g.GetTestContext(), tx, chainID, common.BigToHash(big.NewInt(gofakeit.Int64())), blockNumber)
	Nil(g.T(), err)

	// retrieve it
	txs, err := g.gqlClient.GetTransactions(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	retrievedTx := txs.Response[0]

	// check that the data is equal
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
}

func (g *GQLSuite) buildLog(contractAddress common.Address, blockNumber uint64) types.Log {
	currentIndex := g.logIndex.Load()
	// increment next index
	g.logIndex.Add(1)
	log := types.Log{
		Address:     contractAddress,
		Topics:      []common.Hash{common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64()))},
		Data:        []byte(gofakeit.Sentence(10)),
		BlockNumber: blockNumber,
		TxHash:      common.BigToHash(big.NewInt(gofakeit.Int64())),
		TxIndex:     uint(gofakeit.Uint64()),
		BlockHash:   common.BigToHash(big.NewInt(gofakeit.Int64())),
		Index:       uint(currentIndex),
		Removed:     gofakeit.Bool(),
	}

	return log
}

func (g *GQLSuite) buildReceipt(contractAddress common.Address, blockNumber uint64) types.Receipt {
	receipt := types.Receipt{
		Type:              gofakeit.Uint8(),
		PostState:         []byte(gofakeit.Sentence(10)),
		Status:            gofakeit.Uint64(),
		CumulativeGasUsed: gofakeit.Uint64(),
		Bloom:             types.BytesToBloom([]byte(gofakeit.Sentence(10))),
		TxHash:            common.BigToHash(big.NewInt(gofakeit.Int64())),
		ContractAddress:   contractAddress,
		GasUsed:           gofakeit.Uint64(),
		BlockNumber:       big.NewInt(int64(blockNumber)),
		BlockHash:         common.BigToHash(big.NewInt(gofakeit.Int64())),
		TransactionIndex:  uint(gofakeit.Uint64()),
	}

	return receipt
}

func (g *GQLSuite) buildEthTx() *types.Transaction {
	ethTx := types.NewTx(&types.LegacyTx{
		Nonce:    gofakeit.Uint64(),
		GasPrice: new(big.Int).SetUint64(gofakeit.Uint64()),
		Gas:      gofakeit.Uint64(),
		To:       &common.Address{},
		Value:    new(big.Int).SetUint64(gofakeit.Uint64()),
		Data:     []byte(gofakeit.Paragraph(1, 2, 3, " ")),
	})

	return ethTx
}
