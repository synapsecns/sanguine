package tunnel

import (
	"fmt"
	"golang.ngrok.com/ngrok"
)

type config struct {
	provider     Provider
	ngrokOptions []ngrok.ConnectOption
}

func (c *config) Validate() error {
	if len(c.ngrokOptions) != 0 && c.provider != Ngrok {
		return fmt.Errorf("ngrok options are only valid for ngrok provider")
	}
	return nil
}

// Option is a tunnel option.
type Option func(*config)

// WithNgrokOptions sets the ngrok options to use when using the ngrok provider.
func WithNgrokOptions(opts ...ngrok.ConnectOption) Option {
	return func(c *config) {
		c.ngrokOptions = opts
	}
}

// WithProvider sets the provider to use.
func WithProvider(provider Provider) Option {
	return func(c *config) {
		c.provider = provider
	}
}

func makeConfig(opts []Option) (*config, error) {
	c := &config{
		// TODO: switch to moe once it's ready
		provider: Ngrok,
	}
	for _, opt := range opts {
		opt(c)
	}

	err := c.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid options: %w", err)
	}
	return c, nil
}
