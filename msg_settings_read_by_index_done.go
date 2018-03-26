package sbp

// MsgSettingsReadByIndexDone represents a contents of MSG_SETTINGS_READ_BY_INDEX_DONE.
type MsgSettingsReadByIndexDone struct {
}

func (m *MsgSettingsReadByIndexDone) MsgType() uint16 {
	return TypeMsgSettingsReadByIndexDone
}

func (m *MsgSettingsReadByIndexDone) UnmarshalBinary(bs []byte) error {
	if len(bs) != 0 {
		return ErrInvalidFormat
	}

	return nil
}

func (m *MsgSettingsReadByIndexDone) MarshalBinary() ([]byte, error) {
	return nil, nil
}
