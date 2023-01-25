package utils

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

// CombinedSchema returns the combined schema of only
// schema, metaSchema, resourceMaps and dataSourceMaps

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

	// project is required to start the proxy
	co.Schema["project"].Required = true
	co.Schema["project"].Optional = false
	// zone is required to start the proxy
	co.Schema["zone"].Required = true
	co.Schema["zone"].Optional = false

	co.Schema["service_account"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "service account to proxy through",
	}
	co.Schema["instance"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The name of the instance to start the proxy on",
	}
	co.Schema["interface"] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "The name of the interface to start the proxy on",
		Default:     "nic0",
		// defaults to default
		Optional: true,
	}

	co.Schema["remote_port"] = &schema.Schema{
		Type:        schema.TypeInt,
		Description: "the port to proxy to",
		// defaults to default
		Optional: true,
		// default tinyproxy port
		Default: "8888",
	}

	return co
}
