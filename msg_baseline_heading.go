package sbp

import (
	"encoding/binary"
	"io"
)

// MsgBaselineHeading represents a contents of MSG_BASELINE_HEADING.
type MsgBaselineHeading struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	// Heading (unit:mdeg)
	Heading uint32

	// Number of satellites used in solution
	NSats uint8

	// Status flags
	FixMode uint8
}

// MsgType returns the number representing the type.
func (m *MsgBaselineHeading) MsgType() uint16 {
	return TypeMsgBaselineHeading
}

// UnmarshalBinary parses a byte slice.
func (m *MsgBaselineHeading) UnmarshalBinary(bs []byte) error {
	if len(bs) < 10 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])
	m.Heading = binary.LittleEndian.Uint32(bs[4:8])

	m.NSats = bs[8]

	flags := bs[9]
	m.FixMode = flags & 0x7

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgBaselineHeading) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 10)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)
	binary.LittleEndian.PutUint32(bs[4:8], m.Heading)

	bs[8] = m.NSats

	flags := m.FixMode & 0x7
	bs[9] = flags

	return bs, nil
}
