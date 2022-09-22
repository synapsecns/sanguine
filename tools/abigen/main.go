// Package main provides a command line tool for generating ABIs.
package main

import (
	"github.com/synapsecns/sanguine/tools/abigen/cmd"
	"os"
)

func main() {
	cmd.Run(os.Args)
}
