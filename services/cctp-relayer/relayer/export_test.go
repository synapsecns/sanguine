package relayer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

// HandleLog wraps handleLog for testing.
func (c *CCTPRelayer) HandleLog(parentCtx context.Context, log *types.Log, originChain uint32) (err error) {
	return c.handleLog(parentCtx, log, originChain)
}

// HandleCircleRequestSent wraps fetchAndStoreCircleRequestSent for testing.
func (c *CCTPRelayer) HandleCircleRequestSent(parentCtx context.Context, txhash common.Hash, originChain uint32) (msg *relayTypes.Message, err error) {
	return c.fetchAndStoreCircleRequestSent(parentCtx, txhash, originChain)
}

// StoreCircleRequestFulfilled wraps storeCircleRequestFulfilled for testing.
func (c *CCTPRelayer) StoreCircleRequestFulfilled(parentCtx context.Context, log *types.Log, event *cctp.SynapseCCTPEventsCircleRequestFulfilled, chainID uint32) (err error) {
	return c.storeCircleRequestFulfilled(parentCtx, log, event, chainID)
}

// FetchAttestation wraps fetchAttestation for testing.
func (c *CCTPRelayer) FetchAttestation(parentCtx context.Context, msg *relayTypes.Message) (*relayTypes.Message, error) {
	return c.fetchAttestation(parentCtx, msg)
}

// SubmitReceiveCircleToken wraps submitReceiveCircleToken for testing.
func (c *CCTPRelayer) SubmitReceiveCircleToken(parentCtx context.Context, msg *relayTypes.Message) error {
	return c.submitReceiveCircleToken(parentCtx, msg)
}

// SetOmnirpcClient sets the omnirpc client for testing.
func (c *CCTPRelayer) SetOmnirpcClient(client omniClient.RPCClient) {
	c.omnirpcClient = client
}
