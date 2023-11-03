package api_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/graphql"
	"github.com/synapsecns/sanguine/services/scribe/grpc/client/rest"
	scribeTypes "github.com/synapsecns/sanguine/services/scribe/types"
	"math/big"
)

func (g APISuite) TestRetrieveData() {
	contractAddressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	contractAddressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	chainID := uint32(1)

	// create and store logs, receipts, and txs
	var log types.Log
	var receipt types.Receipt
	var tx *types.Transaction
	var err error
	for blockNumber := 0; blockNumber < 10; blockNumber++ {
		// create and store logs
		log = g.buildLog(contractAddressA, uint64(blockNumber))
		err = g.db.StoreLogs(g.GetTestContext(), chainID, log)
		Nil(g.T(), err)
		log = g.buildLog(contractAddressB, uint64(blockNumber))
		err = g.db.StoreLogs(g.GetTestContext(), chainID, log)
		Nil(g.T(), err)
		// create and store receipts
		receipt = g.buildReceipt(contractAddressA, uint64(blockNumber))
		err = g.db.StoreReceipt(g.GetTestContext(), chainID, receipt)
		Nil(g.T(), err)
		receipt = g.buildReceipt(contractAddressB, uint64(blockNumber))
		err = g.db.StoreReceipt(g.GetTestContext(), chainID, receipt)
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
	Equal(g.T(), 20, len(logs.Response))
	logsRange, err := g.gqlClient.GetLogsRange(g.GetTestContext(), int(chainID), 2, 5, 1, nil)
	Nil(g.T(), err)
	// from 2-5, there were 8 logs created (2 per loop, in a range of 4)
	Equal(g.T(), 8, len(logsRange.Response))

	// Test getting logs in a range in ascending order.
	logsRangeAsc, err := g.gqlClient.GetLogsRange(g.GetTestContext(), int(chainID), 2, 5, 1, core.PtrTo(true))
	Nil(g.T(), err)
	Equal(g.T(), 8, len(logsRangeAsc.Response))
	Equal(g.T(), 2, logsRangeAsc.Response[0].BlockNumber)

	// test get logs and get logs in a range (GRPC)
	grpcLogs, res, err := g.grpcRestClient.ScribeServiceApi.ScribeServiceFilterLogs(g.GetTestContext(), rest.V1FilterLogsRequest{
		Filter: &rest.V1LogFilter{
			ChainId: int64(chainID),
		},
		Page: 1,
	})

	Nil(g.T(), err)
	Equal(g.T(), len(grpcLogs.Logs), 20)
	_ = res.Body.Close()

	// test get receipts and get receipts in a range
	receipts, err := g.gqlClient.GetReceipts(g.GetTestContext(), int(chainID), 1, 8)
	Nil(g.T(), err)
	// there were 20 receipts created (2 per loop, in a loop of 10)
	Equal(g.T(), 2, len(receipts.Response))
	receiptsRange, err := g.gqlClient.GetReceiptsRange(g.GetTestContext(), int(chainID), 1, 7, 1)
	Nil(g.T(), err)
	// from 1-7, there were 14 receipts created (2 per loop, in a range of 7)
	Equal(g.T(), 14, len(receiptsRange.Response))

	// test get transactions and get transactions in a range
	txs, err := g.gqlClient.GetTransactions(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	// there were 20 txs created (2 per loop, in a loop of 10)
	Equal(g.T(), 20, len(txs.Response))
	txsRange, err := g.gqlClient.GetTransactionsRange(g.GetTestContext(), int(chainID), 3, 8, 1)
	Nil(g.T(), err)
	// from 3-8, there were 12 txs created (2 per loop, in a range of 6)
	Equal(g.T(), 12, len(txsRange.Response))
}

func (g APISuite) TestLogDataEquality() {
	// create a log
	chainID := gofakeit.Uint32()
	log := g.buildLog(common.BigToAddress(big.NewInt(gofakeit.Int64())), uint64(gofakeit.Uint32()))

	// store it
	err := g.db.StoreLogs(g.GetTestContext(), chainID, log)
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
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	blockNumber := uint64(gofakeit.Uint32())
	receipt := g.buildReceipt(address, blockNumber)

	// store it
	err := g.db.StoreReceipt(g.GetTestContext(), chainID, receipt)
	Nil(g.T(), err)

	// retrieve it
	receipts, err := g.gqlClient.GetReceipts(g.GetTestContext(), int(chainID), 1, int(blockNumber))
	Nil(g.T(), err)
	retrievedReceipt := receipts.Response[0]

	// check that the data is equal
	Equal(g.T(), int(chainID), retrievedReceipt.ChainID)
	Equal(g.T(), int(receipt.Type), retrievedReceipt.Type)
	Equal(g.T(), string(receipt.PostState), retrievedReceipt.PostState)
	Equal(g.T(), int(receipt.Status), retrievedReceipt.Status)
	Equal(g.T(), int(receipt.CumulativeGasUsed), retrievedReceipt.CumulativeGasUsed)
	Equal(g.T(), common.Bytes2Hex(receipt.Bloom.Bytes()), retrievedReceipt.Bloom)
	Equal(g.T(), receipt.TxHash.String(), retrievedReceipt.TxHash)
	Equal(g.T(), receipt.ContractAddress.String(), retrievedReceipt.ContractAddress)
	Equal(g.T(), int(receipt.GasUsed), retrievedReceipt.GasUsed)
	Equal(g.T(), int(receipt.BlockNumber.Int64()), retrievedReceipt.BlockNumber)
	Equal(g.T(), int(receipt.TransactionIndex), retrievedReceipt.TransactionIndex)

	// retrieve it
	// receiptWithAddress, err := g.gqlClient. .GetReceipts(g.GetTestContext(), int(chainID), 1, address)
	Nil(g.T(), err)

	retrievedReceipt = receipts.Response[0]

	// check that the data is equal
	Equal(g.T(), int(chainID), retrievedReceipt.ChainID)
	Equal(g.T(), int(receipt.Type), retrievedReceipt.Type)
	Equal(g.T(), string(receipt.PostState), retrievedReceipt.PostState)
	Equal(g.T(), int(receipt.Status), retrievedReceipt.Status)
	Equal(g.T(), int(receipt.CumulativeGasUsed), retrievedReceipt.CumulativeGasUsed)
	Equal(g.T(), common.Bytes2Hex(receipt.Bloom.Bytes()), retrievedReceipt.Bloom)
	Equal(g.T(), receipt.TxHash.String(), retrievedReceipt.TxHash)
	Equal(g.T(), receipt.ContractAddress.String(), retrievedReceipt.ContractAddress)
	Equal(g.T(), int(receipt.GasUsed), retrievedReceipt.GasUsed)
	Equal(g.T(), int(receipt.BlockNumber.Int64()), retrievedReceipt.BlockNumber)
	Equal(g.T(), int(receipt.TransactionIndex), retrievedReceipt.TransactionIndex)
}

func (g APISuite) TestTransactionDataEquality() {
	// create a transaction
	chainID := uint32(1)
	blockNumber := uint64(16131419)
	time := uint64(1670398823)

	tx := g.buildEthTx()

	// Store the empty sender.
	msgFrom, _ := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), big.NewInt(1))
	sender := msgFrom.From().String()
	err := g.db.StoreEthTx(g.GetTestContext(), tx, chainID, common.BigToHash(big.NewInt(gofakeit.Int64())), blockNumber, gofakeit.Uint64())
	Nil(g.T(), err)

	// retrieve it
	txs, err := g.gqlClient.GetTransactions(g.GetTestContext(), int(chainID), 1)
	Nil(g.T(), err)
	retrievedTx := txs.Response[0]

	// check that the data is equal
	Equal(g.T(), int(chainID), retrievedTx.ChainID)
	Equal(g.T(), tx.Hash().String(), retrievedTx.TxHash)
	Equal(g.T(), tx.Protected(), retrievedTx.Protected)
	Equal(g.T(), int(tx.Type()), retrievedTx.Type)
	Equal(g.T(), common.Bytes2Hex(tx.Data()), retrievedTx.Data)
	Equal(g.T(), int(tx.Gas()), retrievedTx.Gas)
	Equal(g.T(), int(tx.GasPrice().Uint64()), retrievedTx.GasPrice)
	Equal(g.T(), tx.GasTipCap().String(), retrievedTx.GasTipCap)
	Equal(g.T(), tx.GasFeeCap().String(), retrievedTx.GasFeeCap)
	Equal(g.T(), tx.Value().String(), retrievedTx.Value)
	Equal(g.T(), int(time), retrievedTx.Timestamp)
	Equal(g.T(), sender, retrievedTx.Sender)
	Equal(g.T(), int(tx.Nonce()), retrievedTx.Nonce)
	Equal(g.T(), tx.To().String(), retrievedTx.To)

	dbBlocktime, err := g.db.RetrieveBlockTime(g.GetTestContext(), chainID, blockNumber)
	Nil(g.T(), err)
	Equal(g.T(), time, dbBlocktime)
}

