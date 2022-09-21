// Package main generates a GQL client.
package main

import (
	"fmt"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/Yamashou/gqlgenc/clientgen"
	clientConfig "github.com/Yamashou/gqlgenc/config"
	"github.com/integralist/go-findroot/find"
	"log"
	"os"
	"path/filepath"
)

func main() {
	root, err := find.Repo()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	// ctx := context.Background()
	err = os.Chdir(filepath.Join(root.Path, "services/explorer/consumer/client/"))
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	cfg, err := config.LoadConfig(filepath.Join(root.Path, "services/scribe/graphql/gqlgen.yaml"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}
	gqlgencConfig, err := clientConfig.LoadConfig(filepath.Join(root.Path, "services/explorer/consumer/client/.gqlgenc.yaml"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err.Error())
		os.Exit(2)
	}

	clientGen := clientgen.New(gqlgencConfig.Query, gqlgencConfig.Client, gqlgencConfig.Generate)
	err = api.Generate(cfg,
		api.AddPlugin(clientGen),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
