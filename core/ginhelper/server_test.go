package ginhelper_test

import (
	"github.com/google/uuid"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/temoto/robotstxt"
	_ "github.com/temoto/robotstxt"
	"net/http"
)

func (g *GinHelperSuite) TestRobots() {
	g.runOnAllURLs(func(url string) {
		getRequest, err := http.NewRequestWithContext(g.GetTestContext(), http.MethodGet, url+ginhelper.RobotsTxt, nil)
		Nil(g.T(), err)

		resp, err := http.DefaultClient.Do(getRequest)
		Nil(g.T(), err)
		defer func() {
			_ = resp.Body.Close()
		}()
		_, err = robotstxt.FromResponse(resp)
		Nil(g.T(), err)
	})
}

func (g *GinHelperSuite) TestHealth() {
	g.runOnAllURLs(func(url string) {
		getRequest, err := http.NewRequestWithContext(g.GetTestContext(), http.MethodGet, url+ginhelper.HealthCheck, nil)
		Nil(g.T(), err)

		resp, err := http.DefaultClient.Do(getRequest)
		defer func() {
			_ = resp.Body.Close()
		}()
		Nil(g.T(), err)
		Equal(g.T(), http.StatusOK, resp.StatusCode)
	})
}

func (g *GinHelperSuite) TestRequestID() {
	g.runOnAllURLs(func(url string) {
		getRequest, err := http.NewRequestWithContext(g.GetTestContext(), http.MethodGet, url, nil)
		Nil(g.T(), err)

		resp, err := http.DefaultClient.Do(getRequest)
		defer func() {
			_ = resp.Body.Close()
		}()
		Nil(g.T(), err)

		res := resp.Header.Get(ginhelper.RequestIDHeader)
		_, err = uuid.Parse(res)
		Nil(g.T(), err)
	})
}

// AssertContextWithFallback asserts that the context is set with the fallback value.
// this is required for using opentracing. Please use caution and ask @trajan0x before disabling.
func (g *GinHelperSuite) AssertContextWithFallback() {
	engine := ginhelper.New(g.logger)
	g.True(engine.ContextWithFallback)
}
