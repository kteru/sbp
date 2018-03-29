package sbp

// MsgSettingsReadByIndexDone represents a contents of MSG_SETTINGS_READ_BY_INDEX_DONE.
type MsgSettingsReadByIndexDone struct {
}

// MsgType returns the number representing the type.
func (m *MsgSettingsReadByIndexDone) MsgType() uint16 {
	return TypeMsgSettingsReadByIndexDone
}

// UnmarshalBinary parses a byte slice.
func (m *MsgSettingsReadByIndexDone) UnmarshalBinary(bs []byte) error {
	if len(bs) != 0 {
		return ErrInvalidFormat
	}

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgSettingsReadByIndexDone) MarshalBinary() ([]byte, error) {
	return nil, nil
}
