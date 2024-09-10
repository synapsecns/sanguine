package slackertrace

import (
	"fmt"
	"github.com/slack-io/slacker"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// TracingMiddleware is a middleware that creates a new span for each command.
func TracingMiddleware(handler metrics.Handler) slacker.CommandMiddlewareHandler {
	return func(next slacker.CommandHandler) slacker.CommandHandler {
		return func(cmdCtx *slacker.CommandContext) {
			ctx, span := handler.Tracer().Start(cmdCtx.Context(), fmt.Sprintf("command.%s", cmdCtx.Definition().Command), trace.WithAttributes(
				attribute.String("user_id", cmdCtx.Event().UserID),
				attribute.String("channel_id", retrieveChannelIfExists(cmdCtx.Event())),
			))

			cmdCtx.WithContext(ctx)

			defer func() {
				metrics.EndSpan(span)
			}()

			next(cmdCtx)
		}
	}
}

const unknownChannel = "unknown"

func retrieveChannelIfExists(event *slacker.MessageEvent) string {
	if event != nil && event.Channel != nil {
		return event.Channel.ID
	}
	return unknownChannel
}
