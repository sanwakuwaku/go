package limitreader

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestLimitReader(t *testing.T) {
	r := bytes.NewBuffer([]byte{0x01, 0x1b, 0x5c, 0xff})
	lr := LimitReader(r, 3)

	buf := make([]byte, 4)
	n, err := lr.Read(buf)
	if err != io.EOF {
		t.Errorf("read error. %v\n", err)
	}
	fmt.Printf("read bytes %x, size=%d\n", buf, n)
}
