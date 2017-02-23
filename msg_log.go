package sbp

import "io"

// MsgLog represents a contents of MSG_LOG.
type MsgLog struct {
	// Logging level
	Level uint8

	// Human-readable string
	Text string
}

func (m *MsgLog) FromBytes(bs []byte) error {
	if len(bs) < 1 {
		return io.ErrUnexpectedEOF
	}

	m.Level = bs[0]

	m.Text = string(bs[1:])

	return nil
}

func (m *MsgLog) Bytes() ([]byte, error) {
	bs := make([]byte, 1, 1+len(m.Text))

	bs[0] = m.Level

	bs = append(bs, []byte(m.Text)...)

	return bs, nil
}
