package ringbuf

// RingBuffer1 サイズが2^nであることを仮定して、剰余演算をビット積(&)で代用したもの
type RingBuffer1[T any] struct {
	buf []T
	rx  int
	wx  int
}

func New1[T any](size int) (*RingBuffer1[T], error) {
	return &RingBuffer1[T]{
		buf: make([]T, size),
		rx:  0,
		wx:  0,
	}, nil
}

func (rb *RingBuffer1[T]) Enqueue(v T) error {
	if rb.wx-rb.rx == len(rb.buf) {
		return ErrFull
	}
	rb.buf[rb.wx&(len(rb.buf)-1)] = v
	rb.wx++
	return nil
}

func (rb *RingBuffer1[T]) Dequeue() (v T, err error) {
	if rb.wx == rb.rx {
		return v, ErrEmpty
	}
	v = rb.buf[rb.rx&(len(rb.buf)-1)]
	rb.rx++
	return v, nil
}
