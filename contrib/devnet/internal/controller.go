package internal

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/integralist/go-findroot/find"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"golang.org/x/sync/errgroup"
	"os"
	"os/exec"
	"path"
	"time"
)

const build = false

// Up brings the devnet up.
func Up(ctx context.Context, metricHandler metrics.Handler) error {
	// TODO: figure out if we want to allow this to be passed in w/o using igt
	repoRoot, err := find.Repo()
	if err != nil {
		return fmt.Errorf("failed to find repo root: %w", err)
	}

	devnetPath := path.Join(repoRoot.Path, "docker", "devnet")

	log.Info("Running devnet.")

	args := []string{"docker", "compose", "up", "-d"}
	if build {
		args = append(args, "--build")
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = devnetPath
	cmd.Env = append(os.Environ(), fmt.Sprintf("PWD=%s", devnetPath))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	err = checkLiveness(ctx, metricHandler)
	if err != nil {
		return fmt.Errorf("failed to check liveness: %w", err)
	}

	return nil
}

func checkLiveness(ctx context.Context, metricHandler metrics.Handler) error {
	log.Info("Checking liveness.")

	err := retry.WithBackoff(ctx, func(ctx context.Context) error {
		rpcClient := omnirpcClient.NewOmnirpcClient("http://localhost:9001", metricHandler)

		chainIDs, err := rpcClient.GetChainIDs(ctx)
		if err != nil {
			return fmt.Errorf("failed to get chain IDs: %w", err)
		}

		g, gctx := errgroup.WithContext(ctx)

		for i := range chainIDs {
			chainID := chainIDs[i] // capture func literal

			g.Go(func() error {
				chainClient, err := rpcClient.GetChainClient(gctx, chainID)
				if err != nil {
					return fmt.Errorf("failed to get chain client: %w", err)
				}

				_, err = chainClient.ChainID(gctx)
				if err != nil {
					return fmt.Errorf("failed to connect to chain %d: %w", chainID, err)
				}

				return nil
			})
		}

		err = g.Wait()
		if err != nil {
			return err
		}

		return nil
	}, retry.WithMaxTotalTime(20*time.Second))

	if err != nil {
		return fmt.Errorf("failed to connect to chains: %w", err)
	}

	return nil

}
