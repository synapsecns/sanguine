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

type StipAPIServer struct {
	cfg     stipconfig.Config
	engine  *gin.Engine
	handler metrics.Handler
}

func NewStipAPI(
	ctx context.Context,
	cfg stipconfig.Config,
	handler metrics.Handler,
) (*StipAPIServer, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is nil")
	}
	if handler == nil {
		return nil, fmt.Errorf("handler is nil")
	}

	return &StipAPIServer{
		cfg:     cfg,
		handler: handler,
	}, nil

}

var logger = log.Logger("stip-api")

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

func (h *Handler) GetFeeAndRebateInfo(c *gin.Context) {
	feesAndRebates := ConvertFeesAndRebatesToJSON(h.cfg.FeesAndRebates)
	c.JSON(http.StatusOK, feesAndRebates)
}

const (
	getHealthRoute      = "/health"
	GetFeeAndRebateInfo = "/feesandrebate"
)

// Run runs the rest api server.
func (r *StipAPIServer) Run(ctx context.Context) error {
	// TODO: Use Gin Helper
	engine := ginhelper.New(logger)
	h := NewHandler(r.cfg)

	// Assign GET routes
	engine.GET(getHealthRoute, h.GetHealth)
	engine.GET(GetFeeAndRebateInfo, h.GetFeeAndRebateInfo)

	r.engine = engine

	connection := baseServer.Server{}
	fmt.Printf("starting api at http://localhost:%s\n", r.cfg.StipAPIPort)
	err := connection.ListenAndServe(ctx, fmt.Sprintf(":%s", r.cfg.StipAPIPort), r.engine)
	if err != nil {
		return fmt.Errorf("could not start relayer api server: %w", err)
	}

	return nil
}

func ConvertFeesAndRebatesToJSON(feesAndRebates stipconfig.FeesAndRebates) map[int]interface{} {
	jsonOutput := make(map[int]interface{})

	for toChain, moduleFeeRebate := range feesAndRebates {
		fromChain := determineFromChain(toChain)

		// Initialize the toChainMap if necessary
		if _, exists := jsonOutput[toChain]; !exists {
			jsonOutput[toChain] = make(map[string]interface{})
		}
		toChainMap := jsonOutput[toChain].(map[string]interface{})

		// Initialize the fromChainMap if necessary
		fromChainMap := make(map[string]interface{})
		if existingFromChainMap, exists := toChainMap[fromChain]; exists {
			fromChainMap = existingFromChainMap.(map[string]interface{})
		} else {
			toChainMap[fromChain] = fromChainMap
		}

		for moduleName, tokenFeeRebate := range moduleFeeRebate {
			// Initialize moduleMap if necessary
			if _, exists := fromChainMap[moduleName]; !exists {
				fromChainMap[moduleName] = make(map[string]interface{})
			}
			moduleMap := fromChainMap[moduleName].(map[string]interface{})

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
