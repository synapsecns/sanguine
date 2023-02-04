package utils

import (
	"context"
	"errors"
	provider_diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// WrappedProvider is a provider that wraps another provider. This is used to wrap the underlying provider with the google provider.
type WrappedProvider interface {
	// UnderlyingProvider gets the underlying provider
	UnderlyingProvider() interface{}
	GoogleProvider() interface{}
}

// WrapSchemaResource wraps a schema.resource to extract the underlying provider interface. This way, we can configure context on both
// the underlying provider interfaces without modifying the underlying provider. This allows are only modification to the provider
// itself to be the addition of the proxy_url field and the proxy starter.
// nolint: staticcheck, wrapcheck, gocognit, cyclop
func WrapSchemaResource(resource *schema.Resource) *schema.Resource {
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
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return errors.New("failed to cast meta interface")
			}
			return resource.Create(data, underlyingProvider.UnderlyingProvider())
		}
	}

	if resource.CreateContext != nil {
		resResource.CreateContext = func(ctx context.Context, data *schema.ResourceData, meta interface{}) (_ provider_diag.Diagnostics) {
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.CreateContext(ctx, data, underlyingProvider.UnderlyingProvider())
		}
	}

	if resource.Read != nil {
		resResource.Read = func(data *schema.ResourceData, meta interface{}) error {
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return errors.New("failed to cast meta interface")
			}
			return resource.Read(data, underlyingProvider.UnderlyingProvider())
		}
	}

	if resource.ReadContext != nil {
		resResource.ReadContext = func(ctx context.Context, data *schema.ResourceData, meta interface{}) (_ provider_diag.Diagnostics) {
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.ReadContext(ctx, data, underlyingProvider.UnderlyingProvider())
		}
	}

	if resource.UpdateContext != nil {
		resResource.UpdateContext = func(ctx context.Context, data *schema.ResourceData, meta interface{}) (_ provider_diag.Diagnostics) {
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.UpdateContext(ctx, data, underlyingProvider.UnderlyingProvider())
		}
	}

	if resource.Update != nil {
		resResource.Update = func(data *schema.ResourceData, meta interface{}) error {
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return errors.New("failed to cast meta interface")
			}
			return resource.Update(data, underlyingProvider.UnderlyingProvider())
		}
	}

	if resource.Delete != nil {
		resResource.Delete = func(data *schema.ResourceData, meta interface{}) error {
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return errors.New("failed to cast meta interface")
			}
			return resource.Delete(data, underlyingProvider.UnderlyingProvider())
		}
	}

	if resource.Exists != nil {
		resResource.Exists = func(data *schema.ResourceData, meta interface{}) (bool, error) {
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return false, errors.New("failed to cast meta interface")
			}
			return resource.Exists(data, underlyingProvider.UnderlyingProvider())
		}
	}

	if resource.CreateWithoutTimeout != nil {
		resResource.CreateWithoutTimeout = func(ctx context.Context, data *schema.ResourceData, meta interface{}) provider_diag.Diagnostics {
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.CreateWithoutTimeout(ctx, data, underlyingProvider.UnderlyingProvider())
		}
	}

	if resource.ReadWithoutTimeout != nil {
		resResource.ReadWithoutTimeout = func(ctx context.Context, data *schema.ResourceData, meta interface{}) provider_diag.Diagnostics {
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.ReadWithoutTimeout(ctx, data, underlyingProvider.UnderlyingProvider())
		}
	}

	if resource.UpdateWithoutTimeout != nil {
		resResource.UpdateWithoutTimeout = func(ctx context.Context, data *schema.ResourceData, meta interface{}) provider_diag.Diagnostics {
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.UpdateWithoutTimeout(ctx, data, underlyingProvider.UnderlyingProvider())
		}
	}

	if resource.DeleteWithoutTimeout != nil {
		resResource.DeleteWithoutTimeout = func(ctx context.Context, data *schema.ResourceData, meta interface{}) provider_diag.Diagnostics {
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}

			return resource.DeleteWithoutTimeout(ctx, data, underlyingProvider.UnderlyingProvider())
		}
	}

	if resource.CustomizeDiff != nil {
		resResource.CustomizeDiff = func(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) error {
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return errors.New("failed to cast meta interface")
			}

			return resource.CustomizeDiff(ctx, diff, underlyingProvider.UnderlyingProvider())
		}
	}

	if resource.DeleteContext != nil {
		resResource.DeleteContext = func(ctx context.Context, data *schema.ResourceData, meta interface{}) (_ provider_diag.Diagnostics) {
			underlyingProvider, ok := meta.(WrappedProvider)
			if !ok {
				return provider_diag.Diagnostics{
					{
						Severity: provider_diag.Error,
						Summary:  "failed to cast meta interface",
					},
				}
			}
			return resource.DeleteContext(ctx, data, underlyingProvider.UnderlyingProvider())
		}
	}

	return resResource
}
