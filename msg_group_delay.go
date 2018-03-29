package sbp

import (
	"encoding/binary"
	"io"
)

// MsgGroupDelay represents a contents of MSG_GROUP_DELAY.
type MsgGroupDelay struct {
	// Seconds since start of GPS week (unit: s)
	TOpTow uint32
	// GPS week number (unit: weeks)
	TOpWn uint16

	// Constellation-specific satellite identifier
	SidSat uint8
	// Signal constellation, band and code
	SidCode uint8

	// Validity of the values
	Valid uint8

	// unit: s * 2^-35
	Tgd int16
	// unit: s * 2^-35
	IscL1ca int16
	// unit: s * 2^-35
	IscL2c int16
}

// MsgType returns the number representing the type.
func (m *MsgGroupDelay) MsgType() uint16 {
	return TypeMsgGroupDelay
}

// UnmarshalBinary parses a byte slice.
func (m *MsgGroupDelay) UnmarshalBinary(bs []byte) error {
	if len(bs) < 15 {
		return io.ErrUnexpectedEOF
	}

	m.TOpTow = binary.LittleEndian.Uint32(bs[0:4])
	m.TOpWn = binary.LittleEndian.Uint16(bs[4:6])

	m.SidSat = bs[6]
	m.SidCode = bs[7]

	m.Valid = bs[8]

	m.Tgd = int16(binary.LittleEndian.Uint16(bs[9:11]))
	m.IscL1ca = int16(binary.LittleEndian.Uint16(bs[11:13]))
	m.IscL2c = int16(binary.LittleEndian.Uint16(bs[13:15]))

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgGroupDelay) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 15)

	binary.LittleEndian.PutUint32(bs[0:4], m.TOpTow)
	binary.LittleEndian.PutUint16(bs[4:6], m.TOpWn)

	bs[6] = m.SidSat
	bs[7] = m.SidCode

	bs[8] = m.Valid

	binary.LittleEndian.PutUint16(bs[9:11], uint16(m.Tgd))
	binary.LittleEndian.PutUint16(bs[11:13], uint16(m.IscL1ca))
	binary.LittleEndian.PutUint16(bs[13:15], uint16(m.IscL2c))

	return bs, nil
}
