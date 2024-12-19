package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/tools/abigen/internal"
	cliv2 "github.com/urfave/cli/v2"
	"os"
)

// validateCommandNames ensures command names are unique to prevent registration conflicts.
func validateCommandNames(commands ...*cliv2.Command) {
	seen := make(map[string]bool)
	for _, cmd := range commands {
		if seen[cmd.Name] {
			panic(fmt.Sprintf("duplicate command name detected: %s", cmd.Name))
		}
		seen[cmd.Name] = true
	}
}

var solFlag = &cliv2.StringFlag{
	Name:  "sol",
	Usage: "path to solidity file you want to compile",
}

var pkgFlag = &cliv2.StringFlag{
	Name:  "pkg",
	Usage: "name of the package to create",
}

var filenameFlag = &cliv2.StringFlag{
	Name:  "filename",
	Usage: "[name].abigen.go, [name].contractinfo.json, and [name].metadata.go file",
}

var solVersionFlag = &cliv2.StringFlag{
	Name:  "sol-version",
	Usage: "version of solidity to use to compile the abi. This is pulled from https://hub.docker.com/r/ethereum/solc so version must be present there",
}

var urlFlag = &cliv2.StringFlag{
	Name:  "url",
	Usage: "url of the etherscan api to use",
}

var disableCI = &cliv2.BoolFlag{
	Name:  "disable-ci",
	Usage: "wether or not to disable regeneration on ci",
}

var disableCIEtherscan = &cliv2.BoolFlag{
	Name:  disableCI.Name,
	Usage: "wether or not to disable regeneration on ci, this is disabled on etherscan by default because of api keys",
	Value: true,
}

var optimizerRunsFlags = &cliv2.IntFlag{
	Name:  "optimizer-runs",
	Usage: "number of optimizations to run.",
	Value: 10000,
}

var evmVersionFlags = &cliv2.StringFlag{
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
var GenerateCommand = &cliv2.Command{
	Name:  "generate",
	Usage: "generate abi bindings from a file",
	Flags: []cliv2.Flag{
		solFlag,
		pkgFlag,
		filenameFlag,
		solVersionFlag,
		optimizerRunsFlags,
		evmVersionFlags,
		disableCI,
	},
	Action: func(cliCtx *cliv2.Context) error {
		if cliCtx.Bool(disableCI.Name) && os.Getenv("CI") != "" {
			fmt.Print("skipping generation")
			return nil
		}
		//nolint: wrapcheck
		return internal.BuildTemplates(cliCtx.Context, cliCtx.String(solVersionFlag.Name), cliCtx.String(solFlag.Name), cliCtx.String(pkgFlag.Name), cliCtx.String(filenameFlag.Name), cliCtx.Int(optimizerRunsFlags.Name), strToPt(cliCtx.String(evmVersionFlags.Name)))
	},
}

var addressFlag = &cliv2.StringFlag{
	Name:  "address",
	Usage: "address of the deployed contract",
}

var chainIDFlag = &cliv2.StringFlag{
	Name:  "chainID",
	Usage: "chainID of the deployed contract",
}

// EtherscanCommand is used to pull abi from an etherscan-like api.
var EtherscanCommand = &cliv2.Command{
	Name:  "generate-from-etherscan",
	Usage: "generate abi bindings from a deployed contract on etherscan",
	Flags: []cliv2.Flag{
		addressFlag,
		chainIDFlag,
		pkgFlag,
		filenameFlag,
		solVersionFlag,
		urlFlag,
		disableCIEtherscan,
	},
	// TODO this needs to embed optimizations, etc from the real deployed contract.
	Action: func(cliCtx *cliv2.Context) error {
		if cliCtx.Bool(disableCIEtherscan.Name) && os.Getenv("CI") != "" {
			fmt.Print("skipping generation")
			return nil
		}
		//nolint: wrapcheck
		return internal.GenerateABIFromEtherscan(cliCtx.Context, uint32(cliCtx.Int(chainIDFlag.Name)), cliCtx.String(urlFlag.Name), common.HexToAddress(cliCtx.String(addressFlag.Name)), cliCtx.String(filenameFlag.Name), cliCtx.String(solVersionFlag.Name), cliCtx.String(pkgFlag.Name))
	},
}

func init() {
	// Validate command names are unique before they can be used
	validateCommandNames(GenerateCommand, EtherscanCommand)
}
