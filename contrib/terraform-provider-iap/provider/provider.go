package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/synapsecns/sanguine/contrib/terraform-provider-tunnel/generated/google"
)

func Provider() *schema.Provider {
	underlyingProvider := google.Provider()
	return &schema.Provider{
		Schema:               underlyingProvider.Schema,
		ProviderMetaSchema:   underlyingProvider.ProviderMetaSchema,
		ConfigureContextFunc: underlyingProvider.ConfigureContextFunc,
		DataSourcesMap: map[string]*schema.Resource{
			"tunnel_proxy": dataSourceProxyURL(),
		},
	}
}
