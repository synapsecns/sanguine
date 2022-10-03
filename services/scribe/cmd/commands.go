package cmd

import (
	// used to embed markdown.
	_ "embed"
	"fmt"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hashicorp/consul/sdk/freeport"
	"github.com/jftuga/termsize"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/scribe/api"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/urfave/cli/v2"
	"time"
)

//go:embed cmd.md
var help string

// infoComand gets info about using the scribe service.
var infoCommand = &cli.Command{
	Name:        "info",
	Description: "learn how to use scribe cli",
	Action: func(c *cli.Context) error {
		fmt.Println(string(markdown.Render(help, termsize.Width(), 6)))
		return nil
	},
}

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /Users/synapsecns/config.yaml",
	TakesFile: true,
	Required:  true,
}

var portFlag = &cli.UintFlag{
	Name:  "port",
	Usage: "--port 5121",
	Value: 0,
}

var grpcPortFlag = &cli.UintFlag{
	Name:  "grpc-port",
	Usage: "--port 5121",
	Value: 0,
}

var dbFlag = &cli.StringFlag{
	Name:     "db",
	Usage:    "--db <sqlite> or <mysql>",
	Value:    "sqlite",
	Required: true,
}

var pathFlag = &cli.StringFlag{
	Name:     "path",
	Usage:    "--path <path/to/database> or <database url>",
	Value:    "",
	Required: true,
}

var backfillCommand = &cli.Command{
	Name:        "backfill",
	Description: "backfills up to a block and then halts",
	Flags:       []cli.Flag{configFlag, dbFlag, pathFlag},
	Action: func(c *cli.Context) error {
		decodeConfig, err := config.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not decode config: %w", err)

		}

		db, err := api.InitDB(c.Context, c.String(dbFlag.Name), c.String(pathFlag.Name))
		if err != nil {
			return fmt.Errorf("could not initialize database: %w", err)
		}

		clients := make(map[uint32]backfill.ScribeBackend)
		// TODO: should be resistant to errors on startup from a single chain
		for i := 0; i < len(decodeConfig.Chains); i++ {
			backendClient, errA := ethclient.DialContext(c.Context, decodeConfig.Chains[i].RPCUrl)
			_, errB := backendClient.NetworkID(c.Context)

			// errA: Completely invalid RPC URL, fatal connection
			// errB: URL valid, but cannot make valid connection
			if errA != nil || errB != nil {
				fmt.Println("The RPCurl", decodeConfig.Chains[i].RPCUrl, "is unreachable, skipping chain id: ", decodeConfig.Chains[i].ChainID)
				// remove chain from clients
				delete(clients, decodeConfig.Chains[i].ChainID)
				// check if resulting list of clients is empty (none of the RPC URLS had a good connection)
				if len(clients) == 0 && i+1 == len(decodeConfig.Chains) {
					return fmt.Errorf("no rpc url connection successful - %s", "clients returned a zero length array.")
				}
				// remove chain from decoded chain config
				decodeConfig.Chains = append(decodeConfig.Chains[:i], decodeConfig.Chains[i+1:]...)
				// make sure to iterate over everything
				i--
				continue
			}
			clients[decodeConfig.Chains[i].ChainID] = backendClient
		}
		scribeBackfiller, err := backfill.NewScribeBackfiller(db, clients, decodeConfig)
		if err != nil {
			return fmt.Errorf("could not create scribe backfiller: %w", err)
		}

		// amount of retries for attempting backfill.
		attempts := 1000

		// initial amount of backoff for each retry
		sleep := time.Second
		for i := 0; i < attempts; i++ {
			if i > 0 {
				fmt.Println("[RETRY] retrying after error:", err, "\nNet number of retries: ", i+1)
				time.Sleep(sleep)
				// increase sleep time by *2
				sleep *= 2
			}
			err = scribeBackfiller.Backfill(c.Context)
			if err == nil {
				// backfill successful
				return nil
			}
		}
		fmt.Println("After", attempts, "attempts, backfilling failed. Last error:", err)
		return fmt.Errorf("could not backfill backfiller: %w", err)
	},
}

var serverCommand = &cli.Command{
	Name:        "server",
	Description: "starts a graphql server",
	Flags:       []cli.Flag{portFlag, dbFlag, pathFlag},
	Action: func(c *cli.Context) error {
		err := api.Start(c.Context, api.Config{
			HTTPPort: uint16(c.Uint(portFlag.Name)),
			Database: c.String(dbFlag.Name),
			Path:     c.String(pathFlag.Name),
			GRPCPort: uint16(c.Uint(grpcPortFlag.Name)),
		})
		if err != nil {
			return fmt.Errorf("could not start server: %w", err)
		}

		return nil
	},
}

var omniRPCFlag = &cli.StringFlag{
	Name:     "omnirpc",
	Usage:    "--omnirpc https://omnirpc.url",
	Required: true,
}

var deploymentsPath = &cli.StringFlag{
	Name:     "deployments-dir",
	Usage:    "--deployments-dir /path/to/contracts/deployments",
	Required: true,
}

var confirmationsFlag = &cli.UintFlag{
	Name:  "confirmations",
	Value: 50,
}

var outputPathFlag = &cli.StringFlag{
	Name:      "output-path",
	TakesFile: true,
}

var skippedChainIdsFlag = &cli.IntSliceFlag{
	Name:  "skipped-chains",
	Usage: "--skipped-chains 1,2,3",
	// skip common testnets by default
	Value: cli.NewIntSlice(5, 335, 43113, 1666700000),
}

var generateCommand = &cli.Command{
	Name:        "generate",
	Description: "generates a config from a hardhat deployment folder",
	Flags: []cli.Flag{
		omniRPCFlag,
		deploymentsPath,
		confirmationsFlag,
		skippedChainIdsFlag,
		outputPathFlag,
	},
	Action: func(c *cli.Context) error {
		//nolint: wrapcheck
		return config.GenerateConfig(c.Context, c.String(omniRPCFlag.Name), core.ExpandOrReturnPath(c.String(deploymentsPath.Name)),
			uint32(c.Uint(confirmationsFlag.Name)), core.ExpandOrReturnPath(c.String(outputPathFlag.Name)), c.IntSlice(skippedChainIdsFlag.Name), config.DefaultClientGenerator)
	},
}

func init() {
	ports, err := freeport.Take(1)
	if len(ports) > 0 && err != nil {
		portFlag.Value = uint(ports[0])
	}
}
