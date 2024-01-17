package screener

import (
	"fmt"
	"github.com/synapsecns/sanguine/contrib/screener-api/config"
	"github.com/synapsecns/sanguine/contrib/screener-api/screener/internal"
)

type Screener interface {
}

type screenerImpl struct {
	rulesManager internal.RulesetManager
}

// NewScreener creates a new screener.
func NewScreener(cfg config.Config) (Screener, error) {
	rulesManager, err := setupScreener(cfg.Rules)
	if err != nil {
		return nil, fmt.Errorf("could not setup screener: %w", err)
	}

	return &screenerImpl{
		rulesManager: rulesManager,
	}, nil
}
