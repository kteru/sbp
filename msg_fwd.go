package sbp

import "io"

// MsgFwd represents a contents of MSG_FWD.
type MsgFwd struct {
	// Source identifier
	Source uint8
	// Protocol identifier
	Protocol uint8

	// Variable length wrapped binary message
	FwdPayload []byte
}

func (m *MsgFwd) UnmarshalBinary(bs []byte) error {
	if len(bs) < 2 {
		return io.ErrUnexpectedEOF
	}

	m.Source = bs[0]
	m.Protocol = bs[1]

	m.FwdPayload = bs[2:]

	return nil
}

func (m *MsgFwd) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 2, 2+len(m.FwdPayload))

	bs[0] = m.Source
	bs[1] = m.Protocol

	bs = append(bs, m.FwdPayload...)

	return bs, nil
}
