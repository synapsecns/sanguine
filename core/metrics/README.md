# Metrics

Metrics are supported to be pushed to both datadog and prometheus.

## Datadog:

By default the cpu and heap profile are enabled, however any profile can be enabled
by passing in a comma seperated `DD_PROFILES` using the names [here](https://github.com/DataDog/dd-trace-go/blob/v1.42.1/profiler/profile.go#L78)

## New Relic

Pass in the `NEW_RELIC_LICENSE_KEY` enviornment variable

## Jaeger

Pass in the `JAEGER_ENDPOINT` enviornment variable

## Pyroscope

Pass in the `PYROSCOPE_ENDPOINT` environment variable
