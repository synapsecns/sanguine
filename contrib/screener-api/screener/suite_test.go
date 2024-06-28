package screener_test

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/phayes/freeport"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/contrib/screener-api/chainalysis"
	"github.com/synapsecns/sanguine/contrib/screener-api/client"
	"github.com/synapsecns/sanguine/contrib/screener-api/config"
	"github.com/synapsecns/sanguine/contrib/screener-api/metadata"
	"github.com/synapsecns/sanguine/contrib/screener-api/screener"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
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

func (s *ScreenerSuite) TestScreener() {
	var err error

	s.port, err = freeport.GetFreePort()
	Nil(s.T(), err)

	s.T().Setenv("TRM_URL", "")

	cfg := config.Config{
		AppSecret:    "secret",
		AppID:        "appid",
		BlacklistURL: "https://synapseprotocol.com/blacklist.json", // TODO: mock this out
		Port:         s.port,
		Database: config.DatabaseConfig{
			Type: dbcommon.Sqlite.String(),
			DSN:  filet.TmpDir(s.T(), ""),
		},
		RiskLevels: []string{"Severe", "High"},
	}

	realScreener, err := screener.NewTestScreener(s.GetTestContext(), cfg, s.metrics)
	Nil(s.T(), err)
	NotNil(s.T(), realScreener)
	go func() {
		err = realScreener.Start(s.GetTestContext())
		if !errors.Is(err, context.Canceled) {
			Nil(s.T(), err)
		}
	}()

	m := mockClient{
		risks: []string{"Severe", "High"},
		entityMap: map[string]*Entity{
			"0x123": {
				Address:                "0x123",
				Risk:                   "Severe",
				Cluster:                Cluster{Name: "Example Cluster 2", Category: "benign activity"},
				RiskReason:             "Low risk example",
				AddressType:            "EXCHANGE",
				AddressIdentifications: []interface{}{},
				Exposures: []Exposure{
					{Category: "decentralized exchange", Value: 1234.56, ExposureType: "indirect", Direction: "both_directions"},
					{Category: "mining", Value: 789.01, ExposureType: "direct", Direction: "both_directions"},
				},
				Triggers: []interface{}{},
			},
			"0x456": {
				Address:                "0x456",
				Risk:                   "High",
				Cluster:                Cluster{Name: "High Risk Cluster", Category: "fraud"},
				RiskReason:             "High risk due to fraud",
				AddressType:            "WALLET",
				AddressIdentifications: []interface{}{},
				Exposures: []Exposure{
					{Category: "fee", Value: 5678.90, ExposureType: "indirect", Direction: "outgoing"},
					{Category: "token smart contract", Value: 3456.78, ExposureType: "direct", Direction: "incoming"},
				},
				Triggers: []interface{}{},
			},
		},
	}

	realScreener.SetClient(m)
	time.Sleep(time.Second)

	apiClient, err := client.NewClient(s.metrics, fmt.Sprintf("http://localhost:%d", s.port))
	Nil(s.T(), err)

	// http://localhost:63575/v2/entities/0x123: true
	out, err := apiClient.ScreenAddress(s.GetTestContext(), "0x123")
	Nil(s.T(), err)
	True(s.T(), out)

	out, err = apiClient.ScreenAddress(s.GetTestContext(), "0x456")
	Nil(s.T(), err)
	True(s.T(), out)

	// http://localhost:63575/testrule/address/0x00: false
	out, err = apiClient.ScreenAddress(s.GetTestContext(), "0x00")
	Nil(s.T(), err)
	False(s.T(), out)

	// http://localhost:63575/testrule/address/0x00: false
	out, err = apiClient.ScreenAddress(s.GetTestContext(), "0x00")
	Nil(s.T(), err)
	False(s.T(), out)

	// now test crud screener
	// create a bunch
	statuses, err := blacklistTestWithOperation(s.T(), "create", apiClient, cfg)
	Equal(s.T(), len(statuses), 10)
	all(s.T(), statuses, func(status string) bool {
		return status == success
	})
	Nil(s.T(), err)

	// update a bunch
	statuses, err = blacklistTestWithOperation(s.T(), "update", apiClient, cfg)
	Equal(s.T(), len(statuses), 10)
	all(s.T(), statuses, func(status string) bool {
		return status == success
	})
	Nil(s.T(), err)

	// delete a bunch
	statuses, err = blacklistTestWithOperation(s.T(), "delete", apiClient, cfg)
	Equal(s.T(), len(statuses), 10)
	all(s.T(), statuses, func(status string) bool {
		return status == success
	})
	Nil(s.T(), err)

	// unauthorized, return on err so statuses will be only one
	cfg.AppSecret = "BAD"
	statuses, err = blacklistTestWithOperation(s.T(), "create", apiClient, cfg)
	all(s.T(), statuses, func(status string) bool {
		return status == "ERROR"
	})
	Equal(s.T(), len(statuses), 1)
	NotNil(s.T(), err)

	c := chainalysis.NewClient(s.metrics, []string{"Severe", "High"}, "key", "url")
	NotNil(s.T(), c)

}

