package internal

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/integralist/go-findroot/find"
	"github.com/synapsecns/sanguine/core/retry"
	"os"
	"os/exec"
	"path"
)

// Up brings the devnet up.
func Up(ctx context.Context) error {
	// TODO: figure out if we want to allow this to be passed in w/o using igt
	repoRoot, err := find.Repo()
	if err != nil {
		return fmt.Errorf("failed to find repo root: %w", err)
	}

	devnetPath := path.Join(repoRoot.Path, "docker", "devnet")

	log.Info("Running devnet.")

	cmd := exec.Command("docker", "compose", "up", "-d")
	cmd.Dir = devnetPath
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	retry.WithBackoff(ctx, func(ctx context.Context) error {

	}, retry.WithMaxAttempts())

	return nil
}
