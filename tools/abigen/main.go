// Package main contains a next gen abi generator
package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gen2brain/beeep"
	"github.com/synapsecns/sanguine/tools/abigen/internal"
	"github.com/synapsecns/synapse-node/config"
	"github.com/urfave/cli/v2"
	"os"
)

// appName is the name of the abi generator.
const appName = "abigen"

// TODO use ifacemaker to generate interfaces for these.
func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Description = "abi generator. This extends the standard abi gen by requiring a sol-version and using docker. It also generates metadata for each contract."
	app.Commands = []*cli.Command{
		{
			Name:  "generate",
			Usage: "generate abi bindings from a file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "sol",
					Usage: "path to solidity file you want to compile",
				},
				&cli.StringFlag{
					Name:  "pkg",
					Usage: "name of the package to create",
				},
				&cli.StringFlag{
					Name:  "filename",
					Usage: "[name].abigen.go, [name].contractinfo.json, and [name].metadata.go file",
				},
				&cli.StringFlag{
					Name:  "sol-version",
					Usage: "version of solidity to use to compile the abi. This is pulled from https://hub.docker.com/r/ethereum/solc so version must be present there",
				},
				&cli.IntFlag{
					Name:  "optimizer-runs",
					Usage: "number of optimizations to run.",
					Value: 10000,
				},
			},
			Action: func(context *cli.Context) error {
				//nolint: wrapcheck
				return internal.BuildTemplates(context.String("sol-version"), context.String("sol"), context.String("pkg"), context.String("filename"), context.Int("optimizer-runs"))
			},
		},
		{
			Name:  "generate-from-etherscan",
			Usage: "generate abi bindings from a deployed contract on etherscan",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "address",
					Usage: "address of the deployed contract",
				},
				&cli.StringFlag{
					Name:  "chainID",
					Usage: "chainID of the deployed contract",
				},
				&cli.StringFlag{
					Name:  "pkg",
					Usage: "name of the package to create",
				},
				&cli.StringFlag{
					Name:  "filename",
					Usage: "[name].abigen.go, [name].contractinfo.json, and [name].metadata.go file",
				},
				&cli.StringFlag{
					Name:  "sol-version",
					Usage: "version of solidity to use to compile the abi. This is pulled from https://hub.docker.com/r/ethereum/solc so version must be present there",
				},
			},
			// TODO this needs to embed optimizations, etc from the real deployed contract.
			Action: func(context *cli.Context) error {
				//nolint: wrapcheck
				return internal.GenerateABIFromEtherscan(context.Context, uint(context.Int("chainID")), common.HexToAddress(context.String("address")), context.String("filename"), context.String("sol-version"), context.String("pkg"))
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		// we send an additional alert through beep because go:generate *will* silently fail if ran as
		// go:generate ./...
		logoPath, _ := config.GetLogoPath()

		_ = beeep.Notify("AbiGen Failed", err.Error(), logoPath)
		panic(err)
	}
}
