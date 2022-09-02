package proxy

import (
	"context"
	"github.com/gin-gonic/gin"
)

// ReorderRPCs exports reorderRPCS for testing.
func (r *RPCProxy) ReorderRPCs(ctx context.Context, chainID int) {
	r.reorderRPCs(ctx, chainID)
}

// ServeRPCReq exports serveRPCReq for testing.
func (r *RPCProxy) ServeRPCReq(c *gin.Context, chainID int) {
	r.serveRPCReq(c, chainID)
}

// ParseRPCPayload exports parseRPCPayload for testing
func ParseRPCPayload(body []byte) (method string, err error) {
	return parseRPCPayload(body)
}
