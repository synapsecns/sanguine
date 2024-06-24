// Package screener provides the screener api.
package screener

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/contrib/screener-api/client"
	"github.com/synapsecns/sanguine/contrib/screener-api/config"
	"github.com/synapsecns/sanguine/contrib/screener-api/db"
	"github.com/synapsecns/sanguine/contrib/screener-api/db/sql"
	"github.com/synapsecns/sanguine/contrib/screener-api/docs"
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

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Screener is the interface for the screener.
type Screener interface {
	Start(ctx context.Context) error
}

type screenerImpl struct {
	rulesManager internal.RulesetManager
	thresholds   []config.VolumeThreshold
	db           db.DB
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

	docs.SwaggerInfo.Title = "Screener API"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", cfg.Port)

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
		return nil, fmt.Errorf("could not connect to rules db: %w", err)
	}

	screener.router = ginhelper.New(logger)
	screener.router.Use(screener.metrics.Gin())
	screener.router.Handle(http.MethodGet, "/:ruleset/address/:address", screener.screenAddress)

	screener.router.Handle(http.MethodPost, "/api/data/sync", screener.authMiddleware(cfg), screener.blacklistAddress)
	screener.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

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

// @dev Protected Method
// @Summary blacklist an address
// @Description blacklist an address
// @Param appid header string true "Application ID"
// @Param timestamp header string true "Timestamp of the request"
// @Param nonce header string true "A unique nonce for the request"
// @Param queryString header string true "Query string parameters included in the request"
// @Param signature header string true "Signature for request validation"
// @Param request body db.BlacklistedAddress true "Blacklist request"
// @Accept json
// @Produce json
// @Router /api/data/sync [post].
func (s *screenerImpl) blacklistAddress(c *gin.Context) {
	var err error
	ctx, span := s.metrics.Tracer().Start(c.Request.Context(), "blacklistAddress")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	var blacklistBody client.BlackListBody

	// Grab the body of the JSON request and unmarshal it into the blacklistBody struct.
	if err := c.ShouldBindBodyWith(&blacklistBody, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	span.SetAttributes(attribute.String("type", blacklistBody.Type))
	span.SetAttributes(attribute.String("id", blacklistBody.ID))
	span.SetAttributes(attribute.String("data", blacklistBody.Data))
	span.SetAttributes(attribute.String("network", blacklistBody.Network))
	span.SetAttributes(attribute.String("tag", blacklistBody.Tag))
	span.SetAttributes(attribute.String("remark", blacklistBody.Remark))
	span.SetAttributes(attribute.String("address", blacklistBody.Address))

	blacklistedAddress := db.BlacklistedAddress{
		Type:    blacklistBody.Type,
		ID:      blacklistBody.ID,
		Data:    blacklistBody.Data,
		Network: blacklistBody.Network,
		Tag:     blacklistBody.Tag,
		Remark:  blacklistBody.Remark,
		Address: strings.ToLower(blacklistBody.Address),
	}

	switch blacklistBody.Type {
	case "create":
		if err := s.db.PutBlacklistedAddress(ctx, blacklistedAddress); err != nil {
			span.AddEvent("error", trace.WithAttributes(attribute.String("error", err.Error())))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		span.AddEvent("blacklistedAddress", trace.WithAttributes(attribute.String("address", blacklistBody.Address)))
		c.JSON(http.StatusOK, gin.H{"status": "success"})
		return

	case "update":
		if err := s.db.UpdateBlacklistedAddress(ctx, blacklistedAddress.ID, blacklistedAddress); err != nil {
			span.AddEvent("error", trace.WithAttributes(attribute.String("error", err.Error())))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		span.AddEvent("blacklistedAddress", trace.WithAttributes(attribute.String("address", blacklistBody.Address)))
		c.JSON(http.StatusOK, gin.H{"status": "success"})
		return

	case "delete":
		if err := s.db.DeleteBlacklistedAddress(ctx, blacklistedAddress.Address); err != nil {
			span.AddEvent("error", trace.WithAttributes(attribute.String("error", err.Error())))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		span.AddEvent("blacklistedAddress", trace.WithAttributes(attribute.String("address", blacklistBody.Address)))
		c.JSON(http.StatusOK, gin.H{"status": "success"})
		return

	default:
		span.AddEvent("error", trace.WithAttributes(attribute.String("error", err.Error())))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type"})
		return
	}
}

// This function takes the HTTP headers and the body of the request and reconstructs the signature to
// compare it with the signature provided. If they match, the request is allowed to pass through.
func (s *screenerImpl) authMiddleware(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, span := s.metrics.Tracer().Start(c.Request.Context(), "authMiddleware")
		defer span.End()

		appID := c.Request.Header.Get("X-Signature-appid")
		timestamp := c.Request.Header.Get("X-Signature-timestamp")
		nonce := c.Request.Header.Get("X-Signature-nonce")
		signature := c.Request.Header.Get("X-Signature-signature")
		queryString := c.Request.URL.RawQuery

		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not read request body"})
			c.Abort()
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		bodyStr := string(bodyBytes)

		message := fmt.Sprintf("%s%s%s%s%s%s%s",
			appID, timestamp, nonce, "POST", "/api/data/sync", queryString, bodyStr)

		expectedSignature := client.GenerateSignature(cfg.AppSecret, message)

		if expectedSignature != signature {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
			c.Abort()
			return
		}

		c.Next()
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
		return fmt.Errorf("could not start server: %w", err)
	}
	return nil
}

// screenAddress returns whether an address is risky or not given a ruleset.
// @Summary Screen address for risk
// @Description Assess the risk associated with a given address using specified rulesets.
// @Tags address
// @Accept  json
// @Produce  json
// @Param ruleset query string true "Ruleset to use for screening the address"
// @Param address query string true "Address to be screened"
// @Success 200 {object} map[string]bool "Returns the risk assessment result"
// @Failure 400 {object} map[string]string "Returns error if the required parameters are missing or invalid"
// @Failure 500 {object} map[string]string "Returns error if there are problems processing the indicators"
// @Router /screen/{ruleset}/{address} [get].
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
