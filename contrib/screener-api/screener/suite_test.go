package screener_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/Flaque/filet"
	"github.com/gocarina/gocsv"
	"github.com/phayes/freeport"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/contrib/screener-api/client"
	"github.com/synapsecns/sanguine/contrib/screener-api/config"
	"github.com/synapsecns/sanguine/contrib/screener-api/metadata"
	"github.com/synapsecns/sanguine/contrib/screener-api/screener"
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
	"time"
)

type ScreenerSuite struct {
	*testsuite.TestSuite
	metrics metrics.Handler
	port    int
}

// NewScreenerSuite creates a new screener test suite.
func NewScreenerSuite(tb testing.TB) *ScreenerSuite {
	tb.Helper()
	return &ScreenerSuite{TestSuite: testsuite.NewTestSuite(tb)}
}

func TestScreenerSuite(t *testing.T) {
	suite.Run(t, NewScreenerSuite(t))
}

// TestScreenerSuite runs the screener test suite.
func (s *ScreenerSuite) SetupSuite() {
	s.TestSuite.SetupSuite()

	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(s.GetSuiteContext(), s.T())
		metricsHandler = metrics.Jaeger
	}

	var err error
	s.metrics, err = metrics.NewByType(s.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	Nil(s.T(), err)
}

func (s *ScreenerSuite) makeTestCSV(rules []screener.Set) string {
	content, err := gocsv.MarshalString(rules)
	Nil(s.T(), err)

	file := filet.TmpFile(s.T(), "", content)
	defer func() {
		// _ = Nil(s.T(), file.Close())
	}()

	return file.Name()
}

func (s *ScreenerSuite) TestScreener() {
	var err error

	s.port, err = freeport.GetFreePort()
	Nil(s.T(), err)

	s.T().Setenv("TRM_URL", "")

	cfg := config.Config{
		TRMKey: "",
		Rulesets: map[string]config.RulesetConfig{
			"testrule": {
				Filename: s.makeTestCSV([]screener.Set{
					{
						Enabled:    "true",
						ID:         1,
						Category:   "test_category",
						Name:       "name",
						Severity:   "severity",
						TypeOfRisk: "Risk Type",
					},
				}),
			},
			"testrule2": {
				Filename: s.makeTestCSV([]screener.Set{}),
			},
		},
		BlacklistURL: "https://synapseprotocol.com/blacklist.json", // TODO: mock this out
		CacheTime:    1,
		Port:         s.port,
		Database: config.DatabaseConfig{
			Type: dbcommon.Sqlite.String(),
			DSN:  filet.TmpDir(s.T(), ""),
		},
	}

	realScreener, err := screener.NewTestScreener(s.GetTestContext(), cfg, s.metrics)
	Nil(s.T(), err)

	go func() {
		err = realScreener.Start(s.GetTestContext())
		if !errors.Is(err, context.Canceled) {
			Nil(s.T(), err)
		}
	}()

	m := mockClient{
		responseMap: map[string][]trmlabs.ScreenResponse{
			"0x123": {
				{
					AddressRiskIndicators: []trmlabs.AddressRiskIndicator{
						{
							Category:                    "test_category",
							CategoryID:                  "1",
							CategoryRiskScoreLevel:      1,
							CategoryRiskScoreLevelLabel: "test_category",
							IncomingVolumeUsd:           "1",
						},
					},
				},
			},
		},
	}

	realScreener.SetClient(m)
	time.Sleep(time.Second)

	apiClient, err := client.NewClient(s.metrics, fmt.Sprintf("http://localhost:%d", s.port))
	Nil(s.T(), err)

	// http://localhost:63575/testrule/address/0x123: true
	out, err := apiClient.ScreenAddress(s.GetTestContext(), "testrule", "0x123")
	Nil(s.T(), err)
	True(s.T(), out)

	// http://localhost:63575/testrule/address/0x00: false
	out, err = apiClient.ScreenAddress(s.GetTestContext(), "testrule", "0x00")
	Nil(s.T(), err)
	False(s.T(), out)
}

type mockClient struct {
	responseMap map[string][]trmlabs.ScreenResponse
}

// ScreenAddress mocks the screen address method.
func (m mockClient) ScreenAddress(ctx context.Context, address string) ([]trmlabs.ScreenResponse, error) {
	if m.responseMap == nil {
		return nil, fmt.Errorf("no response map")
	}

	return m.responseMap[address], nil
}

var _ trmlabs.Client = mockClient{}
