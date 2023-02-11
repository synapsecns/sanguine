package kmsmock

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/smithy-go/logging"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/nsmithuk/local-kms/src/config"
	"github.com/nsmithuk/local-kms/src/data"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/awssigner/kmsmock/internal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"k8s.io/apimachinery/pkg/util/wait"
)

// globMux is a global mutex uate configs. Since local-kms uses a global configuration store,
// and we want to be able to run multiple services at once, on each request we acquire a lock
// hot swap the config, handle the request, and then swap the config back.
var globMux sync.Mutex

// MockKMSService is a mocked kms service.
type MockKMSService struct {
	// awsRegion is the region
	awsRegion string
	// awsAccountID is a mock aws account id
	awsAccountID string
	// databasePath is the database path
	databasePath string
	// url gets the server url
	url string
	// database is the leveldb for the mocked kms config
	database *data.Database
	// tb contains the testing (or benchmarking) object
	tb testing.TB
	// ctx is the context
	//nolint: containedctx
	ctx context.Context
}

// NewMockKMS creates a mocked kms server locally. The context must remain alive for the life of the kms service.
func NewMockKMS(ctx context.Context, tb testing.TB) *MockKMSService {
	tb.Helper()

	kmsService := &MockKMSService{}

	kmsService.awsRegion = "us-west-2"
	kmsService.awsAccountID = gofakeit.AchAccount()
	kmsService.databasePath = filet.TmpDir(tb, "")
	kmsService.tb = tb
	kmsService.ctx = ctx

	kmsService.startDB(ctx)
	kmsService.startServer(ctx)
	// wait for the server to boot before returning
	kmsService.waitForServerBoot(ctx)
	return kmsService
}

// Client retrieves a kms client connected to the mock url.
func (k *MockKMSService) Client() *kms.Client {
	httpLogger, err := zap.NewStdLogAt(logger.Desugar(), zapcore.InfoLevel)
	assert.Nil(k.tb, err)

	return kms.New(kms.Options{
		APIOptions:    nil,
		ClientLogMode: aws.LogSigning,
		EndpointOptions: kms.EndpointResolverOptions{
			DisableHTTPS: true,
		},
		EndpointResolver: kms.EndpointResolverFromURL(k.url),
		Logger:           logging.StandardLogger{Logger: httpLogger},
		Region:           k.awsRegion,
		Retryer:          aws.NopRetryer{},
		HTTPClient:       &http.Client{},
	})
}

// startServer starts the server and terminates it when the context is done.
func (k *MockKMSService) startServer(ctx context.Context) {
	// start the serveMux and terminate it when the db config is done
	hostname := fmt.Sprintf("localhost:%d", freeport.GetPort())

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// on every request set the global config & run in lock
		k.runWithConfig(func() {
			internal.HandleRequest(writer, request, k.database)
		})
	})

	var lc net.ListenConfig
	l, err := lc.Listen(ctx, "tcp", hostname)
	assert.Nil(k.tb, err)

	// start the server and terminate it on end
	//nolint: gosec
	server := &http.Server{Addr: hostname, Handler: serveMux, BaseContext: func(listener net.Listener) context.Context {
		return ctx
	}}

	go func() {
		err := server.Serve(l)
		assert.Nil(k.tb, err)
	}()

	//nolint:gosec
	// set the server url
	k.url = fmt.Sprintf("http://%s", hostname)
}

const bootTimeout = time.Second * 10

// waitForServerBoot waits until the server is finished booting. If this isn't done in 10 seconds the test fails.
func (k *MockKMSService) waitForServerBoot(parentCtx context.Context) {
	ctx, cancel := context.WithTimeout(parentCtx, bootTimeout)
	defer cancel()

	// success is set to true on a successful http request.
	// If this is false after the context is expired, the test should fail
	success := false

	var err error

	serverStartedCtx, cancel := context.WithCancel(ctx)
	wait.UntilWithContext(serverStartedCtx, func(ctx context.Context) {
		client := http.Client{}

		var req *http.Request
		req, err = http.NewRequestWithContext(serverStartedCtx, http.MethodGet, k.url, nil)
		assert.Nil(k.tb, err)

		//nolint: bodyclose
		_, err = client.Do(req)
		if err == nil {
			cancel()
			success = true
		}
	}, time.Millisecond)

	if !success {
		// add an error if it doesn't exist since success should never be false
		if err == nil {
			err = fmt.Errorf("successful request never finished")
		}
		k.tb.Errorf("expected server to start successfully, got error: %v", err)
	}
}

// startDB starts the database and terminates it when the context is done. The database path
// must be set before this is called.
func (k *MockKMSService) startDB(ctx context.Context) {
	// start the db and terminate it when the context is done
	k.database = data.NewDatabase(k.databasePath)

	go func(ctx context.Context) {
		<-ctx.Done()
		k.database.Close()
	}(ctx)
}

// runWithConfig runs a function after setting the config. The globMux is locked for the life of the
// config setter.
func (k *MockKMSService) runWithConfig(runFunc func()) {
	globMux.Lock()
	defer globMux.Unlock()

	config.AWSRegion = k.awsRegion
	config.AWSAccountId = k.awsAccountID
	config.DatabasePath = k.databasePath
	runFunc()
}
