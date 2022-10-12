package forker

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log"
	"github.com/phayes/freeport"
	"golang.org/x/sync/errgroup"
	"io"
	"k8s.io/apimachinery/pkg/util/wait"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// TODO: add comment
func Fork(ctx context.Context, rpcURL string, chainID uint64, clientFunc func(client *ethclient.Client)) error {
	// TODO: embed binary/copy or figure out a way to download
	processPort, err := freeport.GetFreePort()
	if err != nil {
		return fmt.Errorf("failed to get free port: %w", err)
	}

	_ = processPort

	cmd := exec.Command("/Users/jake/.foundry/bin/anvil", "--fork-url", rpcURL, "--chain-id", strconv.Itoa(int(chainID)), "--port", strconv.Itoa(processPort))
	cmd.Env = os.Environ()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to get stderr pipe: %w", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}

	go outputLogs(stderr, log.LevelError)
	go outputLogs(stdout, log.LevelInfo)

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("could not start forked node: %w", err)
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		err = cmd.Wait()
		if err != nil {
			return fmt.Errorf("could not start forked node: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		startupCtx, cancel := context.WithCancel(ctx)
		defer cancel()

		var client *ethclient.Client
		wait.UntilWithContext(startupCtx, func(ctx context.Context) {
			client, err = ethclient.DialContext(ctx, fmt.Sprintf("http://localhost:%d", processPort))
			if err != nil {
				logger.Errorf("failed to dial forked node: %s", err)
				return
			}

			rpcChainID, err := client.ChainID(ctx)
			if err != nil {
				return
			}

			if rpcChainID.Uint64() == chainID {
				cancel()
			}
		}, time.Millisecond*10)

		clientFunc(client)
		return nil
	})

	err = g.Wait()
	if err != nil {
		logger.Errorf("waitgroup exited: %s", err)
	}

	return nil
}

func outputLogs(reader io.ReadCloser, logLevel log.LogLevel) {
	s := bufio.NewScanner(reader)

	for s.Scan() {
		// TODO: make this exhautive
		switch logLevel {
		case log.LevelInfo:
			logger.Infof("%s", s.Text())
		case log.LevelError:
			logger.Errorf("%s", s.Text())
		}
	}
}
