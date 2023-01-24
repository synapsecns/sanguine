// Package provider gets the provider for the iap tunnel.
package provider

import (
	"context"
	provider_diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-helm/helm"
	"github.com/synapsecns/sanguine/contrib/tfcore/generated/google"
	"log"
	"os"
	"strings"
)

var combinedSchema map[string]*schema.Schema
var combinedMetaSchema map[string]*schema.Schema

var resourceMap, dataSourceMap map[string]*schema.Resource

// we use this as a static assertion to check for overlap between keys in the two providers.
func init() {
	// schema
	combinedSchema = MustCombineMaps(google.Provider().Schema, helm.Provider().Schema)
	combinedMetaSchema = MustCombineMaps(google.Provider().ProviderMetaSchema, helm.Provider().ProviderMetaSchema)
	resourceMap = make(map[string]*schema.Resource)
	dataSourceMap = make(map[string]*schema.Resource)

	for key, val := range helm.Provider().ResourcesMap {
		resourceMap[strings.Replace(key, "helm", "helmproxy", 1)] = wrapSchemaResource(val)
	}

	for key, val := range helm.Provider().DataSourcesMap {
		dataSourceMap[strings.Replace(key, "helm", "helmproxy", 1)] = wrapSchemaResource(val)
	}

	// project is required to start the proxy
	combinedSchema["project"].Required = true
	combinedSchema["project"].Optional = false
	// zone is required to start the proxy
	combinedSchema["zone"].Required = true
	combinedSchema["zone"].Optional = false

	combinedSchema["service_account"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "service account to proxy through",
	}
	combinedSchema["instance"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The name of the instance to start the proxy on",
	}
	combinedSchema["interface"] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "The name of the interface to start the proxy on",
		Default:     "nic0",
		// defaults to default
		Optional: true,
	}

	combinedSchema["remote_port"] = &schema.Schema{
		Type:        schema.TypeInt,
		Description: "the port to proxy to",
		// defaults to default
		Optional: true,
		// default tinyproxy port
		Default: "8888",
	}
}

type configuredProvider struct {
	// googleIface is the google interface
	googleIface *google.Config
	// helmIface is the helm interface
	helmIface *helm.Meta
}

// Provider gets the provider for the iap tunnel.
func Provider() *schema.Provider {
	underlyingGoogleProvider := google.Provider()
	underlyingHelmProvider := helm.Provider()
	return &schema.Provider{
		Schema:             combinedSchema,
		ProviderMetaSchema: combinedMetaSchema,
		ResourcesMap:       resourceMap,
		DataSourcesMap:     dataSourceMap,
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

			// TODO: the proxy_url needs to be set in here
			proxyURL, err := startTunnel(ctx, data, cp)
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
