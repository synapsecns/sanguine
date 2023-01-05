package cmd

import (
	"context"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/jftuga/termsize"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/mysql"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/agents/executor/src"
	"github.com/synapsecns/sanguine/services/scribe/client"

	// used to embed markdown.
	_ "embed"
	"fmt"
	"github.com/synapsecns/sanguine/agents/agents/executor/config"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/core"
	"github.com/urfave/cli/v2"
	"os"
)

//go:embed cmd.md
var help string

// infoCommand gets info about using the executor agent.
var infoCommand = &cli.Command{
	Name:        "info",
	Description: "learn how to use executor cli",
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

var scribePortFlag = &cli.UintFlag{
	Name:  "scribe-port",
	Usage: "--scribe-port 5121",
	Value: 0,
}

var scribeGrpcPortFlag = &cli.UintFlag{
	Name:  "scribe-grpc-port",
	Usage: "--scribe-grpc-port 5121",
	Value: 0,
}

var scribeURL = &cli.StringFlag{
	Name:  "scribe-url",
	Usage: "--scribe-url <url>",
	Value: "",
}

var runCommand = &cli.Command{
	Name:        "run",
	Description: "runs the executor service",
	Flags:       []cli.Flag{configFlag, dbFlag, pathFlag, scribePortFlag, scribeGrpcPortFlag, scribeURL},
	Action: func(c *cli.Context) error {
		executorConfig, err := config.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("failed to decode config: %w", err)
		}

		executorDB, err := InitDB(c.Context, c.String(dbFlag.Name), c.String(pathFlag.Name))
		if err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		}

		clients := make(map[uint32]src.Backend)
		for _, client := range executorConfig.Chains {
			rpcDial, err := rpc.DialContext(c.Context, fmt.Sprintf("%s/%d/rpc/%d", executorConfig.BaseOmnirpcURL, 1, client.ChainID))
			if err != nil {
				return fmt.Errorf("failed to dial rpc: %w", err)
			}

			ethClient := ethclient.NewClient(rpcDial)
			clients[client.ChainID] = ethClient
		}

		scribeClient := client.NewRemoteScribe(uint16(c.Uint(scribePortFlag.Name)), uint16(c.Uint(scribeGrpcPortFlag.Name)), c.String(scribeURL.Name))

		executor, err := src.NewExecutor(c.Context, executorConfig, executorDB, scribeClient.ScribeClient, clients)
		if err != nil {
			return fmt.Errorf("failed to create executor: %w", err)
		}

		err = executor.Run(c.Context)
		if err != nil {
			return fmt.Errorf("failed to run executor: %w", err)
		}

		return nil
	},
}

// InitDB initializes a database given a database type and path.
func InitDB(ctx context.Context, database string, path string) (db.ExecutorDB, error) {
	switch {
	case database == "sqlite":
		sqliteStore, err := sqlite.NewSqliteStore(ctx, path)
		if err != nil {
			return nil, fmt.Errorf("failed to create sqlite store: %w", err)
		}

		return sqliteStore, nil

	case database == "mysql":
		if os.Getenv("OVERRIDE_MYSQL") != "" {
			dbname := os.Getenv("MYSQL_DATABASE")
			connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", core.GetEnv("MYSQL_USER", "root"), os.Getenv("MYSQL_PASSWORD"), core.GetEnv("MYSQL_HOST", "127.0.0.1"), core.GetEnvInt("MYSQL_PORT", 3306), dbname)
			mysqlStore, err := mysql.NewMysqlStore(ctx, connString)
			if err != nil {
				return nil, fmt.Errorf("failed to create mysql store: %w", err)
			}

			return mysqlStore, nil
		}

		mysqlStore, err := mysql.NewMysqlStore(ctx, path)
		if err != nil {
			return nil, fmt.Errorf("failed to create mysql store: %w", err)
		}

		return mysqlStore, nil

	default:
		return nil, fmt.Errorf("invalid database type: %s", database)
	}
}
