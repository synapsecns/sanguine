package relayer

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/circlecctp"
	db2 "github.com/synapsecns/sanguine/services/cctp-relayer/db"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

type circleCCTPHandler struct {
	cfg              config.Config
	db               db2.CCTPRelayerDB
	omniRPCClient    omniClient.RPCClient
	boundCircleCCTPs map[uint32]*circlecctp.MessageTransmitter
	txSubmitter      submitter.TransactionSubmitter
	handler          metrics.Handler
}

// NewCircleCCTPHandler creates a new CircleCCTPHandler.
func NewCircleCCTPHandler(ctx context.Context, cfg config.Config, db db2.CCTPRelayerDB, omniRPCClient omniClient.RPCClient, txSubmitter submitter.TransactionSubmitter, handler metrics.Handler) (CCTPHandler, error) {
	boundCircleCCTPs := make(map[uint32]*circlecctp.MessageTransmitter)
	for _, chain := range cfg.Chains {
		cl, err := omniRPCClient.GetConfirmationsClient(ctx, int(chain.ChainID), 1)
		if err != nil {
			return nil, fmt.Errorf("could not get client: %w", err)
		}
		boundCircleCCTPs[chain.ChainID], err = circlecctp.NewMessageTransmitter(chain.GetCircleCCTPAddress(), cl)
		if err != nil {
			return nil, fmt.Errorf("could not build bound contract: %w", err)
		}
	}
	return &circleCCTPHandler{
		cfg:              cfg,
		db:               db,
		omniRPCClient:    omniRPCClient,
		boundCircleCCTPs: boundCircleCCTPs,
		txSubmitter:      txSubmitter,
		handler:          handler,
	}, nil
}

func (c *circleCCTPHandler) HandleLog(ctx context.Context, log *types.Log, chainID uint32) (processQueue bool, err error) {
	return false, nil
}

func (c *circleCCTPHandler) FetchAndProcessSentEvent(ctx context.Context, txHash common.Hash, chainID uint32) (*relayTypes.Message, error) {
	return nil, nil
}

func (c *circleCCTPHandler) SubmitReceiveMessage(ctx context.Context, msg *relayTypes.Message) error {
	return nil
}
