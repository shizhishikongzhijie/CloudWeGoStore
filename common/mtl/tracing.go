package mtl

import "github.com/kitex-contrib/obs-opentelemetry/provider"

func InitTracing(serviceName string) provider.OtelProvider {
	// exporter, err := otlptracegrpc.New(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// server.RegisterShutdownHook(func() {
	// 	exporter.Shutdown(context.Background()) //nolint:errcheck
	// })
	// processor := tracesdk.NewBatchSpanProcessor(exporter)
	// res, err := resource.New(context.Background(), resource.WithAttributes(semconv.ServiceNameKey.String(serviceName)))
	// if err != nil {
	// 	res = resource.Default()
	// }
	// TracerProvider = tracesdk.NewTracerProvider(tracesdk.WithSpanProcessor(processor), tracesdk.WithResource(res))
	// otel.SetTracerProvider(TracerProvider)
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
	)
	return p
}
