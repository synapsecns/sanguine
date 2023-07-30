package utils

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

// CombinedSchema returns the combined schema of only
// schema, metaSchema, resourceMaps and dataSourceMaps.
type CombinedSchema struct {
	Schema     map[string]*schema.Schema
	MetaSchema map[string]*schema.Schema

	ResourceMap, DataSourceMap map[string]*schema.Resource
}

// CombineSchemas provides 2 helm schemas
// toReplace and replaceWith are used to specify to the provider to replace
// note: while this does not enforce which provider is used, it will fail if one provider is not a google-like provider (in terms of the fields it has)
func CombineSchemas(googleProvider, underlyingProvider *schema.Provider, toReplace, replaceWith string) (co CombinedSchema) {
	// schema
	co.Schema = MustCombineMaps(googleProvider.Schema, underlyingProvider.Schema)
	co.MetaSchema = MustCombineMaps(googleProvider.ProviderMetaSchema, underlyingProvider.ProviderMetaSchema)
	co.ResourceMap = make(map[string]*schema.Resource)
	co.DataSourceMap = make(map[string]*schema.Resource)

	for key, val := range underlyingProvider.ResourcesMap {
		co.ResourceMap[strings.Replace(key, toReplace, replaceWith, 1)] = WrapSchemaResource(val)
	}

	for key, val := range underlyingProvider.DataSourcesMap {
		co.DataSourceMap[strings.Replace(key, toReplace, replaceWith, 1)] = WrapSchemaResource(val)
	}

	co.Schema = UpdateSchemaWithDefaults(co.Schema)

	return co
}

// UpdateSchemaWithDefaults adds extra fields to the schema needed for the google tunnel.
func UpdateSchemaWithDefaults(smap map[string]*schema.Schema) map[string]*schema.Schema {
	// project is required to start the proxy
	smap["project"].Required = true
	smap["project"].Optional = false
	// zone is required to start the proxy
	smap["zone"].Required = true
	smap["zone"].Optional = false

	smap["instance"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The name of the instance to start the proxy on",
	}
	smap["interface"] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "The name of the interface to start the proxy on",
		Default:     "nic0",
		// defaults to default
		Optional: true,
	}

	smap["remote_port"] = &schema.Schema{
		Type:        schema.TypeInt,
		Description: "the port to proxy to",
		// defaults to default
		Optional: true,
		// default tinyproxy port
		Default: "8888",
	}
	return smap
}

// CombineProtoSchemas combines google schemas and tfproto schemas into a single schema
// this differs from CombineSchemas in that it supports tfproto schemas.
func CombineProtoSchemas(ctx context.Context, googleSchema *schema.Provider, protoSchema *tfprotov5.GetProviderSchemaResponse, toReplace, replaceWith string) (co *tfprotov5.Schema, err error) {
	// add defaults to the terraform schema
	googleSchema.Schema = UpdateSchemaWithDefaults(googleSchema.Schema)
	providerSchema := schema.NewGRPCProviderServer(googleSchema)
	tfProviderSchema, err := providerSchema.GetProviderSchema(ctx, &tfprotov5.GetProviderSchemaRequest{})

	if err != nil {
		return nil, fmt.Errorf("could not get provider schema: %w", err)
	}

	for _, attribute := range tfProviderSchema.Provider.Block.Attributes {
		if hasAttribute(protoSchema, attribute) {
			return nil, fmt.Errorf("cannot override attribute %s", attribute.Name)
		}
		protoSchema.Provider.Block.Attributes = append(protoSchema.Provider.Block.Attributes, attribute)
	}

	protoSchema.Provider.Block.BlockTypes = append(protoSchema.Provider.Block.BlockTypes, tfProviderSchema.Provider.Block.BlockTypes...)

	return protoSchema.Provider, nil
}

func hasAttribute(schema *tfprotov5.GetProviderSchemaResponse, attribute *tfprotov5.SchemaAttribute) bool {
	for _, ogAttribute := range schema.Provider.Block.Attributes {
		if ogAttribute.Name == attribute.Name {
			return true
		}
	}
	return false
}
