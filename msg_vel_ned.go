package sbp

import (
	"encoding/binary"
	"io"
)

// MsgVelNed represents a contents of MSG_VEL_NED.
type MsgVelNed struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	// NED coordinates (unit:mm/s)
	N int32
	E int32
	D int32

	// Horizontal velocity accuracy estimate (unit:mm/s)
	HAccuracy uint16

	// Vertical velocity accuracy estimate (unit:mm/s)
	VAccuracy uint16

	// Number of satellites used in solution
	NSats uint8

	// Status flags
	VelocityMode           uint8
	InertialNavigationMode uint8
}

// MsgType returns the number representing the type.
func (m *MsgVelNed) MsgType() uint16 {
	return TypeMsgVelNed
}

// UnmarshalBinary parses a byte slice.
func (m *MsgVelNed) UnmarshalBinary(bs []byte) error {
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
	m.VelocityMode = flags & 0x7
	m.InertialNavigationMode = flags >> 3 & 0x3

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgVelNed) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 22)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)

	binary.LittleEndian.PutUint32(bs[4:8], uint32(m.N))
	binary.LittleEndian.PutUint32(bs[8:12], uint32(m.E))
	binary.LittleEndian.PutUint32(bs[12:16], uint32(m.D))

	binary.LittleEndian.PutUint16(bs[16:18], m.HAccuracy)
	binary.LittleEndian.PutUint16(bs[18:20], m.VAccuracy)

	bs[20] = m.NSats

	flags := (m.VelocityMode & 0x7) | (m.InertialNavigationMode & 0x3 << 3)
	bs[21] = flags

	return bs, nil
}
