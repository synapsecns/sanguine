package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/core/retry"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestListenAndServe(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	s := &Server{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	port, err := freeport.GetFreePort()
	assert.NoError(t, err)

	go func() {
		err := s.ListenAndServe(ctx, fmt.Sprintf(":%d", port), r)
		assert.NoError(t, err)
	}()

	url := fmt.Sprintf("http://localhost:%d/ping", port)
	// Give some time for server to start
	err = retry.WithBackoff(ctx, func(ctx context.Context) error {
		// Make a request to test if server is running
		resp, err := http.Get(url)
		if err != nil {
			return errors.New("server has not yet started")
		}
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		assert.NoError(t, err)
		assert.Equal(t, "pong", string(body))

		// Wait for the context to cancel
		<-ctx.Done()
		return nil
	})
	assert.NoError(t, err)

	// Make a request to test if server is running
	resp, err := http.Get(url)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, "pong", string(body))

	// Wait for the context to cancel
	<-ctx.Done()

	// make sure cancellation triggers were processed and server is closed
	resp, err = http.Get("http://localhost:9090/ping")
	assert.NotNil(t, err)
}
