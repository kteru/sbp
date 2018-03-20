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

func (m *MsgSettingsReadByIndexReq) UnmarshalBinary(bs []byte) error {
	if len(bs) < 2 {
		return io.ErrUnexpectedEOF
	}

	m.Index = binary.LittleEndian.Uint16(bs[0:2])

	return nil
}

func (m *MsgSettingsReadByIndexReq) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 2)

	binary.LittleEndian.PutUint16(bs[0:2], m.Index)

	return bs, nil
}
