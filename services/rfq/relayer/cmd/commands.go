// Package cmd provides the command line interface for the RFQ relayer service
package cmd

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/charmbracelet/huh/spinner"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
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
var withdrawCommand = &cli.Command{
	Name:        "withdraw",
	Description: "run the withdrawal tool",
	Flags:       []cli.Flag{relayerURLFlag, chainIDFlag, amountFlag, tokenAddressFlag, toFlag, &commandline.LogLevel},
	Action: func(c *cli.Context) (err error) {
		if c.String(relayerURLFlag.Name) == "" {
			return fmt.Errorf("relayer URL is required")
		}

		client := relapi.NewRelayerClient(metrics.Get(), c.String(relayerURLFlag.Name))
		if err != nil {
			return fmt.Errorf("could not create relayer: %w", err)
		}

		chainID := big.NewInt(c.Int64(chainIDFlag.Name))
		if chainID.Cmp(big.NewInt(0)) == 0 {
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
			ChainID:      uint32(chainID.Uint64()),
			Amount:       amount,
			TokenAddress: common.HexToAddress(tokenAddress),
			To:           common.HexToAddress(to),
		}
		res, err := client.Withdraw(c.Context, &withdrawRequest)
		if err != nil {
			return fmt.Errorf("could not start relayer: %w", err)
		}

		var errClient error
		var status *relapi.TxHashByNonceResponse

		ctx, cancel := context.WithTimeout(c.Context, 30*time.Second)
		defer cancel()

		action := func() {
			retry.WithBackoff(ctx, func(ctx context.Context) error {
				status, err = client.GetTxHashByNonce(
					c.Context,
					&relapi.GetTxByNonceRequest{
						ChainID: uint32(chainID.Uint64()),
						Nonce:   res.Nonce,
					})
				if err != nil {
					errClient = err
					return err
				}
				return nil
			})
		}

		err = spinner.New().
			Title("Getting withdrawal tx hash...").
			Action(action).Run()

		if err != nil {
			return fmt.Errorf("could not get withdrawal tx hash: %w", err)
		}
		if errClient != nil {
			return fmt.Errorf("client error: could not get withdrawal tx hash: %w", err)
		}

		fmt.Printf("Withdraw Tx Hash: %s\n", status.Hash)

		return nil
	},
}
