package benchmark

import (
	"context"
	"testing"

	"github.com/ninestems/go-proxy-gen-becnhmark/implementation/fabric"
	"github.com/ninestems/go-proxy-gen-becnhmark/testutils"
)

func BenchmarkConverter_Gen(b *testing.B) {
	f := fabric.NewPureFabric(testutils.InitTestLogger(), testutils.InitTestTracer())
	conv := f.ProxyGen()
	ctx := context.Background()
	id := "123456789"

	// Прогрев (опционально)
	_, _ = conv.Convert(ctx, id)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = conv.Convert(ctx, id)
	}
}
