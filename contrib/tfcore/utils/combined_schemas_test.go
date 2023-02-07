package utils_test

import (
	"context"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/synapsecns/sanguine/contrib/tfcore/utils"
	"testing"
)

func TestCombineSchemas(t *testing.T) {
	googleProvider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Required: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"google_compute_instance": {},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"google_compute_instance": {},
		},
	}
	underlyingProvider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vpc": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"aws_instance": {},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"aws_instance": {},
		},
	}
	toReplace := "aws"
	replaceWith := "google"
	combinedSchema := utils.CombineSchemas(googleProvider, underlyingProvider, toReplace, replaceWith)
	if combinedSchema.Schema["project"].Required != true {
		t.Errorf("Expected project to be required but got %v", combinedSchema.Schema["project"].Required)
	}
	if combinedSchema.Schema["zone"].Required != true {
		t.Errorf("Expected zone to be required but got %v", combinedSchema.Schema["zone"].Required)
	}
	if combinedSchema.ResourceMap["google_instance"] == nil {
		t.Errorf("Expected resource map to have key google_instance but got %v", combinedSchema.ResourceMap)
	}
	if combinedSchema.DataSourceMap["google_instance"] == nil {
		t.Errorf("Expected data source map to have key google_instance but got %v", combinedSchema.DataSourceMap)
	}
}

func TestCombineProtoSchemas(t *testing.T) {
	ctx := context.Background()
	googleSchema := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
		},
	}
	protoSchema := &tfprotov5.GetProviderSchemaResponse{
		Provider: &tfprotov5.Schema{
			Block: &tfprotov5.SchemaBlock{
				Attributes: []*tfprotov5.SchemaAttribute{
					{
						Name: "test_attribute",
						Type: tftypes.String,
					},
				},
				BlockTypes: []*tfprotov5.SchemaNestedBlock{
					{
						TypeName: "test_block",
						Block: &tfprotov5.SchemaBlock{
							Attributes: []*tfprotov5.SchemaAttribute{
								{
									Name: "test_block_attribute",
									Type: tftypes.String,
								},
							},
						},
					},
				},
			},
		},
	}
	co, err := utils.CombineProtoSchemas(ctx, googleSchema, protoSchema, "", "")
	if err != nil {
		t.Fatalf("CombineProtoSchemas returned error: %v", err)
	}
	if co == nil {
		t.Fatalf("CombineProtoSchemas returned nil co")
	}
	if len(co.Block.Attributes) != 6 {
		t.Fatalf("CombineProtoSchemas did not add all google attributes, expected 6 got %d", len(co.Block.Attributes))
	}
	if len(co.Block.BlockTypes) != 1 {
		t.Fatalf("CombineProtoSchemas did not add all proto block types, expected 1 got %d", len(co.Block.BlockTypes))
	}
}
