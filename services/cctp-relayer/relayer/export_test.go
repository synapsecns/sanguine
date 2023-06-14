package relayer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

// HandleLog wraps handleLog for testing.
func (c CCTPRelayer) HandleLog(parentCtx context.Context, log *types.Log, originChain uint32) (err error) {
	return c.handleLog(parentCtx, log, originChain)
}

// HandleCircleRequestSent wraps handleCircleRequestSent for testing.
func (c CCTPRelayer) HandleCircleRequestSent(parentCtx context.Context, txhash common.Hash, originChain uint32) (err error) {
	return c.handleCircleRequestSent(parentCtx, txhash, originChain)
}

// FetchAttestation wraps fetchAttestation for testing.
func (c CCTPRelayer) FetchAttestation(parentCtx context.Context, chainID uint32, msg *relayTypes.Message) error {
	return c.fetchAttestation(parentCtx, chainID, msg)
}

// SubmitReceiveCircleToken wraps submitReceiveCircleToken for testing.
func (c CCTPRelayer) SubmitReceiveCircleToken(parentCtx context.Context, msg *relayTypes.Message) error {
	return c.submitReceiveCircleToken(parentCtx, msg)
}

// SetOmnirpcClient sets the omnirpc client for testing.
func (c *CCTPRelayer) SetOmnirpcClient(client omniClient.RPCClient) {
	c.omnirpcClient = client
}

// RecvUsdcMsg receives a usdc message from the given chain.
func (c *CCTPRelayer) GetUsdcMsgRecvChan(chainID uint32) chan *relayTypes.Message {
	return c.chainListeners[chainID].usdcMsgRecvChan
}

// SendUsdcMsg receives a usdc message from the given chain.
func (c *CCTPRelayer) GetUsdcMsgSendChan(chainID uint32) chan *relayTypes.Message {
	return c.chainListeners[chainID].usdcMsgSendChan
}
