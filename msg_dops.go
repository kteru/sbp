package sbp

import (
	"encoding/binary"
	"io"
)

// MsgDops represents a contents of MSG_DOPS.
type MsgDops struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	//
	// Dilution of Precision (unit:0.01)
	//

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
}

func (m *MsgDops) FromBytes(bs []byte) error {
	if len(bs) < 14 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])

	m.Gdop = binary.LittleEndian.Uint16(bs[4:6])
	m.Pdop = binary.LittleEndian.Uint16(bs[6:8])
	m.Tdop = binary.LittleEndian.Uint16(bs[8:10])
	m.Hdop = binary.LittleEndian.Uint16(bs[10:12])
	m.Vdop = binary.LittleEndian.Uint16(bs[12:14])

	return nil
}

func (m *MsgDops) Bytes() ([]byte, error) {
	bs := make([]byte, 14)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)

	binary.LittleEndian.PutUint16(bs[4:6], m.Gdop)
	binary.LittleEndian.PutUint16(bs[6:8], m.Pdop)
	binary.LittleEndian.PutUint16(bs[8:10], m.Tdop)
	binary.LittleEndian.PutUint16(bs[10:12], m.Hdop)
	binary.LittleEndian.PutUint16(bs[12:14], m.Vdop)

	return bs, nil
}
