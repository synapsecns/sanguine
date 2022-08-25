package main

import "github.com/urfave/cli/v2"

var chainIDFlag = &cli.IntFlag{
	Name:     "chain-id",
	Usage:    "Chain id you'd like tos elect an rpc for",
	Required: true,
}
