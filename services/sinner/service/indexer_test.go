package service_test

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/testutil"
	graphqlModel "github.com/synapsecns/sanguine/services/sinner/graphql/server/graph/model"

	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	indexerConfig "github.com/synapsecns/sanguine/services/sinner/config/indexer"
	"github.com/synapsecns/sanguine/services/sinner/contracts/origin"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/db/model"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/synapsecns/sanguine/services/sinner/service"

	"math/big"
	"time"
)

// TestChainIndexer tests the chain indexer.
func (t *ServiceSuite) TestChainIndexer() {
	t.RunOnAllDBs(func(testDB db.TestEventDB) {
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
			ChainID: t.originChainID,
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
		chainIndexer := service.NewChainIndexer(testDB, parsers, t.scribeFetcher, config, 1*time.Second)
		originEvent := model.OriginSent{}
		indexingCtx, cancelIndexing := context.WithCancel(ctx)
		go func() {
			err = chainIndexer.Index(indexingCtx)
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

// storeTestLogs stores the test logs in the database.
func (t *ServiceSuite) storeTestLog(ctx context.Context, tx *types.Transaction, chainID uint32, blockNumber uint64) (*types.Log, error) {
	t.testBackend.WaitForConfirmation(ctx, tx)
	receipt, err := t.testBackend.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get receipt for transaction %s: %w", tx.Hash().String(), err)
	}
	receipt.Logs[0].BlockNumber = blockNumber
	for _, log := range receipt.Logs {
		err = t.scribeDB.StoreLogs(ctx, chainID, *log)
		if err != nil {
			return nil, fmt.Errorf("error storing swap log: %w", err)
		}
	}
	return receipt.Logs[len(receipt.Logs)-1], nil
}
