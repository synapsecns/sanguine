package rest

import (
	"fmt"

	"github.com/puzpuzpuz/xsync"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

// PubSubManager is a manager for a pubsub system.
type PubSubManager interface {
	AddSubscription(relayerAddr string, params model.SubscriptionParams) error
	RemoveSubscription(relayerAddr string, params model.SubscriptionParams) error
	IsSubscribed(relayerAddr string, origin, dest int) bool
}

type pubSubManagerImpl struct {
	subscriptions *xsync.MapOf[string, map[int]struct{}]
}

// NewPubSubManager creates a new pubsub manager.
func NewPubSubManager() PubSubManager {
	return &pubSubManagerImpl{
		subscriptions: xsync.NewMapOf[map[int]struct{}](),
	}
}

func (p *pubSubManagerImpl) AddSubscription(relayerAddr string, params model.SubscriptionParams) error {
	fmt.Printf("adding subscription for relayer %s with chains %v\n", relayerAddr, params.Chains)
	if params.Chains == nil {
		return fmt.Errorf("chains is nil")
	}

	sub, ok := p.subscriptions.Load(relayerAddr)
	if !ok {
		sub = make(map[int]struct{})
		for _, c := range params.Chains {
			sub[c] = struct{}{}
		}
		p.subscriptions.Store(relayerAddr, sub)
		return nil
	}
	for _, c := range params.Chains {
		sub[c] = struct{}{}
	}
	fmt.Printf("added subscription for relayer %s with chains %v\n", relayerAddr, params.Chains)
	return nil
}

func (p *pubSubManagerImpl) RemoveSubscription(relayerAddr string, params model.SubscriptionParams) error {
	if params.Chains == nil {
		return fmt.Errorf("chains is nil")
	}

	sub, ok := p.subscriptions.Load(relayerAddr)
	if !ok {
		return fmt.Errorf("relayer %s has no subscriptions", relayerAddr)
	}

	for _, c := range params.Chains {
		_, ok := sub[c]
		if !ok {
			return fmt.Errorf("relayer %s is not subscribed to chain %d", relayerAddr, c)
		}
		delete(sub, c)
	}

	return nil
}

func (p *pubSubManagerImpl) IsSubscribed(relayerAddr string, origin, dest int) bool {
	fmt.Printf("checking if relayer %s is subscribed to %d and %d\n", relayerAddr, origin, dest)
	sub, ok := p.subscriptions.Load(relayerAddr)
	if !ok {
		fmt.Printf("relayer %s has no subscriptions\n", relayerAddr)
		return false
	}
	_, ok = sub[origin]
	if !ok {
		fmt.Printf("relayer %s is not subscribed to %d\n", relayerAddr, origin)
		return false
	}
	_, ok = sub[dest]
	if !ok {
		fmt.Printf("relayer %s is not subscribed to %d\n", relayerAddr, dest)
		return false
	}
	fmt.Printf("relayer %s is subscribed to %d and %d\n", relayerAddr, origin, dest)
	return true
}
