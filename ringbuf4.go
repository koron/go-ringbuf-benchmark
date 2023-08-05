package ringbuf

// RingBuffer4 chanをリングバッファにみたて、インターフェースを整えたもの
type RingBuffer4[T any] chan T

func New4[T any](size int) (RingBuffer4[T], error) {
	return make(RingBuffer4[T], size), nil
}

func (rb RingBuffer4[T]) Enqueue(v T) error {
	select {
	case rb<-v:
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
