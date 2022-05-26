package domains

import "github.com/synapsecns/sanguine/core/config"

// DomainClient represents a client used for interacting
// with contracts in a particular domain. The goal of a domain is that
// the particulars of interacting with an eth vs a solana contract are abstracted
// away and can be done through a set of common interfaces.
type DomainClient interface {
	// Name gets the name of the client. This can be used for logging.
	Name() string
	// Config gets the config that was used to create the client.
	Config() config.DomainConfig
}
