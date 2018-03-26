package sbp

import "bytes"

// MsgSettingsReadReq represents a contents of MSG_SETTINGS_READ_REQ.
type MsgSettingsReadReq struct {
	// String with contents
	SectionSetting string
	Setting        string
}

func (m *MsgSettingsReadReq) MsgType() uint16 {
	return TypeMsgSettingsReadReq
}

func (m *MsgSettingsReadReq) UnmarshalBinary(bs []byte) error {
	bss := bytes.Split(bs, []byte{0x00})

	if len(bss) != 3 || len(bss[2]) > 0 {
		return ErrInvalidFormat
	}

	m.SectionSetting = string(bss[0])
	m.Setting = string(bss[1])

	return nil
}

func (m *MsgSettingsReadReq) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 0, len(m.SectionSetting)+1+len(m.Setting)+1)

	bs = append(bs, []byte(m.SectionSetting)...)
	bs = append(bs, 0x00)
	bs = append(bs, []byte(m.Setting)...)
	bs = append(bs, 0x00)

	return bs, nil
}
