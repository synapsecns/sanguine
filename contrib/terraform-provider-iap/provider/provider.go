// Package provider gets the provider for the iap tunnel.
package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	google "github.com/hashicorp/terraform-provider-google/google/provider"
)

// Provider gets the provider for the iap tunnel.
func Provider() *schema.Provider {
	underlyingProvider := google.Provider()
	return &schema.Provider{
		Schema:               underlyingProvider.Schema,
		ProviderMetaSchema:   underlyingProvider.ProviderMetaSchema,
		ConfigureContextFunc: underlyingProvider.ConfigureContextFunc,
		ResourcesMap: map[string]*schema.Resource{
			"iap_tunnel_proxy": dataSourceProxyURL(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"iap_tunnel_keep_alive": keepAlive(),
		},
	}
}
