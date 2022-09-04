package integration_test

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/phayes/freeport"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/scribe/client"
	"github.com/synapsecns/sanguine/services/scribe/server"
)

func (i IntegrationSuite) TestGqlServer() {
	// fill w/ fake data
	// etc

	port := freeport.GetPort()

	go func() {
		Nil(i.T(), server.Start(uint16(port), "sqlite", i.dbPath))
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", port)

	i.Eventually(func() bool {
		// TODO: use context here
		_, err := http.Get(fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint))
		return err == nil
	})

	// TODO: use conext
	gqlClient := client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, server.GraphqlEndpoint))

	res, err := gqlClient.GetLogs(i.GetTestContext())
	Nil(i.T(), err)

	// TODO: this will panic if response is nil
	Equal(i.T(), res.Response[0].BlockNumber, 131)
}

func (i IntegrationSuite) TestRetrieveRangeData() {
	txHashA := common.BigToHash(big.NewInt(gofakeit.Int64()))
	txHashB := common.BigToHash(big.NewInt(gofakeit.Int64()))
	contractAddressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	contractAddressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	chainID := uint32(gofakeit.Uint32())

	// create and store logs, receipts, and txs
	for blockNumber := 0; blockNumber < 10; blockNumber++ {
		i.buildAndStoreLog(txHashA, contractAddressA, uint64(blockNumber), chainID)
		i.buildAndStoreLog(txHashB, contractAddressB, uint64(blockNumber), chainID)
		i.buildAndStoreReceipt(txHashA, contractAddressA, uint64(blockNumber), chainID)
		i.buildAndStoreReceipt(txHashB, contractAddressB, uint64(blockNumber), chainID)
		i.buildAndStoreEthTx(txHashA, contractAddressA, uint64(blockNumber), chainID)
		i.buildAndStoreEthTx(txHashB, contractAddressB, uint64(blockNumber), chainID)
	}

}

func (i *IntegrationSuite) buildAndStoreLog(txHash common.Hash, contractAddress common.Address, blockNumber uint64, chainID uint32) {
	currentIndex := i.logIndex.Load()
	// increment next index
	i.logIndex.Add(1)
	log := types.Log{
		Address:     contractAddress,
		Topics:      []common.Hash{common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64()))},
		Data:        []byte(gofakeit.Sentence(10)),
		BlockNumber: blockNumber,
		TxHash:      txHash,
		TxIndex:     uint(gofakeit.Uint64()),
		BlockHash:   common.BigToHash(big.NewInt(gofakeit.Int64())),
		Index:       uint(currentIndex),
		Removed:     gofakeit.Bool(),
	}
	err := i.db.StoreLog(i.GetTestContext(), log, chainID)
	Nil(i.T(), err)
}

func (i *IntegrationSuite) buildAndStoreReceipt(txHash common.Hash, contractAddress common.Address, blockNumber uint64, chainID uint32) {
	receipt := types.Receipt{
		Type:              gofakeit.Uint8(),
		PostState:         []byte(gofakeit.Sentence(10)),
		Status:            gofakeit.Uint64(),
		CumulativeGasUsed: uint64(gofakeit.Uint64()),
		Bloom:             types.BytesToBloom([]byte(gofakeit.Sentence(10))),
		TxHash:            txHash,
		ContractAddress:   contractAddress,
		GasUsed:           uint64(gofakeit.Uint64()),
		BlockNumber:       big.NewInt(int64(blockNumber)),
		BlockHash:         common.BigToHash(big.NewInt(gofakeit.Int64())),
		TransactionIndex:  uint(gofakeit.Uint64()),
	}
	err := i.db.StoreReceipt(i.GetTestContext(), receipt, chainID)
	Nil(i.T(), err)
}

func (i *IntegrationSuite) buildAndStoreEthTx(txHash common.Hash, contractAddress common.Address, blockNumber uint64, chainID uint32) {
	ethTx := types.NewTx(&types.LegacyTx{
		Nonce:    gofakeit.Uint64(),
		GasPrice: new(big.Int).SetUint64(gofakeit.Uint64()),
		Gas:      gofakeit.Uint64(),
		To:       &contractAddress,
		Value:    new(big.Int).SetUint64(gofakeit.Uint64()),
		Data:     []byte(gofakeit.Paragraph(1, 2, 3, " ")),
	})

	err := i.db.StoreEthTx(i.GetTestContext(), ethTx, chainID, blockNumber)
	Nil(i.T(), err)
}
