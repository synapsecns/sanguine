package internal_test

import (
	"context"
	"github.com/synapsecns/sanguine/contrib/devnet/internal"
	"testing"
	"time"
)

func TestInternal(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	err := internal.Up(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
