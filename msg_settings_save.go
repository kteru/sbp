package sbp

// MsgSettingsSave represents a contents of MSG_SETTINGS_SAVE.
type MsgSettingsSave struct {
}

// MsgType returns the number representing the type.
func (m *MsgSettingsSave) MsgType() uint16 {
	return TypeMsgSettingsSave
}

// UnmarshalBinary parses a byte slice.
func (m *MsgSettingsSave) UnmarshalBinary(bs []byte) error {
	if len(bs) != 0 {
		return ErrInvalidFormat
	}

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgSettingsSave) MarshalBinary() ([]byte, error) {
	return nil, nil
}
