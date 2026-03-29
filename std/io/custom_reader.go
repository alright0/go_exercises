package learn_io

type MyReader struct {
}

func (r *MyReader) Read(p []byte) (n int, err error) {

	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}
