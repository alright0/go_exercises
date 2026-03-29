package learn_io

import (
	"io"
)

// LimitedReader
// Ограничивает чтение из передаваемого буфера размером N
// "Hello world", N=10, lenBuffer=5
// Read -> "Hello", nil
// Read -> " worl", EOF
type LimitedReader struct {
	R io.Reader
	N int
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	if r.N <= 0 {
		return 0, nil
	}
	if len(p) > r.N {
		p = p[:r.N]
	}

	c, err := r.R.Read(p)
	r.N -= c

	if r.N == 0 && err == nil {
		return c, io.EOF
	}

	return c, err
}
