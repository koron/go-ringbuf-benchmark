package ringbuf

// RingBuffer5
type RingBuffer5[T any] struct {
	buf []T
	rx  int
	wn  int
}

func New5[T any](size int) (*RingBuffer5[T], error) {
	return &RingBuffer5[T]{
		buf: make([]T, size),
	}, nil
}

func (rb *RingBuffer5[T]) Enqueue(v T) error {
	if rb.wn == rb.size() {
		return ErrFull
	}
	wx := rb.rx + rb.wn
	if wx >= rb.size() {
		wx -= rb.size()
	}
	rb.buf[wx] = v
	rb.wn++
	return nil
}

func (rb *RingBuffer5[T]) Dequeue() (v T, err error) {
	if rb.wn == 0 {
		return v, ErrEmpty
	}
	v = rb.buf[rb.rx]
	rb.rx++
	if rb.rx == rb.size() {
		rb.rx = 0
	}
	rb.wn--
	return v, nil
}

func (rb *RingBuffer5[T]) size() int {
	return len(rb.buf)
}
