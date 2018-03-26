package sbp

import (
	"encoding/binary"
	"io"
)

// MsgAgeCorrections represents a contents of MSG_AGE_CORRECTIONS.
type MsgAgeCorrections struct {
	// GPS Time of Week (unit:ms)
	Tow uint32

	// Age of the corrections (0xFFFF indicates in-valid)
	Age uint16
}

// MsgType returns the number representing the type.
func (m *MsgAgeCorrections) MsgType() uint16 {
	return TypeMsgAgeCorrections
}

// UnmarshalBinary parses a byte slice.
func (m *MsgAgeCorrections) UnmarshalBinary(bs []byte) error {
	if len(bs) < 6 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])
	m.Age = binary.LittleEndian.Uint16(bs[4:6])

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgAgeCorrections) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 6)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)
	binary.LittleEndian.PutUint16(bs[4:6], m.Age)

	return bs, nil
}
