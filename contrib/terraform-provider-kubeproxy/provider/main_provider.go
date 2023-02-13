package provider

import (
	"context"
	provider_diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-kubernetes/kubernetes"
	"github.com/synapsecns/sanguine/contrib/tfcore/generated/google"
	"github.com/synapsecns/sanguine/contrib/tfcore/utils"
	"log"
)

type configuredKubeProvider struct {
	// googleIface is the google interface
	googleIface *google.Config
	// helmIface is the helm interface
	kubeIface interface{}
}

func (c configuredKubeProvider) GoogleProvider() interface{} {
	return c.googleIface
}

func (c configuredKubeProvider) UnderlyingProvider() interface{} {
	return c.kubeIface
}

var _ utils.WrappedProvider = &configuredKubeProvider{}

// MainProvider creates the main provider for the iap tunnel.
func MainProvider() *schema.Provider {
	combinedSchema := utils.CombineSchemas(google.Provider(), kubernetes.Provider(), "kubernetes", "kubeproxy")
	underlyingGoogleProvider := google.Provider()
	underlyingKubernetesProvider := kubernetes.Provider()
	return &schema.Provider{
		Schema:             combinedSchema.Schema,
		ProviderMetaSchema: combinedSchema.MetaSchema,
		ResourcesMap:       combinedSchema.ResourceMap,
		DataSourcesMap:     combinedSchema.DataSourceMap,
		ConfigureContextFunc: func(ctx context.Context, data *schema.ResourceData) (_ interface{}, dg provider_diag.Diagnostics) {
			cp := &configuredKubeProvider{}
			var gdg, hdg provider_diag.Diagnostics
			var giface interface{}
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

			// TODO: the proxy_url needs to be set in here
			proxyURL, err := utils.StartTunnel(ctx, data, cp.googleIface)
			if err != nil {
				return nil, append(gdg, provider_diag.FromErr(err)[0])
			}

			// set the proxy url
			log.Printf("[INFO] setting proxy url to %s", proxyURL)
			err = data.Set("proxy_url", proxyURL)
			if err != nil {
				return nil, append(gdg, provider_diag.FromErr(err)[0])
			}

			cp.kubeIface, hdg = underlyingKubernetesProvider.ConfigureContextFunc(ctx, data)
			if hdg.HasError() {
				return nil, hdg
			}
			dg = append(dg, hdg...)
			if !ok {
				return nil, append(hdg, provider_diag.Diagnostic{
					Severity: provider_diag.Error,
					Summary:  "failed to cast kubernetes interface",
				})
			}

			return cp, dg
		},
	}
}
