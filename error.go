package ringbuf

import "errors"

var (
	ErrFull  = errors.New("full of ring buffer")
	ErrEmpty = errors.New("empty ring buffer")
)
