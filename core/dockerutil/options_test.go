package dockerutil_test

import (
	"context"
	"errors"
	"github.com/ory/dockertest/v3/docker"
	"github.com/synapsecns/sanguine/core/dockerutil"
	"testing"

	"github.com/ory/dockertest/v3"
)

func TestValidateOptions(t *testing.T) {
	ctx := context.Background()
	resource := &dockertest.Resource{
		Container: &docker.Container{
			ID: "test-container-id",
		},
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("failed to create docker pool: %s", err)
	}

	tests := []struct {
		name    string
		options []dockerutil.Option
		err     error
	}{
		{
			name: "success",
			options: []dockerutil.Option{
				dockerutil.WithContext(ctx),
				dockerutil.WithResource(resource),
				dockerutil.WithPool(pool),
				dockerutil.WithStdout(true),
			},
			err: nil,
		},
		{
			name: "missing context",
			options: []dockerutil.Option{
				dockerutil.WithResource(resource),
				dockerutil.WithPool(pool),
				dockerutil.WithStdout(true),
			},
			err: dockerutil.NewOptionError(dockerutil.NewContextError("context is not provided")),
		},
		{
			name: "missing resource",
			options: []dockerutil.Option{
				dockerutil.WithContext(ctx),
				dockerutil.WithPool(pool),
				dockerutil.WithStdout(true),
			},
			err: dockerutil.NewOptionError(dockerutil.NewResourceError("resource is not provided")),
		},
		{
			name: "missing pool",
			options: []dockerutil.Option{
				dockerutil.WithContext(ctx),
				dockerutil.WithResource(resource),
				dockerutil.WithStdout(true),
			},
			err: dockerutil.NewOptionError(dockerutil.NewPoolError("pool is not provided")),
		},
		{
			name: "missing callback",
			options: []dockerutil.Option{
				dockerutil.WithContext(ctx),
				dockerutil.WithResource(resource),
				dockerutil.WithPool(pool),
				dockerutil.WithStdout(true),
				dockerutil.WithCallback(nil),
			},
			err: dockerutil.NewOptionError(dockerutil.NewCallbackError("callback is not provided")),
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			err := dockerutil.ValidateOptions(tt.options...)
			if !errors.Is(err, tt.err) {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}
		})
	}
}
