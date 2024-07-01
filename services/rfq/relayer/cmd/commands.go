// Package cmd provides the command line interface for the RFQ relayer service
package cmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/service"
	"github.com/urfave/cli/v2"
)

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "path to the config file",
	TakesFile: true,
}

// runCommand runs the rfq relayer.
var runCommand = &cli.Command{
	Name:        "run",
	Description: "run the relayer",
	Flags:       []cli.Flag{configFlag, &commandline.LogLevel},
	Action: func(c *cli.Context) (err error) {
		commandline.SetLogLevel(c)
		cfg, err := relconfig.LoadConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not read config file: %w", err)
		}

		metricsProvider := metrics.Get()

		relayer, err := service.NewRelayer(c.Context, metricsProvider, cfg)
		if err != nil {
			return fmt.Errorf("could not create relayer: %w", err)
		}

		err = relayer.Start(c.Context)
		if err != nil {
			return fmt.Errorf("could not start relayer: %w", err)
		}
		return nil
	},
}

var relayerURLFlag = &cli.StringFlag{
	Name:  "relayer-url",
	Usage: "relayer url",
}

var chainIDFlag = &cli.StringFlag{
	Name:  "chain-id",
	Usage: "chain id",
}

var amountFlag = &cli.StringFlag{
	Name:  "amount",
	Usage: "amount",
}

var tokenAddressFlag = &cli.StringFlag{
	Name:  "token-address",
	Usage: "token address",
}

var toFlag = &cli.StringFlag{
	Name:  "to",
	Usage: "to",
}

// runCommand runs the rfq relayer.
var runWithdrawCommand = &cli.Command{
	Name:        "run",
	Description: "run the withdrawal tool",
	Flags:       []cli.Flag{relayerURLFlag, chainIDFlag, amountFlag, tokenAddressFlag, toFlag, &commandline.LogLevel},
	Action: func(c *cli.Context) (err error) {
		if c.String(relayerURLFlag.Name) == "" {
			return fmt.Errorf("relayer URL is required")
		}

		withdrawer := relapi.NewWithdrawer(relapi.NewRelayerClient(metrics.Get(), c.String(relayerURLFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not create relayer: %w", err)
		}

		chainID := c.Uint(chainIDFlag.Name)
		if chainID == 0 {
			return fmt.Errorf("valid chain ID is required")
		}

		amount := c.String(amountFlag.Name)
		if amount == "" {
			return fmt.Errorf("amount is required")
		}

		tokenAddress := c.String(tokenAddressFlag.Name)
		if !common.IsHexAddress(tokenAddress) {
			return fmt.Errorf("valid token address is required")
		}

		to := c.String(toFlag.Name)
		if !common.IsHexAddress(to) {
			return fmt.Errorf("valid recipient address is required")
		}

		withdrawRequest := relapi.WithdrawRequest{
			ChainID:      uint32(chainID),
			Amount:       amount,
			TokenAddress: common.HexToAddress(tokenAddress),
			To:           common.HexToAddress(to),
		}

		_, err = withdrawer.Withdraw(c.Context, withdrawRequest)
		if err != nil {
			return fmt.Errorf("could not start relayer: %w", err)
		}

		return nil
	},
}
