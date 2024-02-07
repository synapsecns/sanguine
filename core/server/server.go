package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server is a way to access a server listener.
type Server struct {
	listener net.Listener
}

// ListenAndServe provides a way to listen and serve a server with context. The server will terminate if the context is canceled.
func (s Server) ListenAndServe(ctx context.Context, port string, handler *gin.Engine) error {
	var err error
	var lc net.ListenConfig
	s.listener, err = lc.Listen(ctx, "tcp", port)
	if err != nil {
		return fmt.Errorf("could not listen on %s: %w", port, err)
	}

	go func() {
		//nolint:gosec
		// TODO: consider setting timeouts here:  https://ieftimov.com/posts/make-resilient-golang-net-http-servers-using-timeouts-deadlines-context-cancellation/
		err := http.Serve(s.listener, handler)
		if err != nil {
			logger.Errorf(fmt.Sprintf("rpc server got error: %v", err))
		}
	}()

	<-ctx.Done()
	return nil
}
