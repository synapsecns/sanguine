// Package provider gets the provider for the iap tunnel.
package provider

import (
	"context"
	provider_diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-helm/helm"
	"github.com/synapsecns/sanguine/contrib/tfcore/generated/google"
	"github.com/synapsecns/sanguine/contrib/tfcore/utils"
	"log"
	"os"
)

type configuredProvider struct {
	// googleIface is the google interface
	googleIface *google.Config
	// helmIface is the helm interface
	helmIface *helm.Meta
}

func (c configuredProvider) GoogleProvider() interface{} {
	return c.googleIface
}

func (c configuredProvider) UnderlyingProvider() interface{} {
	return c.helmIface
}

var _ utils.WrappedProvider = &configuredProvider{}

// Provider gets the provider for the iap tunnel.
func Provider() *schema.Provider {
	combinedSchema := utils.CombineSchemas(google.Provider(), helm.Provider(), "helm", "helmproxy")
	underlyingGoogleProvider := google.Provider()
	underlyingHelmProvider := helm.Provider()
	return &schema.Provider{
		Schema:             combinedSchema.Schema,
		ProviderMetaSchema: combinedSchema.MetaSchema,
		ResourcesMap:       combinedSchema.ResourceMap,
		DataSourcesMap:     combinedSchema.DataSourceMap,
		ConfigureContextFunc: func(ctx context.Context, data *schema.ResourceData) (_ interface{}, dg provider_diag.Diagnostics) {
			cp := &configuredProvider{}
			var gdg, hdg provider_diag.Diagnostics
			var giface, hiface interface{}
			var ok bool

			giface, gdg = underlyingGoogleProvider.ConfigureContextFunc(ctx, data)
			if gdg.HasError() {
				return nil, gdg
			}
			dg = append(dg, gdg...)
			cp.googleIface, ok = giface.(*google.Config)
			if !ok {
				return nil, append(gdg, provider_diag.Diagnostic{
					Severity: provider_diag.Error,
					Summary:  "failed to cast google interface",
				})
			}

			hiface, hdg = underlyingHelmProvider.ConfigureContextFunc(ctx, data)
			if hdg.HasError() {
				return nil, hdg
			}
			cp.helmIface, ok = hiface.(*helm.Meta)
			if !ok {
				return nil, append(gdg, provider_diag.Diagnostic{
					Severity: provider_diag.Error,
					Summary:  "failed to cast helm interface",
				})
			}

			proxyURL, err := utils.StartTunnel(ctx, data, cp.googleIface)
			if err != nil {
				return nil, append(gdg, provider_diag.FromErr(err)[0])
			}

			// set the proxy url
			log.Printf("[INFO] setting proxy url to %s", proxyURL)
			err = os.Setenv("KUBE_PROXY_URL", proxyURL)
			if err != nil {
				return nil, append(gdg, provider_diag.FromErr(err)[0])
			}

			dg = append(dg, hdg...)
			return cp, dg
		},
	}
}
