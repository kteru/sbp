package sbp

// MsgSettingsSave represents a contents of MSG_SETTINGS_SAVE.
type MsgSettingsSave struct {
}

func (m *MsgSettingsSave) FromBytes(bs []byte) error {
	if len(bs) != 0 {
		return ErrInvalidMsg
	}

	return nil
}

func (m *MsgSettingsSave) Bytes() ([]byte, error) {
	return nil, nil
}
