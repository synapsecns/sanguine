package ginhelper

import (
	"bytes"
	"context"
	"github.com/synapsecns/sanguine/core/metrics/logger"

	// embed is used for importing the robots.txt file.
	_ "embed"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-http-utils/headers"
	"github.com/google/uuid"
	"github.com/ipfs/go-log"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"
)

//go:embed robots.txt
var robots []byte

// New creates a new gin server with some sensible defaults.
// these include:
// - helmet-default handlers
// - request-ids (used for stack tracing)
// - cors (used for requests from the frontend)
// - health-checks
// - restrictive robots.txt.
func New(logger *log.ZapEventLogger) *gin.Engine {
	server := newBase()

	server.Use(ginzap.RecoveryWithZap(logger.Desugar(), true))

	// add the request id to the logger
	server.Use(ginzap.GinzapWithConfig(logger.Desugar(), &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		Context: func(c *gin.Context) (fields []zapcore.Field) {
			requestID := c.GetHeader(RequestIDHeader)
			fields = append(fields, zapcore.Field{
				Key:    "request-id",
				Type:   zapcore.StringType,
				String: requestID,
			})

			return fields
		},
	}))

	return server
}

// NewWithExperimentalLogger creates a new gin server with some sensible defaults.
// See New for more information.
func NewWithExperimentalLogger(ctx context.Context, logger logger.ExperimentalLogger) *gin.Engine {
	server := newBase()

	wrapped := wrappedExperimentalLogger{
		ctx:    ctx,
		logger: logger,
	}

	server.Use(ginzap.RecoveryWithZap(wrapped, true))
	server.Use(ginzap.GinzapWithConfig(wrapped, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		Context: func(c *gin.Context) (fields []zapcore.Field) {
			requestID := c.GetHeader(RequestIDHeader)
			fields = append(fields, zapcore.Field{
				Key:    "request-id",
				Type:   zapcore.StringType,
				String: requestID,
			})

			return fields
		},
	}))

	return server
}

func newBase() *gin.Engine {
	server := gin.New()
	// required for opentracing.
	server.ContextWithFallback = true
	server.Use(helmet.Default())
	server.Use(gin.Recovery())
	server.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"GET", "PUT", "POST", "PATCH", "DELETE", "OPTIONS"},
		MaxAge:          12 * time.Hour,
	}))

	// configure the request id
	server.Use(requestid.New(
		requestid.WithCustomHeaderStrKey(RequestIDHeader),
		requestid.WithGenerator(func() string {
			return uuid.New().String()
		})))

	// set the request id header if the client didn't
	server.Use(func(c *gin.Context) {
		if c.Request.Header.Get(RequestIDHeader) == "" {
			c.Request.Header.Set(RequestIDHeader, c.Writer.Header().Get(RequestIDHeader))
		}
	})

	server.GET(HealthCheck, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	server.GET(RobotsTxt, func(context *gin.Context) {
		reader := bytes.NewReader(robots)
		context.Header(headers.ContentType, gin.MIMEPlain)
		http.ServeContent(context.Writer, context.Request, "robots.txt", bootTime, reader)
	})

	return server
}

// HealthCheck is the health check endpoint.
const HealthCheck string = "/health-check"

// RobotsTxt is used for apis to disallow crawls.
const RobotsTxt string = "/robots.txt"

// RequestIDHeader is used for tracking request ids.
const RequestIDHeader = "X-Request-ID"

// bootTime is used for the mod-time of the robots.txt.
var bootTime = time.Now()

// MetricsEndpoint is used for prometheus metrics.
// Deprecated: use METRICS_PATH instead.
const MetricsEndpoint string = "/metrics"
