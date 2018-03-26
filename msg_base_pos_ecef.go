package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgBasePosEcef represents a contents of MSG_BASE_POS_ECEF.
type MsgBasePosEcef struct {
	// ECEF coordinates (unit:m)
	X float64
	Y float64
	Z float64
}

// MsgType returns the number representing the type.
func (m *MsgBasePosEcef) MsgType() uint16 {
	return TypeMsgBasePosEcef
}

// UnmarshalBinary parses a byte slice.
func (m *MsgBasePosEcef) UnmarshalBinary(bs []byte) error {
	if len(bs) < 24 {
		return io.ErrUnexpectedEOF
	}

	m.X = math.Float64frombits(binary.LittleEndian.Uint64(bs[0:8]))
	m.Y = math.Float64frombits(binary.LittleEndian.Uint64(bs[8:16]))
	m.Z = math.Float64frombits(binary.LittleEndian.Uint64(bs[16:24]))

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgBasePosEcef) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 24)

	binary.LittleEndian.PutUint64(bs[0:8], math.Float64bits(m.X))
	binary.LittleEndian.PutUint64(bs[8:16], math.Float64bits(m.Y))
	binary.LittleEndian.PutUint64(bs[16:24], math.Float64bits(m.Z))

	return bs, nil
}
