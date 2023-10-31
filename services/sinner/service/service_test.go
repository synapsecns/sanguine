package service_test

import (
	"context"
	"github.com/synapsecns/sanguine/services/sinner/db"

	. "github.com/stretchr/testify/assert"
	graphqlModel "github.com/synapsecns/sanguine/services/sinner/graphql/server/graph/model"

	indexerConfig "github.com/synapsecns/sanguine/services/sinner/config/indexer"
	"github.com/synapsecns/sanguine/services/sinner/db/model"

	"github.com/synapsecns/sanguine/services/sinner/service"

	"time"
)

// TestSinner tests Sinner.
func (t *ServiceSuite) TestSinner() {
	t.RunOnAllDBs(func(testDB db.TestEventDB) {
		// Store test logs and txs
		err := t.scribeDB.StoreLogs(t.GetTestContext(), t.originChainID, t.originTestLog)
		Nil(t.T(), err)
		err = t.scribeDB.StoreEthTx(t.GetTestContext(), t.originTestTx, t.originChainID, t.originTestLog.BlockHash, t.originTestLog.BlockNumber, uint64(t.originTestLog.TxIndex))
		Nil(t.T(), err)
		err = t.scribeDB.StoreLastIndexed(t.GetTestContext(), t.originTestLog.Address, t.originChainID, 625782, false)
		Nil(t.T(), err)

		err = t.scribeDB.StoreLogs(t.GetTestContext(), t.destinationChainID, t.destinationTestLog)
		Nil(t.T(), err)
		err = t.scribeDB.StoreEthTx(t.GetTestContext(), t.destinationTestTx, t.destinationChainID, t.destinationTestLog.BlockHash, t.destinationTestLog.BlockNumber, uint64(t.destinationTestLog.TxIndex))
		Nil(t.T(), err)
		err = t.scribeDB.StoreLastIndexed(t.GetTestContext(), t.destinationTestLog.Address, t.destinationChainID, 1975780, false)
		Nil(t.T(), err)

		originConfig := indexerConfig.ChainConfig{
			ChainID: t.originChainID,
			Contracts: []indexerConfig.ContractConfig{{
				ContractType: "origin",
				Address:      t.originTestLog.Address.String(),
				StartBlock:   625780,
			}},
		}

		destinationConfig := indexerConfig.ChainConfig{
			ChainID: t.destinationChainID,
			Contracts: []indexerConfig.ContractConfig{{
				ContractType: "execution_hub",
				Address:      t.destinationTestLog.Address.String(),
				StartBlock:   1975778,
			}},
		}

		config := indexerConfig.Config{
			ScribeURL:          t.scribeFetcherPath,
			DBPath:             t.scribeDBPath,
			DefaultRefreshRate: 1,
			DBType:             "sqlite",
			SkipMigrations:     true,
			Chains:             []indexerConfig.ChainConfig{originConfig, destinationConfig},
		}

		sinner, err := service.NewSinner(testDB, config, t.metrics)
		Nil(t.T(), err)

		originEvent := model.OriginSent{}
		destinationEvent := model.Executed{}

		indexingCtx, cancelIndexing := context.WithCancel(t.GetTestContext())
		go func() {
			err = sinner.Index(indexingCtx)
			Nil(t.T(), err)
		}()

		timeout := 2 * time.Second
		go func() {
			for {
				select {
				case <-t.GetTestContext().Done():
					return
				case <-time.After(timeout):
					// check db
					testDB.UNSAFE_DB().WithContext(t.GetTestContext()).Find(&model.OriginSent{}).First(&originEvent)
					testDB.UNSAFE_DB().WithContext(t.GetTestContext()).Find(&model.Executed{}).First(&destinationEvent)
					if len(originEvent.MessageHash) > 0 && len(destinationEvent.MessageHash) > 0 {
						// cancel if message stored
						cancelIndexing()
					}
				}
			}
		}()
		<-indexingCtx.Done()

		// Check parity of events
		Equal(t.T(), t.originTestLog.TxHash.String(), originEvent.TxHash)
		Equal(t.T(), t.destinationTestLog.TxHash.String(), destinationEvent.TxHash)

		// Get and check the message status
		messageStatus, err := testDB.RetrieveMessageStatus(t.GetTestContext(), originEvent.MessageHash)
		Nil(t.T(), err)
		Equal(t.T(), t.originTestLog.TxHash.String(), *messageStatus.OriginTxHash)
		Equal(t.T(), graphqlModel.MessageStateLastSeenDestination, *messageStatus.LastSeen)
	})
}
