// Package screener provides the screener api.
package screener

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/contrib/screener-api/client"
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
	"golang.org/x/exp/slices"
)

// Screener is the interface for the screener.
type Screener interface {
	Start(ctx context.Context) error
}

type screenerImpl struct {
	rulesManager internal.RulesetManager
	thresholds   []config.VolumeThreshold
	db           db.RuleDB
	router       *gin.Engine
	metrics      metrics.Handler
	cfg          config.Config
	client       trmlabs.Client
	blacklist    []string
	blacklistMux sync.RWMutex
	whitelist    []string
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
	screener.thresholds = cfg.VolumeThresholds

	for _, item := range cfg.Whitelist {
		screener.whitelist = append(screener.whitelist, strings.ToLower(item))
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

	screener.router.Handle(http.MethodPost, "/api/data/sync", screener.blacklistAddress)

	return &screener, nil
}

func (s *screenerImpl) fetchBlacklist(ctx context.Context) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.cfg.BlacklistURL, nil)
	if err != nil {
		logger.Errorf("could not create blacklist request: %s", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Errorf("could not fetch blacklist: %s", err)
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var blacklist []string
	err = json.NewDecoder(resp.Body).Decode(&blacklist)
	if err != nil {
		logger.Errorf("could not decode blacklist: %s", err)
		return
	}

	s.blacklistMux.Lock()
	defer s.blacklistMux.Unlock()

	for _, item := range blacklist {
		s.blacklist = append(s.blacklist, strings.ToLower(item))
	}
}

func (s *screenerImpl) blacklistAddress(c *gin.Context) {
	var blacklistBody client.BlackListBody

	if err := c.ShouldBindJSON(&blacklistBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update DB at this point

	type_req := blacklistBody.TypeReq
	id := blacklistBody.Id
	data := blacklistBody.Data
	address := blacklistBody.Address
	network := blacklistBody.Network
	tag := blacklistBody.Tag
	remark := blacklistBody.Remark

	address = strings.ToLower(address)

	switch type_req {
	case "create":
		// add the user
		// update the new table `blacklisted_users`
		// return the status

	case "update":
		// update the user
		// update the new table `blacklisted_users`
		// return the status

	case "delete":
		// delete the user
		// update the new table `blacklisted_users`
		// return the status

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type"})
		return
	}

}

func (s *screenerImpl) Start(ctx context.Context) error {
	// TODO: potential race condition here, if the blacklist is not fetched before the first request
	// in practice trm will catch
	go func() {
		for {
			if s.cfg.BlacklistURL != "" {
				s.fetchBlacklist(ctx)
				time.Sleep(1 * time.Second * 15)
			}
		}
	}()
	connection := baseServer.Server{}
	err := connection.ListenAndServe(ctx, fmt.Sprintf(":%d", s.cfg.Port), s.router)
	if err != nil {
		return fmt.Errorf("could not start gqlServer: %w", err)
	}
	return nil
}

// screenAddress returns whether an address is risky or not given a ruleset.
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

	s.blacklistMux.RLock()
	if slices.Contains(s.blacklist, address) {
		c.JSON(http.StatusOK, gin.H{"risk": true})
		s.blacklistMux.RUnlock()
		return
	}
	s.blacklistMux.RUnlock()

	if slices.Contains(s.whitelist, address) {
		c.JSON(http.StatusOK, gin.H{"risk": false})
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
	if hasIndicator, err = currentRules.HasAddressIndicators(s.thresholds, indicators...); err != nil {
		c.JSON(http.StatusOK, gin.H{"risk": true})
		return
	}

	c.JSON(http.StatusOK, gin.H{"risk": hasIndicator})
}

func (s *screenerImpl) getIndicators(parentCtx context.Context, address string, goodUntil time.Time) (indicators []trmlabs.AddressRiskIndicator, err error) {
	ctx, span := s.metrics.Tracer().Start(parentCtx, "get-indicators")
	defer func() {
		// nolint: errchkjson
		marshalledIndicators, _ := json.Marshal(indicators)
		span.AddEvent("indicators", trace.WithAttributes(attribute.String("indicators", string(marshalledIndicators))))
		metrics.EndSpanWithErr(span, err)
	}()

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
