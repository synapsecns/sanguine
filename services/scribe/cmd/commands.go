package cmd

import (
	"context"
	"time"

	// used to embed markdown.
	_ "embed"
	"fmt"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/hashicorp/consul/sdk/freeport"
	"github.com/jftuga/termsize"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/scribe/api"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/node"
	"github.com/urfave/cli/v2"
)

//go:embed cmd.md
var help string

var maxConfirmations = 3

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

func createScribeParameters(c *cli.Context) (eventDB db.EventDB, clients map[uint32][]backfill.ScribeBackend, scribeConfig config.Config, err error) {
	scribeConfig, err = config.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
	if err != nil {
		return nil, nil, scribeConfig, fmt.Errorf("could not decode config: %w", err)
	}

	eventDB, err = api.InitDB(c.Context, c.String(dbFlag.Name), c.String(pathFlag.Name))
	if err != nil {
		return nil, nil, scribeConfig, fmt.Errorf("could not initialize database: %w", err)
	}

	clients = make(map[uint32][]backfill.ScribeBackend)
	for _, client := range scribeConfig.Chains {
		for confNum := 1; confNum <= maxConfirmations; confNum++ {
			backendClient, err := backfill.DialBackend(c.Context, fmt.Sprintf("%s/%d/rpc/%d", scribeConfig.RPCURL, confNum, client.ChainID))
			if err != nil {
				return nil, nil, scribeConfig, fmt.Errorf("could not start client for %s", fmt.Sprintf("%s/1/rpc/%d", scribeConfig.RPCURL, client.ChainID))
			}
			clients[client.ChainID] = append(clients[client.ChainID], backendClient)
		}
	}

	return eventDB, clients, scribeConfig, nil
}

var backfillCommand = &cli.Command{
	Name:        "backfill",
	Description: "backfills up to a block and then halts",
	Flags:       []cli.Flag{configFlag, dbFlag, pathFlag},
	Action: func(c *cli.Context) error {
		db, clients, decodeConfig, err := createScribeParameters(c)
		if err != nil {
			return err
		}
		scribeBackfiller, err := backfill.NewScribeBackfiller(db, clients, decodeConfig)
		if err != nil {
			return fmt.Errorf("could not create scribe backfiller: %w", err)
		}

		// TODO delete once livefilling done
		ctx, cancel := context.WithTimeout(c.Context, time.Minute*5)
		cancelVar := cancel
		for {
			err = scribeBackfiller.Backfill(ctx)
			if err != nil {
				cancelVar()
				ctx, cancel = context.WithTimeout(c.Context, time.Minute*5)
				cancelVar = cancel
			}
		}
	},
}

var scribeCommand = &cli.Command{
	Name:        "scribe",
	Description: "scribe runs the scribe, livefilling across all specified chains",
	Flags:       []cli.Flag{configFlag, dbFlag, pathFlag},
	Action: func(c *cli.Context) error {
		db, clients, decodeConfig, err := createScribeParameters(c)
		if err != nil {
			return err
		}
		scribe, err := node.NewScribe(db, clients, decodeConfig)
		if err != nil {
			return fmt.Errorf("could not create scribe: %w", err)
		}
		err = scribe.Start(c.Context)
		if err != nil {
			return fmt.Errorf("could not start scribe: %w", err)
		}
		return nil
	},
}

var serverCommand = &cli.Command{
	Name:        "server",
	Description: "starts a graphql server",
	Flags:       []cli.Flag{portFlag, dbFlag, pathFlag, omniRPCFlag},
	Action: func(c *cli.Context) error {
		err := api.Start(c.Context, api.Config{
			HTTPPort:   uint16(c.Uint(portFlag.Name)),
			Database:   c.String(dbFlag.Name),
			Path:       c.String(pathFlag.Name),
			GRPCPort:   uint16(c.Uint(grpcPortFlag.Name)),
			OmniRPCURL: c.String(omniRPCFlag.Name),
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
