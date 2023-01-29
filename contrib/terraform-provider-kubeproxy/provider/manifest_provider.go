package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-cty/cty/msgpack"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	provider_diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-kubernetes/manifest/test/logging"
	"github.com/synapsecns/sanguine/contrib/tfcore/generated/google"
	"github.com/synapsecns/sanguine/contrib/tfcore/generated/manifest"
	"github.com/synapsecns/sanguine/contrib/tfcore/utils"
	"log"
	"os"
	"sync"
)

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

	return func() tfprotov5.ProviderServer {
		return rawProvider
	}, nil
}

type RawProviderServer struct {
	*manifest.RawProviderServer
	combinedSchema *tfprotov5.Schema
	googleProvider *schema.Provider
	// used for stop context
	stopMu sync.Mutex
	// used for stop context
	stopCh chan struct{}
}

func (r *RawProviderServer) GetProviderSchema(ctx context.Context, req *tfprotov5.GetProviderSchemaRequest) (*tfprotov5.GetProviderSchemaResponse, error) {
	return &tfprotov5.GetProviderSchemaResponse{
		Provider: r.combinedSchema,
		// TODO: keys must be rewritten
		ResourceSchemas:   manifest.GetProviderResourceSchema(),
		DataSourceSchemas: manifest.GetProviderDataSourceSchema(),
	}, nil
}

// ConfigureProvider configures the provider and sets up the tunnel
func (r *RawProviderServer) ConfigureProvider(ctx context.Context, req *tfprotov5.ConfigureProviderRequest) (*tfprotov5.ConfigureProviderResponse, error) {
	updatedSchema := utils.UpdateSchemaWithDefaults(r.googleProvider.Schema)
	googleSchema := schema.InternalMap(updatedSchema).CoreConfigSchema()

	resp := &tfprotov5.ConfigureProviderResponse{}

	configVal, err := msgpack.Unmarshal(req.Config.MsgPack, googleSchema.ImpliedType())
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Could not unmarshal config",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	// Ensure there are no nulls that will cause helper/schema to panic.
	if err := validateConfigNulls(ctx, configVal, nil); err != nil {
		resp.Diagnostics = append(resp.Diagnostics, err...)
		return resp, nil
	}

	config := terraform.NewResourceConfigShimmed(configVal, googleSchema)
	logging.HelperSchemaTrace(ctx, "Calling downstream")
	// TODO: move this to the validation step
	diag := r.googleProvider.Validate(config)
	if diag != nil {
		for _, d := range diag {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
				Severity:  convertDiag(d.Severity),
				Summary:   d.Summary,
				Detail:    d.Detail,
				Attribute: PathToAttributePath(d.AttributePath),
			})
		}
		return resp, nil
	}

	logging.HelperSchemaTrace(ctx, "Called downstream")
	ctxHack := context.WithValue(ctx, schema.StopContextKey, r.StopContext(context.Background()))

	r.googleProvider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (_ interface{}, gdg provider_diag.Diagnostics) {
		gface, googleDiagnostics := google.Provider().ConfigureContextFunc(ctx, d)
		gdg = append(gdg, googleDiagnostics...)
		if gdg.HasError() {
			return nil, diag
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

	r.googleProvider.Configure(ctxHack, config)

	resp, err = r.RawProviderServer.ConfigureProvider(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not configure provider: %w", err)
	}

	return resp, nil
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

// makeRawProvider makes a raw provider
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
