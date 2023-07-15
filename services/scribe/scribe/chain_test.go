package scribe_test

//
//// TestContractBackfill tests using a contractBackfiller for recording receipts and logs in a database.
// func (s *ScribeSuite) TestIndexToBlock() {
//	// Get simulated blockchain, deploy the test contract, and set up test variables.
//	simulatedChain := geth.NewEmbeddedBackendForChainID(s.GetSuiteContext(), s.T(), big.NewInt(142))
//	simulatedClient, err := backend.DialBackend(s.GetTestContext(), simulatedChain.RPCAddress(), s.metrics)
//	Nil(s.T(), err)
//
//	simulatedChain.FundAccount(s.GetTestContext(), s.wallet.Address(), *big.NewInt(params.Ether))
//	testContract, testRef := s.manager.GetTestContract(s.GetTestContext(), simulatedChain)
//	transactOpts := simulatedChain.GetTxContext(s.GetTestContext(), nil)
//
//	// Set config.
//	contractConfig := config.ContractConfig{
//		Address:    testContract.Address().String(),
//		StartBlock: 0,
//	}
//
//	simulatedChainArr := []backend.ScribeBackend{simulatedClient, simulatedClient}
//	chainConfig := config.ChainConfig{
//		ChainID:              142,
//		GetLogsBatchAmount:   1,
//		Confirmations:        0,
//		StoreConcurrency:     1,
//		GetLogsRange:         1,
//		ConcurrencyThreshold: 100,
//		Contracts:            []config.ContractConfig{contractConfig},
//	}
//
//	chainIndexer, err := scribe.NewChainIndexer(s.testDB, simulatedChainArr, chainConfig, s.metrics)
//	Nil(s.T(), err)
//
//	// Emit events for the backfiller to read.
//	tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
//	Nil(s.T(), err)
//	simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)
//
//	tx, err = testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
//	Nil(s.T(), err)
//
//	simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)
//	tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
//	Nil(s.T(), err)
//	simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)
//
//	// Emit two logs in one receipt.
//	tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
//	Nil(s.T(), err)
//
//	simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)
//
//	// Get the block that the last transaction was executed in.
//	txBlockNumber, err := testutil.GetTxBlockNumber(s.GetTestContext(), simulatedChain, tx)
//	Nil(s.T(), err)
//
//	// TODO use no-op meter
//	blockHeightMeter, err := s.metrics.Meter().NewHistogram(fmt.Sprint("scribe_block_meter", chainConfig.ChainID), "block_histogram", "a block height meter", "blocks")
//	Nil(s.T(), err)
//
//	contracts := []common.Address{common.HexToAddress(contractConfig.Address)}
//	indexer, err := indexer.NewIndexer(chainConfig, contracts, s.testDB, simulatedChainArr, s.metrics, blockHeightMeter)
//	Nil(s.T(), err)
//
//	err = chainIndexer.IndexToBlock(s.GetTestContext(), nil, uint64(0), indexer)
//	Nil(s.T(), err)
//
//	// Get all receipts.
//	receipts, err := s.testDB.RetrieveReceiptsWithFilter(s.GetTestContext(), db.ReceiptFilter{}, 1)
//	Nil(s.T(), err)
//
//	// Check to see if 3 receipts were collected.
//	Equal(s.T(), 4, len(receipts))
//
//	// Get all logs.
//	logs, err := s.testDB.RetrieveLogsWithFilter(s.GetTestContext(), db.LogFilter{}, 1)
//	Nil(s.T(), err)
//
//	// Check to see if 4 logs were collected.
//	Equal(s.T(), 5, len(logs))
//
//	// Check to see if the last receipt has two logs.
//	Equal(s.T(), 2, len(receipts[0].Logs))
//
//	// Ensure last indexed block is correct.
//	lastIndexed, err := s.testDB.RetrieveLastIndexed(s.GetTestContext(), testContract.Address(), uint32(testContract.ChainID().Uint64()))
//	Nil(s.T(), err)
//	Equal(s.T(), txBlockNumber, lastIndexed)
//}
//
//// TestChainIndxer tests that the ChainIndxer can backfill events from a chain.
// func (s *ScribeSuite) TestChainIndxer() {
//	const desiredBlockHeight = 20
//
//	var addresses []common.Address
//	var err error
//	var wg sync.WaitGroup
//
//	wg.Add(2)
//	testBackend := geth.NewEmbeddedBackend(s.GetTestContext(), s.T())
//
//	go func() {
//		defer wg.Done()
//		addresses, _, err = testutil.PopulateWithLogs(s.GetTestContext(), testBackend, desiredBlockHeight, s.T(), []*testutil.DeployManager{s.manager})
//		Nil(s.T(), err)
//	}()
//
//	var host string
//	go func() {
//		defer wg.Done()
//		host = testutil.StartOmnirpcServer(s.GetTestContext(), s.T(), testBackend)
//	}()
//
//	wg.Wait()
//
//	scribeBackend, err := backend.DialBackend(s.GetTestContext(), host, s.metrics)
//	Nil(s.T(), err)
//	simulatedChainArr := []backend.ScribeBackend{scribeBackend, scribeBackend}
//
//	chainID, err := scribeBackend.ChainID(s.GetTestContext())
//	Nil(s.T(), err)
//
//	var contractConfigs []config.ContractConfig
//	for _, address := range addresses {
//		contractConfig := config.ContractConfig{
//			Address: address.String(),
//		}
//		contractConfigs = append(contractConfigs, contractConfig)
//	}
//
//	chainConfig := config.ChainConfig{
//		ChainID:            uint32(chainID.Uint64()),
//		Confirmations:      1,
//		GetLogsBatchAmount: 1,
//		StoreConcurrency:   1,
//		GetLogsRange:       1,
//		Contracts:          contractConfigs,
//	}
//	killableContext, _ := context.WithTimeout(s.GetTestContext(), 20*time.Second)
//	chainIndexer, err := scribe.NewChainIndexer(s.testDB, simulatedChainArr, chainConfig, s.metrics)
//	Nil(s.T(), err)
//	_ = chainIndexer.Index(killableContext, nil)
//
//	logs, err := s.testDB.RetrieveLogsWithFilter(s.GetTestContext(), db.LogFilter{}, 1)
//	Nil(s.T(), err)
//	Equal(s.T(), 5, len(logs))
//	receipts, err := s.testDB.RetrieveReceiptsWithFilter(s.GetTestContext(), db.ReceiptFilter{}, 1)
//	Nil(s.T(), err)
//	Equal(s.T(), 5, len(receipts))
//
//}
