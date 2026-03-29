package learn_io

import (
	"io"
)

type TeeReader struct {
	R io.Reader
	W io.Writer
}

func (r *TeeReader) Read(p []byte) (n int, err error) {
	n, err = r.R.Read(p)

	if n > 0 {
		_, werr := r.W.Write(p[:n])
		if werr != nil {
			return n, werr
		}
	}

	return n, err
}
