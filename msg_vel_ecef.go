package sbp

import (
	"encoding/binary"
	"io"
)

// MsgVelEcef represents a contents of MSG_VEL_ECEF.
type MsgVelEcef struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	// ECEF coordinates (unit:mm/s)
	X int32
	Y int32
	Z int32

	// Position accuracy estimate (unit:mm/s)
	Accuracy uint16

	// Number of satellites used in solution
	NSats uint8

	// Status flags
	VelocityMode           uint8
	InertialNavigationMode uint8
}

// MsgType returns the number representing the type.
func (m *MsgVelEcef) MsgType() uint16 {
	return TypeMsgVelEcef
}

// UnmarshalBinary parses a byte slice.
func (m *MsgVelEcef) UnmarshalBinary(bs []byte) error {
	if len(bs) < 20 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])

	m.X = int32(binary.LittleEndian.Uint32(bs[4:8]))
	m.Y = int32(binary.LittleEndian.Uint32(bs[8:12]))
	m.Z = int32(binary.LittleEndian.Uint32(bs[12:16]))

	m.Accuracy = binary.LittleEndian.Uint16(bs[16:18])

	m.NSats = bs[18]

	flags := bs[19]
	m.VelocityMode = flags & 0x7
	m.InertialNavigationMode = flags >> 3 & 0x3

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgVelEcef) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 20)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)

	binary.LittleEndian.PutUint32(bs[4:8], uint32(m.X))
	binary.LittleEndian.PutUint32(bs[8:12], uint32(m.Y))
	binary.LittleEndian.PutUint32(bs[12:16], uint32(m.Z))

	binary.LittleEndian.PutUint16(bs[16:18], m.Accuracy)

	bs[18] = m.NSats

	flags := (m.VelocityMode & 0x7) | (m.InertialNavigationMode & 0x3 << 3)
	bs[19] = flags

	return bs, nil
}
