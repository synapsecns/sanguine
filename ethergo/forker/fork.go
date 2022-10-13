package forker

import (
	"bufio"
	"context"
	// embed is used for anvil binaries
	_ "embed"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log"
	"github.com/phayes/freeport"
	"golang.org/x/sync/errgroup"
	"io"
	"k8s.io/apimachinery/pkg/util/wait"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

// delete me
//
//go:embed anvilbin/anvil_darwin_amd64_bin
var anvilDarwinAmd64Bin []byte

//go:embed anvilbin/anvil_darwin_arm64_bin
var anvilDarwinArm64Bin []byte

//go:embed anvilbin/anvil_linux_amd64_bin
var anvilLinuxAmd64Bin []byte

// Fork will fork the given evm blockchain for the purposes of testing.
func Fork(ctx context.Context, rpcURL string, chainID uint64, clientFunc func(client *ethclient.Client)) error {
	// TODO: embed binary/copy or figure out a way to download
	processPort, err := freeport.GetFreePort()
	if err != nil {
		return fmt.Errorf("failed to get free port: %w", err)
	}

	var filePayload []byte

	osType := runtime.GOOS
	archType := runtime.GOARCH
	osTypeWithArch := osType + "_" + archType
	switch osTypeWithArch {
	case "darwin_amd64":
		filePayload = anvilDarwinAmd64Bin
	case "darwin_arm64":
		filePayload = anvilDarwinArm64Bin
	case "linux_amd64":
		filePayload = anvilLinuxAmd64Bin
	default:
		return fmt.Errorf("unsupported arch %s for os %s", archType, osType)
	}

	now := time.Now().UnixNano()
	rand.Seed(now)

	embeddedFileName := "./anvil_embedded_" + strconv.FormatInt(time.Now().UnixNano(), 16) + "_" + strconv.FormatInt(int64(rand.Intn(math.MaxInt32)), 16)
	err = os.WriteFile(embeddedFileName, filePayload, 0755) //nolint: gosec
	if err != nil {
		return fmt.Errorf("failed to write %s: %w", embeddedFileName, err)
	}
	defer os.Remove(embeddedFileName) //nolint: errcheck

	cmd := exec.Command(embeddedFileName, "--fork-url", rpcURL, "--chain-id", strconv.Itoa(int(chainID)), "--port", strconv.Itoa(processPort)) //nolint: gosec
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
		defer cmd.Process.Kill() //nolint: errcheck
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
