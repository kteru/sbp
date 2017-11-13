package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgPosLlh represents a contents of MSG_POS_LLH.
type MsgPosLlh struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	// Latitude, Longitude (unit:deg)
	Lat float64
	Lon float64

	// Height (unit:m)
	Height float64

	// Horizontal position accuracy estimate (unit:mm)
	HAccuracy uint16

	// Vertical position accuracy estimate (unit:mm)
	VAccuracy uint16

	// Number of satellites used in solution
	NumSats uint8

	// Status flags
	FixMode    uint8
	RaimRepair uint8
}

func (m *MsgPosLlh) FromBytes(bs []byte) error {
	if len(bs) < 34 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])

	m.Lat = math.Float64frombits(binary.LittleEndian.Uint64(bs[4:12]))
	m.Lon = math.Float64frombits(binary.LittleEndian.Uint64(bs[12:20]))
	m.Height = math.Float64frombits(binary.LittleEndian.Uint64(bs[20:28]))

	m.HAccuracy = binary.LittleEndian.Uint16(bs[28:30])
	m.VAccuracy = binary.LittleEndian.Uint16(bs[30:32])

	m.NumSats = bs[32]

	flags := bs[33]
	m.FixMode = flags & 0x7
	m.RaimRepair = flags >> 7 & 0x1

	return nil
}

func (m *MsgPosLlh) Bytes() ([]byte, error) {
	bs := make([]byte, 34)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)

	binary.LittleEndian.PutUint64(bs[4:12], math.Float64bits(m.Lat))
	binary.LittleEndian.PutUint64(bs[12:20], math.Float64bits(m.Lon))
	binary.LittleEndian.PutUint64(bs[20:28], math.Float64bits(m.Height))

	binary.LittleEndian.PutUint16(bs[28:30], m.HAccuracy)
	binary.LittleEndian.PutUint16(bs[30:32], m.VAccuracy)

	bs[32] = m.NumSats

	flags := (m.FixMode & 0x7) | (m.RaimRepair & 0x1 << 7)
	bs[33] = flags

	return bs, nil
}
