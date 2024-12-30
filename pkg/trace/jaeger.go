package trace

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type TraceConfig struct {
	Endpoint   string
	Name       string
	TraceRatio float64 // 1.0 100%采样率
}

func newExporter(ctx context.Context, endpoint string) (trace.SpanExporter, error) {
	return otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithInsecure(),
	)
}

// 设置全局trace
func InitTracer(c *TraceConfig) error {
	if c == nil {
		return errors.New("init tracer config is nil")
	}
	exp, err := newExporter(context.Background(), c.Endpoint)
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		// 将基于父span的采样率设置为100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(c.TraceRatio))),
		// 始终确保在生产中批量处理
		tracesdk.WithBatcher(exp),
		// 在资源中记录有关此应用程序的信息
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(c.Name),
			attribute.String("exporter", "jaeger"),
			// attribute.Float64("float", 312.23),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
