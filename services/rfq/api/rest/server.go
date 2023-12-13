package rest

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/metrics"
	baseServer "github.com/synapsecns/sanguine/core/server"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/api/config"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
)

type APIServer struct {
	cfg           config.Config
	db            db.ApiDB
	engine        *gin.Engine
	omnirpcClient omniClient.RPCClient
	handler       metrics.Handler
}

func NewAPI(
	ctx context.Context,
	cfg config.Config,
	handler metrics.Handler,
	omniRPCClient omniClient.RPCClient,
	store db.ApiDB,
) (*APIServer, error) {
	engine := SetupRouter()
	fmt.Println("Context:", ctx)
	fmt.Println("Config:", cfg)
	fmt.Println("Handler:", handler)
	fmt.Println("OmniRPCClient:", omniRPCClient)
	fmt.Println("Store:", store)
	fmt.Println("engine:", engine)

	if ctx == nil {
		return nil, fmt.Errorf("context is nil")
	}
	if handler == nil {
		return nil, fmt.Errorf("handler is nil")
	}
	if omniRPCClient == nil {
		return nil, fmt.Errorf("omniRPCClient is nil")
	}
	if store == nil {
		return nil, fmt.Errorf("store is nil")
	}
	if engine == nil {
		return nil, fmt.Errorf("engine is nil")
	}
	return &APIServer{
		cfg:           cfg,
		db:            store,
		engine:        engine,
		omnirpcClient: omniRPCClient,
		handler:       handler,
	}, nil
}

// Run runs the rest api server.
func (r *APIServer) Run(ctx context.Context) error {
	connection := baseServer.Server{}
	err := connection.ListenAndServe(ctx, fmt.Sprintf(":%s", r.cfg.Port), r.engine)
	if err != nil {
		return fmt.Errorf("could not start rest api server: %w", err)
	}

	return nil
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	h := NewHandler()

	r.PUT("/quotes", h.ModifyQuote)
	r.GET("/quotes", h.GetQuotes)
	r.GET("/quotes/filter", h.GetFilteredQuotes)

	return r
}
