package main

import (
	"github.com/synapsecns/sanguine/services/omnirpc/cmd"
	_ "go.uber.org/automaxprocs"
	"os"
)

func main() {
	cmd.Start(os.Args)
}
