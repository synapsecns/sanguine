package service_test

import (
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	graphqlModel "github.com/synapsecns/sanguine/services/sinner/graphql/server/graph/model"
	"github.com/synapsecns/sanguine/services/sinner/service"
	"math/big"
	"time"
)

// TODO need to implement notary, guard, and executor to fully test the operation of sinner and its ability to consume
// and parse logs from those events. This will act a foundation for also testing other events emitted by the interchain
// network and sinner's ability to consume and parse those events.

// Was initially trying to just get the logs statically, but implementing a full End to End message lifecycle with the agents
// here would be good. Look into embedded agents.

func (t *ServiceSuite) TestSinner() {
	t.RunOnAllDBs(func(testDB.EventDB) {
		ctx := t.GetTestContext()
		deployManager := testutil.NewDeployManager(t.T())

		_, originContract := deployManager.GetOriginHarness(ctx, t.testBackend)

		wllt, err := wallet.FromRandom()
		Nil(t.T(), err)
		t.testBackend.FundAccount(ctx, wllt.Address(), *big.NewInt(params.Ether))

		txContext := t.testBackend.GetTxContext(ctx, nil)
		paddedRequest := big.NewInt(0)

		tx, err := originContract.SendBaseMessage(txContext.TransactOpts, t.destinationChainID, [32]byte{}, 1, paddedRequest, []byte{})
		Nil(t.T(), err)
		err = t.scribeDB.StoreEthTx(ctx, tx, t.originChainID, common.BigToHash(big.NewInt(int64(3))), 3, uint64(1))
		Nil(t.T(), err)

		sentLog, err := t.storeTestLog(ctx, tx, t.originChainID, 3)
		Nil(t.T(), err)
		err = t.scribeDB.StoreLastIndexed(ctx, originContract.Address(), t.originChainID, sentLog.BlockNumber, false)
		Nil(t.T(), err)

		config := indexerConfig.ChainConfig{
			ChainID:             t.originChainID,
			FetchBlockIncrement: 10000,
			MaxGoroutines:       1,
			Contracts: []indexerConfig.ContractConfig{{
				ContractType: "origin",
				Address:      originContract.Address().String(),
				StartBlock:   0,
			}},
		}

		parsers := service.Parsers{
			ChainID: t.originChainID,
		}

		originParser, err := origin.NewParser(common.HexToAddress(config.Contracts[0].Address), testDB, t.originChainID)
		Nil(t.T(), err)
		parsers.OriginParser = originParser
		chainIndexer := service.NewChainIndexer(testDB, parsers, t.scribeFetcher, config)
		originEvent := model.OriginSent{}
		indexingCtx, cancelIndexing := context.WithCancel(ctx)
		go func() {
			err = chainIndexer.Index(indexingCtx)
			Nil(t.T(), err)
		}()

		timeout := 1 * time.Second
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case <-time.After(timeout):
					testDB.UNSAFE_DB().WithContext(ctx).Find(&model.OriginSent{}).First(&originEvent)
					if len(originEvent.Message) > 0 {
						// cancel if message stored
						cancelIndexing()
					}
				}
			}
		}()
		<-indexingCtx.Done()

		// Check parity of events
		Equal(t.T(), sentLog.TxIndex, originEvent.TxIndex)

		// Get and check the message status
		messageStatus, err := testDB.RetrieveMessageStatus(ctx, originEvent.MessageHash)
		Nil(t.T(), err)
		Equal(t.T(), sentLog.TxHash.String(), *messageStatus.OriginTxHash)
		Equal(t.T(), graphqlModel.MessageStateLastSeenOrigin, *messageStatus.LastSeen)
	})
}
