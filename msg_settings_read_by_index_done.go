package sbp

// MsgSettingsReadByIndexDone represents a contents of MSG_SETTINGS_READ_BY_INDEX_DONE.
type MsgSettingsReadByIndexDone struct {
}

func (m *MsgSettingsReadByIndexDone) FromBytes(bs []byte) error {
	if len(bs) != 0 {
		return ErrInvalidMsg
	}

	return nil
}

func (m *MsgSettingsReadByIndexDone) Bytes() ([]byte, error) {
	return nil, nil
}
