package performance

import (
	"testing"

	"eiyaro/util"
)

// Test rpc call benchmark - 0.1 s/op
func BenchmarkRpcCall(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		util.ClientCall("/net-info")
	}
}
