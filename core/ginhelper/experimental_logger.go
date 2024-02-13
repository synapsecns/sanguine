package ginhelper

import (
	"context"
	ginzap "github.com/gin-contrib/zap"
	"github.com/synapsecns/sanguine/core/metrics/logger"
	"go.uber.org/zap"
)

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
