package forker

import (
	"bufio"
	"context"
	_ "embed"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log"
	"github.com/phayes/freeport"
	"golang.org/x/sync/errgroup"
	"io"
	"k8s.io/apimachinery/pkg/util/wait"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

//go:embed anvilbin/anvil_darwin_bin
var anvilDarwinBin []byte

//go:embed anvilbin/anvil_linux_bin
var anvilLinuxBin []byte

// TODO: add comment
func Fork(ctx context.Context, rpcURL string, chainID uint64, clientFunc func(client *ethclient.Client)) error {
	// TODO: embed binary/copy or figure out a way to download
	processPort, err := freeport.GetFreePort()
	if err != nil {
		return fmt.Errorf("failed to get free port: %w", err)
	}

	var filePayload []byte

	osType := runtime.GOOS
	switch osType {
	case "darwin":
		filePayload = anvilDarwinBin
	case "linux":
		filePayload = anvilLinuxBin
	default:
		return fmt.Errorf("unsupported os %s", osType)
	}

	err = os.WriteFile("./anvil_embedded", filePayload, 0755)
	if err != nil {
		return fmt.Errorf("failed to write anvil_embedded: %w", err)
	}
	defer os.Remove("./anvil_embedded")

	cmd := exec.Command("./anvil_embedded", "--fork-url", rpcURL, "--chain-id", strconv.Itoa(int(chainID)), "--port", strconv.Itoa(processPort))
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
		defer cmd.Process.Kill()
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
