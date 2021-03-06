package sbp

import (
	"encoding/binary"
	"io"
)

// MsgSettingsReadByIndexReq represents a contents of MSG_SETTINGS_READ_BY_INDEX_REQ.
type MsgSettingsReadByIndexReq struct {
	// An index into the device settings, with values ranging from 0 to length (settings)
	Index uint16
}

// MsgType returns the number representing the type.
func (m *MsgSettingsReadByIndexReq) MsgType() uint16 {
	return TypeMsgSettingsReadByIndexReq
}

// UnmarshalBinary parses a byte slice.
func (m *MsgSettingsReadByIndexReq) UnmarshalBinary(bs []byte) error {
	if len(bs) < 2 {
		return io.ErrUnexpectedEOF
	}

	m.Index = binary.LittleEndian.Uint16(bs[0:2])

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgSettingsReadByIndexReq) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 2)

	binary.LittleEndian.PutUint16(bs[0:2], m.Index)

	return bs, nil
}
