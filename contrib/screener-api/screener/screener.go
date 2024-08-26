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

	"github.com/synapsecns/sanguine/core/mapmutex"
	"golang.org/x/sync/errgroup"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/contrib/screener-api/chainalysis"
	"github.com/synapsecns/sanguine/contrib/screener-api/client"
	"github.com/synapsecns/sanguine/contrib/screener-api/config"
	"github.com/synapsecns/sanguine/contrib/screener-api/db"
	"github.com/synapsecns/sanguine/contrib/screener-api/db/sql"
	"github.com/synapsecns/sanguine/contrib/screener-api/docs"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	okResponse  = "OK"
	errResponse = "ERROR"
	meterName   = "github.com/synapsecns/sanguine/contrib/screener-api"
)

// Screener is the interface for the screener.
type Screener interface {
	Start(ctx context.Context) error
}

type screenerImpl struct {
	db                       db.DB
	router                   *gin.Engine
	metrics                  metrics.Handler
	cfg                      config.Config
	client                   chainalysis.Client
	whitelist                map[string]bool
	blacklist                map[string]bool
	blacklistCacheMux        sync.RWMutex
	requestMux               mapmutex.StringMapMutex
	blockedAddressesMetric   metric.Int64Counter
	unblockedAddressesMetric metric.Int64Counter
}

var logger = log.Logger("screener")

// NewScreener creates a new screener.
func NewScreener(ctx context.Context, cfg config.Config, metricHandler metrics.Handler) (_ Screener, err error) {
	screener := screenerImpl{
		metrics:    metricHandler,
		cfg:        cfg,
		requestMux: mapmutex.NewStringMapMutex(),
	}

	docs.SwaggerInfo.Title = "Screener API"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", cfg.Port)

	screener.client = chainalysis.NewClient(metricHandler, cfg.RiskLevels, cfg.ChainalysisKey, core.GetEnv("CHAINALYSIS_URL", cfg.ChainalysisURL))

	screener.blacklist = make(map[string]bool)
	screener.whitelist = make(map[string]bool)
	for _, item := range cfg.Whitelist {
		screener.whitelist[strings.ToLower(item)] = true
	}

	for _, item := range cfg.Blacklist {
		screener.blacklist[strings.ToLower(item)] = true
	}

	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return nil, fmt.Errorf("could not get db type: %w", err)
	}
	screener.db, err = sql.Connect(ctx, dbType, cfg.Database.DSN, metricHandler)
	if err != nil {
		return nil, fmt.Errorf("could not connect to rules db: %w", err)
	}

	meter := metricHandler.Meter(meterName)
	if screener.blockedAddressesMetric, err = meter.Int64Counter("blocked_addresses"); err != nil {
		return nil, fmt.Errorf("could not create blocked addresses metric: %w", err)
	}

	if screener.unblockedAddressesMetric, err = meter.Int64Counter("unblocked_addresses"); err != nil {
		return nil, fmt.Errorf("could not create unblocked addresses metric: %w", err)
	}

	screener.router = ginhelper.New(logger)
	screener.router.Use(screener.metrics.Gin()...)

	// Blacklist route
	screener.router.POST("/api/data/sync", ginhelper.TraceMiddleware(metricHandler.Tracer(), true), screener.authMiddleware(cfg), screener.blacklistAddress)

	// Screening routes
	screener.router.GET("/address/:address", screener.screenAddress)
	// deprecated and ruleset is not used, this is for backwards compatibility
	screener.router.GET("/:ruleset/address/:address", screener.screenAddress)

	// Swagger routes
	screener.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	gin.SetMode(gin.ReleaseMode)

	return &screener, nil
}

const blacklistScreenInterval = 15 * time.Second

func (s *screenerImpl) Start(ctx context.Context) error {
	// TODO: potential race condition here, if the blacklist is not fetched before the first request
	// in practice chainalysis will catch
	go func() {
		// Fetch the blacklist at start.
		s.fetchBlacklist(ctx)

		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(blacklistScreenInterval):
				s.fetchBlacklist(ctx)
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

func (s *screenerImpl) fetchBlacklist(ctx context.Context) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		s.cfg.BlacklistURL,
		nil,
	)
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

	s.blacklistCacheMux.Lock()
	defer s.blacklistCacheMux.Unlock()
	for _, item := range blacklist {
		s.blacklist[strings.ToLower(item)] = true
	}
}

