// Package screener provides the screener api.
package screener

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/contrib/screener-api/config"
	"github.com/synapsecns/sanguine/contrib/screener-api/db"
	"github.com/synapsecns/sanguine/contrib/screener-api/db/sql"
	"github.com/synapsecns/sanguine/contrib/screener-api/screener/internal"
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"strings"
	"time"
)

// Screener is the interface for the screener.
type Screener interface {
	Start(ctx context.Context) error
}

type screenerImpl struct {
	rulesManager internal.RulesetManager
	db           db.RuleDB
	router       *gin.Engine
	metrics      metrics.Handler
	cfg          config.Config
	client       trmlabs.Client
}

var logger = log.Logger("screener")

// NewScreener creates a new screener.
func NewScreener(ctx context.Context, cfg config.Config, metricHandler metrics.Handler) (_ Screener, err error) {
	screener := screenerImpl{
		metrics: metricHandler,
		cfg:     cfg,
	}

	screener.client, err = trmlabs.NewClient(cfg.TRMKey, core.GetEnv("TRM_URL", "https://api.trmlabs.com"))
	if err != nil {
		return nil, fmt.Errorf("could not create trm client: %w", err)
	}

	screener.rulesManager, err = setupScreener(cfg.Rulesets)
	if err != nil {
		return nil, fmt.Errorf("could not setup screener: %w", err)
	}

	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return nil, fmt.Errorf("could not get db type: %w", err)
	}

	screener.db, err = sql.Connect(ctx, dbType, cfg.Database.DSN, metricHandler)
	if err != nil {
		return nil, fmt.Errorf("could not connect to db: %w", err)
	}

	screener.router = ginhelper.New(logger)
	screener.router.Handle(http.MethodGet, "/:ruleset/address/:address", screener.screenAddress)

	return &screener, nil
}

func (s *screenerImpl) Start(ctx context.Context) error {
	connection := baseServer.Server{}
	err := connection.ListenAndServe(ctx, fmt.Sprintf(":%d", s.cfg.Port), s.router)
	if err != nil {
		return fmt.Errorf("could not start gqlServer: %w", err)
	}
	return nil
}

func (s *screenerImpl) screenAddress(c *gin.Context) {
	var err error

	ruleset := strings.ToLower(c.Param("ruleset"))
	if ruleset == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ruleset is required"})
		return
	}

	address := strings.ToLower(c.Param("address"))
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address is required"})
		return
	}

	ctx, span := s.metrics.Tracer().Start(c.Request.Context(), "screenAddress", trace.WithAttributes(attribute.String("address", address)))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	currentRules := s.rulesManager.GetRuleset(ruleset)
	if currentRules == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ruleset not found"})
		return
	}

	goodUntil := time.Now().Add(-1 * s.cfg.GetCacheTime(ruleset))
	var indicators []trmlabs.AddressRiskIndicator
	if indicators, err = s.getIndicators(ctx, address, goodUntil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var hasIndicator bool
	if hasIndicator, err = currentRules.HasAddressIndicators(indicators...); err != nil {
		c.JSON(http.StatusOK, gin.H{"risk": true})
		return
	}

	c.JSON(http.StatusOK, gin.H{"risk": hasIndicator})
}

func (s *screenerImpl) getIndicators(ctx context.Context, address string, goodUntil time.Time) ([]trmlabs.AddressRiskIndicator, error) {
	riskIndicators, err := s.db.GetAddressIndicators(ctx, address, goodUntil)
	if err == nil {
		return riskIndicators, nil
	}

	if !errors.Is(err, db.ErrNoAddressNotCached) {
		return nil, fmt.Errorf("could not get address indicators: %w", err)
	}

	response, err := s.client.ScreenAddress(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("could not screen address: %w", err)
	}

	for _, ri := range response {
		riskIndicators = append(riskIndicators, ri.AddressRiskIndicators...)
	}

	err = s.db.PutAddressIndicators(ctx, address, riskIndicators)
	if err != nil {
		return nil, fmt.Errorf("could not put address indicators: %w", err)
	}

	return riskIndicators, nil
}
