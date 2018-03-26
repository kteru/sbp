package sbp

import (
	"bytes"
	"encoding/binary"
	"io"
)

// MsgSettingsReadByIndexResp represents a contents of MSG_SETTINGS_READ_BY_INDEX_RESP.
type MsgSettingsReadByIndexResp struct {
	// An index into the device settings, with values ranging from 0 to length (settings)
	Index uint16

	// String with contents
	SectionSetting string
	Setting        string
	Value          string
}

func (m *MsgSettingsReadByIndexResp) MsgType() uint16 {
	return TypeMsgSettingsReadByIndexResp
}

func (m *MsgSettingsReadByIndexResp) UnmarshalBinary(bs []byte) error {
	if len(bs) < 2 {
		return io.ErrUnexpectedEOF
	}

	m.Index = binary.LittleEndian.Uint16(bs[0:2])

	bss := bytes.Split(bs[2:], []byte{0x00})

	if len(bss) != 4 || len(bss[3]) > 0 {
		return ErrInvalidFormat
	}

	m.SectionSetting = string(bss[0])
	m.Setting = string(bss[1])
	m.Value = string(bss[2])

	return nil
}

func (m *MsgSettingsReadByIndexResp) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 2, 2+len(m.SectionSetting)+1+len(m.Setting)+1+len(m.Value)+1)

	binary.LittleEndian.PutUint16(bs[0:2], m.Index)

	bs = append(bs, []byte(m.SectionSetting)...)
	bs = append(bs, 0x00)
	bs = append(bs, []byte(m.Setting)...)
	bs = append(bs, 0x00)
	bs = append(bs, []byte(m.Value)...)
	bs = append(bs, 0x00)

	return bs, nil
}
