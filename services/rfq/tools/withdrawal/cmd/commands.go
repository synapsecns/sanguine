package cmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
	"github.com/synapsecns/sanguine/services/rfq/tools/withdrawal/withdraw"
	"github.com/urfave/cli/v2"
)

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "path to the config file",
	TakesFile: true,
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
var runCommand = &cli.Command{
	Name:        "run",
	Description: "run the withdrawal tool",
	Flags:       []cli.Flag{relayerURLFlag, chainIDFlag, amountFlag, tokenAddressFlag, toFlag, &commandline.LogLevel},
	Action: func(c *cli.Context) (err error) {

		metricsProvider := metrics.Get()

		withdrawer := withdraw.NewWithdrawer(metricsProvider, c.String(relayerURLFlag.Name))
		if err != nil {
			return fmt.Errorf("could not create relayer: %w", err)
		}

		withdrawRequest := relapi.WithdrawRequest{
			ChainID:      uint32(c.Uint(chainIDFlag.Name)),
			Amount:       c.String(amountFlag.Name),
			TokenAddress: common.HexToAddress(c.String(tokenAddressFlag.Name)),
			To:           common.HexToAddress(toFlag.Name),
		}

		_, err = withdrawer.Withdraw(c.Context, withdrawRequest)
		if err != nil {
			return fmt.Errorf("could not start relayer: %w", err)
		}

		return nil
	},
}
