package botmd

import (
	"context"
	"fmt"

	"github.com/slack-io/slacker"
	"github.com/synapsecns/sanguine/contrib/opbot/config"
	"github.com/synapsecns/sanguine/contrib/opbot/signoz"
	"github.com/synapsecns/sanguine/contrib/screener-api/screener"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	cctpSql "github.com/synapsecns/sanguine/services/cctp-relayer/db/sql"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"golang.org/x/sync/errgroup"
)

// Bot represents the bot server.
type Bot struct {
	handler       metrics.Handler
	server        *slacker.Slacker
	cfg           config.Config
	signozClient  *signoz.Client
	signozEnabled bool
	rpcClient     omnirpcClient.RPCClient
	signer        signer.Signer
	submitter     submitter.TransactionSubmitter
	screener      screener.Screener
}

// NewBot creates a new bot server.
func NewBot(handler metrics.Handler, cfg config.Config) *Bot {
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

	bot.rpcClient = omnirpcClient.NewOmnirpcClient(cfg.OmniRPCURL, handler, omnirpcClient.WithCaptureReqRes())

	bot.addMiddleware(bot.tracingMiddleware(), bot.metricsMiddleware())
	bot.addCommands(bot.traceCommand(), bot.rfqLookupCommand(), bot.rfqRefund())

	return &bot
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
	var err error
	b.signer, err = signerConfig.SignerFromConfig(ctx, b.cfg.Signer)
	if err != nil {
		return fmt.Errorf("failed to create signer: %w", err)
	}

	dbType, err := dbcommon.DBTypeFromString(b.cfg.Database.Type)
	if err != nil {
		return fmt.Errorf("could not get db type: %w", err)
	}

	store, err := cctpSql.Connect(ctx, dbType, b.cfg.Database.DSN, b.handler)
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}

	b.submitter = submitter.NewTransactionSubmitter(b.handler, b.signer, b.rpcClient, store.SubmitterDB(), &b.cfg.SubmitterConfig)

	screener, err := screener.NewScreener(ctx, b.cfg.ScreenerConfig, b.handler)
	if err != nil {
		return fmt.Errorf("failed to create screener: %w", err)
	}
	b.screener = screener

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return b.submitter.Start(ctx)
	})

	g.Go(func() error {
		return b.server.Listen(ctx)
	})

	g.Go(func() error {
		return b.screener.Start(ctx)
	})

	// nolint: wrapcheck
	return g.Wait()
}
