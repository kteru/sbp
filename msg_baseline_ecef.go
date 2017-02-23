package sbp

import (
	"encoding/binary"
	"io"
)

// MsgBaselineEcef represents a contents of MSG_BASELINE_ECEF.
type MsgBaselineEcef struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	// ECEF coordinates (unit:mm)
	X int32
	Y int32
	Z int32

	// Position accuracy estimate (unit:mm)
	Accuracy uint16

	// Number of satellites used in solution
	NumSats uint8

	// Status flags
	FixMode          uint8
	RaimAvailability uint8
	RaimRepair       uint8
}

func (m *MsgBaselineEcef) FromBytes(bs []byte) error {
	if len(bs) < 20 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])

	m.X = int32(binary.LittleEndian.Uint32(bs[4:8]))
	m.Y = int32(binary.LittleEndian.Uint32(bs[8:12]))
	m.Z = int32(binary.LittleEndian.Uint32(bs[12:16]))

	m.Accuracy = binary.LittleEndian.Uint16(bs[16:18])

	m.NumSats = bs[18]

	flags := bs[19]
	m.FixMode = flags & 0x7
	m.RaimAvailability = (flags & 0x8) >> 3
	m.RaimRepair = (flags & 0x10) >> 4

	return nil
}

func (m *MsgBaselineEcef) Bytes() ([]byte, error) {
	bs := make([]byte, 20)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)

	binary.LittleEndian.PutUint32(bs[4:8], uint32(m.X))
	binary.LittleEndian.PutUint32(bs[8:12], uint32(m.Y))
	binary.LittleEndian.PutUint32(bs[12:16], uint32(m.Z))

	binary.LittleEndian.PutUint16(bs[16:18], m.Accuracy)

	bs[18] = m.NumSats

	flags := (m.FixMode & 0x7) | (m.RaimAvailability << 3 & 0x8) | (m.RaimRepair << 4 & 0x10)
	bs[19] = flags

	return bs, nil
}
