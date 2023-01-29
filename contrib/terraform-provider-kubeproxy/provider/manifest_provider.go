package provider

import (
	"context"
	gojson "encoding/json"
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-cty/cty/json"
	"github.com/hashicorp/go-cty/cty/msgpack"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	provider_diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-kubernetes/manifest/test/logging"
	"github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy/generated/configschema"
	"github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy/generated/convert"
	"github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy/generated/manifest"
	"github.com/synapsecns/sanguine/contrib/tfcore/generated/google"
	"github.com/synapsecns/sanguine/contrib/tfcore/utils"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

// ManifestProvider gets the manifest provider.
func ManifestProvider() (func() tfprotov5.ProviderServer, error) {
	providerSchema, err := manifest.Provider()().GetProviderSchema(context.Background(), &tfprotov5.GetProviderSchemaRequest{})
	if err != nil {
		return nil, fmt.Errorf("could not get provider schema: %w", err)
	}

	rawProvider := makeRawProvider()

	rawProvider.combinedSchema, err = utils.CombineProtoSchemas(context.Background(), google.Provider(), providerSchema, "", "")
	if err != nil {
		return nil, fmt.Errorf("could not combine schemas: %w", err)
	}

	rawProvider.googleProvider = google.Provider()
	rawProvider.googleProvider.Schema = utils.UpdateSchemaWithDefaults(rawProvider.googleProvider.Schema)

	return func() tfprotov5.ProviderServer {
		return rawProvider
	}, nil
}

// RawProviderServer is the raw provider server.
type RawProviderServer struct {
	*manifest.RawProviderServer
	combinedSchema *tfprotov5.Schema
	googleProvider *schema.Provider
	// used for stop context
	stopMu sync.Mutex
	// used for stop context
	stopCh chan struct{}
}

const (
	originalProviderPrefix = "kubernetes"
	replacedProviderPrefix = "kubeproxy"
)

// GetProviderSchema returns the provider schema.
func (r *RawProviderServer) GetProviderSchema(_ context.Context, _ *tfprotov5.GetProviderSchemaRequest) (*tfprotov5.GetProviderSchemaResponse, error) {
	return &tfprotov5.GetProviderSchemaResponse{
		Provider: r.combinedSchema,
		// TODO: keys must be rewritten
		ResourceSchemas:   replaceResourceKeys(manifest.GetProviderResourceSchema(), originalProviderPrefix, replacedProviderPrefix),
		DataSourceSchemas: replaceResourceKeys(manifest.GetProviderDataSourceSchema(), originalProviderPrefix, replacedProviderPrefix),
	}, nil
}

func replaceResourceKeys(keyMap map[string]*tfprotov5.Schema, toReplace, replaceWith string) map[string]*tfprotov5.Schema {
	newKeyMap := make(map[string]*tfprotov5.Schema)
	for key, value := range keyMap {
		newKeyMap[strings.Replace(key, toReplace, replaceWith, 1)] = value
	}
	return newKeyMap
}

// ConfigureProvider configures the provider and sets up the tunnel.
func (r *RawProviderServer) ConfigureProvider(ctx context.Context, req *tfprotov5.ConfigureProviderRequest) (*tfprotov5.ConfigureProviderResponse, error) {
	resp := &tfprotov5.ConfigureProviderResponse{}

	// we start by adding our custom fields to the google schema
	updatedSchema := utils.UpdateSchemaWithDefaults(google.Provider().Schema)
	// and then converting the google schema to an internal map. This can be used by the converter module
	googleSchema := schema.InternalMap(updatedSchema).CoreConfigSchema()

	// we convert the schema to a config schema
	combinedConfigSchema := convert.ProtoToConfigSchema(ctx, r.combinedSchema.Block)

	// we then use that to unmarshall the config
	reqConfig, err := msgpack.Unmarshal(req.Config.MsgPack, combinedConfigSchema.ImpliedType())
	if err != nil {
		resp.Diagnostics = convert.AppendProtoDiag(ctx, resp.Diagnostics, err)
		return resp, nil
	}

	// Ensure there are no nulls that will cause helper/schema to panic.
	if err := validateConfigNulls(ctx, reqConfig, nil); err != nil {
		resp.Diagnostics = append(resp.Diagnostics, err...)
		return resp, nil
	}

	// shim the tfproto config into a terraform config
	config := terraform.NewResourceConfigShimmed(reqConfig, googleSchema)
	ctxHack := context.WithValue(ctx, schema.StopContextKey, r.StopContext(context.Background()))

	logging.HelperSchemaTrace(ctx, "Calling downstream configure google")
	// configure the google provider
	r.googleProvider.ConfigureContextFunc = googConfigureContextFunc
	diag := r.googleProvider.Configure(ctxHack, config)
	if diag.HasError() {
		resp.Diagnostics = convert.AppendProtoDiag(ctx, resp.Diagnostics, diag)
		return resp, nil
	}

	logging.HelperSchemaTrace(ctx, "Called downstream configure google")
	// remove extra fields

	marshalledRequest, err := removeRequestFields(ctx, reqConfig, combinedConfigSchema, maps.Keys(updatedSchema))
	if err != nil {
		resp.Diagnostics = convert.AppendProtoDiag(ctx, resp.Diagnostics, err)
	}
	marshalledRequest.TerraformVersion = req.TerraformVersion

	// end remove extra fields
	resp, err = r.RawProviderServer.ConfigureProvider(ctx, marshalledRequest)
	if err != nil {
		return nil, fmt.Errorf("could not configure provider: %w", err)
	}

	return resp, nil
}

// removeRequestFields removes google specified fields from the configure provider request
// represented as cty.Value and returns a new request that can be used for the provider.
func removeRequestFields(ctx context.Context, reqConfig cty.Value, combinedConfigSchema *configschema.Block, keysToPrune []string) (*tfprotov5.ConfigureProviderRequest, error) {
	// we'll start by marshaling the config to json, we'll then remove extra keys and
	// convert the config back to msgpack
	jsonReq, err := json.Marshal(reqConfig, combinedConfigSchema.ImpliedType())
	if err != nil {
		return nil, fmt.Errorf("could not marshal config: %w", err)
	}
	logging.HelperSchemaTrace(ctx, "pruning google fields from config")
	logging.HelperSchemaTrace(ctx, string(jsonReq))

	var objmap map[string]gojson.RawMessage
	err = gojson.Unmarshal(jsonReq, &objmap)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal config: %w", err)
	}

	// we'll remove the google fields
	for field := range objmap {
		if slices.Contains(keysToPrune, field) {
			delete(objmap, field)
		}
		if field == "proxy_url" {
			objmap["proxy_url"] = gojson.RawMessage(strconv.Quote(os.Getenv("KUBE_PROXY_URL")))
		}
	}

	// we'll then marshal the config back to messagepack
	jsonReq, err = gojson.Marshal(objmap)
	if err != nil {
		return nil, fmt.Errorf("could not marshal config: %w", err)
	}

	logging.HelperSchemaTrace(ctx, "pruned google fields from config")
	logging.HelperSchemaTrace(ctx, string(jsonReq))

	req := &tfprotov5.ConfigureProviderRequest{}
	req.Config = &tfprotov5.DynamicValue{
		JSON: jsonReq,
	}
	return req, nil
}

