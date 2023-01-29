package provider

import (
	"context"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// validateConfigNulls checks a config value for unsupported nulls before
// attempting to shim the value. While null values can mostly be ignored in the
// configuration, since they're not supported in HCL1, the case where a null
// appears in a list-like attribute (list, set, tuple) will present a nil value
// to helper/schema which can panic. Return an error to the user in this case,
// indicating the attribute with the null value.
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
					Attribute: PathToAttributePath(p),
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

// PathToAttributePath takes a cty.Path and converts it to a proto-encoded path.
func PathToAttributePath(p cty.Path) *tftypes.AttributePath {
	if p == nil || len(p) < 1 {
		return nil
	}
	ap := tftypes.NewAttributePath()
	for _, step := range p {
		switch selector := step.(type) {
		case cty.GetAttrStep:
			ap = ap.WithAttributeName(selector.Name)

		case cty.IndexStep:
			key := selector.Key
			switch key.Type() {
			case cty.String:
				ap = ap.WithElementKeyString(key.AsString())
			case cty.Number:
				v, _ := key.AsBigFloat().Int64()
				ap = ap.WithElementKeyInt(int(v))
			default:
				// We'll bail early if we encounter anything else, and just
				// return the valid prefix.
				return ap
			}
		}
	}
	return ap
}

func convertDiag(sev diag.Severity) tfprotov5.DiagnosticSeverity {
	switch sev {
	case diag.Error:
		return tfprotov5.DiagnosticSeverityError
	default:
		return tfprotov5.DiagnosticSeverityWarning
	}
}
