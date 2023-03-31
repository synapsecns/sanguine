package localserver_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/metrics/localserver"
	"github.com/synapsecns/sanguine/core/mocktesting"
)

// make sure jaeger checks work correctly.
func (l *LocalServerSuite) TestJaegerEndpointEmpty() {
	l.T().Setenv(internal.JAEGER_ENDPOINT, gofakeit.URL())
	l.T().Setenv(internal.JAEGER_UI_ENDPOINT, "")

	mockTester := mocktesting.NewMockTester("")
	jaegerTest := localserver.NewTestJaeger(mockTester)

	resource := jaegerTest.StartJaegerServer(l.GetTestContext())

	l.Require().Nil(resource, "jaeger should have errored")
	l.Require().True(mockTester.Failed(), "jaeger should have failed to start")

	l.Require().False(l.hasContainerWithRunID(jaegerTest.GetRunID()), "jaeger should not have started")
}

func (l *LocalServerSuite) TestJaegerEndpointsSet() {
	uiEndpoint := gofakeit.URL()
	l.T().Setenv(internal.JAEGER_ENDPOINT, gofakeit.URL())
	l.T().Setenv(internal.JAEGER_UI_ENDPOINT, uiEndpoint)

	jaegerTest := localserver.NewTestJaeger(l.T())
	retVal := jaegerTest.StartJaegerServer(l.GetTestContext())
	l.Require().NotNil(retVal)
	l.Require().Nil(retVal.GetResource())

	l.Require().Equal(retVal.GetUIURL(), uiEndpoint)

	l.Require().False(l.hasContainerWithRunID(jaegerTest.GetRunID()), "jaeger should not have started")
}
