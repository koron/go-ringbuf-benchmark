package ringbuf_test

import (
	"testing"

	ringbuf "github.com/koron/go-ringbuf-benchmark"
)

func BenchmarkRingBuffer0(b *testing.B) {
	rb, _ := ringbuf.New0[int](2 * 1024 * 1024)

	b.Run("single", func(b *testing.B) {
		b.ResetTimer()
		for n := max(1, b.N/2); n > 0; n -= 1000 {
			m := min(n, 1000)
			for i := 0; i < m; i++ {
				_ = rb.Enqueue(i)
			}
			for i := 0; i < m; i++ {
				_, _ = rb.Dequeue()
			}
		}
	})
}
