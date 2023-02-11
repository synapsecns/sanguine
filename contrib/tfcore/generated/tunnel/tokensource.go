package tunnel

import "golang.org/x/oauth2"

// SetTokenSource sets the token source for the tunnel
func (m *TunnelManager) SetTokenSource(source oauth2.TokenSource) {
	m.ts = source
}
