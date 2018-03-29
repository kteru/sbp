package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgPosEcefCov represents a contents of MSG_POS_ECEF_COV.
type MsgPosEcefCov struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	// ECEF coordinates (unit:m)
	X float64
	Y float64
	Z float64

	// Estimated variance of x (unit:m^2)
	CovXX float32
	// Estimated covariance of x and y (unit:m^2)
	CovXY float32
	// Estimated covariance of x and z (unit:m^2)
	CovXZ float32
	// Estimated variance of y (unit:m^2)
	CovYY float32
	// Estimated covariance of y and z (unit:m^2)
	CovYZ float32
	// Estimated variance of z (unit:m^2)
	CovZZ float32

	// Number of satellites used in solution
	NSats uint8

	// Status flags
	FixMode                uint8
	InertialNavigationMode uint8
}

// MsgType returns the number representing the type.
func (m *MsgPosEcefCov) MsgType() uint16 {
	return TypeMsgPosEcefCov
}

// UnmarshalBinary parses a byte slice.
func (m *MsgPosEcefCov) UnmarshalBinary(bs []byte) error {
	if len(bs) < 54 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])

	m.X = math.Float64frombits(binary.LittleEndian.Uint64(bs[4:12]))
	m.Y = math.Float64frombits(binary.LittleEndian.Uint64(bs[12:20]))
	m.Z = math.Float64frombits(binary.LittleEndian.Uint64(bs[20:28]))

	m.CovXX = math.Float32frombits(binary.LittleEndian.Uint32(bs[28:32]))
	m.CovXY = math.Float32frombits(binary.LittleEndian.Uint32(bs[32:36]))
	m.CovXZ = math.Float32frombits(binary.LittleEndian.Uint32(bs[36:40]))
	m.CovYY = math.Float32frombits(binary.LittleEndian.Uint32(bs[40:44]))
	m.CovYZ = math.Float32frombits(binary.LittleEndian.Uint32(bs[44:48]))
	m.CovZZ = math.Float32frombits(binary.LittleEndian.Uint32(bs[48:52]))

	m.NSats = bs[52]

	flags := bs[53]
	m.FixMode = flags & 0x7
	m.InertialNavigationMode = flags >> 3 & 0x3

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgPosEcefCov) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 54)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)

	binary.LittleEndian.PutUint64(bs[4:12], math.Float64bits(m.X))
	binary.LittleEndian.PutUint64(bs[12:20], math.Float64bits(m.Y))
	binary.LittleEndian.PutUint64(bs[20:28], math.Float64bits(m.Z))

	binary.LittleEndian.PutUint32(bs[28:32], math.Float32bits(m.CovXX))
	binary.LittleEndian.PutUint32(bs[32:36], math.Float32bits(m.CovXY))
	binary.LittleEndian.PutUint32(bs[36:40], math.Float32bits(m.CovXZ))
	binary.LittleEndian.PutUint32(bs[40:44], math.Float32bits(m.CovYY))
	binary.LittleEndian.PutUint32(bs[44:48], math.Float32bits(m.CovYZ))
	binary.LittleEndian.PutUint32(bs[48:52], math.Float32bits(m.CovZZ))

	bs[52] = m.NSats

	flags := (m.FixMode & 0x7) | (m.InertialNavigationMode & 0x3 << 3)
	bs[53] = flags

	return bs, nil
}
