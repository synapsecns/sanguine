package localmetrics

type config struct {
	enablePyroscope       bool
	enablePyroscopeJaeger bool
	keepContainers        bool
	// requiresNetwork is set to true if the test requires a network to be created.
	// this is not exposed and is only used internally to make it clearer pyroscope-jaeger
	// is the only thing that requires a network
	requiresNetwork bool
}

// Option is a configuration option for the local server.
type Option func(*config)

// WithPyroscopeJaeger enables pyroscope jaeger.
func WithPyroscopeJaeger(enabled bool) Option {
	return func(c *config) {
		c.enablePyroscopeJaeger = enabled
		c.requiresNetwork = enabled
	}
}

// WithKeepContainers keeps containers around after the test.
func WithKeepContainers(enabled bool) Option {
	return func(c *config) {
		c.keepContainers = enabled
	}
}

// WithPyroscopeEnabled enables pyroscope.
func WithPyroscopeEnabled(enabled bool) Option {
	return func(c *config) {
		c.enablePyroscope = enabled
	}
}

// makeConfig creates a config from the options.
func makeConfig(options []Option) *config {
	c := &config{}
	// set defaults here
	c.enablePyroscopeJaeger = false
	c.enablePyroscope = true
	for _, opt := range options {
		opt(c)
	}
	return c
}
