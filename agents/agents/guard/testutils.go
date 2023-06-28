package guard

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/metrics"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"time"
)

// NewGuardInjectedBackend creates a new guard with an injected backend, rather than following the omnirpc URL format.
//
//nolint:cyclop
func NewGuardInjectedBackend(ctx context.Context, cfg config.AgentConfig, handler metrics.Handler, rpcURLs map[uint32]string) (_ Guard, err error) {
	guard := Guard{
		refreshInterval: time.Second * time.Duration(cfg.RefreshIntervalSeconds),
	}
	guard.domains = []domains.DomainClient{}

	guard.bondedSigner, err = signerConfig.SignerFromConfig(ctx, cfg.BondedSigner)
	if err != nil {
		return Guard{}, fmt.Errorf("error with bondedSigner, could not create guard: %w", err)
	}

	guard.unbondedSigner, err = signerConfig.SignerFromConfig(ctx, cfg.UnbondedSigner)
	if err != nil {
		return Guard{}, fmt.Errorf("error with unbondedSigner, could not create guard: %w", err)
	}

	for domainName, domain := range cfg.Domains {
		var domainClient domains.DomainClient

		domainClient, err = evm.NewEVM(ctx, domainName, domain, rpcURLs[domain.DomainID])
		if err != nil {
			return Guard{}, fmt.Errorf("failing to create evm for domain, could not create guard for: %w", err)
		}
		guard.domains = append(guard.domains, domainClient)
		if domain.DomainID == cfg.SummitDomainID {
			guard.summitDomain = domainClient
		}
	}

	guard.summitLatestStates = make(map[uint32]types.State, len(guard.domains))
	guard.originLatestStates = make(map[uint32]types.State, len(guard.domains))

	guard.handler = handler

	return guard, nil
}
