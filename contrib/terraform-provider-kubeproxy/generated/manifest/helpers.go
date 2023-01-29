package manifest

import "github.com/hashicorp/go-hclog"

// SetLogger sets the logger on the raw provider
func (s *RawProviderServer) SetLogger(logger hclog.Logger) {
	s.logger = logger
}
