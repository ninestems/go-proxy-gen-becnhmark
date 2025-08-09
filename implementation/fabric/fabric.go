package fabric

import (
	traceprovider "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"

	"github.com/ninestems/go-proxy-gen-benchmark/implementation"
	"github.com/ninestems/go-proxy-gen-benchmark/implementation/proxy"
	proxyhand "github.com/ninestems/go-proxy-gen-benchmark/implementation/proxy_hand"
	"github.com/ninestems/go-proxy-gen-benchmark/implementation/pure"
)

type ConvertFabric struct {
	logger *zap.Logger
	tracer *traceprovider.TracerProvider
}

func NewPureFabric(
	logger *zap.Logger,
	tracer *traceprovider.TracerProvider,
) *ConvertFabric {
	return &ConvertFabric{
		logger: logger,
		tracer: tracer,
	}
}

func (f *ConvertFabric) Pure() implementation.Converter {
	return pure.New()
}

func (f *ConvertFabric) ProxyHand() implementation.Converter {
	var out implementation.Converter
	out = pure.New()
	out = proxyhand.NewConverterProxyLogger(out, f.logger)
	out = proxyhand.NewConverterProxyTracer(out, f.tracer)
	return out
}

func (f *ConvertFabric) ProxyGen() implementation.Converter {
	var out implementation.Converter
	out = pure.New()
	out = proxy.NewConverterProxyLogger(out, f.logger)
	out = proxy.NewConverterProxyTracer(out, f.tracer)
	return out
}
