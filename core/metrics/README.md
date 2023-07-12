# Metrics Package

The metrics package contains standard drivers for opentracing, profiling and metrics.You can enable this in your appplication by calling `metrics.Setup()`. This will start the metric handler defined in the `METRICS_HANDLER` environment variable. The default is `null` which is a no-op handler.

| `METRICS_HANDLER` env | Description                                                                                                                                                                                                                                                                                                                                                                                                                          | Supports Traces | Supports Span Events | Supports Profiling                                                                            |
|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------|----------------------|-----------------------------------------------------------------------------------------------|
| Datadog               | [Datadog go](https://docs.datadoghq.com/integrations/go/) integration using opentracing bridge                                                                                                                                                                                                                                                                                                                                       | ✅               | ❌                    | ✅                                                                                             |
| OTLP                  | [OTLP Exporter](https://opentelemetry.io/docs/specs/otel/protocol/exporter/) protocol. Supported by various external providers including [New Relic](https://docs.newrelic.com/docs/more-integrations/open-source-telemetry-integrations/opentelemetry/opentelemetry-introduction/), [Signoz](https://signoz.io/blog/opentelemetry-collector-complete-guide/), [Grafana](https://grafana.com/docs/opentelemetry/collector/) and more | ✅               | ✅                    | ❌ (but it can through pyroscope, by specifying the `PYROSCOPE_ENDPOINT` enviornment variable) |
| Jaeger                | [Jaeger](https://www.jaegertracing.io/docs/1.46/) Client Clibrary, will soon be deprecated in favor of OTLP exports to jaeger as per [this deprecation notice](https://www.jaegertracing.io/docs/1.46/client-libraries/)                                                                                                                                                                                                             | ✅               | ✅                    | ❌ (but it can through pyroscope, by specifying the `PYROSCOPE_ENDPOINT` enviornment variable) |

## Datadog:

By default the cpu and heap profile are enabled, however any profile can be enabled
by passing in a comma seperated `DD_PROFILES` using the names [here](https://github.com/DataDog/dd-trace-go/blob/v1.42.1/profiler/profile.go#L78)

## OTLP

We do our best to support enviornment variables specified in the [Otel Spec](https://opentelemetry.io/docs/specs/otel/configuration/sdk-environment-variables/) and have added a few of our own. Key ones to note are:


## Jaeger

Pass in the `JAEGER_ENDPOINT` enviornment variable

## Pyroscope

Pass in the `PYROSCOPE_ENDPOINT` environment variable
