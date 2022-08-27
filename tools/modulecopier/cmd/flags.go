package cmd

import "github.com/urfave/cli/v2"

var modulePathFlag = &cli.StringFlag{
	Name:     "module-path",
	Usage:    "module path you'd like to copy. For example github.com/ethereum/go-ethereum/console for https://github.com/ethereum/go-ethereum/tree/master/console",
	Required: false,
}

var filePathFlag = &cli.StringFlag{
	Name:     "file-path",
	Usage:    "file path you'd like to copy. For example github.com/ethereum/go-ethereum/console/console.go for https://github.com/ethereum/go-ethereum/tree/master/console.go",
	Required: false,
}

var packageFlag = &cli.StringFlag{
	Name:     "package-name",
	Usage:    "package name of the new package",
	Required: true,
}
