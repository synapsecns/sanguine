// Package stipapi provides RESTful API services for the STIP relayer
package stipapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/services/stiprelayer/stipconfig"
)

// Server struct holds the configuration, gin engine, and metrics handler.
type Server struct {
	cfg     stipconfig.Config
	engine  *gin.Engine
	handler metrics.Handler
}

// NewStipAPI creates a new instance of Server with the provided configuration and metrics handler.
func NewStipAPI(
	ctx context.Context,
	cfg stipconfig.Config,
	handler metrics.Handler,
) (*Server, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is nil")
	}
	if handler == nil {
		return nil, fmt.Errorf("handler is nil")
	}

	return &Server{
		cfg:     cfg,
		handler: handler,
	}, nil
}

var logger = log.Logger("stip-api")

// Handler is the REST API handler.
type Handler struct {
	cfg stipconfig.Config
}

// NewHandler creates a new REST API handler.
func NewHandler(cfg stipconfig.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

// GetHealth returns a successful response to signify the API is up and running.
func (h *Handler) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// GetFeeAndRebateInfo returns the current STIP Relayer's rebate configuration.
func (h *Handler) GetFeeAndRebateInfo(c *gin.Context) {
	feesAndRebates := ConvertFeesAndRebatesToJSON(h.cfg.FeesAndRebates)
	c.JSON(http.StatusOK, feesAndRebates)
}

const (
	getHealthRoute      = "/health"
	getFeeAndRebateInfo = "/fee-rebate-bps"
)

// Run runs the rest api server.
func (r *Server) Run(ctx context.Context) error {
	// TODO: Use Gin Helper
	engine := ginhelper.New(logger)
	h := NewHandler(r.cfg)

	// Assign GET routes
	engine.GET(getHealthRoute, h.GetHealth)
	engine.GET(getFeeAndRebateInfo, h.GetFeeAndRebateInfo)

	r.engine = engine

	connection := baseServer.Server{}
	fmt.Printf("starting api at http://localhost:%s\n", r.cfg.StipAPIPort)
	err := connection.ListenAndServe(ctx, fmt.Sprintf(":%s", r.cfg.StipAPIPort), r.engine)
	if err != nil {
		return fmt.Errorf("could not start relayer api server: %w", err)
	}

	return nil
}

// ConvertFeesAndRebatesToJSON converts the configured fees and rebates to a JSON that is more consumable.
func ConvertFeesAndRebatesToJSON(feesAndRebates stipconfig.FeesAndRebates) map[int]interface{} {
	jsonOutput := make(map[int]interface{})

	for toChain, moduleFeeRebate := range feesAndRebates {
		fromChain := determineFromChain(toChain)

		// Initialize the toChainMap if necessary
		if _, exists := jsonOutput[toChain]; !exists {
			jsonOutput[toChain] = make(map[string]interface{})
		}
		toChainMap, ok := jsonOutput[toChain].(map[string]interface{})
		if !ok {
			// Instead of logging fatally, we should handle the error gracefully
			// Log the error and continue with an empty map for toChainMap
			fmt.Printf("Type assertion failed: expected map[string]interface{}, got %T\n", jsonOutput[toChain])
			toChainMap = make(map[string]interface{})
		}

		// Initialize the fromChainMap if necessary
		var fromChainMap map[string]interface{}
		existingFromChainMap, exists := toChainMap[fromChain]
		if exists {
			var ok bool
			fromChainMap, ok = existingFromChainMap.(map[string]interface{})
			if !ok {
				fmt.Printf("Type assertion failed: expected map[string]interface{}, got %T\n", existingFromChainMap)
				fromChainMap = make(map[string]interface{})
			}
		} else {
			fromChainMap = make(map[string]interface{})
			toChainMap[fromChain] = fromChainMap
		}

		for moduleName, tokenFeeRebate := range moduleFeeRebate {
			// Initialize moduleMap if necessary
			if _, exists := fromChainMap[moduleName]; !exists {
				fromChainMap[moduleName] = make(map[string]interface{})
			}
			moduleMap, ok := fromChainMap[moduleName].(map[string]interface{})
			if !ok {
				fmt.Printf("Type assertion failed: expected map[string]interface{}, got %T\n", fromChainMap[moduleName])
				moduleMap = make(map[string]interface{})
			}

			for token, feeRebate := range tokenFeeRebate {
				// Convert each FeeRebate into a map with "fee" and "rebate" as keys
				moduleMap[token] = map[string]int{"fee": feeRebate.Fee, "rebate": feeRebate.Rebate}
			}
		}
	}

	return jsonOutput
}

func determineFromChain(toChain int) string {
	switch toChain {
	case 42161:
		return "anyFromChain"
	case 1:
		return "42161"
	case 43114:
		return "42161"
	default:
		return "anyFromChain"
	}
}
