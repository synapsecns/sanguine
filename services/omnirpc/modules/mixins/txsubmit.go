package mixins

import (
	"context"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
	"github.com/synapsecns/sanguine/ethergo/util"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// TxSubmitMixin is a mixin for tracking submitted transactions.
// it can be used to index additional data in otel regarding tx submission status.
func TxSubmitMixin(parentCtx context.Context, handler metrics.Handler, r rpc.Request) {
	if client.RPCMethod(r.Method) != client.SendRawTransactionMethod {
		return
	}

	ctx, span := handler.Tracer().Start(parentCtx, "txsubmit", trace.WithAttributes(attribute.Int("txsubmit", r.ID)))

	var err error
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	tx, err := ReqToTX(r)
	if err != nil {
		handler.ExperimentalLogger().Warnf(ctx, "could not convert request to transaction: %v", err)
		return
	}

	span.SetAttributes(util.TxToAttributes(tx)...)
}
