// Terraform provider iap is a provider that allows you to create and manage long lived iap tunnels in terraform.
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/synapsecns/sanguine/contrib/terraform-provider-iap/provider"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
// this is temporarily disabled until tfexec compatibility issue is fixed (this was removed in 0.16.0)
// we can do this manually for now
// go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider})
}