func (g APISuite) TestBlockTimeDataEquality() {
	// create data for storing a block time
	chainID := uint32(1)
	blockNumber := uint64(1000000)
	blockTime := uint64(1455404053)

	// store block time
	err := g.db.StoreBlockTime(g.GetTestContext(), chainID, blockNumber, blockTime)
	Nil(g.T(), err)

	// retrieve block time
	retrievedBlockTime, err := g.gqlClient.GetBlockTime(g.GetTestContext(), int(chainID), int(blockNumber))
	Nil(g.T(), err)

	// check that the data is equal
	Equal(g.T(), int(blockTime), *retrievedBlockTime.Response)

	// check that the last stored block is correct
	lastBlock, err := g.gqlClient.GetLastStoredBlockNumber(g.GetTestContext(), int(chainID))
	Nil(g.T(), err)
	Equal(g.T(), int(blockNumber), *lastBlock.Response)

	// check that the first stored block is correct
	firstBlock, err := g.gqlClient.GetFirstStoredBlockNumber(g.GetTestContext(), int(chainID))
	Nil(g.T(), err)
	Equal(g.T(), int(blockNumber), *firstBlock.Response)
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
	err := g.db.StoreLastIndexed(g.GetTestContext(), contractAddress, chainID, blockNumber, scribeTypes.IndexingConfirmed)
	Nil(g.T(), err)

	// retrieve last indexed
	retrievedBlockTime, err := g.gqlClient.GetLastIndexed(g.GetTestContext(), int(chainID), contractAddress.String())
	Nil(g.T(), err)

	// check that the data is equal
	Equal(g.T(), int(blockNumber), *retrievedBlockTime.Response)
}

