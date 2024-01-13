package screener

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/patrickmn/go-cache"
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type SimpleScreener struct {
	client  trmlabs.Client
	cache   *cache.Cache
	risks   map[string]bool
	metrics metrics.Handler
}

func NewSimpleScreener(apiKey, blacklistFile string, handler metrics.Handler) (*SimpleScreener, error) {
	risks, err := parseBlacklist(blacklistFile)
	if err != nil {
		return nil, fmt.Errorf("could not parse blacklist: %w", err)
	}
	c := cache.New(time.Minute*30, time.Second*10)
	client, err := trmlabs.NewClient(apiKey, "https://api.trmlabs.com")
	if err != nil {
		return nil, fmt.Errorf("could not create client: %w", err)
	}
	return &SimpleScreener{
		client:  client,
		cache:   c,
		metrics: handler,
		risks:   risks,
	}, nil
}

func parseBlacklist(blacklistFile string) (map[string]bool, error) {
	fileHandle, err := os.Open(blacklistFile)
	if err != nil {
		return nil, fmt.Errorf("could not open blacklist file: %w", err)
	}

	defer func() {
		_ = fileHandle.Close()
	}()

	r := csv.NewReader(fileHandle)

	i := 0
	// TRMID->RISK_TYPE
	risks := make(map[string]bool)
	for {
		// skip first row
		if i == 0 {
			i++
			continue
		}
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("could not read blacklist file: %w", err)
		}

		shoudBlock := false
		if record[7] != "" {
			shoudBlock = true
		}

		i++
		risks[strings.ToLower(fmt.Sprintf("%s_%s", record[2], record[4]))] = shoudBlock
	}

	return risks, nil
}

func (s *SimpleScreener) ScreenAddress(parentCtx context.Context, address common.Address) (blocked bool, err error) {
	ctx, span := s.metrics.Tracer().Start(parentCtx, "ScreenAddress", trace.WithAttributes(
		attribute.String("address", address.String()),
	))

	defer func() {
		span.AddEvent("result", trace.WithAttributes(attribute.Bool("result", blocked)))
		metrics.EndSpanWithErr(span, err)
	}()

	unmarshalledBlocked, found := s.cache.Get(address.String())
	if found {
		span.AddEvent("used_cache")
		blocked, ok := unmarshalledBlocked.(bool)
		if ok {
			return blocked, nil
		}
	}

	res, err := s.client.ScreenAddress(ctx, address.String())
	if err != nil {
		return false, fmt.Errorf("could not screen address: %w", err)
	}

	for _, addressRisk := range res {
		for _, ri := range addressRisk.AddressRiskIndicators {
			// TODO: can do in json converter
			incoming, err := strconv.ParseFloat(ri.IncomingVolumeUsd, 10)
			if err != nil {
				return false, fmt.Errorf("could not parse incoming volume: %w", err)
			}

			outgoing, err := strconv.ParseFloat(ri.OutgoingVolumeUsd, 10)
			if err != nil {
				return false, fmt.Errorf("could not parse outgoing volume: %w", err)
			}

			riskParam := strings.ToLower(fmt.Sprintf("%s_%s", ri.Category, ri.RiskType))
			isBlocked, found := s.risks[riskParam]
			if isBlocked && found && (incoming > 0 || outgoing > 0) {
				span.AddEvent(fmt.Sprintf("%s found", riskParam))
				s.cache.Set(address.String(), true, cache.DefaultExpiration)
				return true, nil
			}
		}
	}

	s.cache.Set(address.String(), false, cache.DefaultExpiration)

	return false, nil

}
