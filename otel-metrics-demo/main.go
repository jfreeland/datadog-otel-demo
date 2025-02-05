package main

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func main() {
	ctx := context.TODO()

	metricExporter, err := otlpmetrichttp.New(ctx, otlpmetrichttp.WithEndpoint("opentelemetry-collector.platform.svc.cluster.local:4318"), otlpmetrichttp.WithInsecure())
	if err != nil {
		panic(err)
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter, sdkmetric.WithInterval(15))),
	)
	defer meterProvider.Shutdown(context.Background())

	meter := meterProvider.Meter("dumb")

	counterRequest, err := meter.Int64Counter("joey_requests_home_total", metric.WithDescription("home requests"))
	if err != nil {
		panic(err)
	}
	counterHealth, err := meter.Int64Counter("joey_requests_health_total", metric.WithDescription("health endpoint requests"))
	if err != nil {
		panic(err)
	}

	homeAttr := attribute.String("endpoint", "/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counterRequest.Add(ctx, 1, metric.WithAttributes(homeAttr))
		w.Write([]byte("hello world"))
	})

	healthAttr := attribute.String("endpoint", "/health")
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		counterHealth.Add(ctx, 1, metric.WithAttributes(healthAttr))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":8080", nil)
}
