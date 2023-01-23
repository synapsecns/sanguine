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

// we use this as a static assertion to check for overlap between keys in the two providers
func init() {
	// schema
	combinedSchema = make(map[string]*schema.Schema)
	combinedMetaSchema = make(map[string]*schema.Schema)

	googleSchema := google.Provider().Schema
	helmSchema := helm.Provider().Schema
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
}

type configuredProvider struct {
	// googleIface is the google interface
	googleIface interface{}
	// helmIface is the helm interface
	helmIface interface{}
}

// Provider gets the provider for the iap tunnel.
func Provider() *schema.Provider {
	underlyingGoogleProvider := google.Provider()
	underlyingHelmProvider := helm.Provider()
	return &schema.Provider{
		Schema:             combinedSchema,
		ProviderMetaSchema: combinedMetaSchema,
		ConfigureContextFunc: func(ctx context.Context, data *schema.ResourceData) (_ interface{}, dg provider_diag.Diagnostics) {
			cp := configuredProvider{}
			var gdg, hdg provider_diag.Diagnostics
			cp.googleIface, gdg = underlyingGoogleProvider.ConfigureContextFunc(ctx, data)
			if gdg.HasError() {
				return nil, gdg
			}
			dg = append(dg, gdg...)

			cp.helmIface, hdg = underlyingHelmProvider.ConfigureContextFunc(ctx, data)
			if hdg.HasError() {
				return nil, hdg
			}

			dg = append(dg, hdg...)
			return cp, dg
		},
	}
}
