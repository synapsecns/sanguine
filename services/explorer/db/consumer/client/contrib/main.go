// Package main generates a GQL client.
package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/api"
	"github.com/Yamashou/gqlgenc/clientgen"
	"github.com/Yamashou/gqlgenc/clientgenv2"
	clientConfig "github.com/Yamashou/gqlgenc/config"
	"github.com/Yamashou/gqlgenc/generator"
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
	ctx := context.Background()
	err = os.Chdir(filepath.Join(root.Path, "services/explorer/db/consumer/client/"))
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	cfg, err := clientConfig.LoadConfig(filepath.Join(root.Path, "services/explorer/db/consumer/client/.gqlgenc.yaml"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err.Error())
		os.Exit(2)
	}

	clientGen := api.AddPlugin(clientgen.New(cfg.Query, cfg.Client, cfg.Generate))
	if cfg.Generate != nil {
		if cfg.Generate.ClientV2 {
			clientGen = api.AddPlugin(clientgenv2.New(cfg.Query, cfg.Client, cfg.Generate))
		}
	}

	if err := generator.Generate(ctx, cfg, clientGen); err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err.Error())
		os.Exit(4)
	}
}
