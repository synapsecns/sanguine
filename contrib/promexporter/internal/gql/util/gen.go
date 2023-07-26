package util

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/api"
	"github.com/Yamashou/gqlgenc/clientgen"
	"github.com/Yamashou/gqlgenc/clientgenv2"
	clientConfig "github.com/Yamashou/gqlgenc/config"
	"github.com/Yamashou/gqlgenc/generator"
	"github.com/integralist/go-findroot/find"
	"path/filepath"
)

// GenerateGQLFromLocalServer wraps gqlgenc generation.
func GenerateGQLFromLocalServer(ctx context.Context, configPath string, endpointURL string) error {
	root, err := find.Repo()
	if err != nil {
		return fmt.Errorf("could not find repo root: %w", err)
	}

	gqlgenConfig, err := clientConfig.LoadConfig(filepath.Join(root.Path, configPath))
	if err != nil {
		panic(fmt.Errorf("could not find config: %w", err))
	}

	gqlgenConfig.Endpoint.URL = endpointURL

	clientGen := api.AddPlugin(clientgen.New(gqlgenConfig.Query, gqlgenConfig.Client, gqlgenConfig.Generate))
	if gqlgenConfig.Generate != nil {
		if gqlgenConfig.Generate.ClientV2 {
			clientGen = api.AddPlugin(clientgenv2.New(gqlgenConfig.Query, gqlgenConfig.Client, gqlgenConfig.Generate))
		}
	}

	err = generator.Generate(ctx, gqlgenConfig, clientGen)
	if err != nil {
		return fmt.Errorf("could not generate client: %w", err)
	}
	return nil
}
