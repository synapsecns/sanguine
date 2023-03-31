package localmetrics_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
)

func (l *LocalServerSuite) TestPyroscopeEndpointSet() {
	uiEndpoint := gofakeit.URL()
	l.T().Setenv(internal.PyroscopeEndpoint, uiEndpoint)

	jaegerTest := localmetrics.NewTestJaeger(l.T())
	retVal := jaegerTest.StartPyroscopeServer(l.GetTestContext())
	l.Require().NotNil(retVal)
	l.Require().Nil(retVal.GetResource())
	l.Require().Equal(retVal.GetUIURL(), uiEndpoint)

	l.Require().False(l.hasContainerWithRunID(jaegerTest.GetRunID()), "pyroscope should not have started")
}
