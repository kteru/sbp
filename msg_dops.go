package sbp

import (
	"encoding/binary"
	"io"
)

// MsgDops represents a contents of MSG_DOPS.
type MsgDops struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	////
	//// Dilution of Precision (unit:0.01)
	////
	// Geometric
	Gdop uint16
	// Position
	Pdop uint16
	// Time
	Tdop uint16
	// Horizontal
	Hdop uint16
	// Vertical
	Vdop uint16

	// Status flags
	FixMode    uint8
	RaimRepair uint8
}

func (m *MsgDops) UnmarshalBinary(bs []byte) error {
	if len(bs) < 15 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])

	m.Gdop = binary.LittleEndian.Uint16(bs[4:6])
	m.Pdop = binary.LittleEndian.Uint16(bs[6:8])
	m.Tdop = binary.LittleEndian.Uint16(bs[8:10])
	m.Hdop = binary.LittleEndian.Uint16(bs[10:12])
	m.Vdop = binary.LittleEndian.Uint16(bs[12:14])

	flags := bs[14]
	m.FixMode = flags & 0x7
	m.RaimRepair = flags >> 7 & 0x1

	return nil
}

func (m *MsgDops) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 15)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)

	binary.LittleEndian.PutUint16(bs[4:6], m.Gdop)
	binary.LittleEndian.PutUint16(bs[6:8], m.Pdop)
	binary.LittleEndian.PutUint16(bs[8:10], m.Tdop)
	binary.LittleEndian.PutUint16(bs[10:12], m.Hdop)
	binary.LittleEndian.PutUint16(bs[12:14], m.Vdop)

	flags := (m.FixMode & 0x7) | (m.RaimRepair & 0x1 << 7)
	bs[14] = flags

	return bs, nil
}
