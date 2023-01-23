// Package provider gets the provider for the iap tunnel.
package provider

import (
	"context"
	"fmt"
	provider_diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-helm/helm"
	"github.com/synapsecns/sanguine/contrib/terraform-provider-iap/generated/google"
)

var combinedSchema map[string]*schema.Schema
var combinedMetaSchema map[string]*schema.Schema

var resourceMap, dataSourceMap map[string]*schema.Resource

// we use this as a static assertion to check for overlap between keys in the two providers.
func init() {
	// schema
	combinedSchema = make(map[string]*schema.Schema)
	combinedMetaSchema = make(map[string]*schema.Schema)
	resourceMap = make(map[string]*schema.Resource)
	dataSourceMap = make(map[string]*schema.Resource)

	googleSchema := google.Provider().Schema
	helmSchema := helm.Provider().Schema
	// TODO: remove proxy_url or make inaccessible
	for item, val := range googleSchema {
		combinedSchema[item] = val
		if helmSchema[item] != nil {
			panic(fmt.Errorf("key overlap between google and helm providers on key %s", item))
		}
	}

	// metaschema
	googleSchema = google.Provider().ProviderMetaSchema
	helmSchema = helm.Provider().ProviderMetaSchema
	for item, val := range googleSchema {
		combinedMetaSchema[item] = val
		if helmSchema[item] != nil {
			panic(fmt.Errorf("key overlap between google and helm providers on key %s", item))
		}
	}

	for key, val := range helm.Provider().ResourcesMap {
		resourceMap[key] = wrapSchemaResource(val)
	}

	for key, val := range helm.Provider().DataSourcesMap {
		dataSourceMap[key] = wrapSchemaResource(val)
	}

	// project is required to start the proxy
	combinedSchema["project"].Required = true
	// zone is required to start the proxy
	combinedSchema["zone"].Required = true
	combinedSchema["instance"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The name of the instance to start the proxy on",
	}
	combinedSchema["interface"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The name of the interface to start the proxy on",
		Default:     "nic0",
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

			err = data.Set("proxy_url", proxyURL)
			if err != nil {
				return nil, append(gdg, provider_diag.FromErr(err)[0])
			}

			dg = append(dg, hdg...)
			return cp, dg
		},
	}
}
