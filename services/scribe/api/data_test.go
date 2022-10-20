package api_test

import (
	"github.com/synapsecns/sanguine/services/scribe/graphql"
	"github.com/synapsecns/sanguine/services/scribe/grpc/client/rest"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
)

func (g APISuite) TestRetrieveData() {
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
		err = g.db.StoreEthTx(g.GetTestContext(), tx, chainID, common.BigToHash(big.NewInt(gofakeit.Int64())), uint64(blockNumber), gofakeit.Uint64())
		Nil(g.T(), err)
		tx = g.buildEthTx()
		err = g.db.StoreEthTx(g.GetTestContext(), tx, chainID, common.BigToHash(big.NewInt(gofakeit.Int64())), uint64(blockNumber), gofakeit.Uint64())
		Nil(g.T(), err)
	}

	// test get logs and get logs in a range (Graphql)
	logs, err := g.gqlClient.GetLogs(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	// there were 20 logs created (2 per loop, in a loop of 10)
	Equal(g.T(), len(logs.Response), 20)
	logsRange, err := g.gqlClient.GetLogsRange(g.GetTestContext(), int(chainID), 2, 5, 1)
	Nil(g.T(), err)
	// from 2-5, there were 8 logs created (2 per loop, in a range of 4)
	Equal(g.T(), len(logsRange.Response), 8)

	// test get logs and get logs in a range (GRPC)
	grpcLogs, res, err := g.grpcClient.ScribeServiceApi.ScribeServiceFilterLogs(g.GetTestContext(), rest.V1FilterLogsRequest{
		Filter: &rest.V1LogFilter{
			ChainId: int64(chainID),
		},
		Page: 1,
	})

	Nil(g.T(), err)
	Equal(g.T(), len(grpcLogs.Logs), 20)
	_ = res.Body.Close()

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

func (g APISuite) TestLogDataEquality() {
	// create a log
	chainID := gofakeit.Uint32()
	log := g.buildLog(common.BigToAddress(big.NewInt(gofakeit.Int64())), uint64(gofakeit.Uint32()))

	// store it
	err := g.db.StoreLog(g.GetTestContext(), log, chainID)
	Nil(g.T(), err)

	// retrieve it using gql
	logs, err := g.gqlClient.GetLogs(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)

	parsedLog, err := graphql.ParseLog(*logs.Response[0])
	Nil(g.T(), err)

	// check equality
	Equal(g.T(), *parsedLog, log)
}

func (g APISuite) TestReceiptDataEquality() {
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
	Equal(g.T(), retrievedReceipt.Bloom, common.Bytes2Hex(receipt.Bloom.Bytes()))
	Equal(g.T(), retrievedReceipt.TxHash, receipt.TxHash.String())
	Equal(g.T(), retrievedReceipt.ContractAddress, receipt.ContractAddress.String())
	Equal(g.T(), retrievedReceipt.GasUsed, int(receipt.GasUsed))
	Equal(g.T(), retrievedReceipt.BlockNumber, int(receipt.BlockNumber.Int64()))
	Equal(g.T(), retrievedReceipt.TransactionIndex, int(receipt.TransactionIndex))
}

func (g APISuite) TestTransactionDataEquality() {
	// create a transaction
	chainID := gofakeit.Uint32()
	blockNumber := uint64(gofakeit.Uint32())
	tx := g.buildEthTx()

	// store it
	err := g.db.StoreEthTx(g.GetTestContext(), tx, chainID, common.BigToHash(big.NewInt(gofakeit.Int64())), blockNumber, gofakeit.Uint64())
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
	Equal(g.T(), retrievedTx.Data, common.Bytes2Hex(tx.Data()))
	Equal(g.T(), retrievedTx.Gas, int(tx.Gas()))
	Equal(g.T(), retrievedTx.GasPrice, int(tx.GasPrice().Uint64()))
	Equal(g.T(), retrievedTx.GasTipCap, tx.GasTipCap().String())
	Equal(g.T(), retrievedTx.GasFeeCap, tx.GasFeeCap().String())
	Equal(g.T(), retrievedTx.Value, tx.Value().String())
	Equal(g.T(), retrievedTx.Nonce, int(tx.Nonce()))
	Equal(g.T(), retrievedTx.To, tx.To().String())
}

func (g APISuite) TestBlockTimeDataEquality() {
	// create data for storing a block time
	chainID := gofakeit.Uint32()
	blockNumber := uint64(gofakeit.Uint32())
	blockTime := uint64(gofakeit.Uint32())

	// store block time
	err := g.db.StoreBlockTime(g.GetTestContext(), chainID, blockNumber, blockTime)
	Nil(g.T(), err)

	// retrieve block time
	retrievedBlockTime, err := g.gqlClient.GetBlockTime(g.GetTestContext(), int(chainID), int(blockNumber))
	Nil(g.T(), err)

	// check that the data is equal
	Equal(g.T(), *retrievedBlockTime.Response, int(blockTime))

	// check that the last stored block is correct
	lastBlock, err := g.gqlClient.GetLastStoredBlockNumber(g.GetTestContext(), int(chainID))
	Nil(g.T(), err)
	Equal(g.T(), *lastBlock.Response, int(blockNumber))

	// check that the first stored block is correct
	firstBlock, err := g.gqlClient.GetFirstStoredBlockNumber(g.GetTestContext(), int(chainID))
	Nil(g.T(), err)
	Equal(g.T(), *firstBlock.Response, int(blockNumber))
}

func (g *APISuite) buildLog(contractAddress common.Address, blockNumber uint64) types.Log {
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

func (g *APISuite) buildReceipt(contractAddress common.Address, blockNumber uint64) types.Receipt {
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

func (g *APISuite) buildEthTx() *types.Transaction {
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

func (g APISuite) TestLastContractIndexed() {
	// create data for storing a block time
	chainID := gofakeit.Uint32()
	blockNumber := uint64(gofakeit.Uint32())
	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64()))

	// store last indexed
	err := g.db.StoreLastIndexed(g.GetTestContext(), contractAddress, chainID, blockNumber)
	Nil(g.T(), err)

	// retrieve last indexed
	retrievedBlockTime, err := g.gqlClient.GetLastIndexed(g.GetTestContext(), int(chainID), contractAddress.String())
	Nil(g.T(), err)

	// check that the data is equal
	Equal(g.T(), *retrievedBlockTime.Response, int(blockNumber))
}
