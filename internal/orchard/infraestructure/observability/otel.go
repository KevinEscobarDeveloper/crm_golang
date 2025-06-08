package observability

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type Config struct {
	ServiceName  string
	OtlpEndpoint string
}

func Init(cfg Config) (func(context.Context) error, http.Handler, error) {
	// Tracing
	traceExp, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpoint(cfg.OtlpEndpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExp),
		sdktrace.WithResource(resource.Default()),
	)
	otel.SetTracerProvider(tp)

	// Metrics
	metricsExp, err := prometheus.New()
	if err != nil {
		return nil, nil, err
	}
	mp := metric.NewMeterProvider(metric.WithReader(metricsExp))
	otel.SetMeterProvider(mp)

	return tp.Shutdown, promhttp.Handler(), nil

}
