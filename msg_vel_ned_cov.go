package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgVelNedCov represents a contents of MSG_VEL_NED_COV.
type MsgVelNedCov struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	// NED coordinates (unit:mm/s)
	N int32
	E int32
	D int32

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
	VelocityMode           uint8
	InertialNavigationMode uint8
}

func (m *MsgVelNedCov) MsgType() uint16 {
	return TypeMsgVelNedCov
}

func (m *MsgVelNedCov) UnmarshalBinary(bs []byte) error {
	if len(bs) < 42 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])

	m.N = int32(binary.LittleEndian.Uint32(bs[4:8]))
	m.E = int32(binary.LittleEndian.Uint32(bs[8:12]))
	m.D = int32(binary.LittleEndian.Uint32(bs[12:16]))

	m.CovNN = math.Float32frombits(binary.LittleEndian.Uint32(bs[16:20]))
	m.CovNE = math.Float32frombits(binary.LittleEndian.Uint32(bs[20:24]))
	m.CovND = math.Float32frombits(binary.LittleEndian.Uint32(bs[24:28]))
	m.CovEE = math.Float32frombits(binary.LittleEndian.Uint32(bs[28:32]))
	m.CovED = math.Float32frombits(binary.LittleEndian.Uint32(bs[32:36]))
	m.CovDD = math.Float32frombits(binary.LittleEndian.Uint32(bs[36:40]))

	m.NSats = bs[40]

	flags := bs[41]
	m.VelocityMode = flags & 0x7
	m.InertialNavigationMode = flags >> 3 & 0x3

	return nil
}

func (m *MsgVelNedCov) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 42)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)

	binary.LittleEndian.PutUint32(bs[4:8], uint32(m.N))
	binary.LittleEndian.PutUint32(bs[8:12], uint32(m.E))
	binary.LittleEndian.PutUint32(bs[12:16], uint32(m.D))

	binary.LittleEndian.PutUint32(bs[16:20], math.Float32bits(m.CovNN))
	binary.LittleEndian.PutUint32(bs[20:24], math.Float32bits(m.CovNE))
	binary.LittleEndian.PutUint32(bs[24:28], math.Float32bits(m.CovND))
	binary.LittleEndian.PutUint32(bs[28:32], math.Float32bits(m.CovEE))
	binary.LittleEndian.PutUint32(bs[32:36], math.Float32bits(m.CovED))
	binary.LittleEndian.PutUint32(bs[36:40], math.Float32bits(m.CovDD))

	bs[40] = m.NSats

	flags := (m.VelocityMode & 0x7) | (m.InertialNavigationMode & 0x3 << 3)
	bs[41] = flags

	return bs, nil
}
