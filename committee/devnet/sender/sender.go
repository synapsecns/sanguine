package sender

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/committee/contracts/interchaindb"
	"github.com/synapsecns/sanguine/committee/db"
	"github.com/synapsecns/sanguine/committee/devnet/config"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	ethergoClient "github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"

	"github.com/synapsecns/sanguine/committee/db/connect"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
)

const (
	InterchainDBSepoliaAddress  = "0x8d50e833331A0D01d6F286881ce2C3A5DAD12e26"
	SynapseModuleSepoliaAddress = "0x93391bD1De68aFBAB10BB94BF3d36a4484B60eA2"
)

// TODO: make this an interface
type Sender struct {
	originDB          *interchaindb.InterchainDB
	originAnvilClient *anvil.Client
	originEVMClient   ethergoClient.EVM

	submitter submitter.TransactionSubmitter
	address   common.Address
	db        db.Service
}

func NewSender(ctx context.Context, cfg *config.SenderConfig, handler metrics.Handler) (*Sender, error) {
	s := &Sender{}

	originAnvilClient, err := anvil.Dial(ctx, "http://localhost:9001/rpc/42")
	if err != nil {
		return nil, fmt.Errorf("could not dial origin client: %w", err)
	}

	// set up omnirpc
	client := omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, handler, omnirpcClient.WithCaptureReqRes())
	chainClient, err := client.GetChainClient(ctx, 42)
	if err != nil {
		return nil, fmt.Errorf("could not get chain client: %w", err)
	}

	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return nil, fmt.Errorf("could not get db type: %w", err)
	}

	s.db, err = connect.Connect(ctx, dbType, cfg.Database.DSN, handler)
	if err != nil {
		return nil, fmt.Errorf("could not make db: %w", err)
	}

	// get the origin DB
	originInterchainDB, err := interchaindb.NewInterchainDB(
		common.HexToAddress(InterchainDBSepoliaAddress), chainClient,
	)
	if err != nil {
		return nil, fmt.Errorf("could not create interchaindb: %w", err)
	}

	signer, err := signerConfig.SignerFromConfig(ctx, cfg.SignerConfig)
	if err != nil {
		return nil, fmt.Errorf("could not create signer: %w", err)
	}
	submitter := submitter.NewTransactionSubmitter(
		handler,
		signer,
		client,
		s.db.SubmitterDB(),
		&cfg.SubmitterConfig,
	)

	e, err := ethergoClient.DialBackend(ctx, "http://localhost:9001/rpc/42", handler)
	if err != nil {
		return nil, fmt.Errorf("could not dial ethergo backend: %w", err)
	}
	return &Sender{
		originDB:          originInterchainDB,
		originAnvilClient: originAnvilClient,
		originEVMClient:   e,
		submitter:         submitter,
		address:           signer.Address(),
	}, nil
}

func (s *Sender) Start(ctx context.Context, cfg *config.SenderConfig) error {
	bal := params.Ether * 100_000_000
	s.originAnvilClient.SetBalance(
		ctx,
		s.address,
		uint64(bal),
	)
	go func() {
		err := s.submitter.Start(ctx)
		if err != nil {
			fmt.Printf("submitter error: %v", err)
		}
	}()

	tx, err := s.sendWriteEntryWithVerification(ctx, cfg)
	if err != nil {
		return fmt.Errorf("could not send write entry with verification: %w", err)
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	fmt.Println("waiting for the tx.... 🕰️")
	receipt, err := bind.WaitMined(ctxWithTimeout, s.originEVMClient, tx)
	if err != nil {
		return fmt.Errorf("could not get transaction receipt: %w", err)
	}

	fmt.Printf("Tx %s status: %d", receipt.TxHash.Hex(), receipt.Status)

	return nil
}

func (s *Sender) sendWriteEntryWithVerification(
	ctx context.Context,
	cfg *config.SenderConfig,
) (*types.Transaction, error) {
	fee, err := s.getInterchainFee(ctx, uint64(cfg.DestinationChainID))
	if err != nil {
		return nil, fmt.Errorf("could not get interchain fee: %w", err)
	}

	_, err = s.submitter.SubmitTransaction(
		ctx,
		big.NewInt(int64(cfg.OriginChainID)),
		func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
			transactor.Value = fee
			return s.originDB.WriteEntryWithVerification(
				transactor,
				uint64(cfg.DestinationChainID),
				sha256.Sum256([]byte("fat")),
				[]common.Address{common.HexToAddress(SynapseModuleSepoliaAddress)},
			)
		},
	)
	if err != nil {
		return nil, fmt.Errorf("could not submit transaction: %w", err)
	}

	fmt.Printf("sent transaction %s\n", tx.Hash().Hex())
	return tx, nil
}

// gets the interchain fee for the destination chain
func (s *Sender) getInterchainFee(ctx context.Context, destChainID uint64) (*big.Int, error) {
	fee, err := s.originDB.GetInterchainFee(
		&bind.CallOpts{Context: ctx},
		destChainID,
		[]common.Address{common.HexToAddress(SynapseModuleSepoliaAddress)},
	)
	if err != nil {
		return nil, fmt.Errorf("could not get interchain fee: %w", err)
	}
	return fee, nil
}
