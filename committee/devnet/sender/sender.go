package sender

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/committee/contracts/interchaindb"
	"github.com/synapsecns/sanguine/committee/devnet/config"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

const (
	InterchainDBSepoliaAddress  = "0x8d50e833331A0D01d6F286881ce2C3A5DAD12e26"
	SynapseModuleSepoliaAddress = "0x93391bD1De68aFBAB10BB94BF3d36a4484B60eA2"
)

// TODO: make this an interface
type Sender struct {
	originDB     *interchaindb.InterchainDB
	transactOpts *bind.TransactOpts
	client       omnirpcClient.RPCClient
}

func NewSender(ctx context.Context, cfg config.SenderConfig, handler metrics.Handler) (*Sender, error) {

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

	return &Sender{
		client:   client,
		originDB: originInterchainDB,
	}, nil
}

func (s *Sender) Start(ctx context.Context, cfg config.SenderConfig) error {
	// send verificationrequests every 5 seconds
	for {
		err := s.sendWriteEntryWithVerification(ctx, cfg)
		if err != nil {
			return fmt.Errorf("could not send write entry with verification: %w", err)

		}
		// wait for it
		// todo; better way to do it
		time.Sleep(5 * time.Second)
	}
}

func (s *Sender) sendWriteEntryWithVerification(ctx context.Context, cfg config.SenderConfig) error {
	if fee, err := s.getInterchainFee(ctx, uint64(cfg.DestinationChainID)); err != nil {
		s.transactOpts.Value = core.CopyBigInt(fee)
	} else {
		return fmt.Errorf("could not get interchain fee: %w", err)
	}

	//test a single verification
	_, err := s.originDB.WriteEntryWithVerification(
		s.transactOpts,
		uint64(cfg.DestinationChainID),
		sha256.Sum256([]byte("fat")),
		[]common.Address{common.HexToAddress(SynapseModuleSepoliaAddress)},
	)
	if err != nil {
		return fmt.Errorf("could not write entry with verification: %w", err)
	}

	return nil

}

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