func blacklistTestWithOperation(t *testing.T, operation string, apiClient client.ScreenerClient, cfg config.Config) (statuses []string, err error) {
	t.Helper()
	for range 10 {
		randomNumber, err := rand.Int(rand.Reader, big.NewInt(1000))
		if err != nil {
			return statuses, fmt.Errorf("error generating random number: %w", err)
		}

		var body client.BlackListBody

		if operation == "create" || operation == "update" {
			body = client.BlackListBody{
				Type: operation,
				ID:   fmt.Sprintf("unique-id-%d", randomNumber),
				Data: client.Data{
					Address: fmt.Sprintf("address-%d", randomNumber),
					Network: fmt.Sprintf("network-%d", randomNumber),
					Tag:     fmt.Sprintf("tag-%d", randomNumber),
					Remark:  "remark",
				},
			}
		} else {
			body = client.BlackListBody{
				Type: operation,
				ID:   fmt.Sprintf("unique-id-%d", randomNumber),
			}
		}
		status, err := apiClient.BlacklistAddress(context.Background(), cfg.AppSecret, cfg.AppID, body)
		statuses = append(statuses, status)
		if err != nil {
			return statuses, fmt.Errorf("error blacklisting address: %w", err)
		}
	}
	return statuses, nil
}

type mockClient struct {
	risks     []string
	entityMap map[string]*Entity
}

// ScreenAddress mocks the screen address method.
func (m mockClient) ScreenAddress(ctx context.Context, address string) (bool, error) {
	if m.entityMap == nil {
		return false, fmt.Errorf("no response map")
	}
	entity, ok := m.entityMap[address]
	if !ok {
		err := m.RegisterAddress(ctx, address)
		if err != nil {
			return false, fmt.Errorf("could not register address: %w", err)
		}
		entity = m.entityMap[address]
	}

	if slices.Contains(m.risks, entity.Risk) {
		return true, nil
	}

	return false, nil
}

// RegisterAddress mocks the register address method.
func (m mockClient) RegisterAddress(ctx context.Context, address string) error {
	m.entityMap[address] = &Entity{
		Address:                "0x1234abcdef1234abcdef1234abcdef1234abcd",
		Risk:                   "Critical",
		Cluster:                Cluster{Name: "Critical Risk Cluster", Category: "money laundering"},
		RiskReason:             "Involved in money laundering",
		AddressType:            "PRIVATE_WALLET",
		AddressIdentifications: []interface{}{},
		Exposures: []Exposure{
			{Category: "smart contract", Value: 9876.54, ExposureType: "indirect", Direction: "both_directions"},
			{Category: "stolen funds", Value: 1234.56, ExposureType: "direct", Direction: "both_directions"},
		},
		Triggers: []interface{}{},
	}
	return nil
}

var _ chainalysis.Client = mockClient{}

type Exposure struct {
	Category     string  `json:"category"`
	Value        float64 `json:"value"`
	ExposureType string  `json:"exposureType"`
	Direction    string  `json:"direction"`
}

type Cluster struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Entity struct {
	Address                string        `json:"address"`
	Risk                   string        `json:"risk"`
	Cluster                Cluster       `json:"cluster"`
	RiskReason             string        `json:"riskReason"`
	AddressType            string        `json:"addressType"`
	AddressIdentifications []interface{} `json:"addressIdentifications"`
	Exposures              []Exposure    `json:"exposures"`
	Triggers               []interface{} `json:"triggers"`
}

func all(t *testing.T, statuses []string, f func(string) bool) {
	t.Helper()
	for _, status := range statuses {
		if !f(status) {
			t.Fail()
		}
	}
}

const success = "OK"
