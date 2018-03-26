package sbp

import "io"

// MsgInsStatus represents a contents of MSG_INS_STATUS.
type MsgInsStatus struct {
	// Status flags
	InsStatus             uint8
	InitializationRoutine uint8
	InsError              uint8
}

func (m *MsgInsStatus) MsgType() uint16 {
	return TypeMsgInsStatus
}

func (m *MsgInsStatus) UnmarshalBinary(bs []byte) error {
	if len(bs) < 4 {
		return io.ErrUnexpectedEOF
	}

	m.InsStatus = bs[0] & 0x7
	m.InitializationRoutine = bs[0] >> 3 & 0x7
	m.InsError = bs[0] >> 6 & 0x3

	return nil
}

func (m *MsgInsStatus) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 4)

	bs[0] = (m.InsStatus & 0x7) | (m.InitializationRoutine & 0x7 << 3) | (m.InsError & 0x3 << 6)

	return bs, nil
}
