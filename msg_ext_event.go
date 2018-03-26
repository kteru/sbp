package sbp

import (
	"encoding/binary"
	"io"
)

// MsgExtEvent represents a contents of MSG_EXT_EVENT.
type MsgExtEvent struct {
	// GPS week number (unit:weeks)
	Wn uint16

	// GPS time of week rounded to the nearest millisecond (unit:ms)
	Tow uint32

	// Nanosecond residual of millisecond-rounded TOW (unit:ns)
	NsResidual uint32

	// Status flags
	NewLevelOfPin uint8
	TimeQuality   uint8

	// Pin number (0..9)
	Pin uint8
}

func (m *MsgExtEvent) MsgType() uint16 {
	return TypeMsgExtEvent
}

func (m *MsgExtEvent) UnmarshalBinary(bs []byte) error {
	if len(bs) < 12 {
		return io.ErrUnexpectedEOF
	}

	m.Wn = binary.LittleEndian.Uint16(bs[0:2])
	m.Tow = binary.LittleEndian.Uint32(bs[2:6])
	m.NsResidual = binary.LittleEndian.Uint32(bs[6:10])

	flags := bs[10]
	m.NewLevelOfPin = flags & 0x1
	m.TimeQuality = flags >> 1 & 0x1

	m.Pin = bs[11]

	return nil
}

func (m *MsgExtEvent) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 12)

	binary.LittleEndian.PutUint16(bs[0:2], m.Wn)
	binary.LittleEndian.PutUint32(bs[2:6], m.Tow)
	binary.LittleEndian.PutUint32(bs[6:10], m.NsResidual)

	flags := (m.NewLevelOfPin & 0x1) | (m.TimeQuality & 0x1 << 1)
	bs[10] = flags

	bs[11] = m.Pin

	return bs, nil
}
