package sbp

import (
	"encoding/binary"
	"io"
)

// MsgSvConfigurationGps represents a contents of MSG_SV_CONFIGURATION_GPS.
type MsgSvConfigurationGps struct {
	// Seconds since start of GPS week (unit: s)
	TNmctTow uint32
	// GPS week number (unit: weeks)
	TNmctWn uint16

	// L2C capability mask, SV32 bit being MSB, SV1 bit being LSB
	L2cMask uint32
}

// MsgType returns the number representing the type.
func (m *MsgSvConfigurationGps) MsgType() uint16 {
	return TypeMsgSvConfigurationGps
}

// UnmarshalBinary parses a byte slice.
func (m *MsgSvConfigurationGps) UnmarshalBinary(bs []byte) error {
	if len(bs) < 10 {
		return io.ErrUnexpectedEOF
	}

	m.TNmctTow = binary.LittleEndian.Uint32(bs[0:4])
	m.TNmctWn = binary.LittleEndian.Uint16(bs[4:6])

	m.L2cMask = binary.LittleEndian.Uint32(bs[6:10])

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgSvConfigurationGps) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 10)

	binary.LittleEndian.PutUint32(bs[0:4], m.TNmctTow)
	binary.LittleEndian.PutUint16(bs[4:6], m.TNmctWn)

	binary.LittleEndian.PutUint32(bs[6:10], m.L2cMask)

	return bs, nil
}
