package integration_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
)

// func (i IntegrationSuite) TestGqlServer() {
// 	// fill w/ fake data
// 	// etc

// 	port := freeport.GetPort()

// 	go func() {
// 		Nil(i.T(), server.Start(uint16(port), "sqlite", i.dbPath))
// 	}()

// 	baseURL := fmt.Sprintf("http://127.0.0.1:%d", port)

// 	i.Eventually(func() bool {
// 		// TODO: use context here
// 		_, err := http.Get(fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint))
// 		return err == nil
// 	})

// 	// TODO: use conext
// 	gqlClient := client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, server.GraphqlEndpoint))

// 	res, err := gqlClient.GetLogs(i.GetTestContext())
// 	Nil(i.T(), err)

// 	// TODO: this will panic if response is nil
// 	Equal(i.T(), res.Response[0].BlockNumber, 131)
// }

func (i IntegrationSuite) TestRetrieveData() {
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
		log = i.buildLog(contractAddressA, uint64(blockNumber))
		err = i.db.StoreLog(i.GetTestContext(), log, chainID)
		Nil(i.T(), err)
		log = i.buildLog(contractAddressB, uint64(blockNumber))
		err = i.db.StoreLog(i.GetTestContext(), log, chainID)
		Nil(i.T(), err)
		// create and store receipts
		receipt = i.buildReceipt(contractAddressA, uint64(blockNumber))
		err = i.db.StoreReceipt(i.GetTestContext(), receipt, chainID)
		Nil(i.T(), err)
		receipt = i.buildReceipt(contractAddressB, uint64(blockNumber))
		err = i.db.StoreReceipt(i.GetTestContext(), receipt, chainID)
		Nil(i.T(), err)
		// create and store txs
		tx = i.buildEthTx()
		err = i.db.StoreEthTx(i.GetTestContext(), tx, chainID, uint64(blockNumber))
		Nil(i.T(), err)
		tx = i.buildEthTx()
		err = i.db.StoreEthTx(i.GetTestContext(), tx, chainID, uint64(blockNumber))
		Nil(i.T(), err)
	}

	// test get logs and get logs in a range
	logs, err := i.gqlClient.GetLogs(i.GetTestContext(), int(chainID))
	Nil(i.T(), err)
	// there were 20 logs created (2 per loop, in a loop of 10)
	Equal(i.T(), len(logs.Response), 20)
	logsRange, err := i.gqlClient.GetLogsRange(i.GetTestContext(), int(chainID), 2, 5)
	Nil(i.T(), err)
	// from 2-5, there were 8 logs created (2 per loop, in a range of 4)
	Equal(i.T(), len(logsRange.Response), 8)

	// test get receipts and get receipts in a range
	receipts, err := i.gqlClient.GetReceipts(i.GetTestContext(), int(chainID))
	Nil(i.T(), err)
	// there were 20 receipts created (2 per loop, in a loop of 10)
	Equal(i.T(), len(receipts.Response), 20)
	receiptsRange, err := i.gqlClient.GetReceiptsRange(i.GetTestContext(), int(chainID), 1, 7)
	Nil(i.T(), err)
	// from 1-7, there were 14 receipts created (2 per loop, in a range of 7)
	Equal(i.T(), len(receiptsRange.Response), 14)

	// test get transactions and get transactions in a range
	txs, err := i.gqlClient.GetTransactions(i.GetTestContext(), int(chainID))
	Nil(i.T(), err)
	// there were 20 txs created (2 per loop, in a loop of 10)
	Equal(i.T(), len(txs.Response), 20)
	txsRange, err := i.gqlClient.GetTransactionsRange(i.GetTestContext(), int(chainID), 3, 8)
	Nil(i.T(), err)
	// from 3-8, there were 12 txs created (2 per loop, in a range of 6)
	Equal(i.T(), len(txsRange.Response), 12)
}

