// Package config provides a simple way to read and write configuration files.
package config

import (
	"errors"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	"strings"
)

// Config represents the configuration of the application.
type Config struct {
	// SlackBotToken is the token of the slack bot.
	SlackBotToken string `yaml:"slack_bot_token"`
	// SlackAppToken is the token of the slack app.
	SlackAppToken string `yaml:"slack_app_token"`
	// Email is the email address of the user.
	// this is used to authenticate with signoz and should be a READ ONLY KEY.
	SignozEmail string `yaml:"signoz_email"`
	// Password is the password of the user.
	// this is used to authenticate with signoz and should be a READ ONLY KEY.
	// inject only with init container!
	SignozPassword string `yaml:"signoz_password"`
	// SignozBaseURL is the base url of the signoz instance.
	SignozBaseURL string `yaml:"signoz_base_url"`
	// RelayerURLS is the list of RFQ relayer URLs.
	RelayerURLS []string `yaml:"rfq_relayer_urls"`
	// RFQApiURL is the URL of the RFQ API.
	RFQApiURL string `yaml:"rfq_api_url"`
	// OmniRPCURL is the URL of the Omni RPC.
	OmniRPCURL string `yaml:"omnirpc_url"`
	// Signer is the signer config.
	Signer config.SignerConfig `yaml:"signer"`
	// SubmitterConfig is the submitter config.
	SubmitterConfig submitterConfig.Config `yaml:"submitter_config"`
	// Database is the database config.
	Database DatabaseConfig `yaml:"database"`
}

// DatabaseConfig represents the configuration for the database.
type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"` // Data Source Name
}

func (c *Config) Validate() error {
	if c.SlackBotToken == "" {
		return errors.New("slack_bot_token is required")
	}
	if !strings.HasPrefix(c.SlackBotToken, "xoxb-") {
		return errors.New("slack_bot_token must start with xoxb-")
	}
	if c.SlackAppToken == "" {
		return errors.New("slack_app_token is required")
	}
	if !strings.HasPrefix(c.SlackAppToken, "xapp-") {
		return errors.New("slack_app_token must start with xapp-")
	}
	return nil
}
