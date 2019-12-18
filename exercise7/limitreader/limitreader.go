package limitreader

import (
	"errors"
	"io"
)

type limitReader struct {
	n      int64
	reader io.Reader
}

func (r limitReader) Read(p []byte) (n int, err error) {
	if len(p) < n {
		return 0, errors.New("slice length is short.")
	}

	buf := make([]byte, r.n)
	size, err := io.ReadFull(r.reader, buf)
	if err != nil {
		return size, err
	}

	copy(p, buf)

	return size, io.EOF
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return newReader(r, n)
}

func newReader(r io.Reader, n int64) io.Reader {
	return limitReader{n: n, reader: r}
}