// googConfigureContextFunc configures the context function for google.
func googConfigureContextFunc(ctx context.Context, d *schema.ResourceData) (_ interface{}, gdg provider_diag.Diagnostics) {
	gface, googleDiagnostics := google.Provider().ConfigureContextFunc(ctx, d)
	gdg = append(gdg, googleDiagnostics...)
	if gdg.HasError() {
		return nil, gdg
	}

	googleConfig, ok := gface.(*google.Config)
	if !ok {
		return nil, append(gdg, provider_diag.Diagnostic{
			Severity: provider_diag.Error,
			Summary:  "failed to cast google interface",
		})
	}
	// TODO: the proxy_url needs to be set in here
	proxyURL, err := utils.StartTunnel(ctx, d, googleConfig)
	if err != nil {
		return nil, append(gdg, provider_diag.FromErr(err)[0])
	}

	// set the proxy url
	log.Printf("[INFO] setting proxy url to %s", proxyURL)
	err = os.Setenv("KUBE_PROXY_URL", proxyURL)
	if err != nil {
		return nil, append(gdg, provider_diag.FromErr(err)[0])
	}
	return gface, gdg
}

// StopContext derives a new context from the passed in grpc context.
// It creates a goroutine to wait for the server stop and propagates
// cancellation to the derived grpc context.
func (r *RawProviderServer) StopContext(ctx context.Context) context.Context {
	ctx = logging.InitContext(ctx)
	r.stopMu.Lock()
	defer r.stopMu.Unlock()

	stoppable, cancel := context.WithCancel(ctx)
	go mergeStop(stoppable, cancel, r.stopCh)
	return stoppable
}

// mergeStop is called in a goroutine and waits for the global stop signal
// and propagates cancellation to the passed in ctx/cancel func. The ctx is
// also passed to this function and waited upon so no goroutine leak is caused.
func mergeStop(ctx context.Context, cancel context.CancelFunc, stopCh chan struct{}) {
	select {
	case <-ctx.Done():
		return
	case <-stopCh:
		cancel()
	}
}

// makeRawProvider makes a raw provider.
func makeRawProvider() *RawProviderServer {
	var logLevel string
	var ok = false
	for _, ev := range []string{"TF_LOG_PROVIDER_KUBERNETES", "TF_LOG_PROVIDER", "TF_LOG"} {
		logLevel, ok = os.LookupEnv(ev)
		if ok {
			break
		}
	}
	if !ok {
		logLevel = "off"
	}

	rawProvider := &manifest.RawProviderServer{}
	rawProvider.SetLogger(hclog.New(&hclog.LoggerOptions{
		Level:  hclog.LevelFromString(logLevel),
		Output: os.Stderr,
	}))

	return &RawProviderServer{RawProviderServer: rawProvider}
}
