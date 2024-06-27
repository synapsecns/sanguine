package botmd

import (
	"context"
	"github.com/slack-io/slacker"
	"github.com/synapsecns/sanguine/contrib/opbot/config"
	"github.com/synapsecns/sanguine/contrib/opbot/signoz"
	"github.com/synapsecns/sanguine/core/metrics"
)

// Bot represents the bot server.
type Bot struct {
	handler       metrics.Handler
	server        *slacker.Slacker
	cfg           config.Config
	signozClient  *signoz.Client
	signozEnabled bool
}

// NewBot creates a new bot server.
func NewBot(handler metrics.Handler, cfg config.Config) Bot {
	server := slacker.NewClient(cfg.SlackBotToken, cfg.SlackAppToken)
	bot := Bot{
		handler: handler,
		cfg:     cfg,
		server:  server,
	}

	// you should be able to run opbot even without signoz.
	if cfg.SignozPassword != "" && cfg.SignozEmail != "" && cfg.SignozBaseURL != "" {
		bot.signozClient = signoz.NewClientFromUser(handler, cfg.SignozBaseURL, cfg.SignozEmail, cfg.SignozPassword)
		bot.signozEnabled = true
	}

	bot.addMiddleware(bot.tracingMiddleware(), bot.metricsMiddleware())
	bot.addCommands(bot.traceCommand(), bot.rfqLookupCommand())

	return bot
}

func (b *Bot) addMiddleware(middlewares ...slacker.CommandMiddlewareHandler) {
	for _, middleware := range middlewares {
		b.server.AddCommandMiddleware(middleware)
	}
}

func (b *Bot) addCommands(commands ...*slacker.CommandDefinition) {
	for _, command := range commands {
		b.server.AddCommand(command)
	}
}

// Start starts the bot server.
// nolint: wrapcheck
func (b *Bot) Start(ctx context.Context) error {
	// nolint: wrapcheck
	return b.server.Listen(ctx)
}
