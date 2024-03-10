package screener

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/contrib/screener-api/config"
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"github.com/synapsecns/sanguine/core/metrics"
)

type TestScreener interface {
	Screener
	SetClient(client trmlabs.Client)
}

func NewTestScreener(ctx context.Context, cfg config.Config, metricHandler metrics.Handler) (_ TestScreener, err error) {
	screener, err := NewScreener(ctx, cfg, metricHandler)
	if err != nil {
		return nil, fmt.Errorf("could not create screener: %w", err)
	}

	ts, ok := screener.(TestScreener)
	if !ok {
		return nil, fmt.Errorf("could not cast screener to test screener")
	}

	return ts, nil
}

func (s *screenerImpl) SetClient(client trmlabs.Client) {
	s.client = client
}

func SplitCSV(file string) (map[string][]Set, error) {
	return splitCSV(file)
}
