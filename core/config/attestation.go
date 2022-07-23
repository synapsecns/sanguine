package config

import "context"

// AttestationConfig stores the config defining where attestations are stored
type AttestationConfig struct {
	// DomainID is the id of the domain
	DomainID uint32 `toml:"domain_id"`
	// CollectorAddress is the address of the attestation collector
	CollectorAddress string `toml:"address"`
}

// IsValid checks the validity of an attestation
func (a AttestationConfig) IsValid(_ context.Context) (ok bool, err error) {
	// TODO
	return true, nil
}
