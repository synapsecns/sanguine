// Package cmd provides the command line interface for the withdrawal tool. It takes in five
// flags: relayer-url, chain-id, amount, token-address, and to to withdraw <amount> of <token-address>
// to <to> on the <chain-id> chain.
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

		if c.String(relayerURLFlag.Name) == "" {
			return fmt.Errorf("relayer URL is required")
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
