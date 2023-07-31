package moe

type config struct {
	serverURL  string
	remotePort int
}

// Option is a tunnel option.
type Option func(*config)

// WithServerURL sets the server URL to use.
func (c *config) WithServerURL(serverURL string) Option {
	return func(c *config) {
		c.serverURL = serverURL
	}
}

// WithRemotePort sets the remote port to use.
func (c *config) WithRemotePort(remotePort int) Option {
	return func(c *config) {
		c.remotePort = remotePort
	}
}

// moeServer is the remote moe server.
const moeServer = "remote.moe"
const remotePort = 80

func makeConfig(opts []Option) *config {
	c := &config{
		serverURL:  moeServer,
		remotePort: remotePort,
	}
	for _, opt := range opts {
		opt(c)
	}

	return c
}
