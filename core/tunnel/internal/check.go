package internal

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

// CheckPathFunc is called when the checkPath is hit.
type CheckPathFunc func()

// VerifiableProxy is a proxy that has an overidable path. The path calls checkFunc when hit.
// this is used to make sure a proxy has started.
func VerifiableProxy(ctx context.Context, backendURL, checkPath string, listener net.Listener, pathFunc CheckPathFunc) (_ <-chan error) {
	errChan := make(chan error, 1)

	proxy, err := newCheckableProxy(backendURL, checkPath, pathFunc)
	if err != nil {
		select {
		case <-ctx.Done():
			return errChan
		case errChan <- err:
			return errChan
		}
	}

	go func() {
		err = http.Serve(listener, proxy)
		if err != nil {
			select {
			case errChan <- fmt.Errorf("could not serve ngrok: %w", err):
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return errChan
}
