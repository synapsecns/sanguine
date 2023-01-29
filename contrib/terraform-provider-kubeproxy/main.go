// Package main provides a Terraform wrapper around kubernetes for using an IAP (Identity-Aware Proxy) when interacting with GCP resources.
// The provider wraps the original provider (e.g Helm or Kubernetes) and adds the necessary fields and functionality for configuring and using the IAP proxy.
// This allows for more fine-grained authentication and authorization of access to resources, and is especially useful for short-lived Terraform resources.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy/provider"
	"log"
	"os"
	"time"

	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform-exec/tfexec"
	tf5server "github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	tf5muxserver "github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
)

const providerName = "registry.terraform.io/hashicorp/kubernetes"

// Generate docs for website
// go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	debugFlag := flag.Bool("debug", false, "Start provider in stand-alone debug mode.")
	flag.Parse()

	mainProvider := provider.MainProvider().GRPCProvider
	manifestProvider, err := provider.ManifestProvider()
	if err != nil {
		panic(err)
	}
	// note: manifest provider is not currently supported

	ctx := context.Background()
	muxer, err := tf5muxserver.NewMuxServer(ctx, mainProvider, manifestProvider)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	opts := []tf5server.ServeOpt{}
	if *debugFlag {
		reattachConfigCh := make(chan *plugin.ReattachConfig)
		go func() {
			reattachConfig, err := waitForReattachConfig(reattachConfigCh)
			if err != nil {
				fmt.Printf("Error getting reattach config: %s\n", err)
				return
			}
			printReattachConfig(reattachConfig)
		}()
		opts = append(opts, tf5server.WithDebug(ctx, reattachConfigCh, nil))
	}

	_ = tf5server.Serve(providerName, muxer.ProviderServer, opts...)
}

// convertReattachConfig converts plugin.ReattachConfig to tfexec.ReattachConfig.
func convertReattachConfig(reattachConfig *plugin.ReattachConfig) tfexec.ReattachConfig {
	return tfexec.ReattachConfig{
		Protocol: string(reattachConfig.Protocol),
		Pid:      reattachConfig.Pid,
		Test:     true,
		Addr: tfexec.ReattachConfigAddr{
			Network: reattachConfig.Addr.Network(),
			String:  reattachConfig.Addr.String(),
		},
	}
}

// printReattachConfig prints the line the user needs to copy and paste
// to set the TF_REATTACH_PROVIDERS variable.
func printReattachConfig(config *plugin.ReattachConfig) {
	reattachStr, err := json.Marshal(map[string]tfexec.ReattachConfig{
		"kubernetes": convertReattachConfig(config),
	})
	if err != nil {
		fmt.Printf("Error building reattach string: %s", err)
		return
	}
	fmt.Printf("# Provider server started\nexport TF_REATTACH_PROVIDERS='%s'\n", string(reattachStr))
}

// waitForReattachConfig blocks until a ReattachConfig is received on the
// supplied channel or times out after 2 seconds.
func waitForReattachConfig(ch chan *plugin.ReattachConfig) (*plugin.ReattachConfig, error) {
	select {
	case config := <-ch:
		return config, nil
	case <-time.After(2 * time.Second):
		return nil, fmt.Errorf("timeout while waiting for reattach configuration")
	}
}
