package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	listener net.Listener
}

func (s Server) ListenAndServe(ctx context.Context, port string, handler *gin.Engine) error {
	var err error
	var lc net.ListenConfig
	s.listener, err = lc.Listen(ctx, "tcp", port)
	if err != nil {
		return fmt.Errorf("could not listen on %s: %w", port, err)
	}

	go func() {
		err := http.Serve(s.listener, handler)
		if err != nil {
			logger.Errorf(fmt.Sprintf("rpc server got error: %v", err))
		}
	}()

	select {
	case <-ctx.Done():
		return nil
	}
}
