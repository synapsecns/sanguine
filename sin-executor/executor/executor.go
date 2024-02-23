package executor

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/sin-executor/config"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
	"github.com/synapsecns/sanguine/sin-executor/db"
	"github.com/synapsecns/sanguine/sin-executor/db/connect"
	"github.com/synapsecns/sanguine/sin-executor/listener"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

type Executor struct {
	signer          signer.Signer
	submitter       submitter.TransactionSubmitter
	client          omnirpcClient.RPCClient
	metrics         metrics.Handler
	db              db.Service
	cfg             config.Config
	chainListeners  map[int]listener.ContractListener
	clientContracts map[int]*interchainclient.InterchainClientRef
}

// NewExecutor creates a new executor.
func NewExecutor(ctx context.Context, handler metrics.Handler, cfg config.Config) (executor *Executor, err error) {
	executor = &Executor{
		cfg:     cfg,
		metrics: handler,
	}

	executor.client = omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, handler, omnirpcClient.WithCaptureReqRes())

	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return nil, fmt.Errorf("could not get db type: %w", err)
	}

	executor.db, err = connect.Connect(ctx, dbType, cfg.Database.DSN, handler)
	if err != nil {
		return nil, fmt.Errorf("could not make db: %w", err)
	}

	executor.chainListeners = make(map[int]listener.ContractListener)
	executor.clientContracts = make(map[int]*interchainclient.InterchainClientRef)

	for chainID, chainCFG := range cfg.Chains {
		synapseModule := common.HexToAddress(chainCFG.SynapseClientAddress)
		chainClient, err := executor.client.GetChainClient(ctx, chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get chain client: %w", err)
		}

		chainListener, err := listener.NewChainListener(chainClient, executor.db, synapseModule, handler)
		if err != nil {
			return nil, fmt.Errorf("could not get chain listener: %w", err)
		}
		executor.chainListeners[chainID] = chainListener

		executor.clientContracts[chainID], err = interchainclient.NewInterchainClientRef(synapseModule, chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not get synapse module ref: %w", err)
		}
	}

	executor.signer, err = signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not get signer: %w", err)
	}

	executor.submitter = submitter.NewTransactionSubmitter(handler, executor.signer, executor.client, executor.db.SubmitterDB(), &executor.cfg.SubmitterConfig)
	if err != nil {
		return nil, fmt.Errorf("could not get submitter: %w", err)
	}

	return executor, nil
}

func (e *Executor) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		// nolint: errcheck
		return e.submitter.Start(ctx)
	})

	g.Go(func() error {
		// nolint: errcheck
		return e.startChainIndexers(ctx)
	})

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(defaultDBInterval * time.Second):
				err := e.runDBSelector(ctx)
				if err != nil {
					e.metrics.ExperimentalLogger().Errorf(ctx, "could not cleanup: %w", err)
				}
			}
		}
	})

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("error starting executor: %w", err)
	}
	return nil
}

func (e *Executor) runDBSelector(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(defaultDBInterval * time.Second):
			dbItems, err := e.db.GetInterchainTXsByStatus(ctx, db.Seen, db.Ready)
			if err != nil {
				return fmt.Errorf("could not cleanup: %w", err)
			}

			for _, request := range dbItems {
				switch request.Status {
				case db.Seen:
					err := e.checkReady(ctx, request)
					if err != nil {
						e.metrics.ExperimentalLogger().Errorf(ctx, "could not sign and broadcast: %v", err)
					}
				case db.Ready:
					err := e.executeTransaction(ctx, request)
					if err != nil {
						e.metrics.ExperimentalLogger().Errorf(ctx, "could not sign and broadcast: %v", err)
					}
				case db.Executed:
					fmt.Println("executed")

				default:
					panic("unhandled default case")
				}
			}
		}
	}
}

func (e *Executor) executeTransaction(ctx context.Context, request db.TransactionSent) error {
	contract, ok := e.clientContracts[int(request.DstChainID.Int64())]
	if !ok {
		return fmt.Errorf("could not get contract for chain %d", request.SrcChainID.Int64())
	}

	nonce, err := e.submitter.SubmitTransaction(ctx, request.DstChainID, func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		transactor.GasLimit = request.Options.GasLimit.Uint64()
		transactor.Value = request.Options.GasAirdrop

		return contract.InterchainExecute(transactor, request.EncodedTX)
	})
	if err != nil {
		return fmt.Errorf("could not submit transaction: %w", err)
	}

	go func() {
		for {
			time.Sleep(time.Second * 5)
			status, _ := e.submitter.GetSubmissionStatus(ctx, request.DstChainID, nonce)

			if status.TxHash().String() != (common.Hash{}).String() {
				fmt.Println("STATUS")
				fmt.Println(status)
				fmt.Println("hash")
				fmt.Println(status.TxHash().String())
			}
		}
	}()

	err = e.db.UpdateInterchainTransactionStatus(ctx, request.TransactionID, db.Executed)
	if err != nil {
		return fmt.Errorf("could not update transaction status: %w", err)
	}

	return nil
}

