package ringbuf

import "sync/atomic"

// RingBuffer3 sync/atomicを使ってインデックスをキャッシュしたもの
type RingBuffer3[T any] struct {
	buf []T
	rx  atomic.Int64
	wx  atomic.Int64
	rc  int64
	wc  int64
}

func New3[T any](size int) (*RingBuffer3[T], error) {
	return &RingBuffer3[T]{
		buf: make([]T, size),
	}, nil
}

func (rb *RingBuffer3[T]) Enqueue(v T) error {
	wx := rb.wx.Load()
	if wx-rb.rc == rb.size() {
		rb.rc = rb.rx.Load()
		if wx-rb.rc == rb.size() {
			return ErrFull
		}
	}
	rb.buf[wx&(rb.size()-1)] = v
	rb.wx.Add(1)
	return nil
}

func (rb *RingBuffer3[T]) Dequeue() (v T, err error) {
	rx := rb.rx.Load()
	if rb.wc == rx {
		rb.wc = rb.wx.Load()
		if rb.wc == rx {
			return v, ErrEmpty
		}
	}
	v = rb.buf[rx&(rb.size()-1)]
	rb.rx.Add(1)
	return v, nil
}

func (rb *RingBuffer3[T]) size() int64 {
	return int64(len(rb.buf))
}

// RingBuffer3B RingBuffer3のインデックスキャッシュは使わない版
type RingBuffer3B[T any] struct {
	buf []T
	rx  atomic.Int64
	wx  atomic.Int64
}

func New3B[T any](size int) (*RingBuffer3B[T], error) {
	return &RingBuffer3B[T]{
		buf: make([]T, size),
	}, nil
}

func (rb *RingBuffer3B[T]) Enqueue(v T) error {
	wx := rb.wx.Load()
	if wx-rb.rx.Load() == rb.size() {
		return ErrFull
	}
	rb.buf[wx&(rb.size()-1)] = v
	rb.wx.Add(1)
	return nil
}

func (rb *RingBuffer3B[T]) Dequeue() (v T, err error) {
	rx := rb.rx.Load()
	if rb.wx.Load() == rx {
		return v, ErrEmpty
	}
	v = rb.buf[rx&(rb.size()-1)]
	rb.rx.Add(1)
	return v, nil
}

func (rb *RingBuffer3B[T]) size() int64 {
	return int64(len(rb.buf))
}
