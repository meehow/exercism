package paasio

import (
	"io"
	"sync"
)

type rwCounter struct {
	reader io.Reader
	writer io.Writer
	sync.Mutex
	rn    int64
	wn    int64
	rnops int
	wnops int
}

func (c *rwCounter) Read(p []byte) (int, error) {
	n, err := c.reader.Read(p)
	c.Lock()
	c.rn += int64(n)
	c.rnops++
	c.Unlock()
	return n, err
}

func (c *rwCounter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	c.Lock()
	c.wn += int64(n)
	c.wnops++
	c.Unlock()
	return n, err
}

func (c *rwCounter) ReadCount() (int64, int) {
	c.Lock()
	n, nops := c.rn, c.rnops
	c.Unlock()
	return n, nops
}

func (c *rwCounter) WriteCount() (int64, int) {
	c.Lock()
	n, nops := c.wn, c.wnops
	c.Unlock()
	return n, nops
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &rwCounter{reader: r}
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &rwCounter{writer: w}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &rwCounter{reader: rw, writer: rw}
}
