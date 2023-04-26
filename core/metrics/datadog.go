package metrics

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/config"
	"go.opentelemetry.io/contrib/propagators/b3"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentelemetry"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	"k8s.io/apimachinery/pkg/util/sets"
	"strings"
)

type datadogHandler struct {
	*baseHandler
	profilerOptions []profiler.Option
	buildInfo       config.BuildInfo
}

const ddCommitTag = "git.commit.sha"
const ddEnvTag = "DD_ENV"
const ddServiceTag = "DD_SERVICE"
const ddVersionTag = "DD_VERSION"
const defaultEnv = "default"

// NewDatadogMetricsHandler creates a new datadog metrics handler.
func NewDatadogMetricsHandler(buildInfo config.BuildInfo) Handler {
	handler := datadogHandler{
		buildInfo: buildInfo,
	}

	datadogBuildInfo := config.NewBuildInfo(core.GetEnv(ddVersionTag, buildInfo.Version()), buildInfo.Commit(), core.GetEnv(ddServiceTag, buildInfo.Name()), buildInfo.Date())

	// This is a no-op handler to prevent panics. it gets set in start!
	handler.baseHandler = newBaseHandler(datadogBuildInfo)

	handler.profilerOptions = []profiler.Option{
		profiler.WithService(datadogBuildInfo.Name()),
		profiler.WithEnv(core.GetEnv(ddEnvTag, defaultEnv)),
		profiler.WithVersion(datadogBuildInfo.Version()),
		profiler.WithTags(
			fmt.Sprintf("%s:%s", ddCommitTag, datadogBuildInfo.Commit()),
		),
		profiler.WithLogStartup(true),
		profiler.WithProfileTypes(getProfileTypesFromEnv()...),
	}

	return &handler
}

func (d *datadogHandler) Type() HandlerType {
	return DataDog
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

	propagator := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader | b3.B3SingleHeader))

	ddPrpopgator := tracer.NewPropagator(&tracer.PropagatorConfig{B3: true})
	tracerProvider := opentelemetry.NewTracerProvider(tracer.WithRuntimeMetrics(), tracer.WithProfilerEndpoints(true), tracer.WithAnalytics(true),
		tracer.WithPropagator(ddPrpopgator), tracer.WithEnv(core.GetEnv(ddEnvTag, defaultEnv)), tracer.WithService(d.buildInfo.Name()), tracer.WithServiceVersion(d.buildInfo.Version()))

	d.baseHandler = newBaseHandlerWithTracerProvider(d.buildInfo, tracerProvider, propagator)

	// stop on context cancellation
	go func() {
		<-ctx.Done()
		profiler.Stop()
		tracer.Stop()
	}()
	return nil
}

// DDProfileEnv is the data daog profile neviornment variable.
const DDProfileEnv = "DD_PROFILES"

// getProfileTypesFromEnv gets a list of enabled profile types from environment variables.
func getProfileTypesFromEnv() (profiles []profiler.ProfileType) {
	profileEnv := core.GetEnv(DDProfileEnv, strings.Join([]string{profiler.CPUProfile.String(), profiler.MetricsProfile.String(), profiler.GoroutineProfile.String(), profiler.MutexProfile.String(), profiler.HeapProfile.String()}, ","))
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
