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
	NumSats uint8

	// Status flags
	FixMode    uint8
	RaimRepair uint8
}

func (m *MsgBaselineHeading) UnmarshalBinary(bs []byte) error {
	if len(bs) < 10 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])
	m.Heading = binary.LittleEndian.Uint32(bs[4:8])

	m.NumSats = bs[8]

	flags := bs[9]
	m.FixMode = flags & 0x7
	m.RaimRepair = flags >> 7 & 0x1

	return nil
}

func (m *MsgBaselineHeading) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 10)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)
	binary.LittleEndian.PutUint32(bs[4:8], m.Heading)

	bs[8] = m.NumSats

	flags := (m.FixMode & 0x7) | (m.RaimRepair & 0x1 << 7)
	bs[9] = flags

	return bs, nil
}