// nolint:dupl
func (g APISuite) TestLogCount() {
	// create data for storing a block time
	chainID := gofakeit.Uint32()
	contractAddressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	contractAddressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))

	// create and store logs, receipts, and txs
	var log types.Log
	var err error
	for blockNumber := 0; blockNumber < 10; blockNumber++ {
		// create and store logs
		if blockNumber%2 == 0 {
			log = g.buildLog(contractAddressA, uint64(blockNumber))
			err = g.db.StoreLogs(g.GetTestContext(), chainID, log)
			Nil(g.T(), err)
		} else {
			log = g.buildLog(contractAddressB, uint64(blockNumber))
			err = g.db.StoreLogs(g.GetTestContext(), chainID, log)
			Nil(g.T(), err)
		}
	}

	// test get logs and get logs in a range (Graphql)
	logCountA, err := g.gqlClient.GetLogCount(g.GetTestContext(), int(chainID), contractAddressA.String())
	Nil(g.T(), err)
	Equal(g.T(), 5, *logCountA.Response)
	// store last indexed
	logCountB, err := g.gqlClient.GetLogCount(g.GetTestContext(), int(chainID), contractAddressB.String())
	Nil(g.T(), err)
	Equal(g.T(), 5, *logCountB.Response)
}

// nolint:dupl
func (g APISuite) TestReceiptCount() {
	// create data for storing a block time
	chainIDA := gofakeit.Uint32()
	chainIDB := gofakeit.Uint32()

	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64()))

	// create and store logs, receipts, and txs
	var receipt types.Receipt
	var err error
	for blockNumber := 0; blockNumber < 10; blockNumber++ {
		receipt = g.buildReceipt(contractAddress, uint64(blockNumber))
		err = g.db.StoreReceipt(g.GetTestContext(), chainIDA, receipt)
		Nil(g.T(), err)
		err = g.db.StoreReceipt(g.GetTestContext(), chainIDB, receipt)
		Nil(g.T(), err)
	}

	// test get logs and get logs in a range (Graphql)
	receiptCountA, err := g.gqlClient.GetReceiptCount(g.GetTestContext(), int(chainIDA))
	Nil(g.T(), err)
	Equal(g.T(), 10, *receiptCountA.Response)
	// store last indexed
	receiptCountB, err := g.gqlClient.GetReceiptCount(g.GetTestContext(), int(chainIDB))
	Nil(g.T(), err)
	Equal(g.T(), 10, *receiptCountB.Response)
}

