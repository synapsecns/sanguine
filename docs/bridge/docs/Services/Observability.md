All off-chain systems are by default observable and configured through the [metrics](https://pkg.go.dev/github.com/synapsecns/sanguine/core/metrics#section-readme) package. The observability stack is built around [open telemetry](https://opentelemetry.io/) with metrics also being exported using this standard.

"Metrics" themselves are divided into 3 different types of metrics:

1. **Traces**: These are used to trace the flow of a request through the system. Our code uses traces for all sorts of events, from the start of a request to the end of a request. Traces are used to debug and profile the system. Traces are made up of spans, which are individual events that occur during the lifetime of a trace. Traces are aggregated ex-post facto by observability solutions. Please see [this article](https://grafana.com/docs/tempo/latest/introduction/) for a good overview.
2. **Metrics**: Metrics consist of things like gauges, counters, and histograms. These are used to monitor the health of the system. Metrics are aggregated in real-time by observability solutions. Please see [this article](https://grafana.com/docs/grafana-cloud/introduction/) for a good overview. We use the Prometheus exporter for metrics and by default make them available on `METRICS_PORT` at `METRICS_PATH`.
3. **APM (Application Performance Monitoring)**: APM is a combination of traces and metrics. It is used to monitor the performance of the system. We use Pyroscope for APM and by default make it available on `PYROSCOPE_ENDPOINT`.

:::tip
Sanguine is an open-source repo and tries to be as unopinionated as possible when it comes to which libraries it uses for observability. If you're interested in maintaining an unsupported module as part of the [metrics](https://pkg.go.dev/github.com/synapsecns/sanguine/core/metrics#section-readme) package, pull requests are always welcome!
:::

### Configuring Observability

**Configuring Tracing**:

1. Pick a vendor that supports the [OTLP](https://opentelemetry.io/ecosystem/vendors/) standard. We recommend Signoz for OSS.
2. Set the `METRICS_HANDLER` environment variable to `OTLP`. If you're using a legacy jaeger instance, you can also use `JAEGER`
3. Set the `OTLP_ENDPOINT` environment variable to the endpoint of your observability solution.

**Configuring Metrics**

1. Pick a vendor that supports the [prometheus](https://prometheus.io/docs/instrumenting/exporters/) standard. We recommend Grafana for OSS.
2. Set the `METRICS_PORT_ENABLED` environment variable to `true` to enable the metrics endpoint.
3. Find your metrics on `METRICS_PORT` at `METRICS_PATH` and configure scrapers.

**Configuring APM**

Set the `PYROSCOPE_ENDPOINT` environment variable to the endpoint of your [pyroscope instance](https://pyroscope.io/).
