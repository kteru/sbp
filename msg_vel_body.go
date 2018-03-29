package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgVelBody represents a contents of MSG_VEL_BODY.
type MsgVelBody struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	// Velocity (unit:mm/s)
	X int32
	Y int32
	Z int32

	// Estimated variance of x (unit:m^2)
	CovXX float32
	// Covariance of x and y (unit:m^2)
	CovXY float32
	// Covariance of x and z (unit:m^2)
	CovXZ float32
	// Estimated variance of y (unit:m^2)
	CovYY float32
	// Covariance of y and z (unit:m^2)
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
func (m *MsgVelBody) MsgType() uint16 {
	return TypeMsgVelBody
}

// UnmarshalBinary parses a byte slice.
func (m *MsgVelBody) UnmarshalBinary(bs []byte) error {
	if len(bs) < 42 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])

	m.X = int32(binary.LittleEndian.Uint32(bs[4:8]))
	m.Y = int32(binary.LittleEndian.Uint32(bs[8:12]))
	m.Z = int32(binary.LittleEndian.Uint32(bs[12:16]))

	m.CovXX = math.Float32frombits(binary.LittleEndian.Uint32(bs[16:20]))
	m.CovXY = math.Float32frombits(binary.LittleEndian.Uint32(bs[20:24]))
	m.CovXZ = math.Float32frombits(binary.LittleEndian.Uint32(bs[24:28]))
	m.CovYY = math.Float32frombits(binary.LittleEndian.Uint32(bs[28:32]))
	m.CovYZ = math.Float32frombits(binary.LittleEndian.Uint32(bs[32:36]))
	m.CovZZ = math.Float32frombits(binary.LittleEndian.Uint32(bs[36:40]))

	m.NSats = bs[40]

	flags := bs[41]
	m.FixMode = flags & 0x7
	m.InertialNavigationMode = flags >> 3 & 0x3

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgVelBody) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 42)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)

	binary.LittleEndian.PutUint32(bs[4:8], uint32(m.X))
	binary.LittleEndian.PutUint32(bs[8:12], uint32(m.Y))
	binary.LittleEndian.PutUint32(bs[12:16], uint32(m.Z))

	binary.LittleEndian.PutUint32(bs[16:20], math.Float32bits(m.CovXX))
	binary.LittleEndian.PutUint32(bs[20:24], math.Float32bits(m.CovXY))
	binary.LittleEndian.PutUint32(bs[24:28], math.Float32bits(m.CovXZ))
	binary.LittleEndian.PutUint32(bs[28:32], math.Float32bits(m.CovYY))
	binary.LittleEndian.PutUint32(bs[32:36], math.Float32bits(m.CovYZ))
	binary.LittleEndian.PutUint32(bs[36:40], math.Float32bits(m.CovZZ))

	bs[40] = m.NSats

	flags := (m.FixMode & 0x7) | (m.InertialNavigationMode & 0x3 << 3)
	bs[41] = flags

	return bs, nil
}