func (g APISuite) TestRetrieveBlockTimesCountForChain() {
	chainIDA := gofakeit.Uint32()
	chainIDB := gofakeit.Uint32()
	blockTime := uint64(gofakeit.Uint32())
	// Store 10 blocks for both chains.
	for i := uint64(0); i < 10; i++ {
		err := g.db.StoreBlockTime(g.GetTestContext(), chainIDA, i, blockTime+i)
		Nil(g.T(), err)
		err = g.db.StoreBlockTime(g.GetTestContext(), chainIDB, i, blockTime+(i*2))
		Nil(g.T(), err)
	}

	blockTimeCountA, err := g.gqlClient.GetBlockTimeCount(g.GetTestContext(), int(chainIDA))
	Nil(g.T(), err)
	Equal(g.T(), 10, *blockTimeCountA.Response)
	blockTimeCountB, err := g.gqlClient.GetBlockTimeCount(g.GetTestContext(), int(chainIDB))
	Nil(g.T(), err)
	Equal(g.T(), 10, *blockTimeCountB.Response)
}

func (g APISuite) TestLastConfirmedBlock() {
	// create data for storing a block time
	chainID := gofakeit.Uint32()
	blockNumber := uint64(gofakeit.Uint32())

	// store last indexed
	err := g.db.StoreLastConfirmedBlock(g.GetTestContext(), chainID, blockNumber)
	Nil(g.T(), err)

	// retrieve last indexed
	retrievedBlockTime, err := g.gqlClient.GetLastConfirmedBlockNumber(g.GetTestContext(), int(chainID))
	Nil(g.T(), err)

	// check that the data is equal
	Equal(g.T(), int(blockNumber), *retrievedBlockTime.Response)
}

// nolint:dupl
func (g APISuite) TestReceiptEmptyBlock() {
	// create data for storing a block time
	chainID := gofakeit.Uint32()

	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64()))

	// create and store logs, receipts, and txs
	var receipt types.Receipt
	var err error
	for blockNumber := 0; blockNumber < 10; blockNumber++ {
		if blockNumber%2 == 0 {
			receipt = g.buildReceipt(contractAddress, uint64(blockNumber))
			err = g.db.StoreReceipt(g.GetTestContext(), chainID, receipt)
			Nil(g.T(), err)
		}
	}

	receiptFilter := db.ReceiptFilter{
		ChainID:     chainID,
		BlockNumber: 9,
	}

	receipts, err := g.db.RetrieveReceiptsWithFilter(g.GetTestContext(), receiptFilter, 1)
	Nil(g.T(), err)
	Equal(g.T(), 0, len(receipts))
}
