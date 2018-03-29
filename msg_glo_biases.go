package sbp

import (
	"encoding/binary"
	"io"
)

// MsgGloBiases represents a contents of MSG_GLO_BIASES.
type MsgGloBiases struct {
	// GLONASS FDMA signals mask (unit:boolean)
	Mask uint8

	// GLONASS L1 C/A Code-Phase Bias (unit: m * 0.02)
	L1CaBias int16
	// GLONASS L1 P Code-Phase Bias (unit: m * 0.02)
	L1PBias int16
	// GLONASS L2 C/A Code-Phase Bias (unit: m * 0.02)
	L2CaBias int16
	// GLONASS L2 P Code-Phase Bias (unit: m * 0.02)
	L2PBias int16
}

// MsgType returns the number representing the type.
func (m *MsgGloBiases) MsgType() uint16 {
	return TypeMsgGloBiases
}

// UnmarshalBinary parses a byte slice.
func (m *MsgGloBiases) UnmarshalBinary(bs []byte) error {
	if len(bs) < 9 {
		return io.ErrUnexpectedEOF
	}

	m.Mask = bs[0]

	m.L1CaBias = int16(binary.LittleEndian.Uint16(bs[1:3]))
	m.L1PBias = int16(binary.LittleEndian.Uint16(bs[3:5]))
	m.L2CaBias = int16(binary.LittleEndian.Uint16(bs[5:7]))
	m.L2PBias = int16(binary.LittleEndian.Uint16(bs[7:9]))

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgGloBiases) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 9)

	bs[0] = m.Mask

	binary.LittleEndian.PutUint16(bs[1:3], uint16(m.L1CaBias))
	binary.LittleEndian.PutUint16(bs[3:5], uint16(m.L1PBias))
	binary.LittleEndian.PutUint16(bs[5:7], uint16(m.L2CaBias))
	binary.LittleEndian.PutUint16(bs[7:9], uint16(m.L2PBias))

	return bs, nil
}
