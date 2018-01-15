package sbp

import (
	"io"
	"sync"
)

// FrameWriter writes a Frame.
type FrameWriter struct {
	wr     io.Writer
	wrLock *sync.Mutex
}

// NewFrameWriter creates a new FrameWriter instance.
func NewFrameWriter(wr io.Writer) *FrameWriter {
	fwr := &FrameWriter{
		wr:     wr,
		wrLock: new(sync.Mutex),
	}

	return fwr
}

// WriteFrame writes a single Frame.
func (fwr *FrameWriter) WriteFrame(f *Frame) (int, error) {
	bs, err := f.Bytes()
	if err != nil {
		return 0, nil
	}

	fwr.wrLock.Lock()
	defer fwr.wrLock.Unlock()

	return fwr.wr.Write(bs)
}
