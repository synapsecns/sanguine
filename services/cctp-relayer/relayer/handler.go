package relayer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"
)

// CCTPHandler is an interface for interacting with CCTP contracts.
type CCTPHandler interface {
	HandleLog(ctx context.Context, log *types.Log, chainID uint32) (processQueue bool, err error)
	FetchAndProcessSentEvent(ctx context.Context, txHash common.Hash, chainID uint32) (*relayTypes.Message, error)
	SubmitReceiveMessage(ctx context.Context, msg *relayTypes.Message) error
}
