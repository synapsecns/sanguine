package metrics

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/config"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	ddhttp "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/util/sets"
	"net/http"
	"strings"
)

type datadogHandler struct {
	*baseHandler
	profilerOptions []profiler.Option
	buildInfo       config.BuildInfo
}

// NewDatadogMetricsHandler creates a new datadog metrics handler.
func NewDatadogMetricsHandler(buildInfo config.BuildInfo) Handler {
	handler := datadogHandler{
		buildInfo: buildInfo,
	}

	// TODO: these need to be bridged
	handler.baseHandler = newBaseHandler(buildInfo)
	logger.Warn("datadog metrics handler is not fully implemented, please see: https://docs.datadoghq.com/tracing/trace_collection/open_standards/go/")

	handler.profilerOptions = []profiler.Option{
		profiler.WithService(buildInfo.Name()),
		profiler.WithEnv(core.GetEnv("ENVIRONMENT", "default")),
		profiler.WithVersion(buildInfo.Version()),
		profiler.WithTags(
			fmt.Sprintf("commit:%s", buildInfo.Commit()),
		),
		profiler.WithLogStartup(true),
		profiler.WithProfileTypes(getProfileTypesFromEnv()...),
	}

	return &handler
}

func (d *datadogHandler) AddGormCallbacks(db *gorm.DB) {
	// TODO: implement, see:  https://github.com/DataDog/dd-trace-go/blob/main/contrib/jinzhu/gorm/example_test.go
}

// Gin gets a gin middleware for datadog tracing.
func (d *datadogHandler) Gin() gin.HandlerFunc {
	return gintrace.Middleware(d.buildInfo.Name())
}

// Start starts the handler and stops it when context is canceled.
func (d *datadogHandler) Start(ctx context.Context) error {
	err := profiler.Start(d.profilerOptions...)
	if err != nil {
		return fmt.Errorf("could not start profiler: %w", err)
	}

	tracer.Start(tracer.WithRuntimeMetrics(), tracer.WithProfilerEndpoints(true), tracer.WithAnalytics(true))

	// stop on context cancellation
	go func() {
		<-ctx.Done()
		profiler.Stop()
		tracer.Stop()
	}()
	return nil
}

// ConfigureHTTPClient wraps the Transport of an http.Client with a datadog tracer.
func (d *datadogHandler) ConfigureHTTPClient(client *http.Client) {
	wrappedTransport := ddhttp.WrapClient(client).Transport
	client.Transport = wrappedTransport
}

// DDProfileEnv is the data daog profile neviornment variable.
const DDProfileEnv = "DD_PROFILES"

// getProfileTypesFromEnv gets a list of enabled profile types from environment variables.
func getProfileTypesFromEnv() (profiles []profiler.ProfileType) {
	profileEnv := core.GetEnv(DDProfileEnv, strings.Join([]string{profiler.CPUProfile.String(), profiler.MetricsProfile.String()}, ","))
	profilesStr := strings.Split(strings.ToLower(profileEnv), ",")

	// strip duplicates by using a set
	profileSet := sets.NewString(profilesStr...)

	for _, profile := range profileSet.List() {
		usedProfile, ok := allProfileTypes[strings.ToLower(profile)]
		if !ok {
			logger.Errorf("profile %s not found, please use an existent profile type (one of: %s)", profile, allProfilesString())
			continue
		}

		profiles = append(profiles, usedProfile)
	}

	return profiles
}

// allProfileTypes is a list of all profile types supported by datadog.
// these are used to allow enabling by environment variable.
var allProfileTypes map[string]profiler.ProfileType

// allProfilesString gets all profiles as a string for errores.
func allProfilesString() (res string) {
	allProfiles := make([]string, len(allProfileTypes))

	i := 0
	for _, profile := range allProfileTypes {
		allProfiles[i] = profile.String()
		i++
	}
	return strings.Join(allProfiles, ",")
}

func init() {
	allProfileTypes = make(map[string]profiler.ProfileType)

	i := 0
	for {
		profileType := profiler.ProfileType(i)
		if profileType.String() == "unknown" {
			break
		}
		allProfileTypes[strings.ToLower(profileType.String())] = profileType
		i++
	}
}
