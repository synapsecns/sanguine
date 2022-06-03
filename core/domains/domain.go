package domains

import (
	"context"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/types"
)

// DomainClient represents a client used for interacting
// with contracts in a particular domain. The goal of a domain is that
// the particulars of interacting with an eth vs a solana contract are abstracted
// away and can be done through a set of common interfaces.
type DomainClient interface {
	// Name gets the name of the client. This can be used for logging.
	Name() string
	// Config gets the config that was used to create the client.
	Config() config.DomainConfig
	// BlockNumber gets the latest block
	BlockNumber(ctx context.Context) (uint32, error)
	// Home retreives a handle for the home contract
	Home() HomeContract
}

// HomeContract represents the home contract on a particular chain.
type HomeContract interface {
	// FetchSortedMessages fetches all messages in order form lowest->highest in a given block range
	FetchSortedMessages(ctx context.Context, from uint32, to uint32) (messages []types.CommittedMessage, err error)
}
