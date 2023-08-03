package ringbuf

// RingBuffer0 ナイーブなリングバッファ
type RingBuffer0[T any] struct {
	buf []T
	rx  int
	wx  int
}

func New0[T any](size int) (*RingBuffer0[T], error) {
	return &RingBuffer0[T]{
		buf: make([]T, size),
		rx:  0,
		wx:  0,
	}, nil
}

func (rb *RingBuffer0[T]) Enqueue(v T) error {
	if rb.wx-rb.rx == rb.size() {
		return ErrFull
	}
	rb.buf[rb.wx%rb.size()] = v
	rb.wx++
	return nil
}

func (rb *RingBuffer0[T]) Dequeue() (v T, err error) {
	if rb.wx == rb.rx {
		return v, ErrEmpty
	}
	v = rb.buf[rb.rx%rb.size()]
	rb.rx++
	return v, nil
}

func (rb *RingBuffer0[T]) size() int {
	return len(rb.buf)
}
