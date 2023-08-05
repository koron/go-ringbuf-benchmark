package ringbuf

// RingBuffer4 chanをリングバッファにみたて、インターフェースを整えたもの
type RingBuffer4[T any] chan T

func New4[T any](size int) (RingBuffer4[T], error) {
	return make(RingBuffer4[T], size), nil
}

func (rb RingBuffer4[T]) Enqueue(v T) error {
	select {
	case rb <- v:
		return nil
	default:
		return ErrFull
	}
}

func (rb RingBuffer4[T]) Dequeue() (v T, err error) {
	select {
	case v = <-rb:
		return v, nil
	default:
		return v, ErrEmpty
	}
}

// RingBuffer4B Enqueue/Dequeueでエラーを返さずにブロックするもの
type RingBuffer4B[T any] chan T

func New4B[T any](size int) (RingBuffer4B[T], error) {
	return make(RingBuffer4B[T], size), nil
}

func (rb RingBuffer4B[T]) Enqueue(v T) error {
	rb <- v
	return nil
}

func (rb RingBuffer4B[T]) Dequeue() (v T, err error) {
	v = <-rb
	return v, nil
}
