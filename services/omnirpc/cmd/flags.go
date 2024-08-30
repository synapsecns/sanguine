package cmd

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var chainIDFlag = &cli.IntFlag{
	Name:     "chain-id",
	Usage:    "Chain id you'd like to select an rpc for",
	Required: true,
}

var portFlag = &cli.IntFlag{
	Name:  "port",
	Usage: "port to run the omniproxy on",
}

var maxSubmitAhead = &cli.IntFlag{
	Name:  "max-submit-ahead",
	Usage: "max number of blocks to submit ahead",
	Value: 0,
}

// set the default dir to the users home path/omnirpc.yaml.
func init() {
	// user must set manually if this errors anyway
	homeDir, _ := os.UserHomeDir()
	defaultConfig := filepath.Join(homeDir, "omnirpc.yaml")
	outputFlag.Value = defaultConfig
	configFlag.Value = defaultConfig
}

var outputFlag = &cli.StringFlag{
	Name:  "output",
	Usage: "path to output the new config file",
}

var configFlag = &cli.StringFlag{
	Name:  "config",
	Usage: "path to output the new config file",
}

var fileFlag = &cli.StringFlag{
	Name:  "file",
	Usage: "path to json file to debug",
}

var rpcFlag = &cli.StringFlag{
	Name:  "rpc",
	Usage: "rpc url to rewrite requests from",
	// --rpc-url is used by cast so we alias it here.
	Aliases: []string{"rpc-url"},
}

var backupRPCFlag = &cli.StringFlag{
	Name:  "backup-rpc",
	Usage: "backup rpc url to use if the primary rpc fails",
}

var recieptsTimeoutFlag = &cli.DurationFlag{
	Name:  "receipts-timeout",
	Usage: "timeout to use for hanging receipts requests",
}

var omnirpcURLFlag = &cli.StringFlag{
	Name:  "omnirpc-url",
	Usage: "Omnirpc rul flag",
}

var confirmationsFlag = &cli.StringFlag{
	Name:  "confirmations",
	Usage: "confirmations flag",
}