// screenAddress godoc
// @Summary Screen an address for risk
// @Description Screen an address using Chainalysis API to determine if it's high risk
// @Tags Address Screening
// @Accept  json
// @Produce  json
// @Param   address path string true "Address to be screened"
// @Accept json
// @Produce json
// @Router /screen/{address} [get].
func (s *screenerImpl) screenAddress(c *gin.Context) {
	address := strings.ToLower(c.Param("address"))
	if address == "" {
		logger.Errorf("address is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "address is required"})
		return
	}

	// Check if the address is in the blacklist.
	if s.isBlacklistedCache(address) {
		c.JSON(http.StatusOK, gin.H{"risk": true})
		return
	}

	// Check if the address is in the whitelist.
	if _, whitelisted := s.whitelist[address]; whitelisted {
		c.JSON(http.StatusOK, gin.H{"risk": false})
		return
	}

	// prevent a single address from saturating the server.
	// the only case this is useful is with a bad client that continuously sends requests for the same address.
	// due to a goroutine leak, etc.
	unlocker := s.requestMux.Lock(address)
	defer unlocker.Unlock()

	g, ctx := errgroup.WithContext(c.Request.Context())
	var isAPIBlocked, isDBBlocked bool
	g.Go(func() (err error) {
		// If not, check db & Chainalysis for the risk assessment.
		isAPIBlocked, err = s.client.ScreenAddress(ctx, address)
		if err != nil {
			return fmt.Errorf("error screening address: %w", err)
		}
		return nil
	})

	g.Go(func() (err error) {
		isDBBlocked, err = s.isDBBlacklisted(ctx, address)
		if err != nil {
			return fmt.Errorf("error checking db: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		logger.Errorf("error screening address: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if isAPIBlocked || isDBBlocked {
		s.blockedAddressesMetric.Add(ctx, 1)
	} else {
		s.unblockedAddressesMetric.Add(ctx, 1)
	}

	c.JSON(http.StatusOK, gin.H{"risk": isAPIBlocked || isDBBlocked})
}

func (s *screenerImpl) isDBBlacklisted(ctx context.Context, address string) (bool, error) {
	_, err := s.db.GetBlacklistedAddress(ctx, address)
	if err != nil && !errors.Is(err, db.ErrNoAddressNotFound) {
		return false, fmt.Errorf("could not get blacklisted address: %w", err)
	}

	if errors.Is(err, db.ErrNoAddressNotFound) {
		return false, nil
	}

	return true, nil
}

func (s *screenerImpl) isBlacklistedCache(address string) bool {
	s.blacklistCacheMux.RLock()
	defer s.blacklistCacheMux.RUnlock()
	return s.blacklist[address]
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
	defer metrics.EndSpanWithErr(span, err)

	var blacklistBody client.BlackListBody

	// Grab the body of the JSON request and unmarshal it into the blacklistBody struct.
	if err := c.ShouldBindBodyWith(&blacklistBody, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errResponse})
		return
	}

	span.SetAttributes(
		attribute.String("type", blacklistBody.Type),
		(attribute.String("id", blacklistBody.ID)),
		(attribute.String("network", blacklistBody.Data.Network)),
		(attribute.String("tag", blacklistBody.Data.Tag)),
		(attribute.String("remark", blacklistBody.Data.Remark)),
		(attribute.String("address", blacklistBody.Data.Address)),
	)

	blacklistedAddress := db.BlacklistedAddress{
		Type:    blacklistBody.Type,
		ID:      blacklistBody.ID,
		Network: blacklistBody.Data.Network,
		Tag:     blacklistBody.Data.Tag,
		Remark:  blacklistBody.Data.Remark,
		Address: strings.ToLower(blacklistBody.Data.Address),
	}

	s.blacklistCacheMux.Lock()
	defer s.blacklistCacheMux.Unlock()
	s.blacklist[blacklistBody.Data.Address] = true

	switch blacklistBody.Type {
	case "create":
		if err := s.db.PutBlacklistedAddress(ctx, blacklistedAddress); err != nil {
			span.AddEvent("error", trace.WithAttributes(attribute.String("error", err.Error())))
			c.JSON(http.StatusInternalServerError, gin.H{"error": errResponse})
			return
		}

		span.AddEvent("blacklistedAddress", trace.WithAttributes(attribute.String("address", blacklistBody.Data.Address)))
		c.JSON(http.StatusOK, gin.H{"status": okResponse})
		return

	case "update":
		if err := s.db.UpdateBlacklistedAddress(ctx, blacklistedAddress.ID, blacklistedAddress); err != nil {
			span.AddEvent("error", trace.WithAttributes(attribute.String("error", err.Error())))
			c.JSON(http.StatusInternalServerError, gin.H{"error": errResponse})
			return
		}

		span.AddEvent("blacklistedAddress", trace.WithAttributes(attribute.String("address", blacklistBody.Data.Address)))
		c.JSON(http.StatusOK, gin.H{"status": okResponse})
		return

	case "delete":
		if err := s.db.DeleteBlacklistedAddress(ctx, blacklistedAddress.ID); err != nil {
			span.AddEvent("error", trace.WithAttributes(attribute.String("error", err.Error())))
			c.JSON(http.StatusInternalServerError, gin.H{"error": errResponse})
			return
		}

		span.AddEvent("blacklistedAddress", trace.WithAttributes(attribute.String("address", blacklistBody.Data.Address)))
		c.JSON(http.StatusOK, gin.H{"status": okResponse})
		return

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": errResponse})
		return
	}
}

// This function takes the HTTP headers and the body of the request and reconstructs the signature to
// compare it with the signature provided. If they match, the request is allowed to pass through.
// nolint: canonicalheader
func (s *screenerImpl) authMiddleware(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, span := s.metrics.Tracer().Start(c.Request.Context(), "authMiddleware")
		defer span.End()

		appID := c.Request.Header.Get("X-Signature-appid")
		timestamp := c.Request.Header.Get("X-Signature-timestamp")
		nonce := c.Request.Header.Get("X-Signature-nonce")
		signature := c.Request.Header.Get("X-Signature-signature")
		queryString := c.Request.URL.RawQuery

		bodyBz, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errResponse})
			c.Abort()
			return
		}
		// Put it back so we can read it again for DB operations.
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBz))

		bodyStr, err := core.BytesToJSONString(bodyBz)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errResponse})
			c.Abort()
			return
		}

		var message string
		if len(queryString) > 0 {
			message = fmt.Sprintf(
				"%s;%s;%s;%s;%s;%s;%s",
				appID, timestamp, nonce, "POST", "/api/data/sync", queryString, bodyStr,
			)
		} else {
			message = fmt.Sprintf(
				"%s;%s;%s;%s;%s;%s",
				appID, timestamp, nonce, "POST", "/api/data/sync", bodyStr,
			)
		}

		expectedSignature := client.GenerateSignature(cfg.AppSecret, message)

		span.SetAttributes(
			attribute.String("appid", appID),
			attribute.String("timestamp", timestamp),
			attribute.String("nonce", nonce),
			attribute.String("signature", signature),
			attribute.String("queryString", queryString),
			attribute.String("body", bodyStr),
			attribute.String("expectedSignature", expectedSignature),
			attribute.String("message", message),
		)

		if expectedSignature != signature {
			span.AddEvent(
				"error",
				trace.WithAttributes(attribute.String("error", "Invalid signature"+expectedSignature)),
			)
			c.JSON(http.StatusUnauthorized, gin.H{"error": errResponse})
			c.Abort()
			return
		}
		span.AddEvent("success", trace.WithAttributes(attribute.String("message", "Valid signature"+expectedSignature)))
		c.Next()
	}
}
