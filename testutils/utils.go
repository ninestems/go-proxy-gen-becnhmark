package testutils

import (
	"context"
	"io"

	"go.opentelemetry.io/otel"
	traceprovider "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitTestLogger() *zap.Logger {
	nullCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(io.Discard), // -> null
		zapcore.DebugLevel,          // log level
	)
	return zap.New(nullCore)
}

type noopExporter struct{}

func (n *noopExporter) ExportSpans(_ context.Context, _ []traceprovider.ReadOnlySpan) error {
	return nil
}
func (n *noopExporter) Shutdown(_ context.Context) error {
	return nil
}

func InitTestTracer() *traceprovider.TracerProvider {
	noOpExporter := traceprovider.NewSimpleSpanProcessor(&noopExporter{})
	tp := traceprovider.NewTracerProvider(
		traceprovider.WithSpanProcessor(noOpExporter),
	)
	otel.SetTracerProvider(tp)
	return tp
}
