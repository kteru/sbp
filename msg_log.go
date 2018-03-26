package sbp

import "io"

// MsgLog represents a contents of MSG_LOG.
type MsgLog struct {
	// Logging level
	Level uint8

	// Human-readable string
	Text string
}

// MsgType returns the number representing the type.
func (m *MsgLog) MsgType() uint16 {
	return TypeMsgLog
}

// UnmarshalBinary parses a byte slice.
func (m *MsgLog) UnmarshalBinary(bs []byte) error {
	if len(bs) < 1 {
		return io.ErrUnexpectedEOF
	}

	level := bs[0]
	m.Level = level & 0x7

	m.Text = string(bs[1:])

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgLog) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 1, 1+len(m.Text))

	level := m.Level & 0x7
	bs[0] = level

	bs = append(bs, []byte(m.Text)...)

	return bs, nil
}
