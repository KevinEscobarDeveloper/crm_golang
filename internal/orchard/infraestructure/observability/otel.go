package observability

import (
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"log"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type Config struct {
	ServiceName  string
	OTLPEndpoint string
	MetricsAddr  string
}

func Init(cfg Config) (shutdown func(context.Context) error) {
	/* ---------- Tracing ---------- */
	traceExp, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpoint(cfg.OTLPEndpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExp),
		sdktrace.WithResource(resource.Default()),
	)
	otel.SetTracerProvider(tp)

	/* ---------- Metrics ---------- */
	metricsExp, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}

	mp := metric.NewMeterProvider(metric.WithReader(metricsExp))
	otel.SetMeterProvider(mp)

	/* ---------- /metrics ---------- */
	go func() {
		http.Handle("/metrics", metricsExp)
		log.Println("Prometheus /metrics at", cfg.MetricsAddr)
		log.Fatal(http.ListenAndServe(cfg.MetricsAddr, nil))
	}()

	return tp.Shutdown
}
