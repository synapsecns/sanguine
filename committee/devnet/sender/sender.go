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
	"github.com/synapsecns/sanguine/committee/contracts/interchaindb"
	"github.com/synapsecns/sanguine/committee/devnet/config"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	ethergoClient "github.com/synapsecns/sanguine/ethergo/client"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

const (
	InterchainDBSepoliaAddress  = "0x8d50e833331A0D01d6F286881ce2C3A5DAD12e26"
	SynapseModuleSepoliaAddress = "0x93391bD1De68aFBAB10BB94BF3d36a4484B60eA2"
	RichGuy                     = "0xaBD7F11c1BF37c966D59924553F0437B33C791dc"
)

// TODO: make this an interface
type Sender struct {
	originDB     *interchaindb.InterchainDB
	originClient *anvil.Client
	client       omnirpcClient.RPCClient
	evmClient    ethergoClient.EVM
}

func NewSender(ctx context.Context, cfg *config.SenderConfig, handler metrics.Handler) (*Sender, error) {
	originClient, err := anvil.Dial(ctx, "http://localhost:9001/rpc/42")
	if err != nil {
		return nil, fmt.Errorf("could not dial origin client: %w", err)
	}

	// set up omnirpc
	client := omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, handler, omnirpcClient.WithCaptureReqRes())
	chainClient, err := client.GetChainClient(ctx, cfg.OriginChainID)
	if err != nil {
		return nil, fmt.Errorf("could not get chain client: %w", err)
	}

	// get the origin DB
	originInterchainDB, err := interchaindb.NewInterchainDB(
		common.HexToAddress(InterchainDBSepoliaAddress), chainClient,
	)
	if err != nil {
		return nil, fmt.Errorf("could not create interchaindb: %w", err)
	}

	e, err := ethergoClient.DialBackend(ctx, "http://localhost:9001/rpc/42", handler)
	if err != nil {
		return nil, fmt.Errorf("could not dial ethergo backend: %w", err)
	}
	return &Sender{
		client:       client,
		originDB:     originInterchainDB,
		originClient: originClient,
		evmClient:    e,
	}, nil
}

func (s *Sender) Start(ctx context.Context, cfg *config.SenderConfig) error {
	// send verificationrequests every 5 seconds
	for {
		tx, err := s.sendWriteEntryWithVerification(ctx, cfg)
		if err != nil {
			return fmt.Errorf("could not send write entry with verification: %w", err)

		}
		fmt.Println("Verification Request Tx sent", tx.Hash())
		// wait for it
		// todo; better way to do it
		time.Sleep(5 * time.Second)
		receipt, err := s.evmClient.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return fmt.Errorf("could not get transaction receipt: %w", err)
		}
		fmt.Println("Verification Request Tx receipt", receipt)
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

	richAddr := common.HexToAddress(RichGuy)

	err = s.originClient.ImpersonateAccount(ctx, richAddr)
	if err != nil {
		return nil, fmt.Errorf("could not impersonate account: %w", err)
	}
	defer s.originClient.StopImpersonatingAccount(ctx, richAddr)

	hash := sha256.Sum256([]byte("fat"))
	modules := []common.Address{common.HexToAddress(SynapseModuleSepoliaAddress)}
	tx, err := s.originDB.WriteEntryWithVerification(
		&bind.TransactOpts{
			From:   richAddr,
			Value:  fee,
			NoSend: true,
			Signer: anvil.ImpersonatedSigner,
		},
		uint64(cfg.DestinationChainID),
		hash,
		modules,
	)
	if err != nil {
		return nil, fmt.Errorf("could not create entry with verification: %w", err)
	}

	err = s.originClient.SendUnsignedTransaction(ctx, richAddr, tx)
	if err != nil {
		return nil, fmt.Errorf("could not write entry with verification: %w", err)
	}

	return tx, nil
}

// gets the interchain fee for the destination chain
func (s *Sender) getInterchainFee(ctx context.Context, destChainID uint64) (*big.Int, error) {
	addy := common.HexToAddress(SynapseModuleSepoliaAddress)
	fee, err := s.originDB.GetInterchainFee(
		&bind.CallOpts{Context: ctx},
		destChainID,
		[]common.Address{addy},
	)
	if err != nil {
		return nil, fmt.Errorf("could not get interchain fee: %w", err)
	}
	return fee, nil
}
