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
	NumSats uint8

	// Status flags
	Flags uint8
}

func (m *MsgVelNed) FromBytes(bs []byte) error {
	if len(bs) < 22 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])

	m.N = int32(binary.LittleEndian.Uint32(bs[4:8]))
	m.E = int32(binary.LittleEndian.Uint32(bs[8:12]))
	m.D = int32(binary.LittleEndian.Uint32(bs[12:16]))

	m.HAccuracy = binary.LittleEndian.Uint16(bs[16:18])
	m.VAccuracy = binary.LittleEndian.Uint16(bs[18:20])

	m.NumSats = bs[20]
	m.Flags = bs[21]

	return nil
}

func (m *MsgVelNed) Bytes() ([]byte, error) {
	bs := make([]byte, 22)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)

	binary.LittleEndian.PutUint32(bs[4:8], uint32(m.N))
	binary.LittleEndian.PutUint32(bs[8:12], uint32(m.E))
	binary.LittleEndian.PutUint32(bs[12:16], uint32(m.D))

	binary.LittleEndian.PutUint16(bs[16:18], m.HAccuracy)
	binary.LittleEndian.PutUint16(bs[18:20], m.VAccuracy)

	bs[20] = m.NumSats
	bs[21] = m.Flags

	return bs, nil
}
