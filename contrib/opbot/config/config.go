// Package config provides a simple way to read and write configuration files.
package config

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
	SignozBaseURL string   `yaml:"signoz_base_url"`
	RelayerURLS   []string `yaml:"rfq_relayer_urls"`
}
