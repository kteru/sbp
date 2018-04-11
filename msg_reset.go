package sbp

import "io"

// MsgReset represents a contents of MSG_RESET.
type MsgReset struct {
	// Status flags
	ResoreDefaultSettings bool
}

// MsgType returns the number representing the type.
func (m *MsgReset) MsgType() uint16 {
	return TypeMsgReset
}

// UnmarshalBinary parses a byte slice.
func (m *MsgReset) UnmarshalBinary(bs []byte) error {
	if len(bs) < 4 {
		return io.ErrUnexpectedEOF
	}

	if bs[0]&0x1 > 0 {
		m.ResoreDefaultSettings = true
	}

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgReset) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 4)

	if m.ResoreDefaultSettings {
		bs[0] |= 0x1
	}

	return bs, nil
}
