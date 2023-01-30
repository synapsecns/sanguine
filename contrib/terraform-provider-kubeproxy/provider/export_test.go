package provider

import (
	"context"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ValidateConfigNulls(ctx context.Context, v cty.Value, path cty.Path) []*tfprotov5.Diagnostic {
	return validateConfigNulls(ctx, v, path)
}

func (r *RawProviderServer) CombinedSchema() *tfprotov5.Schema {
	return r.combinedSchema
}

func (r *RawProviderServer) GoogleProvider() *schema.Provider {
	return r.googleProvider
}
