package cmd

import (
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
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
}
