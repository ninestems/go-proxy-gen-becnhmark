package proxyhand

import (
	"context"

	zap "go.uber.org/zap"

	source "github.com/ninestems/go-proxy-gen-benchmark/implementation"
)

// ConverterProxyLogger proxy wrapper for source.Converter.
type ConverterProxyLogger struct {
	src source.Converter
	log *zap.Logger
}

// NewConverterProxyLogger creates a new proxy logger for source.Converter.
func NewConverterProxyLogger(
	src source.Converter,
	log *zap.Logger,
) *ConverterProxyLogger {
	return &ConverterProxyLogger{
		src: src,
		log: log.Named("ConverterProxyLogger"),
	}
}

// Convert is proxy method for source.Converter.Convert.
func (p *ConverterProxyLogger) Convert(ctx context.Context, id string) (out0 int64, out1 error) {
	p.log.Info(
		"Convert() started",
		zap.Any("trace_id", ctx.Value("traceID")),
		zap.String("id", id),
	)
	out0, out1 = p.src.Convert(ctx, id)
	if out1 != nil {
		p.log.Info(
			"Convert() ends with error",
			zap.Error(out1),
		)
	} else {
		p.log.Info(
			"Convert() ends with success",
			zap.Int64("out_alias", out0),
		)
	}

	return
}
