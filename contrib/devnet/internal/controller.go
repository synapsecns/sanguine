package internal

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/log"
	"github.com/integralist/go-findroot/find"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"golang.org/x/sync/errgroup"
	"os"
	"os/exec"
	"path"
	"time"
)

const build = false

type devnetInfo struct {
	repoRoot     string
	configDir    string
	devnetPath   string
	contractsDir string
}

// Up brings the devnet up.
func Up(ctx context.Context, metricHandler metrics.Handler) error {
	// make the info
	info := devnetInfo{}

	// TODO: figure out if we want to allow this to be passed in w/o using igt
	repoRoot, err := find.Repo()
	if err != nil {
		return fmt.Errorf("failed to find repo root: %w", err)
	}
	info.repoRoot = repoRoot.Path

	info.configDir = path.Join(info.repoRoot, ".devnet")
	err = os.MkdirAll(info.configDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create config dir: %w", err)
	}

	// TODO: make sure the user has run yarn install

	info.devnetPath = path.Join(repoRoot.Path, "docker", "devnet")

	log.Info("Running devnet.")

	// spin up our chains in a docker env
	args := []string{"docker", "compose", "up", "-d"}
	if build {
		args = append(args, "--build")
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = info.devnetPath
	cmd.Env = append(os.Environ(), fmt.Sprintf("PWD=%s", info.devnetPath))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	err = checkLiveness(ctx, metricHandler)
	if err != nil {
		return fmt.Errorf("failed to check liveness: %w", err)
	}

	deployContracts(ctx, info)

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

// should match dockerfile
const mnemonic = "tag volcano eight thank tide danger coast health above argue embrace heavy"

// deploy some contracts
func deployContracts(ctx context.Context, info devnetInfo) error {
	// first up, let's define the deployer address in forge
	deployer, err := wallet.FromSeedPhrase(mnemonic, accounts.DefaultBaseDerivationPath)
	if err != nil {
		return fmt.Errorf("could not create wallet")
	}

	forgeEnv := append(os.Environ(), fmt.Sprintf("MESSAGING_DEPLOYER_PRIVATE_KEY=%s", deployer.PrivateKeyHex()))
	fmt.Println(deployer.PrivateKeyHex())
	fmt.Println(deployer.Address())
	_ = forgeEnv

	// do this next. TODO: a clean command should clean out everything in contracts-core/deployments/chain_a, contracts-core/deployments/chain_b, and contracts-core/deployments/chain_c
	//  forge script script/DeployMessaging003SynChain.s.sol  --ffi -f chain_a --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	//  forge script script/DeployMessaging003LightChain.s.sol  --ffi -f chain_b --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	//  forge script script/DeployMessaging003LightChain.s.sol  --ffi -f chain_c --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	//
	// forge script script/DeployClients003.s.sol  --ffi -f chain_a --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	// forge script script/DeployClients003.s.sol  --ffi -f chain_b --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	// forge script script/DeployClients003.s.sol  --ffi -f chain_c --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	return nil
}
