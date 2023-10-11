package service_test

// TODO need to implement notary, guard, and executor to fully test the operation of sinner and its ability to consume
// and parse logs from those events. This will act a foundation for also testing other events emitted by the interchain
// network and sinner's ability to consume and parse those events.

// Was initially trying to just get the logs statically, but implementing a full End to End message lifecycle with the agents
// here would be good. Look into embedded agents.

// func (t *ServiceSuite) TestSinnerEndtoEnd() {
//	t.RunOnAllDBs(func(testDB db.EventDB) {
//
//		ctx := t.GetTestContext()
//
//		deployManager := testutil.NewDeployManager(t.T())
//
//		_, originContract := deployManager.GetOriginHarness(ctx, t.testBackend)
//		_, destinationContract := deployManager.GetDestinationHarness(ctx, t.destinationTestBackend)
//		wllt, err := wallet.FromRandom()
//		Nil(t.T(), err)
//		t.testBackend.FundAccount(ctx, wllt.Address(), *big.NewInt(params.Ether))
//
//		txContext := t.testBackend.GetTxContext(ctx, nil)
//		paddedRequest := big.NewInt(0)
//		//sentSink := make(chan *originharness.OriginHarnessSent)
//		//sub, err := originContract.WatchSent(&bind.WatchOpts{Context: ctx}, sentSink, [][32]byte{}, []uint32{}, []uint32{})
//		//Nil(t.T(), err)
//		chain, err := t.testBackend.ChainID(ctx)
//		Nil(t.T(), err)
//		tx, err := originContract.SendBaseMessage(txContext.TransactOpts, t.destinationChainID, [32]byte{}, 1, paddedRequest, []byte{})
//		Nil(t.T(), err)
//		err = t.scribeDB.StoreEthTx(ctx, tx, t.originChainID, common.BigToHash(big.NewInt(int64(3))), 3, uint64(1))
//		Nil(t.T(), err)
//
//		sentLog, err := t.storeTestLog(ctx, tx, t.originChainID, 3)
//		Nil(t.T(), err)
//		err = t.scribeDB.StoreLastIndexed(ctx, originContract.Address(), t.originChainID, sentLog.BlockNumber, false)
//		Nil(t.T(), err)
//
//
//		//header := agentsTypes.NewHeader(agentsTypes.MessageFlagManager, t.originChainID, 1, destinationID, 1)
//		//message := agentsTypes.NewMessage(header, nil, []byte{byte(gofakeit.Uint32())})
//		//messageBytes, err := message.ToLeaf()
//		//Nil(t.T(), err)
//
//		config := indexerConfig.ChainConfig{
//			ChainID:             t.originChainID,
//			FetchBlockIncrement: 10000,
//			MaxGoroutines:       1,
//			Contracts: []indexerConfig.ContractConfig{{
//				ContractType: "origin",
//				Address:      originContract.Address().String(),
//				StartBlock:   0,
//			}},
//		}
//
//		desConfig := indexerConfig.ChainConfig{
//			ChainID:             t.destinationChainID,
//			FetchBlockIncrement: 10000,
//			MaxGoroutines:       1,
//			Contracts: []indexerConfig.ContractConfig{{
//				ContractType: "execution_hub",
//				Address:      destinationContract.Address().String(),
//				StartBlock:   0,
//			}},
//		}
//
//		parsers := service.Parsers{
//			ChainID: t.originChainID,
//		}
//
//		originParser, err := origin.NewParser(common.HexToAddress(config.Contracts[0].Address), testDB, t.originChainID)
//		Nil(t.T(), err)
//		parsers.OriginParser = originParser
//		chainIndexer := service.NewChainIndexer(testDB, parsers, t.scribeFetcher, config)
//
//		// destination
//		desParsers := service.Parsers{
//			ChainID: t.destinationChainID,
//		}
//		destinationParser, err := destination.NewParser(common.HexToAddress(desConfig.Contracts[0].Address), testDB, t.destinationChainID)
//		Nil(t.T(), err)
//		desParsers.DestinationParser = destinationParser
//		desChainIndexer := service.NewChainIndexer(testDB, desParsers, t.scribeFetcher, desConfig)
//
//		desTxHash := common.HexToHash(big.NewInt(gofakeit.Int64()).String())
//		executeTopic := common.HexToHash("0x39c48fd1b2185b07007abc7904a8cdf782cfe449fd0e9bba1c2223a691e15f0b")
//
//		desLog := types.Log{
//			Address:     destinationContract.Address(),
//			BlockNumber: 4,
//			Topics:      []common.Hash{executeTopic},
//			Data:        []byte{},
//			TxHash:      desTxHash,
//			TxIndex:     1,
//			BlockHash:   common.HexToHash(big.NewInt(gofakeit.Int64()).String()),
//			Index:       1,
//			Removed:     false,
//		}
//		desData := agentsDestination.DestinationExecuted{
//			RemoteDomain: t.destinationChainID,
//			MessageHash:  [32]byte{},
//			Success:      true,
//			Raw:          desLog,
//		}
//		desDataBytes, err := rlp.EncodeToBytes(&desData)
//		Nil(t.T(), err)
//		desLog.Data = desDataBytes
//
//		desTx := types.NewTx(&types.LegacyTx{
//			Nonce:    gofakeit.Uint64(),
//			GasPrice: new(big.Int).SetUint64(gofakeit.Uint64()),
//			Gas:      gofakeit.Uint64(),
//			To:       &common.Address{},
//			Value:    new(big.Int).SetUint64(gofakeit.Uint64()),
//			Data:     []byte(gofakeit.Paragraph(1, 2, 3, " ")),
//		})
//		err = t.scribeDB.StoreEthTx(ctx, desTx, t.originChainID, common.BigToHash(big.NewInt(int64(4))), 5, uint64(1))
//		Nil(t.T(), err)
//		err = t.scribeDB.StoreLogs(ctx, t.destinationChainID, desLog)
//		Nil(t.T(), err)
//		err = t.scribeDB.StoreLastIndexed(ctx, originContract.Address(), t.originChainID, desLog.BlockNumber, false)
//		Nil(t.T(), err)
//		indexingCtx, cancelIndexing := context.WithCancel(ctx)
//		go func() {
//			// Call the Index function.
//
//			err = chainIndexer.Index(indexingCtx)
//			err = desChainIndexer.Index(indexingCtx)
//			Nil(t.T(), err)
//		}()
//		timeout := 1 * time.Second
//		go func() {
//			for {
//				select {
//				case <-ctx.Done():
//					return
//				case <-time.After(timeout):
//
//
//						cancelIndexing()
//
//					}
//				}
//			}
//		}()
//
//		<-indexingCtx.Done()
//
//		//
//		//tx, err = destinationContract.Execute(txContext.TransactOpts, messageBytes[:], [][32]byte{}, [][32]byte{}, paddedRequest, 10000000)
//		//Nil(t.T(), err)
//		//
//		//err = t.scribeDB.StoreEthTx(ctx, tx, t.originChainID, common.BigToHash(big.NewInt(int64(4))), 5, uint64(1))
//		//Nil(t.T(), err)
//		//destinationLog, err := t.testDBstoreTestLog(ctx, tx, t.originChainID, 3)
//		//Nil(t.T(), err)
//
//	})
//}
//
//// storeTestLogs stores the test logs in the database.
//func (t *ServiceSuite) storeTestLog(ctx context.Context, tx *types.Transaction, chainID uint32, blockNumber uint64) (*types.Log, error) {
//	t.testBackend.WaitForConfirmation(ctx, tx)
//	receipt, err := t.testBackend.TransactionReceipt(ctx, tx.Hash())
//	if err != nil {
//		return nil, fmt.Errorf("failed to get receipt for transaction %s: %w", tx.Hash().String(), err)
//	}
//	receipt.Logs[0].BlockNumber = blockNumber
//	for _, log := range receipt.Logs {
//		err = t.scribeDB.StoreLogs(ctx, chainID, *log)
//		if err != nil {
//			return nil, fmt.Errorf("error storing swap log: %w", err)
//		}
//	}
//	return receipt.Logs[len(receipt.Logs)-1], nil
//}
