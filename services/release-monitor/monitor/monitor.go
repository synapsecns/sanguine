package monitor

import (
	"context"
	"github.com/google/go-github/v37/github"
	"github.com/synapsecns/sanguine/services/release-monitor/config"
)

// Monitor monitors for new releases
type Monitor interface {
}

type monitorImpl struct {
	cfg    config.Config
	client *github.Client
}

// NewMonitor creates a new monitor
func NewMonitor(ctx context.Context, cfg config.Config) Monitor {
	return &monitorImpl{
		cfg:    cfg,
		client: newGithubClient(ctx, cfg.GithubAPIKey),
	}
}

func (m *monitorImpl) Start(ctx context.Context) {
	for _, chain := range m.cfg.Chains {
		_ = chain
	}
}
