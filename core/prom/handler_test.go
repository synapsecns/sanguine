package prom_test

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/grafana-tools/sdk"
	"github.com/ipfs/go-log"
	"github.com/phayes/freeport"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/route"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"
	api_v1 "github.com/prometheus/pushgateway/api/v1"
	"github.com/prometheus/pushgateway/asset"
	"github.com/prometheus/pushgateway/handler"
	"github.com/prometheus/pushgateway/storage"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/prom"
	"k8s.io/apimachinery/pkg/util/wait"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var logger = log.Logger("synapse-metrics-test")

// make sure push gateway metrics are correctly pushed.
func (m MetricsSuite) TestPushGateway() {
	gatewayURL := NewMockGateway(m.GetTestContext(), m.T())
	metricHandler := prom.NewMetricHandler("", &prom.Config{
		PushGateway: gatewayURL,
	})
	err := metricHandler.Start(m.GetTestContext())
	Nil(m.T(), err)
}

// TestBoardHandler makes sure the dashboard returns a json dashboard in http.
func (m MetricsSuite) TestBoardHandler() {
	board := sdk.NewBoard("hi")
	req, err := http.NewRequestWithContext(m.GetTestContext(), http.MethodGet, "/doesntmatter", nil)
	Nil(m.T(), err)

	response := httptest.NewRecorder()
	server := prom.CreateBoardJSONHandler(board)
	server.ServeHTTP(response, req)

	Equal(m.T(), response.Code, http.StatusOK)
	True(m.T(), govalidator.IsJSON(response.Body.String()))
}

func NewMockGateway(ctx context.Context, tb testing.TB) string {
	tb.Helper()

	mockRegistery := prometheus.NewRegistry()
	err := mockRegistery.Register(version.NewCollector("pushgateway"))
	Nil(tb, err)

	promLogger := prom.NewPromLogger(logger)

	ms := storage.NewDiskMetricStore("", time.Minute*5, mockRegistery, promLogger)

	buildInfo := map[string]string{
		"version":   version.Version,
		"revision":  version.Revision,
		"branch":    version.Branch,
		"buildUser": version.BuildUser,
		"buildDate": version.BuildDate,
		"goVersion": version.GoVersion,
	}

	r := route.New()

	mux := http.NewServeMux()
	mux.Handle("/", r)

	apiv1 := api_v1.New(promLogger, ms, map[string]string{}, buildInfo)
	r.Get("/-/healthy", handler.Healthy(ms).ServeHTTP)
	r.Get("/-/ready", handler.Ready(ms).ServeHTTP)
	r.Get("/static/*filepath", handler.Static(asset.Assets, "").ServeHTTP)

	pushAPIPath := "/metrics"

	for _, suffix := range []string{"", handler.Base64Suffix} {
		jobBase64Encoded := suffix == handler.Base64Suffix
		r.Put(pushAPIPath+"/job"+suffix+"/:job/*labels", handler.Push(ms, true, false, jobBase64Encoded, promLogger))
		r.Post(pushAPIPath+"/job"+suffix+"/:job/*labels", handler.Push(ms, false, false, jobBase64Encoded, promLogger))
		r.Del(pushAPIPath+"/job"+suffix+"/:job/*labels", handler.Delete(ms, jobBase64Encoded, promLogger))
		r.Put(pushAPIPath+"/job"+suffix+"/:job", handler.Push(ms, true, false, jobBase64Encoded, promLogger))
		r.Post(pushAPIPath+"/job"+suffix+"/:job", handler.Push(ms, false, false, jobBase64Encoded, promLogger))
		r.Del(pushAPIPath+"/job"+suffix+"/:job", handler.Delete(ms, jobBase64Encoded, promLogger))
	}
	apiv1.Register(r)

	hostname := fmt.Sprintf("localhost:%d", freeport.GetPort())
	var lc net.ListenConfig
	l, err := lc.Listen(ctx, "tcp", hostname)
	Nil(tb, err)

	go func() {
		//nolint: gosec
		err = web.Serve(l, &http.Server{Addr: hostname, Handler: mux, BaseContext: func(listener net.Listener) context.Context {
			return ctx
		}}, "", promLogger)
	}()

	url := fmt.Sprintf("http://%s", hostname)
	// wait for server start
	serverStartedCtx, cancel := context.WithCancel(ctx)
	wait.UntilWithContext(serverStartedCtx, func(ctx context.Context) {
		client := http.Client{}
		req, err := http.NewRequestWithContext(serverStartedCtx, http.MethodGet, url, nil)
		Nil(tb, err)

		//nolint: bodyclose
		_, err = client.Do(req)
		if err == nil {
			cancel()
		}
	}, time.Millisecond)

	return url
}