func (e *Executor) checkReady(ctx context.Context, request db.TransactionSent) error {
	contract, ok := e.clientContracts[int(request.DstChainID.Int64())]
	if !ok {
		return fmt.Errorf("could not get contract for chain %d", request.DstChainID.Int64())
	}

	// TODO REMOVE ME
	// e.db.UpdateInterchainTransactionStatus(ctx, request.TransactionId, db.Ready)
	// TODO REMOVE ME

	isExecutable, err := contract.IsExecutable(&bind.CallOpts{Context: ctx}, request.EncodedTX)
	if err != nil {
		return fmt.Errorf("could not check if executable: %w", err)
	}

	if isExecutable {
		err = e.db.UpdateInterchainTransactionStatus(ctx, request.TransactionID, db.Ready)
		if err != nil {
			return fmt.Errorf("could not update transaction status: %w", err)
		}
	}
	return nil
}

// startChainIndexers starts the chain indexers for each bridge.
// these listen on the chain using the chain listeners and then handle the events.
func (e *Executor) startChainIndexers(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	// TODO: good chance we wanna prepare these chain listeners up front and then listen later.
	for chainID := range e.cfg.Chains {
		chainID := chainID // capture func literal

		g.Go(func() error {
			err := e.runChainIndexer(ctx, chainID)
			if err != nil {
				return fmt.Errorf("could not runChainIndexer chain indexer for chain %d: %w", chainID, err)
			}
			return nil
		})
	}

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("error starting chain indexers: %w", err)
	}
	return nil
}

const defaultDBInterval = 3

// runChainIndexer runs the chain indexer for a given chain.
// any events that an action exists for are indexed.
// nolint: cyclop
func (e *Executor) runChainIndexer(parentCtx context.Context, chainID int) (err error) {
	chainListener := e.chainListeners[chainID]

	parser, err := interchainclient.NewParser(chainListener.Address())
	if err != nil {
		return fmt.Errorf("could not parse: %w", err)
	}

	err = chainListener.Listen(parentCtx, func(parentCtx context.Context, log types.Log) (err error) {
		et, parsedEvent, ok := parser.ParseEvent(log)
		// handle unknown event
		if !ok {
			if len(log.Topics) != 0 {
				e.metrics.ExperimentalLogger().Warnf(parentCtx, "unknown event %s", log.Topics[0])
			}
			return nil
		}

		ctx, span := e.metrics.Tracer().Start(parentCtx, fmt.Sprintf("handleLog-%s", et), trace.WithAttributes(
			attribute.String(metrics.TxHash, log.TxHash.String()),
			attribute.Int(metrics.Origin, chainID),
			attribute.String(metrics.Contract, log.Address.String()),
			attribute.String("block_hash", log.BlockHash.String()),
			attribute.Int64("block_number", int64(log.BlockNumber)),
		))

		defer func() {
			metrics.EndSpanWithErr(span, err)
		}()

		switch event := parsedEvent.(type) {
		case *interchainclient.InterchainClientV1InterchainTransactionSent:
			encodedTX, err := e.clientContracts[chainID].EncodeTransaction(&bind.CallOpts{Context: ctx}, interchainclient.InterchainTransaction{
				SrcChainId:  big.NewInt(int64(chainID)),
				SrcSender:   event.SrcSender,
				DstChainId:  core.CopyBigInt(event.DstChainId),
				DstReceiver: event.DstReceiver,
				Nonce:       event.ClientNonce.Uint64(),
				DbNonce:     event.DbNonce,
				Options:     event.Options,
				Message:     event.Message,
			})

			decodedOptions, err := e.clientContracts[chainID].EncodeOptionsV1()

			if err != nil {
				return fmt.Errorf("could not encode transaction: %w", err)
			}

			err = e.db.StoreInterchainTransaction(ctx, &db.TransactionSent{
				TransactionID: event.TransactionId,
			})

		}

		// stop the world.
		if err != nil {
			return fmt.Errorf("could not handle event: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("listener failed: %w", err)
	}
	return nil
}
