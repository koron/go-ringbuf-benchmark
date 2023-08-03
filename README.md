# koron-go/ringbuf

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron-go/ringbuf)](https://pkg.go.dev/github.com/koron-go/ringbuf)
[![Actions/Go](https://github.com/koron-go/ringbuf/workflows/Go/badge.svg)](https://github.com/koron-go/ringbuf/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/ringbuf)](https://goreportcard.com/report/github.com/koron-go/ringbuf)

See <https://kumagi.hatenablog.com/entry/ring-buffer>

## Benchmark results

```console
$ go test -bench . -benchmem
goos: windows
goarch: amd64
pkg: github.com/koron-go/ringbuf
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkRingBuffer0/single-16          130561113                9.071 ns/op           0 B/op          0 allocs/op
BenchmarkRingBuffer1/single-16          493192504                2.340 ns/op           0 B/op          0 allocs/op
BenchmarkRingBuffer2/single-16          83521255                14.77 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer2/multi-16           62141391                19.87 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer3/single-16          132964179                8.926 ns/op           0 B/op          0 allocs/op
BenchmarkRingBuffer3/multi-16           45100573                24.89 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer4/single-16          59892492                19.70 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer4/multi-16           57575494                21.68 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer5/single-16          515780518                2.313 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/koron-go/ringbuf     13.304s
```

* 複数goroutine環境での評価
    * sync.Mutex(2)は意外と安定
    * sync/atomic(3)は性能劣化が著しい
    * chan(4)で良いのでは?
* 1 goroutineでの評価
    * index+書き込まれ量の管理方式(5)は、意外と良い。ただし複数goroutine化の見通しが立たないか、もしくは性能が悪そう。
