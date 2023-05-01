package internal

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func newCheckableProxy(backendURL string, checkPath string, checkFunc CheckPathFunc) (*checkableProxy, error) {
	parsedURL, err := url.Parse(backendURL)
	if err != nil {
		return nil, fmt.Errorf("could not parse backend URL: %w", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(parsedURL)

	return &checkableProxy{
		proxy:     proxy,
		checkPath: checkPath,
		checkFunc: checkFunc,
	}, nil
}

type checkableProxy struct {
	// proxy is the underlying proxy.
	proxy *httputil.ReverseProxy
	// checkPath is the checkPath
	checkPath string
	// checkFunc is called when checkPath is hit
	checkFunc CheckPathFunc
}

func (c checkableProxy) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if c.checkPath == request.URL.Path {
		c.checkFunc()
		writer.WriteHeader(http.StatusOK)
		return
	}

	c.proxy.ServeHTTP(writer, request)
}

var _ http.Handler = &checkableProxy{}
