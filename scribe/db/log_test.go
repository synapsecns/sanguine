package db_test

import (
	"fmt"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/scribe/db"
)

func (t *DBSuite) TestStoreRetrieveLog() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		txHashRandom := gofakeit.Int64()
		chainID := gofakeit.Uint32()
		nullHash := common.Hash{}

		// Store two logs with different tx hashes.
		txHashA := common.BigToHash(big.NewInt(txHashRandom))
		logA := MakeRandomLog(txHashA)
		err := testDB.StoreLog(t.GetTestContext(), *logA, chainID)
		Nil(t.T(), err)

		txHashB := common.BigToHash(big.NewInt(txHashRandom + 1))
		logB := MakeRandomLog(txHashB)
		err = testDB.StoreLog(t.GetTestContext(), *logB, chainID+1)
		Nil(t.T(), err)

		// Ensure the logs from the database match the ones stored.
		retrievedLogA, err := testDB.RetrieveLogByTxHash(t.GetTestContext(), txHashA)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLogA.Address(), logA.Address)
		Equal(t.T(), retrievedLogA.ChainID(), chainID)
		Equal(t.T(), retrievedLogA.PrimaryTopic(), logA.Topics[0])
		if retrievedLogA.TopicA() != nullHash {
			Equal(t.T(), retrievedLogA.TopicA(), logA.Topics[1])
		}
		if retrievedLogA.TopicB() != nullHash {
			Equal(t.T(), retrievedLogA.TopicB(), logA.Topics[2])
		}
		fmt.Println(retrievedLogA.TopicC().String())
		if retrievedLogA.TopicC() != nullHash {
			Equal(t.T(), retrievedLogA.TopicC(), logA.Topics[3])
		}
		Equal(t.T(), retrievedLogA.Data(), logA.Data)
		Equal(t.T(), retrievedLogA.BlockNumber(), logA.BlockNumber)
		Equal(t.T(), retrievedLogA.TxHash(), logA.TxHash)
		Equal(t.T(), uint(retrievedLogA.TxIndex()), logA.TxIndex)
		Equal(t.T(), retrievedLogA.BlockHash(), logA.BlockHash)
		Equal(t.T(), uint(retrievedLogA.Index()), logA.Index)
		Equal(t.T(), retrievedLogA.Removed(), logA.Removed)

		retrievedLogB, err := testDB.RetrieveLogByTxHash(t.GetTestContext(), txHashB)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLogB.Address(), logB.Address)
		Equal(t.T(), retrievedLogB.ChainID(), chainID+1)
		if retrievedLogB.TopicA() != nullHash {
			Equal(t.T(), retrievedLogB.TopicA(), logB.Topics[1])
		}
		if retrievedLogB.TopicB() != nullHash {
			Equal(t.T(), retrievedLogB.TopicB(), logB.Topics[2])
		}
		if retrievedLogB.TopicC() != nullHash {
			Equal(t.T(), retrievedLogB.TopicC(), logB.Topics[3])
		}
		Equal(t.T(), retrievedLogB.Data(), logB.Data)
		Equal(t.T(), retrievedLogB.BlockNumber(), logB.BlockNumber)
		Equal(t.T(), retrievedLogB.TxHash(), logB.TxHash)
		Equal(t.T(), uint(retrievedLogB.TxIndex()), logB.TxIndex)
		Equal(t.T(), retrievedLogB.BlockHash(), logB.BlockHash)
		Equal(t.T(), uint(retrievedLogB.Index()), logB.Index)
		Equal(t.T(), retrievedLogB.Removed(), logB.Removed)
	})
}

func MakeRandomLog(txHash common.Hash) *ethTypes.Log {
	return &ethTypes.Log{
		Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
		Topics:      []common.Hash{common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64()))},
		Data:        []byte(gofakeit.Sentence(10)),
		BlockNumber: gofakeit.Uint64(),
		TxHash:      txHash,
		TxIndex:     uint(gofakeit.Uint64()),
		BlockHash:   common.BigToHash(big.NewInt(gofakeit.Int64())),
		Index:       uint(gofakeit.Uint64()),
		Removed:     gofakeit.Bool(),
	}
}
