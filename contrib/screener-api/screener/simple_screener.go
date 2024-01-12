package screener

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/patrickmn/go-cache"
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"time"
)

type SimpleScreener struct {
	client trmlabs.Client
	cache  *cache.Cache
}

func NewSimpleScreener(apiKey string) (*SimpleScreener, error) {
	c := cache.New(time.Minute*10, time.Second*10)
	client, err := trmlabs.NewClient(apiKey, "https://api.trmlabs.com")
	if err != nil {
		return nil, fmt.Errorf("could not create client: %w", err)
	}
	return &SimpleScreener{
		client: client,
		cache:  c,
	}, nil
}

func (s *SimpleScreener) ScreenAddress(ctx context.Context, address common.Address) (blocked bool, err error) {
	unmarshalledBlocked, found := s.cache.Get(address.String())
	if found {
		blocked, ok := unmarshalledBlocked.(bool)
		if ok {
			return blocked, nil
		}
	}

	res, err := s.client.ScreenAddress(ctx, address.String())
	if err != nil {
		return false, fmt.Errorf("could not screen address: %w", err)
	}

	for _, addressRisk := range res {
		for _, ri := range addressRisk.AddressRiskIndicators {
			if (ri.CategoryRiskScoreLevel >= 10 && ri.RiskType == "OWNERSHIP") || (ri.CategoryRiskScoreLevel >= 15 && ri.RiskType == "COUNTERPARTY") {
				s.cache.Set(address.String(), true, cache.DefaultExpiration)
				return true, nil
			}

		}
	}

	s.cache.Set(address.String(), false, cache.DefaultExpiration)

	return false, nil

}
