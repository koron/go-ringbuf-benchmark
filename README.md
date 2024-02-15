# koron/go-ringbuf-benchmark

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron/go-ringbuf-benchmark)](https://pkg.go.dev/github.com/koron/go-ringbuf-benchmark)
[![Actions/Go](https://github.com/koron/go-ringbuf-benchmark/workflows/Go/badge.svg)](https://github.com/koron/go-ringbuf-benchmark/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron/go-ringbuf-benchmark)](https://goreportcard.com/report/github.com/koron/go-ringbuf-benchmark)

See <https://kumagi.hatenablog.com/entry/ring-buffer>

## Variations

Type         | Desc.
-------------|-------
RingBuffer0  | Naive implementation with write and read indexes with modulo
RingBuffer1  | Use bit **and** op instead of modulo (based RingBuffer0)
RingBuffer2  | Support goroutine with `sync.Mutex` (based RingBuffer1)
RingBuffer3  | Use sync/atomic for multi-goroutine (based RingBuffer1)
RingBuffer3B | Same as RingBuffer3 but w/o index caches
RingBuffer4  | Use chan for multi-goroutine
RingBuffer4B | Use chan for multi-goroutine with blocking
RingBuffer5  | Use read index and capacity length w/o multi-goroutine supports

## Benchmark results

```console
goos: windows
goarch: amd64
pkg: github.com/koron/go-ringbuf-benchmark
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkRingBuffer0/single-16          178371265                6.526 ns/op
BenchmarkRingBuffer1/single-16          786034266                1.505 ns/op
BenchmarkRingBuffer2/single-16          85309674                12.78 ns/op
BenchmarkRingBuffer2/multi-16           69087810                16.26 ns/op
BenchmarkRingBuffer3/single-16          140011688                8.641 ns/op
BenchmarkRingBuffer3/multi-16           37608711                30.85 ns/op
BenchmarkRingBuffer3B/single-16         139669322                8.494 ns/op
BenchmarkRingBuffer3B/multi-16          56262482                33.03 ns/op
BenchmarkRingBuffer4/single-16          57601750                20.16 ns/op
BenchmarkRingBuffer4/multi-16           57848888                21.95 ns/op
BenchmarkRingBuffer4B/single-16         58735713                19.57 ns/op
BenchmarkRingBuffer4B/multi-16          55689105                21.48 ns/op
BenchmarkRingBuffer5/single-16          665201374                1.753 ns/op
PASS
ok      github.com/koron/go-ringbuf-benchmark   19.022s
```

* 複数goroutine環境での評価
    * sync.Mutex(2)は意外と安定
    * sync/atomic(3)は性能劣化が著しい
    * chan(4)で良いのでは?
* 1 goroutineでの評価
    * index+書き込まれ量の管理方式(5)は、意外と良い。ただし複数goroutine化の見通しが立たないか、もしくは性能が悪そう。
