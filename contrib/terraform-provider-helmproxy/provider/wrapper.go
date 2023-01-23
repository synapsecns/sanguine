package provider

import (
	"context"
	"errors"
	provider_diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// wrapSchemaResource wraps a schema.resource to extract the helm interface. This way, we can configure context on both
// the helm and google interfaces without modifying the underlying provider. This allows are only modification to the provider
// itself to be the addition of the proxy_url field and the proxy starter.
// nolint: staticcheck, wrapcheck, gocognit, cyclop
func wrapSchemaResource(resource *schema.Resource) *schema.Resource {
	resResource := &schema.Resource{
		Schema:         resource.Schema,
		SchemaVersion:  resource.SchemaVersion,
		MigrateState:   resource.MigrateState,
		StateUpgraders: resource.StateUpgraders,
		Importer:       resource.Importer,
		Description:    resource.Description,
		UseJSONNumber:  resource.UseJSONNumber,
	}

	if resource.Create != nil {
		resResource.Create = func(data *schema.ResourceData, meta interface{}) error {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return errors.New("failed to cast meta interface")
			}
			return resource.Create(data, cp.helmIface)
		}
	}

	if resource.CreateContext != nil {
		resResource.CreateContext = func(ctx context.Context, data *schema.ResourceData, meta interface{}) (_ provider_diag.Diagnostics) {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.CreateContext(ctx, data, cp.helmIface)
		}
	}

	if resource.Read != nil {
		resResource.Read = func(data *schema.ResourceData, meta interface{}) error {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return errors.New("failed to cast meta interface")
			}
			return resource.Read(data, cp.helmIface)
		}
	}

	if resource.ReadContext != nil {
		resResource.ReadContext = func(ctx context.Context, data *schema.ResourceData, meta interface{}) (_ provider_diag.Diagnostics) {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.ReadContext(ctx, data, cp.helmIface)
		}
	}

	if resource.UpdateContext != nil {
		resResource.UpdateContext = func(ctx context.Context, data *schema.ResourceData, meta interface{}) (_ provider_diag.Diagnostics) {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.UpdateContext(ctx, data, cp.helmIface)
		}
	}

	if resource.Update != nil {
		resResource.Update = func(data *schema.ResourceData, meta interface{}) error {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return errors.New("failed to cast meta interface")
			}
			return resource.Update(data, cp.helmIface)
		}
	}

	if resource.Delete != nil {
		resResource.Delete = func(data *schema.ResourceData, meta interface{}) error {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return errors.New("failed to cast meta interface")
			}
			return resource.Delete(data, cp.helmIface)
		}
	}

	if resource.Exists != nil {
		resResource.Exists = func(data *schema.ResourceData, meta interface{}) (bool, error) {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return false, errors.New("failed to cast meta interface")
			}
			return resource.Exists(data, cp.helmIface)
		}
	}

	if resource.CreateWithoutTimeout != nil {
		resResource.CreateWithoutTimeout = func(ctx context.Context, data *schema.ResourceData, meta interface{}) provider_diag.Diagnostics {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.CreateWithoutTimeout(ctx, data, cp.helmIface)
		}
	}

	if resource.ReadWithoutTimeout != nil {
		resResource.ReadWithoutTimeout = func(ctx context.Context, data *schema.ResourceData, meta interface{}) provider_diag.Diagnostics {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.ReadWithoutTimeout(ctx, data, cp.helmIface)
		}
	}

	if resource.UpdateWithoutTimeout != nil {
		resResource.UpdateWithoutTimeout = func(ctx context.Context, data *schema.ResourceData, meta interface{}) provider_diag.Diagnostics {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.UpdateWithoutTimeout(ctx, data, cp.helmIface)
		}
	}

	if resource.DeleteWithoutTimeout != nil {
		resResource.DeleteWithoutTimeout = func(ctx context.Context, data *schema.ResourceData, meta interface{}) provider_diag.Diagnostics {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}

			return resource.DeleteWithoutTimeout(ctx, data, cp.helmIface)
		}
	}

	if resource.CustomizeDiff != nil {
		resResource.CustomizeDiff = func(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) error {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return errors.New("failed to cast meta interface")
			}

			return resource.CustomizeDiff(ctx, diff, cp.helmIface)
		}
	}

	if resource.DeleteContext != nil {
		resResource.DeleteContext = func(ctx context.Context, data *schema.ResourceData, meta interface{}) (_ provider_diag.Diagnostics) {
			cp, ok := meta.(*configuredProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.DeleteContext(ctx, data, cp.helmIface)
		}
	}

	return resResource
}
