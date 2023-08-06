package ringbuf

import "sync"

// RingBuffer2 sync.Mutexを用いてマルチスレッド対応したもの
type RingBuffer2[T any] struct {
	buf []T
	rx  int
	wx  int
	mu  sync.Mutex
}

func New2[T any](size int) (*RingBuffer2[T], error) {
	return &RingBuffer2[T]{
		buf: make([]T, size),
		rx:  0,
		wx:  0,
	}, nil
}

func (rb *RingBuffer2[T]) Enqueue(v T) error {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	if rb.wx-rb.rx == len(rb.buf) {
		return ErrFull
	}
	rb.buf[rb.wx&rb.mask()] = v
	rb.wx++
	return nil
}

func (rb *RingBuffer2[T]) Dequeue() (v T, err error) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	if rb.wx == rb.rx {
		return v, ErrEmpty
	}
	v = rb.buf[rb.rx&rb.mask()]
	rb.rx++
	return v, nil
}

func (rb *RingBuffer2[T]) mask() int {
	return len(rb.buf) - 1
}
