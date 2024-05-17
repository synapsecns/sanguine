package sender

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"crypto/sha256"

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
	RichGuy                     = "0xc6e2459991BfE27cca6d86722F35da23A1E4Cb97"
)

// TODO: make this an interface
type Sender struct {
	originDB          *interchaindb.InterchainDB
	originAnvilClient *anvil.Client
	originEVMClient   ethergoClient.EVM
}

func NewSender(ctx context.Context, cfg *config.SenderConfig, handler metrics.Handler) (*Sender, error) {
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
		originDB:          originInterchainDB,
		originAnvilClient: originAnvilClient,
		originEVMClient:   e,
	}, nil
}

func (s *Sender) Start(ctx context.Context, cfg *config.SenderConfig) error {
	// send verificationrequests every 5 seconds
	for {
		fmt.Println("sending entry with verification")
		_, err := s.sendWriteEntryWithVerification(ctx, cfg)
		if err != nil {
			return fmt.Errorf("could not send write entry with verification: %w", err)

		}

		// ctxWithTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
		// defer cancel()

		// fmt.Println("waiting for the tx.... 🕰️")
		// receipt, err := bind.WaitMined(ctxWithTimeout, s.originEVMClient, tx)
		// if err != nil {
		// 	return fmt.Errorf("could not get transaction receipt: %w", err)
		// }

		// fmt.Printf("Tx %s status: %d", receipt.TxHash.Hex(), receipt.Status)
		time.Sleep(5 * time.Second)
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

	err = s.originAnvilClient.ImpersonateAccount(ctx, richAddr)
	if err != nil {
		return nil, fmt.Errorf("could not impersonate account: %w", err)
	}
	defer s.originAnvilClient.StopImpersonatingAccount(ctx, richAddr)

	tx, err := s.originDB.WriteEntryWithVerification(
		&bind.TransactOpts{
			From:     richAddr,
			Value:    fee,
			NoSend:   true,
			Signer:   anvil.ImpersonatedSigner,
			GasLimit: 10_000_000,
			GasPrice: big.NewInt(100000000),
		},
		uint64(cfg.DestinationChainID),
		sha256.Sum256([]byte("fat")),
		[]common.Address{common.HexToAddress(SynapseModuleSepoliaAddress)},
	)
	if err != nil {
		return nil, fmt.Errorf("could not create entry with verification: %w", err)
	}

	err = s.originAnvilClient.SendUnsignedTransaction(ctx, richAddr, tx)
	if err != nil {
		return nil, fmt.Errorf("could not send unsigned transaction: %w", err)
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
