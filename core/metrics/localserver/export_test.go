package localserver

import (
	"context"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

// UIResource exports ui resource for testing.
type UIResource interface {
	GetResource() *dockertest.Resource
	GetUIURL() string
}

func (u uiResource) GetResource() *dockertest.Resource {
	return u.Resource
}

func (u uiResource) GetUIURL() string {
	return u.uiURL
}

var _ UIResource = &uiResource{}

// TestJaeger exports jaeger for testing.
type TestJaeger interface {
	StartJaegerServer(ctx context.Context) UIResource
	StartPyroscopeServer(ctx context.Context) UIResource
	GetRunID() string
	GetPool() *dockertest.Pool
}

// StartTestServer starts a local jaeger server for testing.
// this will return a TestJaeger that can be used to start other servers.
func StartTestServer(ctx context.Context, tb testing.TB) TestJaeger {
	tj := startServer(ctx, tb)
	return &exportedJaeger{
		tj: tj,
	}
}

type exportedJaeger struct {
	tj *testJaeger
}

func NewTestJaeger(tb testing.TB) TestJaeger {
	logDir := filet.TmpDir(tb, "")
	pool, err := dockertest.NewPool("")
	assert.NoError(tb, err)

	return &exportedJaeger{
		tj: &testJaeger{
			tb:     tb,
			logDir: logDir,
			pool:   pool,
			runID:  gofakeit.UUID(),
		},
	}
}

func (e *exportedJaeger) GetPool() *dockertest.Pool {
	return e.tj.pool
}

func (e *exportedJaeger) GetRunID() string {
	return e.tj.runID
}

func (e *exportedJaeger) StartJaegerServer(ctx context.Context) UIResource {
	return e.tj.StartJaegerServer(ctx)
}

func (e *exportedJaeger) StartPyroscopeServer(ctx context.Context) UIResource {
	return e.tj.StartPyroscopeServer(ctx)
}

const RunIDLabel = runIDLabel
const AppLabel = appLabel

// ogDebugLocal is the original debugLocal value.
var ogDebugLocal = debugLocal

func GetOriginalDebugLocal() bool {
	return ogDebugLocal
}

func SetDebugLocal(enabled bool) {
	debugLocal = enabled
}
