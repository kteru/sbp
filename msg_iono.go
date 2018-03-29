package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgIono represents a contents of MSG_IONO.
type MsgIono struct {
	// Seconds since start of GPS week (unit: s)
	TNmctTow uint32
	// GPS week number (unit: weeks)
	TNmctWn uint16

	// unit: s
	A0 float64
	// unit: s/semi-circle
	A1 float64
	// unit: s/(semi-circle)^2
	A2 float64
	// unit: s/(semi-circle)^3
	A3 float64
	// unit: s
	B0 float64
	// unit: s/semi-circle
	B1 float64
	// unit: s/(semi-circle)^2
	B2 float64
	// unit: s/(semi-circle)^3
	B3 float64
}

// MsgType returns the number representing the type.
func (m *MsgIono) MsgType() uint16 {
	return TypeMsgIono
}

// UnmarshalBinary parses a byte slice.
func (m *MsgIono) UnmarshalBinary(bs []byte) error {
	if len(bs) < 70 {
		return io.ErrUnexpectedEOF
	}

	m.TNmctTow = binary.LittleEndian.Uint32(bs[0:4])
	m.TNmctWn = binary.LittleEndian.Uint16(bs[4:6])

	m.A0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[6:14]))
	m.A1 = math.Float64frombits(binary.LittleEndian.Uint64(bs[14:22]))
	m.A2 = math.Float64frombits(binary.LittleEndian.Uint64(bs[22:30]))
	m.A3 = math.Float64frombits(binary.LittleEndian.Uint64(bs[30:38]))
	m.B0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[38:46]))
	m.B1 = math.Float64frombits(binary.LittleEndian.Uint64(bs[46:54]))
	m.B2 = math.Float64frombits(binary.LittleEndian.Uint64(bs[54:62]))
	m.B3 = math.Float64frombits(binary.LittleEndian.Uint64(bs[62:70]))

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgIono) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 70)

	binary.LittleEndian.PutUint32(bs[0:4], m.TNmctTow)
	binary.LittleEndian.PutUint16(bs[4:6], m.TNmctWn)

	binary.LittleEndian.PutUint64(bs[6:14], math.Float64bits(m.A0))
	binary.LittleEndian.PutUint64(bs[14:22], math.Float64bits(m.A1))
	binary.LittleEndian.PutUint64(bs[22:30], math.Float64bits(m.A2))
	binary.LittleEndian.PutUint64(bs[30:38], math.Float64bits(m.A3))
	binary.LittleEndian.PutUint64(bs[38:46], math.Float64bits(m.B0))
	binary.LittleEndian.PutUint64(bs[46:54], math.Float64bits(m.B1))
	binary.LittleEndian.PutUint64(bs[54:62], math.Float64bits(m.B2))
	binary.LittleEndian.PutUint64(bs[62:70], math.Float64bits(m.B3))

	return bs, nil
}
