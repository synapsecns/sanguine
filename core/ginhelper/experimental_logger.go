package ginhelper

import (
	"context"
	ginzap "github.com/gin-contrib/zap"
	"github.com/synapsecns/sanguine/core/metrics/logger"
	"go.uber.org/zap"
)

// TODO: should be replaced with custom methods.
// this is currently insufficient and will not eep logged errors from gin in line w/ spans.
// the middleware itself needs to pass the context to the logger so this requires some copy/paste actoin
// from the gin-contrib for zap.
type wrappedExperimentalLogger struct {
	// nolint: containedctx
	ctx    context.Context
	logger logger.ExperimentalLogger
}

func (w wrappedExperimentalLogger) Info(msg string, fields ...zap.Field) {
	w.logger.WithOptions(zap.Fields(fields...)).Infof(w.ctx, msg)
}

func (w wrappedExperimentalLogger) Error(msg string, fields ...zap.Field) {
	w.logger.WithOptions(zap.Fields(fields...)).Errorw(w.ctx, msg)
}

var _ ginzap.ZapLogger = &wrappedExperimentalLogger{}
