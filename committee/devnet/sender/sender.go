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
	"github.com/synapsecns/sanguine/committee/devnet/config"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	ethergoClient "github.com/synapsecns/sanguine/ethergo/client"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"

	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

const (
	InterchainDBSepoliaAddress  = "0x8d50e833331A0D01d6F286881ce2C3A5DAD12e26"
	SynapseModuleSepoliaAddress = "0x93391bD1De68aFBAB10BB94BF3d36a4484B60eA2"
)

// TODO: make this an interface
type Sender struct {
	originDB          *interchaindb.InterchainDB
	originAnvilClient *anvil.Client

	signer  signer.Signer
	address common.Address

	handler metrics.Handler
}

func NewSender(ctx context.Context, cfg *config.SenderConfig, handler metrics.Handler) (*Sender, error) {

	// Get the origin chain's Anvil Client
	originAnvilClient, err := anvil.Dial(ctx, "http://localhost:9001/rpc/42")
	if err != nil {
		return nil, fmt.Errorf("could not dial origin client: %w", err)
	}

	// Set up OmniRPC to interact with the chains
	client := omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, handler, omnirpcClient.WithCaptureReqRes())
	chainClient, err := client.GetChainClient(ctx, cfg.OriginChainID)
	if err != nil {
		return nil, fmt.Errorf("could not get chain client: %w", err)
	}

	// Get bindings for InterchainDB on the Origin chain
	originInterchainDB, err := interchaindb.NewInterchainDB(
		common.HexToAddress(InterchainDBSepoliaAddress), chainClient,
	)
	if err != nil {
		return nil, fmt.Errorf("could not create interchaindb: %w", err)
	}

	// Get the signer to send Verification Requests
	signer, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not create signer: %w", err)
	}

	return &Sender{
		originDB:          originInterchainDB,
		originAnvilClient: originAnvilClient,
		signer:            signer,
		address:           signer.Address(),
		handler:           handler,
	}, nil
}

func (s *Sender) Start(ctx context.Context, cfg *config.SenderConfig) error {
	bal := params.Ether * 100_000_000
	s.originAnvilClient.SetBalance(ctx, s.address, uint64(bal))

	for {
		tx, err := s.sendWriteEntryWithVerification(ctx, cfg)
		if err != nil {
			return fmt.Errorf("could not send write entry with verification: %w", err)
		}

		if receipt, err := s.waitForReceipt(ctx, tx, s.handler); err != nil {
			return fmt.Errorf("could not wait for receipt: %w", err)
		} else {
			fmt.Println("Verification Request sent! 🎉")
			fmt.Println("Transaction hash:", receipt.TxHash.Hex())
		}
	}

}

func (s *Sender) sendWriteEntryWithVerification(
	ctx context.Context,
	cfg *config.SenderConfig,
) (*types.Transaction, error) {
	fee, err := s.getInterchainFee(ctx, uint64(cfg.DestinationChainID))
	if err != nil {
		return nil, fmt.Errorf("could not get interchain fee: %w", err)
	}

	txOpts, err := s.signer.GetTransactor(ctx, big.NewInt(int64(cfg.OriginChainID)))
	if err != nil {
		return nil, fmt.Errorf("could not get transactor: %w", err)
	}
	txOpts.Value = fee

	tx, err := s.originDB.WriteEntryWithVerification(
		txOpts,
		uint64(cfg.DestinationChainID),
		sha256.Sum256([]byte("fat")),
		[]common.Address{common.HexToAddress(SynapseModuleSepoliaAddress)},
	)
	if err != nil {
		return nil, fmt.Errorf("could not write entry with verification: %w", err)
	}

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

func (s *Sender) waitForReceipt(parentCtx context.Context, tx *types.Transaction, handler metrics.Handler) (*types.Receipt, error) {
	fmt.Println("waiting for the tx.... 🕰️")

	// wait a couple blocks at most
	ctxWithTimeout, cancel := context.WithTimeout(parentCtx, 30*time.Second)
	defer cancel()

	// Get the EVM client
	e, err := ethergoClient.DialBackend(parentCtx, "http://localhost:9001/rpc/42", handler)
	if err != nil {
		return nil, fmt.Errorf("could not dial ethergo backend: %w", err)
	}

	receipt, err := bind.WaitMined(ctxWithTimeout, e, tx)
	if err != nil {
		return nil, fmt.Errorf("could not get transaction receipt: %w", err)
	}

	return receipt, nil
}
