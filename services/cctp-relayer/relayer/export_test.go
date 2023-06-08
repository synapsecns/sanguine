package relayer

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

// HandleSendRequest wraps handlessendrequest for testing.
func (c CCTPRelayer) HandleSendRequest(parentCtx context.Context, txhash common.Hash, originChain uint32) (err error) {
	return c.handleSendRequest(parentCtx, txhash, originChain)
}

func (c *CCTPRelayer) SetOmnirpcClient(client omniClient.RPCClient) {
	c.omnirpcClient = client
}
