package tenderly

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/debug"
	"github.com/tenderly/tenderly-cli/config"
	"github.com/tenderly/tenderly-cli/ethereum"
	tEVM "github.com/tenderly/tenderly-cli/ethereum/evm"
	tenderlyTypes "github.com/tenderly/tenderly-cli/ethereum/types"
	"github.com/tenderly/tenderly-cli/model"
	"github.com/tenderly/tenderly-cli/providers"
	"github.com/tenderly/tenderly-cli/rest"
	"github.com/tenderly/tenderly-cli/rest/call"
	"github.com/tenderly/tenderly-cli/rest/payloads"
	"github.com/tenderly/tenderly-cli/userError"
)

func init() {
	// Add the LocalTXDeploymentProvider to the list of customproviders
	providers.AllProviders = append(providers.AllProviders, LocalTXDeploymentProviderName)

	optUsed := true
	optCount := 10000
	evmVersion := "istanbul"

	compilerConfig = payloads.Config{
		OptimizationsUsed:  &optUsed,
		OptimizationsCount: &optCount,
		EvmVersion:         &evmVersion,
	}
}

// TODO this needs to be parsed.
var compilerConfig payloads.Config
var globMux sync.Mutex

// Tenderly provides utilities for using tenderly for debugging local tests using
// tenderly http://tenderly.co/
type Tenderly struct {
	*debug.ContractSource
	// rest is used for communicating with tenderly
	rest *rest.Rest
	// projectSlug is the slug of the project on tenderly
	projectSlug string
	// structMux for local contract upload
	structMux sync.RWMutex
	// ctx is the context object
	//nolint: containedctx
	ctx context.Context
}

// NewTenderly creates a new tenderly object and sets up the basic config.
func NewTenderly(ctx context.Context) (_ *Tenderly, err error) {
	t := &Tenderly{}
	t.ctx = ctx
	t.ContractSource = debug.NewContractSource()

	// initialize the global config
	globMux.Lock()
	defer globMux.Unlock()
	config.Init()
	// setup rest
	t.rest = t.setupRest()
	// make sure the user is logged in
	if !config.IsLoggedIn() {
		logger.Warn("Could not find tenderly account, ignoring.")
		return nil, fmt.Errorf("could not log into tenderly account")
	}

	// create a new deployment provider
	deploymentProvider := NewDeploymentProvider()

	// get projects for the account
	accountID := config.GetAccountId()
	projectsResponse, err := t.rest.Project.GetProjects(accountID)
	if err != nil || projectsResponse.Error != nil {
		return nil, userError.NewUserError(err, "Fetching projects for account failed.")
	}
	t.projectSlug = core.GetEnv("TENDERLY_PROJECT", "project")
	var accountProject *model.Project

	// make sure the project exists
	for _, potentialProject := range projectsResponse.Projects {
		if potentialProject.Slug == t.projectSlug {
			accountProject = potentialProject

			t.projectSlug = fmt.Sprintf("%s/%s", accountProject.OwnerInfo.Username, accountProject.Slug)
		}
	}

	config.SetProjectConfig(config.ProjectSlug, t.projectSlug)
	config.SetProjectConfig(config.AccountID, accountID)
	config.SetProjectConfig(config.Provider, deploymentProvider.GetProviderName())

	return t, nil
}

// StartListener starts listening for events on all chains and exports them to tenderly.
//
//nolint:staticcheck
func (t *Tenderly) StartListener(chn chain.Chain) error {
	watcher := chn.GetHeightWatcher()
	heightSubscription := watcher.Subscribe()

	chainID := strconv.Itoa(int(chn.GetChainID()))

	t.structMux.Lock()
	chainConfig, err := config.DefaultChainConfig.Config()
	if err != nil {
		logger.Errorf("could not get default chain: %s", err)
	}
	chainConfig.ChainID = chn.GetBigChainID()

	// setup chain config network
	client, err := debug.MakeClient(chn.RPCAddress(), chainID, t.projectSlug, chainConfig)
	if err != nil {
		return fmt.Errorf("could not make client: %w", err)
	}
	t.structMux.Unlock()

	go func(t *Tenderly) {
		defer watcher.Unsubscribe(heightSubscription)
		for {
			select {
			case <-t.ctx.Done():
				return
			case height := <-heightSubscription:
				if t == nil {
					fmt.Println("hi")
				}
				fmt.Println(t)
				// wait to add newly mined blocks so that VerifyContracts() can be called
				t.processBlock(chn, height, client, chainConfig)
			}
		}
	}(t)
	return nil
}

// getCompilerConfig gets the compiler config for a transaction, this tries to get the compiler
// for the tranaction (using the to field) and falls back to the default compiler.
func (t *Tenderly) getCompilerConfig(tx tenderlyTypes.Transaction) *payloads.Config {
	// we can't get a contract on a contract creation request
	if tx.To() == nil {
		return &compilerConfig
	}

	contractDetails, err := t.ContractSource.GetContract(tx.To().String())
	if err != nil {
		return &compilerConfig
	}

	var metadata debug.ContractMetadata
	err = json.Unmarshal([]byte(contractDetails.ContractType.ContractInfo().Info.Metadata), &metadata)
	if err != nil {
		return &compilerConfig
	}
	return &payloads.Config{
		OptimizationsUsed:  &metadata.Settings.Optimizer.Enabled,
		OptimizationsCount: &metadata.Settings.Optimizer.Runs,
		EvmVersion:         &metadata.Settings.EvmVersion,
	}
}

// processBlock uploads all txes in a block to tenderly.
//
//nolint:staticcheck
func (t *Tenderly) processBlock(chn chain.Chain, height uint64, client *ethereum.Client, chainConfig *params.ChainConfig) {
	// get the block
	blockData, err := chn.BlockByNumber(t.ctx, big.NewInt(int64(height)))
	if err != nil {
		logger.Errorf("error getting height for block: %d", height)
		return
	}

	// iterate through each tx in the block to submit
	for _, rawTx := range blockData.Transactions() {
		tx, err := client.GetTransaction(rawTx.Hash().String())
		if err != nil {
			logger.Errorf("could not get tx: %s: %s", rawTx.Hash().String(), err)
			continue
		}

		state, err := tEVM.NewProcessor(client, chainConfig).ProcessTransaction(tx.Hash().String(), true)
		if err != nil {
			logger.Errorf("could not prrocess tx %s: %s", tx.Hash().String(), err)
			continue
		}

		res, err := t.rest.Export.ExportTransaction(payloads.ExportTransactionRequest{
			NetworkData: payloads.NetworkData{
				Name:        strconv.Itoa(int(chn.GetChainID())),
				NetworkId:   strconv.Itoa(int(chn.GetChainID())),
				ChainConfig: chainConfig,
			},
			TransactionData: payloads.TransactionData{
				Transaction: tx,
				State:       state,
				Status:      state.Status,
			},
			ContractsData: payloads.UploadContractsRequest{
				Contracts: t.FetchLocalContracts(),
				Config:    t.getCompilerConfig(tx),
			},
		}, t.projectSlug)

		if res.Error != nil || err != nil {
			logger.Warn(err)
		}
	}
}

// setupRest sets up a tenderly rest cliuent.
func (t *Tenderly) setupRest() *rest.Rest {
	t.rest = rest.NewRest(
		call.NewAuthCalls(),
		call.NewUserCalls(),
		call.NewProjectCalls(),
		call.NewContractCalls(),
		call.NewExportCalls(),
		call.NewNetworkCalls(),
		call.NewActionCalls(),
	)
	return t.rest
}
