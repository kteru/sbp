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

func (m *MsgBasePosEcef) FromBytes(bs []byte) error {
	if len(bs) < 24 {
		return io.ErrUnexpectedEOF
	}

	m.X = math.Float64frombits(binary.LittleEndian.Uint64(bs[0:8]))
	m.Y = math.Float64frombits(binary.LittleEndian.Uint64(bs[8:16]))
	m.Z = math.Float64frombits(binary.LittleEndian.Uint64(bs[16:24]))

	return nil
}

func (m *MsgBasePosEcef) Bytes() ([]byte, error) {
	bs := make([]byte, 24)

	binary.LittleEndian.PutUint64(bs[0:8], math.Float64bits(m.X))
	binary.LittleEndian.PutUint64(bs[8:16], math.Float64bits(m.Y))
	binary.LittleEndian.PutUint64(bs[16:24], math.Float64bits(m.Z))

	return bs, nil
}