package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgBasePosLlh represents a contents of MSG_BASE_POS_LLH.
type MsgBasePosLlh struct {
	// Latitude, Longitude (unit:deg)
	Lat float64
	Lon float64

	// Height (unit:m)
	Height float64
}

func (m *MsgBasePosLlh) UnmarshalBinary(bs []byte) error {
	if len(bs) < 24 {
		return io.ErrUnexpectedEOF
	}

	m.Lat = math.Float64frombits(binary.LittleEndian.Uint64(bs[0:8]))
	m.Lon = math.Float64frombits(binary.LittleEndian.Uint64(bs[8:16]))
	m.Height = math.Float64frombits(binary.LittleEndian.Uint64(bs[16:24]))

	return nil
}

func (m *MsgBasePosLlh) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 24)

	binary.LittleEndian.PutUint64(bs[0:8], math.Float64bits(m.Lat))
	binary.LittleEndian.PutUint64(bs[8:16], math.Float64bits(m.Lon))
	binary.LittleEndian.PutUint64(bs[16:24], math.Float64bits(m.Height))

	return bs, nil
}
