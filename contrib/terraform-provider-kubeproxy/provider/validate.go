// Package provider is the Terraform provider for Kubernetes
// nolint
package provider

import (
	"context"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy/generated/convert"
)

// validateConfigNulls checks a config value for unsupported nulls before
// attempting to shim the value. While null values can mostly be ignored in the
// configuration, since they're not supported in HCL1, the case where a null
// appears in a list-like attribute (list, set, tuple) will present a nil value
// to helper/schema which can panic. Return an error to the user in this case,
// indicating the attribute with the null value.
// this function is copied from the grpc provider server with some modifications to avoid use of internals
func validateConfigNulls(ctx context.Context, v cty.Value, path cty.Path) []*tfprotov5.Diagnostic {
	var diags []*tfprotov5.Diagnostic
	if v.IsNull() || !v.IsKnown() {
		return diags
	}

	switch {
	case v.Type().IsListType() || v.Type().IsSetType() || v.Type().IsTupleType():
		it := v.ElementIterator()
		for it.Next() {
			kv, ev := it.Element()
			if ev.IsNull() {
				// if this is a set, the kv is also going to be null which
				// isn't a valid path element, so we can't append it to the
				// diagnostic.
				p := path
				if !kv.IsNull() {
					p = append(p, cty.IndexStep{Key: kv})
				}

				diags = append(diags, &tfprotov5.Diagnostic{
					Severity:  tfprotov5.DiagnosticSeverityError,
					Summary:   "Null value found in list",
					Detail:    "Null values are not allowed for this attribute value.",
					Attribute: convert.PathToAttributePath(p),
				})
				continue
			}

			d := validateConfigNulls(ctx, ev, append(path, cty.IndexStep{Key: kv}))
			diags = append(diags, d...)
		}

	case v.Type().IsMapType() || v.Type().IsObjectType():
		it := v.ElementIterator()
		for it.Next() {
			kv, ev := it.Element()
			var step cty.PathStep
			switch {
			case v.Type().IsMapType():
				step = cty.IndexStep{Key: kv}
			case v.Type().IsObjectType():
				step = cty.GetAttrStep{Name: kv.AsString()}
			}
			d := validateConfigNulls(ctx, ev, append(path, step))
			diags = append(diags, d...)
		}
	}

	return diags
}
