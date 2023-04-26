package metrics

import (
	"context"
	"fmt"
	"github.com/ImVexed/fasturl"
	"github.com/ethereum/go-ethereum/rpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
	"sync"
)

// logOnce ensures that log messages related to unsuppported clients are used only once.
var logOnce = sync.Once{}

const httpScheme = "http"
const httpsScheme = "https"

// RPCClient is a wrapper around rpc.Client that adds metrics/tracing.
func RPCClient(ctx context.Context, metrics Handler, url string, client *http.Client, opts ...otelhttp.Option) (*rpc.Client, error) {
	u, err := fasturl.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("could not parse url: %w", err)
	}

	switch u.Protocol {
	case httpScheme, httpsScheme:
		metrics.ConfigureHTTPClient(client, opts...)

		rpcclient, err := rpc.DialHTTPWithClient(url, client)
		if err != nil {
			return nil, fmt.Errorf("could not dial http: %w", err)
		}

		return rpcclient, nil
	default:
		logOnce.Do(func() {
			logger.Warnf("unsupported protocol: %s: only %s and %s are supported for metrics, future warnings will be surprssed", u.Protocol, httpScheme, httpsScheme)
		})
		//nolint: wrapcheck
		return rpc.DialContext(ctx, url)
	}
}