func (i IntegrationSuite) TestLogDataEquality() {
	// create a log
	chainID := gofakeit.Uint32()
	log := i.buildLog(common.BigToAddress(big.NewInt(gofakeit.Int64())), uint64(gofakeit.Uint32()))

	// store it
	err := i.db.StoreLog(i.GetTestContext(), log, chainID)
	Nil(i.T(), err)

	// retrieve it
	logs, err := i.gqlClient.GetLogs(i.GetTestContext(), int(chainID))
	Nil(i.T(), err)
	retrievedLog := logs.Response[0]

	// convert topics
	var topics []string
	for _, topic := range log.Topics {
		topics = append(topics, topic.String())
	}

	// check that the data is equal
	Equal(i.T(), retrievedLog.ContractAddress, log.Address.String())
	Equal(i.T(), retrievedLog.ChainID, int(chainID))
	Equal(i.T(), retrievedLog.Topics, topics)
	Equal(i.T(), retrievedLog.Data, common.BytesToHash(log.Data).String())
	Equal(i.T(), retrievedLog.BlockNumber, int(log.BlockNumber))
	Equal(i.T(), retrievedLog.TxHash, log.TxHash.String())
	Equal(i.T(), retrievedLog.TxIndex, int(log.TxIndex))
	Equal(i.T(), retrievedLog.BlockHash, log.BlockHash.String())
	Equal(i.T(), retrievedLog.Index, int(log.Index))
	Equal(i.T(), retrievedLog.Removed, log.Removed)
}

func (i IntegrationSuite) TestReceiptDataEquality() {
	// create a receipt
	chainID := gofakeit.Uint32()
	receipt := i.buildReceipt(common.BigToAddress(big.NewInt(gofakeit.Int64())), uint64(gofakeit.Uint32()))

	// store it
	err := i.db.StoreReceipt(i.GetTestContext(), receipt, chainID)
	Nil(i.T(), err)

	// retrieve it
	receipts, err := i.gqlClient.GetReceipts(i.GetTestContext(), int(chainID))
	Nil(i.T(), err)
	retrievedReceipt := receipts.Response[0]

	// check that the data is equal
	Equal(i.T(), retrievedReceipt.ChainID, int(chainID))
	Equal(i.T(), retrievedReceipt.Type, int(receipt.Type))
	Equal(i.T(), retrievedReceipt.PostState, string(receipt.PostState))
	Equal(i.T(), retrievedReceipt.Status, int(receipt.Status))
	Equal(i.T(), retrievedReceipt.CumulativeGasUsed, int(receipt.CumulativeGasUsed))
	Equal(i.T(), retrievedReceipt.Bloom, common.BytesToHash(receipt.Bloom.Bytes()).String())
	Equal(i.T(), retrievedReceipt.TxHash, receipt.TxHash.String())
	Equal(i.T(), retrievedReceipt.ContractAddress, receipt.ContractAddress.String())
	Equal(i.T(), retrievedReceipt.GasUsed, int(receipt.GasUsed))
	Equal(i.T(), retrievedReceipt.BlockNumber, int(receipt.BlockNumber.Int64()))
	Equal(i.T(), retrievedReceipt.TransactionIndex, int(receipt.TransactionIndex))
}

func (i IntegrationSuite) TestTransactionDataEquality() {
	// create a transaction
	chainID := gofakeit.Uint32()
	blockNumber := uint64(gofakeit.Uint32())
	tx := i.buildEthTx()

	// store it
	err := i.db.StoreEthTx(i.GetTestContext(), tx, chainID, blockNumber)
	Nil(i.T(), err)

	// retrieve it
	txs, err := i.gqlClient.GetTransactions(i.GetTestContext(), int(chainID))
	Nil(i.T(), err)
	retrievedTx := txs.Response[0]

	// check that the data is equal
	Equal(i.T(), retrievedTx.ChainID, int(chainID))
	Equal(i.T(), retrievedTx.TxHash, tx.Hash().String())
	Equal(i.T(), retrievedTx.Protected, tx.Protected())
	Equal(i.T(), retrievedTx.Type, int(tx.Type()))
	Equal(i.T(), retrievedTx.Data, common.BytesToHash(tx.Data()).String())
	Equal(i.T(), retrievedTx.Gas, int(tx.Gas()))
	Equal(i.T(), retrievedTx.GasPrice, int(tx.GasPrice().Uint64()))
	Equal(i.T(), retrievedTx.GasTipCap, tx.GasTipCap().String())
	Equal(i.T(), retrievedTx.GasFeeCap, tx.GasFeeCap().String())
	Equal(i.T(), retrievedTx.Value, tx.Value().String())
	Equal(i.T(), retrievedTx.Nonce, int(tx.Nonce()))
	Equal(i.T(), retrievedTx.To, tx.To().String())
}

func (i *IntegrationSuite) buildLog(contractAddress common.Address, blockNumber uint64) types.Log {
	currentIndex := i.logIndex.Load()
	// increment next index
	i.logIndex.Add(1)
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

func (i *IntegrationSuite) buildReceipt(contractAddress common.Address, blockNumber uint64) types.Receipt {
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

func (i *IntegrationSuite) buildEthTx() *types.Transaction {
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
