package ringbuf_test

import (
	"sync"
	"testing"

	"github.com/koron-go/ringbuf"
)

type RingBuffer[T any] interface {
	Enqueue(v T) error
	Dequeue() (T, error)
}

func benchmarkSingle(b *testing.B, rb RingBuffer[int]) {
	n := b.N / 2
	Iq := n / 1000
	Ir := n % 1000
	b.ResetTimer()
	for i := 0; i < Iq; i++ {
		for j := 0; j < 1000; j++ {
			_ = rb.Enqueue(j)
		}
		for j := 0; j < 1000; j++ {
			_, _ = rb.Dequeue()
		}
	}
	if Ir > 0 {
		for j := 0; j < Ir; j++ {
			_ = rb.Enqueue(j)
		}
		for j := 0; j < Ir; j++ {
			_, _ = rb.Dequeue()
		}
	}
}

func benchmarkMulti(b *testing.B, rb RingBuffer[int]) {
	benchmarkMulti2(b, rb)
}

func benchmarkMulti1(b *testing.B, rb RingBuffer[int]) {
	n := b.N / 2
	Iq := n / 1000
	Ir := n % 1000
	var wg sync.WaitGroup
	wg.Add(2)
	b.ResetTimer()
	//var readFail, writeFail, readTotal, writeTotal int
	go func() {
		defer wg.Done()
		for i := 0; i < Iq; i++ {
			count := 1000
			for count > 0 {
				if rb.Enqueue(count) == nil {
					count--
					//} else {
					//	writeFail++
				}
				//writeTotal++
			}
		}
		if Ir > 0 {
			count := Ir
			for count > 0 {
				if rb.Enqueue(count) == nil {
					count--
					//} else {
					//	writeFail++
				}
				//writeTotal++
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < Iq; i++ {
			count := 1000
			for count > 0 {
				if _, err := rb.Dequeue(); err == nil {
					count--
					//} else {
					//	readFail++
				}
				//readTotal++
			}
		}
		if Ir > 0 {
			count := Ir
			for count > 0 {
				if _, err := rb.Dequeue(); err == nil {
					count--
					//} else {
					//	readFail++
				}
				//readTotal++
			}
		}
	}()
	wg.Wait()
	/*
		if writeFail > 0 || readFail > 0 {
			b.Logf("N=%-9d write=%d/%d (%f) read=%d/%d (%f)", b.N,
				writeFail, writeTotal, float64(writeFail)/float64(writeTotal),
				readFail, readTotal, float64(readFail)/float64(readTotal))
		}
	*/
}

func benchmarkMulti2(b *testing.B, rb RingBuffer[int]) {
	n := b.N/2
	if n == 0 {
		n = 1
	}
	var wg sync.WaitGroup
	wg.Add(2)
	b.ResetTimer()
	go func() {
		defer wg.Done()
		count := n
		for count > 0 {
			if rb.Enqueue(count) == nil {
				count--
			}
		}
	}()
	go func() {
		defer wg.Done()
		count := n
		for count > 0 {
			if _, err := rb.Dequeue(); err == nil {
				count--
			}
		}
	}()
	wg.Wait()
}

func BenchmarkRingBuffer0(b *testing.B) {
	rb, _ := ringbuf.New0[int](2 * 1024 * 1024)
	b.Run("single", func(b *testing.B) {
		benchmarkSingle(b, rb)
		// about 12Kops/msec
	})
}

func BenchmarkRingBuffer1(b *testing.B) {
	rb, _ := ringbuf.New1[int](2 * 1024 * 1024)
	b.Run("single", func(b *testing.B) {
		benchmarkSingle(b, rb)
		// about 42Kops/msec
	})
}

func BenchmarkRingBuffer2(b *testing.B) {
	rb, _ := ringbuf.New2[int](2 * 1024 * 1024)
	b.Run("single", func(b *testing.B) {
		benchmarkSingle(b, rb)
	})
	b.Run("multi", func(b *testing.B) {
		benchmarkMulti(b, rb)
	})
}

func BenchmarkRingBuffer3(b *testing.B) {
	rb, _ := ringbuf.New3[int](2 * 1024 * 1024)
	b.Run("single", func(b *testing.B) {
		benchmarkSingle(b, rb)
	})
	b.Run("multi", func(b *testing.B) {
		benchmarkMulti(b, rb)
	})
}

func BenchmarkRingBuffer3B(b *testing.B) {
	rb, _ := ringbuf.New3B[int](2 * 1024 * 1024)
	b.Run("single", func(b *testing.B) {
		benchmarkSingle(b, rb)
	})
	b.Run("multi", func(b *testing.B) {
		benchmarkMulti(b, rb)
	})
}

func BenchmarkRingBuffer4(b *testing.B) {
	rb, _ := ringbuf.New4[int](2 * 1024 * 1024)
	b.Run("single", func(b *testing.B) {
		benchmarkSingle(b, rb)
	})
	b.Run("multi", func(b *testing.B) {
		benchmarkMulti(b, rb)
	})
}

func BenchmarkRingBuffer4B(b *testing.B) {
	rb, _ := ringbuf.New4B[int](2 * 1024 * 1024)
	b.Run("single", func(b *testing.B) {
		benchmarkSingle(b, rb)
	})
	b.Run("multi", func(b *testing.B) {
		benchmarkMulti(b, rb)
	})
}

func BenchmarkRingBuffer5(b *testing.B) {
	rb, _ := ringbuf.New5[int](2 * 1024 * 1024)
	b.Run("single", func(b *testing.B) {
		benchmarkSingle(b, rb)
	})
}
