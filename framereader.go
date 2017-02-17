package sbp

import (
	"io"
)

// FrameReader reads Frame.
type FrameReader struct {
	rd io.Reader
}

// NewFrameReader creates a new FrameReader instance.
func NewFrameReader(rd io.Reader) *FrameReader {
	frd := &FrameReader{
		rd: rd,
	}

	return frd
}

// Next returns a next Frame.
func (frd *FrameReader) Next() (*Frame, error) {
	bs := make([]byte, 6, 6+2)

	for {
		if _, err := io.ReadFull(frd.rd, bs[:1]); err != nil {
			return nil, err
		}
		if bs[0] == FramePreamble {
			break
		}
	}

	if _, err := io.ReadFull(frd.rd, bs[1:]); err != nil {
		return nil, err
	}

	plen := int(bs[5])

	bs = append(bs, make([]byte, plen+2)...)

	if _, err := io.ReadFull(frd.rd, bs[6:]); err != nil {
		return nil, err
	}

	return NewFrame(bs)
}
