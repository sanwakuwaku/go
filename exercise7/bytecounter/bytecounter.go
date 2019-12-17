package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type ByteCounter int

type WordCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(string(p)))
	sc.Split(bufio.ScanWords)

	count := 0
	for sc.Scan() {
		count++
	}

	*c += WordCounter(count)
	if err := sc.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	writer := bufio.NewWriter(w)
	count := int64(writer.Buffered())

	return writer, &count
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	var w WordCounter
	var words = "Lorem ipsum dolor sit amet"
	fmt.Fprint(&w, words)
	fmt.Println(w)
}
