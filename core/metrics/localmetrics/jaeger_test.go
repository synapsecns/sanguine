package localmetrics_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/mocktesting"
)

// make sure jaeger checks work correctly.
func (l *LocalServerSuite) TestJaegerEndpointEmpty() {
	l.T().Setenv(internal.JaegerEndpoint, gofakeit.URL())
	l.T().Setenv(internal.JaegerUIEndpoint, "")

	mockTester := mocktesting.NewMockTester("")
	jaegerTest := localmetrics.NewTestJaeger(mockTester)

	jaegerTest.StartJaegerServer(l.GetTestContext())

	l.Require().True(mockTester.Failed(), "jaeger should have failed to start")

	l.Require().False(l.hasContainerWithRunID(jaegerTest.GetRunID()), "jaeger should not have started")
}

func (l *LocalServerSuite) TestJaegerEndpointsSet() {
	uiEndpoint := gofakeit.URL()
	l.T().Setenv(internal.JaegerEndpoint, gofakeit.URL())
	l.T().Setenv(internal.JaegerUIEndpoint, uiEndpoint)

	jaegerTest := localmetrics.NewTestJaeger(l.T())
	retVal := jaegerTest.StartJaegerServer(l.GetTestContext())
	l.Require().NotNil(retVal)
	l.Require().Nil(retVal.GetResource())

	l.Require().Equal(retVal.GetUIURL(), uiEndpoint)

	l.Require().False(l.hasContainerWithRunID(jaegerTest.GetRunID()), "jaeger should not have started")
}
