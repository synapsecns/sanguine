package proxy

import (
	"context"
	"fmt"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/services/omnirpc/chainmanager"
	"github.com/synapsecns/sanguine/services/omnirpc/collection"
	"github.com/synapsecns/sanguine/services/omnirpc/config"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"go.uber.org/zap/zapcore"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// RPCProxy proxies rpc request to the fastest endpoint. Requests fallback in cases where data is not available.
type RPCProxy struct {
	// chainManager contains a list of chains and latency ordered rpcs
	chainManager chainmanager.ChainManager
	// config contains the config for each chain
	refreshInterval time.Duration
	// port is the por the server is run on
	port uint16
	// forwarderPool is used for allocating forwarders
	forwarderPool sync.Pool
	// client contains the http client
	client omniHTTP.Client
}

// NewProxy creates a new rpc proxy.
func NewProxy(config config.Config, clientType omniHTTP.ClientType) *RPCProxy {
	return &RPCProxy{
		chainManager:    chainmanager.NewChainManagerFromConfig(config),
		refreshInterval: time.Second * time.Duration(config.RefreshInterval),
		port:            config.Port,
		client:          omniHTTP.NewClient(clientType),
	}
}

// Run runs the rpc server until context cancellation.
func (r *RPCProxy) Run(ctx context.Context) {
	tracer.Start()
	defer tracer.Stop()

	go r.startProxyLoop(ctx)

	router := gin.New()
	router.Use(gintrace.Middleware("omnirpc"))
	router.Use(requestid.New(
		requestid.WithCustomHeaderStrKey(requestid.HeaderStrKey(omniHTTP.XRequestIDString)),
		requestid.WithGenerator(func() string {
			return uuid.New().String()
		})))

	router.Use(helmet.Default())
	router.Use(gin.Recovery())
	log.SetAllLoggers(log.LevelDebug)
	router.Use(ginzap.GinzapWithConfig(logger.Desugar(), &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		Context: func(c *gin.Context) (fields []zapcore.Field) {
			requestID := c.GetHeader(omniHTTP.XRequestIDString)
			fields = append(fields, zapcore.Field{
				Key:    "request-id",
				Type:   zapcore.StringType,
				String: requestID,
			})

			return fields
		},
	}))
	router.Use(ginzap.RecoveryWithZap(logger.Desugar(), true))

	router.Use(func(c *gin.Context) {
		// set on request as well
		if c.Request.Header.Get(omniHTTP.XRequestIDString) == "" {
			c.Request.Header.Set(omniHTTP.XRequestIDString, c.Writer.Header().Get(omniHTTP.XRequestIDString))
		}
	})

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	router.POST("/rpc/:id", func(c *gin.Context) {
		chainID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("chainid must be a number: %d", chainID),
			})
		}
		r.Forward(c, uint32(chainID), nil)
	})

	router.POST("/confirmations/:confirmations/rpc/:id", func(c *gin.Context) {
		chainID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("chainid must be a number: %d", chainID),
			})
		}
		realConfs, err := strconv.Atoi(c.Param("confirmations"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("confirmations must be a number: %d", chainID),
			})
		}

		confirmations := uint16(realConfs)

		r.Forward(c, uint32(chainID), &confirmations)
	})

	router.GET("/collection.json", func(c *gin.Context) {
		res, err := collection.CreateCollection()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("could not parse collection: %v", err),
			})
		}

		c.Data(http.StatusOK, gin.MIMEJSON, res)
	})

	logger.Infof("running on port %d", r.port)
	err := router.Run(fmt.Sprintf("0.0.0.0:%d", r.port))
	if err != nil {
		logger.Warn(err)
	}
}

// scanInterval is how long to wait between latency scans.
const scanInterval = time.Second * 60

func (r *RPCProxy) startProxyLoop(ctx context.Context) {
	// TODO(trajan): jitter if not first run
	var waitTime time.Duration

	for {
		select {
		// parent loop terminated
		case <-ctx.Done():
			return
		case <-time.After(waitTime):
			var wg sync.WaitGroup

			for _, chainID := range r.chainManager.GetChainIDs() {
				wg.Add(1)

				go func(chainID uint32) {
					r.chainManager.RefreshRPCInfo(ctx, chainID)

					wg.Done()
				}(chainID)
			}

			wg.Wait()

			waitTime = scanInterval
		}
	}
}
