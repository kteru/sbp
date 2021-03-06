package sbp

import (
	"encoding/binary"
	"io"
)

// MsgBaselineNed represents a contents of MSG_BASELINE_NED.
type MsgBaselineNed struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	// NED coordinates (unit:mm)
	N int32
	E int32
	D int32

	// Horizontal position accuracy estimate (unit:mm)
	HAccuracy uint16

	// Vertical position accuracy estimate (unit:mm)
	VAccuracy uint16

	// Number of satellites used in solution
	NSats uint8

	// Status flags
	FixMode uint8
}

// MsgType returns the number representing the type.
func (m *MsgBaselineNed) MsgType() uint16 {
	return TypeMsgBaselineNed
}

// UnmarshalBinary parses a byte slice.
func (m *MsgBaselineNed) UnmarshalBinary(bs []byte) error {
	if len(bs) < 22 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])

	m.N = int32(binary.LittleEndian.Uint32(bs[4:8]))
	m.E = int32(binary.LittleEndian.Uint32(bs[8:12]))
	m.D = int32(binary.LittleEndian.Uint32(bs[12:16]))

	m.HAccuracy = binary.LittleEndian.Uint16(bs[16:18])
	m.VAccuracy = binary.LittleEndian.Uint16(bs[18:20])

	m.NSats = bs[20]

	flags := bs[21]
	m.FixMode = flags & 0x7

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgBaselineNed) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 22)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)

	binary.LittleEndian.PutUint32(bs[4:8], uint32(m.N))
	binary.LittleEndian.PutUint32(bs[8:12], uint32(m.E))
	binary.LittleEndian.PutUint32(bs[12:16], uint32(m.D))

	binary.LittleEndian.PutUint16(bs[16:18], m.HAccuracy)
	binary.LittleEndian.PutUint16(bs[18:20], m.VAccuracy)

	bs[20] = m.NSats

	flags := m.FixMode & 0x7
	bs[21] = flags

	return bs, nil
}
