package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/tools/abigen/internal"
	"github.com/urfave/cli/v2"
)

var solFlag = &cli.StringFlag{
	Name:  "sol",
	Usage: "path to solidity file you want to compile",
}

var pkgFlag = &cli.StringFlag{
	Name:  "pkg",
	Usage: "name of the package to create",
}

var filenameFlag = &cli.StringFlag{
	Name:  "filename",
	Usage: "[name].abigen.go, [name].contractinfo.json, and [name].metadata.go file",
}

var solVersionFlag = &cli.StringFlag{
	Name:  "sol-version",
	Usage: "version of solidity to use to compile the abi. This is pulled from https://hub.docker.com/r/ethereum/solc so version must be present there",
}

var urlFlag = &cli.StringFlag{
	Name:  "url",
	Usage: "url of the etherscan api to use",
}

var optimizerRunsFlags = &cli.IntFlag{
	Name:  "optimizer-runs",
	Usage: "number of optimizations to run.",
	Value: 10000,
}

var evmVersionFlags = &cli.StringFlag{
	Name:  "evm-version",
	Usage: "evm version to target",
}

// strToPt converts a string to a pointer
// crucially, will return nil if stirng is empty
func strToPt(str string) *string {
	if str == "" {
		return nil
	}
	return core.PtrTo(str)
}

// GenerateCommand generates abi using flags.
var GenerateCommand = &cli.Command{
	Name:  "generate",
	Usage: "generate abi bindings from a file",
	Flags: []cli.Flag{
		solFlag,
		pkgFlag,
		filenameFlag,
		solVersionFlag,
		optimizerRunsFlags,
		evmVersionFlags,
	},
	Action: func(context *cli.Context) error {
		//nolint: wrapcheck
		return internal.BuildTemplates(context.String(solVersionFlag.Name), context.String(solFlag.Name), context.String(pkgFlag.Name), context.String(filenameFlag.Name), context.Int(optimizerRunsFlags.Name), strToPt(context.String(evmVersionFlags.Name)))
	},
}

var addressFlag = &cli.StringFlag{
	Name:  "address",
	Usage: "address of the deployed contract",
}

var chainIDFlag = &cli.StringFlag{
	Name:  "chainID",
	Usage: "chainID of the deployed contract",
}

// EtherscanCommand is used to pull abi from an etherscan-like api.
var EtherscanCommand = &cli.Command{
	Name:  "generate-from-etherscan",
	Usage: "generate abi bindings from a deployed contract on etherscan",
	Flags: []cli.Flag{
		addressFlag,
		chainIDFlag,
		pkgFlag,
		filenameFlag,
		solVersionFlag,
		urlFlag,
	},
	// TODO this needs to embed optimizations, etc from the real deployed contract.
	Action: func(context *cli.Context) error {
		//nolint: wrapcheck
		return internal.GenerateABIFromEtherscan(context.Context, uint32(context.Int(chainIDFlag.Name)), context.String(urlFlag.Name), common.HexToAddress(context.String(addressFlag.Name)), context.String(filenameFlag.String()), context.String(solVersionFlag.Name), context.String(pkgFlag.Name))
	},
}
