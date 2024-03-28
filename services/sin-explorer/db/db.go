package db

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/listener/db"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
)

// Service is the database service.
type Service interface {
	db.ChainListenerDB
	PutInterchainTransactionSent(ctx context.Context, sent *interchainclient.InterchainClientV1InterchainTransactionSent) error
	PutInterchainTransactionReceived(ctx context.Context, received *interchainclient.InterchainClientV1InterchainTransactionReceived) error
}
