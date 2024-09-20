package rest

import (
	"github.com/puzpuzpuz/xsync"
)

// SubscriptionParams are the parameters for a subscription.
// A nil slice means wildcard, an empty slice means no chains
type SubscriptionParams struct {
	OriginChains map[int]struct{}    // filter by origin chain
	DestChains   map[int]struct{}    // filter by destination chain
	Routes       map[[2]int]struct{} // specific routes
}

func (s *SubscriptionParams) merge(other SubscriptionParams) {
	if s.OriginChains == nil {
		s.OriginChains = other.OriginChains
	} else if other.OriginChains == nil {
		s.OriginChains = nil
	} else {
		for chain := range other.OriginChains {
			s.OriginChains[chain] = struct{}{}
		}
	}
	if s.DestChains == nil {
		s.DestChains = other.DestChains
	} else if other.DestChains == nil {
		s.DestChains = nil
	} else {
		for chain := range other.DestChains {
			s.DestChains[chain] = struct{}{}
		}
	}
	if s.Routes == nil {
		s.Routes = other.Routes
	} else if other.Routes == nil {
		s.Routes = nil
	} else {
		for route := range other.Routes {
			s.Routes[route] = struct{}{}
		}
	}
}

// PubSubManager is a manager for a pubsub system.
type PubSubManager interface {
	AddSubscription(relayerAddr string, params SubscriptionParams) error
	RemoveSubscription(relayerAddr string, params SubscriptionParams) error
	IsSubscribed(relayerAddr string, origin, dest int) bool
}

type pubSubManagerImpl struct {
	subscriptions *xsync.MapOf[string, *SubscriptionParams]
}

// NewPubSubManager creates a new pubsub manager.
func NewPubSubManager() PubSubManager {
	return &pubSubManagerImpl{
		subscriptions: xsync.NewMapOf[*SubscriptionParams](),
	}
}

func (p *pubSubManagerImpl) AddSubscription(relayerAddr string, params SubscriptionParams) error {
	sub, ok := p.subscriptions.Load(relayerAddr)
	if !ok {
		sub = &params
		p.subscriptions.Store(relayerAddr, sub)
		return nil
	}
	sub.merge(params)
	return nil
}

func (p *pubSubManagerImpl) RemoveSubscription(relayerAddr string, params SubscriptionParams) error {
	return nil
}

func (p *pubSubManagerImpl) IsSubscribed(relayerAddr string, origin, dest int) bool {
	return false
}
