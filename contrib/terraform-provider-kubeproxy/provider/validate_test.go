package provider_test

import (
	"context"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy/generated/convert"
	"github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy/provider"
	"strconv"
	"testing"
)

// Copied from grpc provider to test parity.
func TestValidateNulls(t *testing.T) {
	for i, tc := range []struct {
		Cfg cty.Value
		Err bool
	}{
		{
			Cfg: cty.ObjectVal(map[string]cty.Value{
				"list": cty.ListVal([]cty.Value{
					cty.StringVal("string"),
					cty.NullVal(cty.String),
				}),
			}),
			Err: true,
		},
		{
			Cfg: cty.ObjectVal(map[string]cty.Value{
				"map": cty.MapVal(map[string]cty.Value{
					"string": cty.StringVal("string"),
					"null":   cty.NullVal(cty.String),
				}),
			}),
			Err: false,
		},
		{
			Cfg: cty.ObjectVal(map[string]cty.Value{
				"object": cty.ObjectVal(map[string]cty.Value{
					"list": cty.ListVal([]cty.Value{
						cty.StringVal("string"),
						cty.NullVal(cty.String),
					}),
				}),
			}),
			Err: true,
		},
		{
			Cfg: cty.ObjectVal(map[string]cty.Value{
				"object": cty.ObjectVal(map[string]cty.Value{
					"list": cty.ListVal([]cty.Value{
						cty.StringVal("string"),
						cty.NullVal(cty.String),
					}),
					"list2": cty.ListVal([]cty.Value{
						cty.StringVal("string"),
						cty.NullVal(cty.String),
					}),
				}),
			}),
			Err: true,
		},
		{
			Cfg: cty.ObjectVal(map[string]cty.Value{
				"object": cty.ObjectVal(map[string]cty.Value{
					"list": cty.SetVal([]cty.Value{
						cty.StringVal("string"),
						cty.NullVal(cty.String),
					}),
				}),
			}),
			Err: true,
		},
	} {
		tc := tc // capture func literal
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			d := provider.ValidateConfigNulls(context.Background(), tc.Cfg, nil)
			diags := convert.ProtoToDiags(d)
			switch {
			case tc.Err:
				if !diags.HasError() {
					t.Fatal("expected error")
				}
			default:
				for _, d := range diags {
					if d.Severity == diag.Error {
						t.Fatalf("unexpected error: %q", d)
					}
				}
			}
		})
	}
}
