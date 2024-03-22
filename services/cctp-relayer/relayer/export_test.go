package relayer

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

// HandleLog wraps handleLog for testing.
//
//nolint:wrapcheck
func (c *CCTPRelayer) HandleLog(parentCtx context.Context, log *types.Log, originChain uint32) (shouldProcess bool, err error) {
	return c.cctpHandler.HandleLog(parentCtx, log, originChain)
}

// FetchAndProcessSentEvent wraps FetchAndProcessSentEvent for testing.
//
//nolint:wrapcheck
func (c *CCTPRelayer) FetchAndProcessSentEvent(parentCtx context.Context, txhash common.Hash, originChain uint32) (msg *relayTypes.Message, err error) {
	return c.cctpHandler.FetchAndProcessSentEvent(parentCtx, txhash, originChain)
}

// StoreCircleRequestFulfilled wraps storeCircleRequestFulfilled for testing.
func (c *CCTPRelayer) StoreCircleRequestFulfilled(parentCtx context.Context, log *types.Log, event *cctp.SynapseCCTPEventsCircleRequestFulfilled, chainID uint32) (err error) {
	handler, ok := c.cctpHandler.(*synapseCCTPHandler)
	if !ok {
		return fmt.Errorf("handler is not a synapseCCTPHandler")
	}
	return handler.storeCircleRequestFulfilled(parentCtx, log, event, chainID)
}

// FetchAttestation wraps fetchAttestation for testing.
func (c *CCTPRelayer) FetchAttestation(parentCtx context.Context, msg *relayTypes.Message) (*relayTypes.Message, error) {
	return c.fetchAttestation(parentCtx, msg)
}

// SubmitReceiveMessage wraps SubmitReceiveMessage for testing.
//
//nolint:wrapcheck
func (c *CCTPRelayer) SubmitReceiveMessage(parentCtx context.Context, msg *relayTypes.Message) error {
	return c.cctpHandler.SubmitReceiveMessage(parentCtx, msg)
}

// SetOmnirpcClient sets the omnirpc client for testing.
func (c *CCTPRelayer) SetOmnirpcClient(client omniClient.RPCClient) {
	c.omnirpcClient = client
}
