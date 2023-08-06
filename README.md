# koron-go/ringbuf

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron-go/ringbuf)](https://pkg.go.dev/github.com/koron-go/ringbuf)
[![Actions/Go](https://github.com/koron-go/ringbuf/workflows/Go/badge.svg)](https://github.com/koron-go/ringbuf/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/ringbuf)](https://goreportcard.com/report/github.com/koron-go/ringbuf)

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
$ go test -bench . -benchmem
goos: windows
goarch: amd64
pkg: github.com/koron-go/ringbuf
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkRingBuffer0/single-16          134190121                8.990 ns/op           0 B/op          0 allocs/op
BenchmarkRingBuffer1/single-16          515369830                2.357 ns/op           0 B/op          0 allocs/op
BenchmarkRingBuffer2/single-16          71773768                14.63 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer2/multi-16           74173429                19.48 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer3/single-16          134209962                8.842 ns/op           0 B/op          0 allocs/op
BenchmarkRingBuffer3/multi-16           52277178                23.99 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer3B/single-16         129035656                9.440 ns/op           0 B/op          0 allocs/op
BenchmarkRingBuffer3B/multi-16          47464974                26.73 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer4/single-16          60202279                19.94 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer4/multi-16           68165551                22.11 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer4B/single-16         59372432                19.20 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer4B/multi-16          57555612                21.16 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer5/single-16          511431560                2.278 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/koron-go/ringbuf     20.594s
```

* 複数goroutine環境での評価
    * sync.Mutex(2)は意外と安定
    * sync/atomic(3)は性能劣化が著しい
    * chan(4)で良いのでは?
* 1 goroutineでの評価
    * index+書き込まれ量の管理方式(5)は、意外と良い。ただし複数goroutine化の見通しが立たないか、もしくは性能が悪そう。
