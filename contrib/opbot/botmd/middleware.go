package botmd

import (
	"fmt"
	"github.com/slack-io/slacker"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (b *Bot) tracingMiddleware() slacker.CommandMiddlewareHandler {
	return func(next slacker.CommandHandler) slacker.CommandHandler {
		return func(cmdCtx *slacker.CommandContext) {
			// TODO: context is not inherited here.
			_, span := b.handler.Tracer().Start(cmdCtx.Context(), fmt.Sprintf("command.%s", cmdCtx.Definition().Command), trace.WithAttributes(
				attribute.String("user_id", cmdCtx.Event().UserID),
				attribute.String("channel_id", cmdCtx.Event().Channel.ID),
			))

			defer func() {
				metrics.EndSpan(span)
			}()

			next(cmdCtx)
		}
	}
}

func (b *Bot) metricsMiddleware() slacker.CommandMiddlewareHandler {
	return func(handler slacker.CommandHandler) slacker.CommandHandler {
		return func(cmdCtx *slacker.CommandContext) {

		}
	}
}
