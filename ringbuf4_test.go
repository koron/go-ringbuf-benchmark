package ringbuf_test

import (
	"sync"
	"testing"

	"github.com/koron-go/ringbuf"
)

func BenchmarkRingBuffer4(b *testing.B) {
	rb, _ := ringbuf.New4[int](2 * 1024 * 1024)

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

	b.Run("multi", func(b *testing.B) {
		n := max(1, b.N/2)
		var wg sync.WaitGroup
		wg.Add(2)
		b.ResetTimer()
		go func(m int) {
			for m > 0 {
				if rb.Enqueue(m) == nil {
					m--
				}
			}
			wg.Done()
		}(n)
		go func(m int) {
			for m > 0 {
				if _, err := rb.Dequeue(); err == nil {
					m--
				}
			}
			wg.Done()
		}(n)
		wg.Wait()
	})
}

func BenchmarkRingBuffer4B(b *testing.B) {
	rb, _ := ringbuf.New4B[int](2 * 1024 * 1024)

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

	b.Run("multi", func(b *testing.B) {
		n := max(1, b.N/2)
		var wg sync.WaitGroup
		wg.Add(2)
		b.ResetTimer()
		go func(m int) {
			for m > 0 {
				if rb.Enqueue(m) == nil {
					m--
				}
			}
			wg.Done()
		}(n)
		go func(m int) {
			for m > 0 {
				if _, err := rb.Dequeue(); err == nil {
					m--
				}
			}
			wg.Done()
		}(n)
		wg.Wait()
	})
}
