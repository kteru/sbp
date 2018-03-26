package sbp

import "bytes"

// MsgSettingsWrite represents a contents of MSG_SETTINGS_WRITE.
type MsgSettingsWrite struct {
	// String with contents
	SectionSetting string
	Setting        string
	Value          string
}

func (m *MsgSettingsWrite) MsgType() uint16 {
	return TypeMsgSettingsWrite
}

func (m *MsgSettingsWrite) UnmarshalBinary(bs []byte) error {
	bss := bytes.Split(bs, []byte{0x00})

	if len(bss) != 4 || len(bss[3]) > 0 {
		return ErrInvalidFormat
	}

	m.SectionSetting = string(bss[0])
	m.Setting = string(bss[1])
	m.Value = string(bss[2])

	return nil
}

func (m *MsgSettingsWrite) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 0, len(m.SectionSetting)+1+len(m.Setting)+1+len(m.Value)+1)

	bs = append(bs, []byte(m.SectionSetting)...)
	bs = append(bs, 0x00)
	bs = append(bs, []byte(m.Setting)...)
	bs = append(bs, 0x00)
	bs = append(bs, []byte(m.Value)...)
	bs = append(bs, 0x00)

	return bs, nil
}
