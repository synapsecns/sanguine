package metrics

import (
	"context"
	"fmt"
	"github.com/ImVexed/fasturl"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"net/http"
	"sync"
)

// logOnce ensures that log messages related to unsuppported clients are used only once
var logOnce = sync.Once{}

const httpScheme = "http"
const httpsScheme = "https"

// EthClient is a wrapper around ethclient.Client that adds metrics/tracing
func EthClient(ctx context.Context, metrics Handler, url string) (*ethclient.Client, error) {
	u, err := fasturl.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("could not parse url: %w", err)
	}

	switch u.Protocol {
	case httpScheme, httpsScheme:
		client := new(http.Client)
		metrics.ConfigureHttpClient(client)
		rpcclient, err := rpc.DialHTTPWithClient(url, client)
		if err != nil {
			return nil, fmt.Errorf("could not dial http: %w", err)
		}
		return ethclient.NewClient(rpcclient), nil
	default:
		logOnce.Do(func() {
			logger.Warnf("unsupported protocol: %s: only %s and %s are supported for metrics, future warnings will be surprssed", u.Protocol, httpScheme, httpsScheme)
		})
		return ethclient.DialContext(ctx, url)
	}
}
