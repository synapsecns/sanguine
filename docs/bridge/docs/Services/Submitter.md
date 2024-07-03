# Submitter

This section is still in progress, please see [here](https://pkg.go.dev/github.com/synapsecns/sanguine/ethergo/submitter#section-readme) for details.

# Ethergo Submitter

## Overview

The Ethergo Submitter module is designed to submit transactions to an EVM-based blockchain. It handles gas bumping and confirmation checking to ensure that transactions are eventually confirmed. This module is essential because the EVM does not specify transaction submission or consensus, and rate limits can affect transaction submission.

## Key Features

The module is the `SubmitTransaction` method, which returns a nonce and ensures that the transaction will eventually be confirmed. The nonce may then be used in the `GetSubmissionStatus` method to check the status: `Pending`, `Stored`, `Submitted`, `FailedSubmit`, `ReplacedOrConfirmed`, `Replaced`, `Confirmed`. [More about the `Status` enum](#status-enum)

- **Gas Bumping**: Automatically adjusts the gas price to ensure timely transaction confirmation.
- **Confirmation Checking**: Continuously checks the status of submitted transactions to confirm their inclusion in the blockchain.
- **Reaper Functionality**: Flushes old entries in the database that have reached a terminal state.

### Reaper

The Submitter also has "reaper" functionality, which flushes old entries in the database that have reached a terminal state (`Replaced`, `ReplacedOrConfirmed`, `Confirmed`). By default, entries are flushed after a week, but this functionality is configurable by the `MaxRecordAge` config value.

### Submitter Config

Config contains configuration for the Submitter. It can be loaded from a YAML file.
Chain-specific configuration items can be provided via the `Chains` map, which overrides the global config
for each chain. If a chain-specific item is not provided, the global config is used.

#### Example config

```yaml
submitter_config:
  chains:
    1:
      # MaxBatchSize is the maximum number of transactions to send in a batch.
      # If this is zero, the default will be used.
      # This field is ignored if batching is disabled.
      max_batch_size: 50
      # Batch is whether or not to batch transactions at the rpc level.
      skip_batching: false
      # MaxGasPrice is the maximum gas price to use for transactions.
      max_gas_price: 200000000000 # 200 Gwei
      # MinGasPrice is the gas price that will be used if 0 is returned
      # from the gas price oracle.
      min_gas_price: 1000000000 # 1 Gwei
      # BumpIntervalSeconds is the number of seconds to
      # wait before bumping a transaction.
      bump_interval_seconds: 120
      # GasBumpPercentages is the percentage to bump the gas price by.
      # This is applied to the greatrer of the chainprice or the last price.
      gas_bump_percentage: 10
      # GasEstimate is the gas estimate to use for transactions if
      # dynamic gas estimation is enabled.
      # This is only used as a default if the estimate fails.
      gas_estimate: 1000000
      # DynamicGasEstimate is whether or not to use dynamic gas estimation.
      dynamic_gas_estimate: true
      # SupportsEIP1559 is whether or not this chain supports EIP1559.
      supports_eip_1559: true
    43114:
      max_gas_price: 100000000000 # 100 Gwei
    10:
      max_gas_price: 90000000000 # 90 Gwei
      min_gas_price: 100000000 # 0.1 Gwei
  # ReaperInterval is the interval at which scan for transactions to flush
  reaper_interval: 604800000000000 # int64(7 * 24 * time.Hour)
  # MaxRecordAge is the maximum age of a record before it is flushed
  max_record_age: 86400000000000 # int64(1 * 24 * time.Hour)
```

Please see [here](https://pkg.go.dev/github.com/synapsecns/sanguine/ethergo@v0.9.0/submitter/config) for details on the configuration.

## Overview

`SubmitTransaction` abstracts many of the complexities of on-chain transaction submission such as nonce management and gas bumping. In addition, sent transactions are stored in the database for easy indexing of older transactions.

## Learning how to use Submitter via Go Tests and Anvil via Docker

We will make a small test suite in Go which will send transactions to a local Anvil chain via Docker.
This should leave you understanding both how to use Submitter and familiarzing yourself with some of the tools and idioms commonly seen in Ethergo.

### 1: Make sure you have Docker

First, check if Docker is installed:

```
docker --version
```

If Docker is not installed, you can install it using the following steps:

For **Mac**:

```
brew install --cask docker
```

For **Ubuntu**:

```
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
```

For **Windows**, download the installer from the official Docker website and follow the installation instructions.

### 2: Initialize project

```
mkdir submitter_example
cd submitter_example
go mod init submitter_example
```

This will create a new Go module for our test.

### 3: create test files

```
touch suite_test.go submitter_test.go
```

`suite_test.go` will contain the setup and hold the actual suite, while `submitter_test.go` will contain the test.

### 4. Write code

#### 4.1 Submitter Suite

We will first setup the Submitter Suite, which handles creating the Anvil backends, database, metrics, and more. This can be seen as the actual blockchain test environment.

<details>
<summary><code>submitter_example/suite_test.go</code></summary>

```go title = "suite_test.go"
package submitter_example_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ipfs/go-log"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/config"
	cmn "github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/processlog"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/examples/contracttests"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/submitter/db/txdb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var buildInfo = config.NewBuildInfo(
	config.DefaultVersion,
	config.DefaultCommit,
	"submitter",
	config.DefaultDate,
)
var tenEth = new(big.Int).Mul(
	new(big.Int).SetUint64(uint64(params.Ether)),
	big.NewInt(10),
)

type SubmitterSuite struct {
	// TestSuite is the base test suite.
	*testsuite.TestSuite
	// TestBackends are the backends to use for the test.
	testBackends []backends.SimulatedTestBackend
	// metrics is the metrics handler to use for the test.
	metrics metrics.Handler
	// deployer is the deployer to use for the counter contracts.
	deployer *manager.DeployerManager
	// store is the store to use for the test
	store db.Service
	// signer is the signer to use for the Submitter.
	signer signer.Signer
	// localAccount is the account used to construct the signer.
	localAccount *keystore.Key
}

// SetupSuite sets up 3 backends and metrics.
func (s *SubmitterSuite) SetupSuite() {
	s.TestSuite.SetupSuite()

	// Create three simulated backends with chain IDs 1, 3, and 4.
	testChainIDs := []uint64{1, 3, 4}
	s.testBackends = make(
		[]backends.SimulatedTestBackend,
		len(testChainIDs),
	)
	s.deployer = manager.NewDeployerManager(
		s.T(),
		contracttests.NewCounterDeployer,
	)

	var wg sync.WaitGroup
	// Wait for all the backends to be created,
	// add 1 to the wait group for the metrics
	wg.Add(len(testChainIDs) + 1)

	logDir := filet.TmpDir(s.T(), "")

	// create the jaeger instance
	go func() {
		defer wg.Done()
		var err error
		// don't use metrics on ci for integration tests
		isCI := core.GetEnvBool("CI", false)
		useMetrics := !isCI
		metricsHandler := metrics.Null

		if useMetrics {
			localmetrics.SetupTestJaeger(s.GetSuiteContext(), s.T())
			metricsHandler = metrics.Jaeger
		}
		s.metrics, err = metrics.NewByType(
			s.GetSuiteContext(),
			buildInfo,
			metricsHandler,
		)
		s.Require().NoError(err)
	}()

	// create the backends
	for i, chainID := range testChainIDs {
		go func(index int, chainID uint64) {
			defer wg.Done()
			options := anvil.NewAnvilOptionBuilder()
			options.SetChainID(chainID)
			// make sure all the docker containers log to the same directory
			options.SetProcessLogOptions(
				processlog.WithLogFileName(
					fmt.Sprintf("chain-%d.log", chainID),
				),
				processlog.WithLogDir(logDir),
			)

			s.testBackends[index] = anvil.NewAnvilBackend(
				s.GetSuiteContext(),
				s.T(),
				options,
			)
			s.deployer.Get(
				s.GetSuiteContext(),
				s.testBackends[index],
				contracttests.CounterType,
			)
		}(i, chainID)
	}
	wg.Wait()
}

// SetupTest sets up the signer and
// funds the account with 10 eth on each backend.
func (s *SubmitterSuite) SetupTest() {
	s.TestSuite.SetupTest()
	s.localAccount = mocks.MockAccount(s.T())
	// create the local signer
	s.signer = localsigner.NewSigner(s.localAccount.PrivateKey)
	var wg sync.WaitGroup
	wg.Add(len(s.testBackends) + 1)

	// setup the db
	go func() {
		defer wg.Done()
		var err error
		s.store, err = NewSqliteStore(
			s.GetTestContext(),
			filet.TmpDir(s.T(), ""),
			s.metrics,
		)
		s.Require().NoError(err)
	}()

	// fund the account on each chain
	for i := range s.testBackends {
		go func(index int) {
			defer wg.Done()

			s.testBackends[index].FundAccount(
				s.GetTestContext(),
				s.signer.Address(),
				*tenEth,
			)
		}(i)
	}
	wg.Wait()
}

// NewGasSuite creates a new chain testing suite.
func NewSubmitterSuite(tb testing.TB) *SubmitterSuite {
	tb.Helper()
	return &SubmitterSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestSubmitterSuite(t *testing.T) {
	suite.Run(t, NewSubmitterSuite(t))
}

// Store wraps the store. Since tx submitter is a library and not a standalone
// service, we simulate db creation here and then proceed as we would with
// any other db test.
type Store struct {
	*txdb.Store
}

func (s SubmitterSuite) GetClient(
	ctx context.Context,
	chainID *big.Int,
) (client.EVM, error) {
	for _, backend := range s.testBackends {
		if backend.GetBigChainID().Cmp(chainID) == 0 {
			//nolint: wrapcheck
			return client.DialBackend(ctx, backend.RPCAddress(), s.metrics)
		}
	}
	return nil, fmt.Errorf("could not find client for chain id %v", chainID)
}

// NewSqliteStore creates a new sqlite data store.
func NewSqliteStore(
	parentCtx context.Context,
	dbPath string,
	handler metrics.Handler,
) (_ *Store, err error) {
	logger := log.Logger("sqlite-store")

	logger.Debugf("creating sqlite store at %s", dbPath)

	ctx, span := handler.Tracer().Start(parentCtx, "start-sqlite")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// create the directory to the store if it doesn't exist
	err = os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("could not create sqlite store")
	}

	logger.Warnf("submitter database is at %s/synapse.db", dbPath)

	namingStrategy := schema.NamingStrategy{
		TablePrefix: fmt.Sprintf(
			"test%d_%d_",
			gofakeit.Int64(),
			time.Now().Unix(),
		),
	}

	gdb, err := gorm.Open(
		sqlite.Open(fmt.Sprintf("%s/%s", dbPath, "synapse.db")),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   cmn.GetGormLogger(logger),
			FullSaveAssociations:                     true,
			SkipDefaultTransaction:                   true,
			NamingStrategy:                           namingStrategy,
		})
	if err != nil {
		return nil, fmt.Errorf("could not connect to db %s: %w", dbPath, err)
	}

	handler.AddGormCallbacks(gdb)

	err = gdb.WithContext(ctx).AutoMigrate(txdb.GetAllModels()...)
	if err != nil {
		return nil, fmt.Errorf("could not migrate models: %w", err)
	}
	return &Store{txdb.NewTXStore(gdb, handler)}, nil
}
```

</details>

#### 4.2 Submitter test

Now for the actual test. Copy and paste the following in `submitter_test.go`.
The comments walk you through the test to help you understand how to use
Submitter as a library.

<details>
<summary><code>submitter_example/submitter_test.go</code></summary>

```go title = "submitter_test.go
package submitter_example_test

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/ethergo/examples/contracttests"
	"github.com/synapsecns/sanguine/ethergo/examples/contracttests/counter"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/ethergo/submitter/config"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
)

// This test below is a good example of how the Submitter works.
// It shows off the database/queue aspect of the Submitter, and
// how it can be used to queue up transactions to be submitted later.
func (s SubmitterSuite) TestSubmitTransaction() {
	// Get the counter contract binding.
	// If you do not want to use manager, you can also use the go-ethereum
	// bindings directly, like
	// `contract, err := someContact.NewSomeContract(contractAddress, backend)`
	_, cntr := manager.GetContract[*counter.CounterRef](
		s.GetTestContext(), s.T(),
		s.deployer, s.testBackends[0], contracttests.CounterType,
	)

	// Grab the origin chain ID.
	chainID := s.testBackends[0].GetBigChainID()

	// Get the current count.
	startingCount, err := cntr.GetCount(&bind.CallOpts{
		Context: s.GetTestContext(),
	})
	s.Require().NoError(err)

	// Get the legacy and dynamic chain IDs and set the gas prices.
	legacyChainID := s.testBackends[0].GetBigChainID()
	dynamicChainID := s.testBackends[1].GetBigChainID()

	maxGasPrice := big.NewInt(1000 * params.GWei)
	minGasPrice := big.NewInt(1 * params.GWei)

	cfg := &config.Config{
		Chains: map[int]config.ChainConfig{
			int(legacyChainID.Int64()): {
				MinGasPrice:     minGasPrice,
				MaxGasPrice:     maxGasPrice,
				SupportsEIP1559: false,
			},
			int(dynamicChainID.Int64()): {
				MinGasPrice:     minGasPrice,
				MaxGasPrice:     maxGasPrice,
				SupportsEIP1559: true,
			},
		},
	}

	// Create a new transaction submitter.
	ts := submitter.NewTransactionSubmitter(s.metrics, s.signer, s, s.store, cfg)

	// Submit a transaction to increment the counter.
	// Notice that a nonce is returned from the SubmitTransaction method, which
	// can be used to check the status of the transaction.
	nonce, err := ts.SubmitTransaction(
		s.GetTestContext(),
		chainID,
		func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
			tx, err = cntr.IncrementCounter(transactor)
			if err != nil {
				return nil, fmt.Errorf("failed to increment counter: %w", err)
			}

			return tx, nil
		})
	s.Require().NoError(err)

	// Check the status of the transaction.
	submissionStatus, err := ts.GetSubmissionStatus(
		s.GetTestContext(),
		chainID,
		nonce,
	)
	s.Require().NoError(err)
	// The transaction should be in the pending state,
	// because we never called `Start` method.
	s.Require().Equal(submissionStatus.State(), submitter.Pending)

	// Get the current count.
	currentCount, err := cntr.GetCount(&bind.CallOpts{
		Context: s.GetTestContext(),
	})
	s.Require().NoError(err)

	// The original transaction should not be submitted yet because remember,
	// we never called `Start` method.
	s.Equal(startingCount.Uint64(), currentCount.Uint64())

	// There is, however one transaction queued in the Submitter database,
	// ready to be fired off.
	txs, err := s.store.GetTXS(
		s.GetTestContext(),
		s.signer.Address(),
		chainID,
		db.Stored,
	)
	s.Require().NoError(err)
	s.Require().NotNil(txs[0])
	s.Require().Equal(len(txs), 1)

	go func() {
		// Now we'll start a new submitter with a new signer and submit the tx.
		err = ts.Start(s.GetTestContext())
		s.Require().NoError(err)
	}()

	// Eventually, two things should happen:
	// 1. The transaction should be submitted and confirmed.
	// 2. The counter should be incremented.
	s.Eventually(func() bool {
		currentCount, err = cntr.GetCount(&bind.CallOpts{
			Context: s.GetTestContext(),
		})
		s.Require().NoError(err)

		submissionStatus, err = ts.GetSubmissionStatus(
			s.GetTestContext(),
			chainID,
			nonce,
		)
		s.Require().NoError(err)

		return currentCount.Uint64() > startingCount.Uint64() &&
			submissionStatus.State() == submitter.Confirmed
	})

}
```

</details>

### 5: Run tests

In the console run `go test ./...`. You should see the output

```zsh
$ go test ./...
ok      submitter_example       95.372s
```

Congrats! You should now have learned how to use Submitter in your own projects while
also learning about Ethergo, our comprehensive embedded test suite.

## Nonce Management, Database, Internals

### Nonce Management and Multichain

Submitter was designed with multiple chains in mind by keeping track of a thread-safe `map[chainid]nonce`. When we build the transaction opts, we lock on the chainid until we finish firing off the transaction.
We also keep a `map[txHash]txStatus` with the same thread-safe mechanism.
This allows us to concurrently fire off transactions on different chains while ensuring our nonces are correct.
The [Queue](https://github.com/synapsecns/sanguine/blob/ethergo/v0.9.0/ethergo/submitter/chain_queue.go) has a selector loop running at all times which calls the `processQueue` method, concurrently processing and storing confirmed txs, or using the [chain queue](https://github.com/synapsecns/sanguine/blob/ethergo/v0.9.0/ethergo/submitter/chain_queue.go) to fire off and store pending txs on chain.

### DB Configurability

#### Customizing DB Behavior

The Chain Queue db interface, [Service](https://github.com/synapsecns/sanguine/blob/ethergo/v0.9.0/ethergo/submitter/db/service.go), allows a user to customize their Transaction DB behavior. The concrete implementation is in [store.go](https://github.com/synapsecns/sanguine/blob/ethergo/v0.9.0/ethergo/submitter/db/txdb/store.go).

#### Transaction DB Schema

The schema for a transaction to be stored in the Transaction DB is:

```go title="submitter/db/txdb/model.go"
// ETHTX contains a raw evm transaction that is unsigned.
type ETHTX struct {
  ID uint64 `gorm:"column:id;primaryKey;autoIncrement:true"`
  // UUID is a unique ID for this transaction that will persist across retries.
  UUID string `gorm:"column:uuid;index"`
  // CreatedAt is the time the transaction was created
  CreatedAt time.Time
  // TXHash is the hash of the transaction
  TXHash string `gorm:"column:tx_hash;uniqueIndex;size:256"`
  // From is the sender of the transaction
  From string `gorm:"column:from;index"`
  // ChainID is the chain id the transaction hash will be sent on
  ChainID uint64 `gorm:"column:chain_id;index"`
  // Nonce is the nonce of the raw evm tx
  Nonce uint64 `gorm:"column:nonce;index"`
  // RawTx is the raw serialized transaction
  RawTx []byte `gorm:"column:raw_tx"`
  // Status is the status of the transaction
  Status db.Status `gorm:"column:status;index"`
}
```

Using [GORM.db](https://pkg.go.dev/gorm.io/gorm), you can use whatever database you'd like, MySQL, Sqlite, etc.

#### MySQL Example

```go
gdb, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{
    Logger:               common_base.GetGormLogger(logger),
    FullSaveAssociations: true,
    NamingStrategy:       NamingStrategy,
    NowFunc:              time.Now,
})
```

### Status Enum

In the DB, `Status`, is an enum, represented as a uint8. It is important to know what number indicates which status.

```go title="submitter/db/service.go"
type Status uint8

// Important: do not modify the order of these constants.
// if one needs to be removed, replace it with a no-op status.
// additionally, due to the GetMaxNoncestatus function, statuses are currently assumed to be in order.
// if you need to modify this functionality, please update that function. to reflect that the highest status
// is no longer the expected end status.
const (
	// Pending is the status of a tx that has not been processed yet.
	Pending Status = iota + 1 // Pending
	// Stored is the status of a tx that has been stored.
	Stored // Stored
	// Submitted is the status of a tx that has been submitted.
	Submitted // Submitted
	// FailedSubmit is the status of a tx that has failed to submit.
	FailedSubmit // Failed
	// ReplacedOrConfirmed is the status of a tx that has been replaced by a new tx or confirmed. The actual status will be set later.
	ReplacedOrConfirmed // ReplacedOrConfirmed
	// Replaced is the status of a tx that has been replaced by a new tx.
	Replaced // Replaced
	// Confirmed is the status of a tx that has been confirmed.
	Confirmed // Confirmed
)
```

## Observability

Submitter exposes metrics for Prometheus. The metrics are:

- `num_pending_txs`: The number of pending transactions
- `current_nonce`: The current nonce
- `oldest_pending_tx`: The age of the oldest pending transaction
- `confirmed_queue`: The number of confirmed transactions
- `gas_balance`: The current gas balance

The metrics can be used in a dashboard [here](https://raw.githubusercontent.com/synapsecns/sanguine/master/ethergo/dashboard.json). It looks like this:

![Submitter Dashboard](img/submitter/metrics.png)
