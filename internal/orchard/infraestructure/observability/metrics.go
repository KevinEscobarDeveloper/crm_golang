package observability

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	meter          = otel.Meter("crm-golang")
	requestCounter metric.Int64Counter
	successCounter metric.Int64Counter
	errorCounter   metric.Int64Counter
)

func init() {
	var err error
	requestCounter, err = meter.Int64Counter(
		"http_server_requests",
		metric.WithDescription("Total Request count"),
	)
	if err != nil {
		panic(err)
	}

	successCounter, err = meter.Int64Counter(
		"http_server_responses_success",
		metric.WithDescription("Total Response success count"),
	)
	if err != nil {
		panic(err)
	}

	errorCounter, err = meter.Int64Counter(
		"http_server_responses_errors",
		metric.WithDescription("Total responses with error"),
	)
	if err != nil {
		panic(err)
	}
}

func IncRequest(ctx context.Context, method, route string) {
	requestCounter.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("method", method),
			attribute.String("route", route),
		),
	)
}

func IncSuccess(ctx context.Context, method, route string, status int) {
	successCounter.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("method", method),
			attribute.String("route", route),
			attribute.Int("status_code", status),
		),
	)
}

func IncError(ctx context.Context, method, route string, status int) {
	errorCounter.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("method", method),
			attribute.String("route", route),
			attribute.Int("status_code", status),
		),
	)
}
