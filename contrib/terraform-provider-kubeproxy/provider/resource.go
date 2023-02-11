// Package provider is a shim for the Package method of the underlying provider.
// nolint: wrapcheck
package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"strings"
)

// ReadDataSource is a shim for the ReadDataSource method of the underlying provider.
func (r *RawProviderServer) ReadDataSource(ctx context.Context, req *tfprotov5.ReadDataSourceRequest) (*tfprotov5.ReadDataSourceResponse, error) {
	req.TypeName = strings.Replace(req.TypeName, replacedProviderPrefix, originalProviderPrefix, 1)

	return r.RawProviderServer.ReadDataSource(ctx, req)
}

// ValidateDataSourceConfig is a shim for the ValidateDataSourceConfig method of the underlying provider.
func (r *RawProviderServer) ValidateDataSourceConfig(ctx context.Context, req *tfprotov5.ValidateDataSourceConfigRequest) (*tfprotov5.ValidateDataSourceConfigResponse, error) {
	req.TypeName = strings.Replace(req.TypeName, replacedProviderPrefix, originalProviderPrefix, 1)

	return r.RawProviderServer.ValidateDataSourceConfig(ctx, req)
}

// ValidateResourceTypeConfig is a shim for the ValidateProviderConfig method of the underlying provider.
func (r *RawProviderServer) ValidateResourceTypeConfig(ctx context.Context, req *tfprotov5.ValidateResourceTypeConfigRequest) (*tfprotov5.ValidateResourceTypeConfigResponse, error) {
	req.TypeName = strings.Replace(req.TypeName, replacedProviderPrefix, originalProviderPrefix, 1)

	return r.RawProviderServer.ValidateResourceTypeConfig(ctx, req)
}

// UpgradeResourceState is a shim for the UpgradeResourceState method of the underlying provider.
func (r *RawProviderServer) UpgradeResourceState(ctx context.Context, req *tfprotov5.UpgradeResourceStateRequest) (*tfprotov5.UpgradeResourceStateResponse, error) {
	req.TypeName = strings.Replace(req.TypeName, replacedProviderPrefix, originalProviderPrefix, 1)

	return r.RawProviderServer.UpgradeResourceState(ctx, req)
}

// ReadResource is a shim for the ReadResource method of the underlying provider.
func (r *RawProviderServer) ReadResource(ctx context.Context, req *tfprotov5.ReadResourceRequest) (*tfprotov5.ReadResourceResponse, error) {
	req.TypeName = strings.Replace(req.TypeName, replacedProviderPrefix, originalProviderPrefix, 1)

	return r.RawProviderServer.ReadResource(ctx, req)
}

// PlanResourceChange is a shim for the PlanResourceChange method of the underlying provider.
func (r *RawProviderServer) PlanResourceChange(ctx context.Context, req *tfprotov5.PlanResourceChangeRequest) (*tfprotov5.PlanResourceChangeResponse, error) {
	req.TypeName = strings.Replace(req.TypeName, replacedProviderPrefix, originalProviderPrefix, 1)

	return r.RawProviderServer.PlanResourceChange(ctx, req)
}

// ApplyResourceChange is a shim for the ApplyResourceChange method of the underlying provider.
func (r *RawProviderServer) ApplyResourceChange(ctx context.Context, req *tfprotov5.ApplyResourceChangeRequest) (*tfprotov5.ApplyResourceChangeResponse, error) {
	req.TypeName = strings.Replace(req.TypeName, replacedProviderPrefix, originalProviderPrefix, 1)

	return r.RawProviderServer.ApplyResourceChange(ctx, req)
}

// ImportResourceState is a shim for the ImportResourceState method of the underlying provider.
func (r *RawProviderServer) ImportResourceState(ctx context.Context, req *tfprotov5.ImportResourceStateRequest) (*tfprotov5.ImportResourceStateResponse, error) {
	req.TypeName = strings.Replace(req.TypeName, replacedProviderPrefix, originalProviderPrefix, 1)

	return r.RawProviderServer.ImportResourceState(ctx, req)
}
