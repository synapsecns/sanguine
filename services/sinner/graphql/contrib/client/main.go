// generate gql schema
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/Yamashou/gqlgenc/clientgen"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	clientConfig "github.com/Yamashou/gqlgenc/config"
	"github.com/integralist/go-findroot/find"
)

func main() {
	root, err := find.Repo()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	log.SetOutput(io.Discard)

	err = os.Chdir(filepath.Join(root.Path, "services/sinner/graphql/"))
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	cfg, err := config.LoadConfig(filepath.Join(root.Path, "services/sinner/graphql/gqlgen.yaml"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}
	gqlgencConfig, err := clientConfig.LoadConfig(filepath.Join(root.Path, "services/sinner/graphql/.gqlgenc.yaml"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load client config", err.Error())
		os.Exit(2)
	}

	// copy over relevant config
	gqlgencConfig.Models = cfg.Models
	gqlgencConfig.Model = cfg.Model
	gqlgencConfig.SchemaFilename = clientConfig.StringList(cfg.SchemaFilename)

	clientPlugin := clientgen.New(gqlgencConfig.Query, gqlgencConfig.Client, gqlgencConfig.Generate)
	err = api.Generate(cfg,
		api.AddPlugin(clientPlugin),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
