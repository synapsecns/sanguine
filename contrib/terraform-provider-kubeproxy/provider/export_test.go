package provider

import (
	"context"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ValidateConfigNulls is a wrapper around validateConfigNulls for testing.
func ValidateConfigNulls(ctx context.Context, v cty.Value, path cty.Path) []*tfprotov5.Diagnostic {
	return validateConfigNulls(ctx, v, path)
}

// CombinedSchema is a wrapper around combinedSchema for testing.
func (r *RawProviderServer) CombinedSchema() *tfprotov5.Schema {
	return r.combinedSchema
}

// GoogleProvider is a wrapper around googleProvider for testing.
func (r *RawProviderServer) GoogleProvider() *schema.Provider {
	return r.googleProvider
}
