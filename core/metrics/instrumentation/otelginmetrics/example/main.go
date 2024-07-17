package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/metrics/instrumentation/otelginmetrics"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func initMetrics() {
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		panic(err)
	}
	reader := metric.NewPeriodicReader(metricExporter, metric.WithInterval(1*time.Second))

	fmt.Println(metricExporter)
	res, err := resource.New(context.Background(),
		resource.WithAttributes(semconv.ServiceNameKey.String("PG2")),
		resource.WithAttributes(semconv.ServiceNamespaceKey.String("Spark")),
		resource.WithSchemaURL(semconv.SchemaURL),
	)
	if err != nil {
		panic(err)
	}
	metricProvider := metric.NewMeterProvider(metric.WithReader(reader), metric.WithResource(res))
	otel.SetMeterProvider(metricProvider)
}

func main() {
	router := gin.New()
	initMetrics()
	router.Use(otelginmetrics.Middleware(
		"TEST-SERVICE",
		// Custom attributes
		otelginmetrics.WithAttributes(func(serverName, route string, request *http.Request) []attribute.KeyValue {
			return append(otelginmetrics.DefaultAttributes(serverName, route, request), attribute.String("Custom-attribute", "value"))
		}),
	))

	logic := func(ctx *gin.Context, sleep int) {
		// xxx, _ := strconv.Atoi(ctx.Param("xxx"))
		time.Sleep(time.Duration(sleep) * time.Second)
	}

	router.GET("/test/:xxx", func(ctx *gin.Context) {
		logic(ctx, 1)
		ctx.JSON(200, map[string]string{
			"productId": ctx.Param("xxx"),
		})
	})

	go func() {
		for {
			_, _ = http.DefaultClient.Get("http://localhost:9199/test/1")
		}
	}()

	_ = router.Run(":9199")

}
