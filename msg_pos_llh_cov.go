package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgPosLlhCov represents a contents of MSG_POS_LLH_COV.
type MsgPosLlhCov struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	// Latitude, Longitude (unit:deg)
	Lat float64
	Lon float64

	// Height (unit:m)
	Height float64

	// Estimated variance of northing (unit:m^2)
	CovNN float32
	// Covariance of northing and easting (unit:m^2)
	CovNE float32
	// Covariance of northing and downward measurement (unit:m^2)
	CovND float32
	// Estimated variance of easting (unit:m^2)
	CovEE float32
	// Covariance of easting and downward measurement (unit:m^2)
	CovED float32
	// Estimated variance of downward measurement (unit:m^2)
	CovDD float32

	// Number of satellites used in solution
	NSats uint8

	// Status flags
	FixMode                uint8
	InertialNavigationMode uint8
}

func (m *MsgPosLlhCov) MsgType() uint16 {
	return TypeMsgPosLlhCov
}

func (m *MsgPosLlhCov) UnmarshalBinary(bs []byte) error {
	if len(bs) < 54 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])

	m.Lat = math.Float64frombits(binary.LittleEndian.Uint64(bs[4:12]))
	m.Lon = math.Float64frombits(binary.LittleEndian.Uint64(bs[12:20]))
	m.Height = math.Float64frombits(binary.LittleEndian.Uint64(bs[20:28]))

	m.CovNN = math.Float32frombits(binary.LittleEndian.Uint32(bs[28:32]))
	m.CovNE = math.Float32frombits(binary.LittleEndian.Uint32(bs[32:36]))
	m.CovND = math.Float32frombits(binary.LittleEndian.Uint32(bs[36:40]))
	m.CovEE = math.Float32frombits(binary.LittleEndian.Uint32(bs[40:44]))
	m.CovED = math.Float32frombits(binary.LittleEndian.Uint32(bs[44:48]))
	m.CovDD = math.Float32frombits(binary.LittleEndian.Uint32(bs[48:52]))

	m.NSats = bs[52]

	flags := bs[53]
	m.FixMode = flags & 0x7
	m.InertialNavigationMode = flags >> 3 & 0x3

	return nil
}

func (m *MsgPosLlhCov) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 54)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)

	binary.LittleEndian.PutUint64(bs[4:12], math.Float64bits(m.Lat))
	binary.LittleEndian.PutUint64(bs[12:20], math.Float64bits(m.Lon))
	binary.LittleEndian.PutUint64(bs[20:28], math.Float64bits(m.Height))

	binary.LittleEndian.PutUint32(bs[28:32], math.Float32bits(m.CovNN))
	binary.LittleEndian.PutUint32(bs[32:36], math.Float32bits(m.CovNE))
	binary.LittleEndian.PutUint32(bs[36:40], math.Float32bits(m.CovND))
	binary.LittleEndian.PutUint32(bs[40:44], math.Float32bits(m.CovEE))
	binary.LittleEndian.PutUint32(bs[44:48], math.Float32bits(m.CovED))
	binary.LittleEndian.PutUint32(bs[48:52], math.Float32bits(m.CovDD))

	bs[52] = m.NSats

	flags := (m.FixMode & 0x7) | (m.InertialNavigationMode & 0x3 << 3)
	bs[53] = flags

	return bs, nil
}
