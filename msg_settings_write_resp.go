package sbp

import (
	"bytes"
	"io"
)

// MsgSettingsWriteResp represents a contents of MSG_SETTINGS_WRITE_RESP.
type MsgSettingsWriteResp struct {
	// Write status
	Status uint8

	// String with contents
	SectionSetting string
	Setting        string
	Value          string
}

func (m *MsgSettingsWriteResp) MsgType() uint16 {
	return TypeMsgSettingsWriteResp
}

func (m *MsgSettingsWriteResp) UnmarshalBinary(bs []byte) error {
	if len(bs) < 1 {
		return io.ErrUnexpectedEOF
	}

	m.Status = bs[0]

	bss := bytes.Split(bs[1:], []byte{0x00})

	if len(bss) != 4 || len(bss[3]) > 0 {
		return ErrInvalidFormat
	}

	m.SectionSetting = string(bss[0])
	m.Setting = string(bss[1])
	m.Value = string(bss[2])

	return nil
}

func (m *MsgSettingsWriteResp) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 1, 1+len(m.SectionSetting)+1+len(m.Setting)+1+len(m.Value)+1)

	bs[0] = m.Status

	bs = append(bs, []byte(m.SectionSetting)...)
	bs = append(bs, 0x00)
	bs = append(bs, []byte(m.Setting)...)
	bs = append(bs, 0x00)
	bs = append(bs, []byte(m.Value)...)
	bs = append(bs, 0x00)

	return bs, nil
}
