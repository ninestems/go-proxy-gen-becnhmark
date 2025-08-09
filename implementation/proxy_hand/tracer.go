package proxyhand

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	traceprovider "go.opentelemetry.io/otel/sdk/trace"
	trace "go.opentelemetry.io/otel/trace"

	source "github.com/ninestems/go-proxy-gen-becnhmark/implementation"
)

// ConverterProxyTracer proxy wrapper for source.Converter.
type ConverterProxyTracer struct {
	src   source.Converter
	trace trace.Tracer
}

// NewConverterProxyTracer creates a new proxy trace for source.Converter.
func NewConverterProxyTracer(
	src source.Converter,
	tp *traceprovider.TracerProvider,
) *ConverterProxyTracer {
	return &ConverterProxyTracer{
		src:   src,
		trace: tp.Tracer("ConverterProxyTracer"),
	}
}

// Convert is proxy method for source.Converter.Convert.
func (p *ConverterProxyTracer) Convert(ctx context.Context, id string) (out0 int64, out1 error) {
	ctx, span := p.trace.Start(ctx, "ConverterProxyTracer.Convert() started")
	defer span.End()
	traceID, _ := ctx.Value("traceID").(string)
	span.SetAttributes(
		attribute.String("trace_id", traceID),
		attribute.String("id", id),
	)
	out0, out1 = p.src.Convert(ctx, id)
	if out1 != nil {
		span.RecordError(out1)
		span.SetStatus(codes.Error, out1.Error())
	} else {
		span.SetAttributes(
			attribute.Int64("out_alias", out0),
		)
		span.SetStatus(codes.Ok, "ConverterProxyTracer.Convert() success")
	}

	return
}
