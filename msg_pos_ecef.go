package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgPosEcef represents a contents of MSG_POS_ECEF.
type MsgPosEcef struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	// ECEF coordinates (unit:m)
	X float64
	Y float64
	Z float64

	// Position accuracy estimate (unit:mm)
	Accuracy uint16

	// Number of satellites used in solution
	NSats uint8

	// Status flags
	FixMode                uint8
	InertialNavigationMode uint8
}

func (m *MsgPosEcef) MsgType() uint16 {
	return TypeMsgPosEcef
}

func (m *MsgPosEcef) UnmarshalBinary(bs []byte) error {
	if len(bs) < 32 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])

	m.X = math.Float64frombits(binary.LittleEndian.Uint64(bs[4:12]))
	m.Y = math.Float64frombits(binary.LittleEndian.Uint64(bs[12:20]))
	m.Z = math.Float64frombits(binary.LittleEndian.Uint64(bs[20:28]))

	m.Accuracy = binary.LittleEndian.Uint16(bs[28:30])

	m.NSats = bs[30]

	flags := bs[31]
	m.FixMode = flags & 0x7
	m.InertialNavigationMode = flags >> 3 & 0x3

	return nil
}

func (m *MsgPosEcef) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 32)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)

	binary.LittleEndian.PutUint64(bs[4:12], math.Float64bits(m.X))
	binary.LittleEndian.PutUint64(bs[12:20], math.Float64bits(m.Y))
	binary.LittleEndian.PutUint64(bs[20:28], math.Float64bits(m.Z))

	binary.LittleEndian.PutUint16(bs[28:30], m.Accuracy)

	bs[30] = m.NSats

	flags := (m.FixMode & 0x7) | (m.InertialNavigationMode & 0x3 << 3)
	bs[31] = flags

	return bs, nil
}
