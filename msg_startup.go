package sbp

import (
	"encoding/binary"
	"io"
)

// MsgStartup represents a contents of MSG_STARTUP.
type MsgStartup struct {
	// Cause of startup
	Cause uint8

	// Startup type
	StartupType uint8

	// Reserved
	Reserved uint16
}

func (m *MsgStartup) MsgType() uint16 {
	return TypeMsgStartup
}

func (m *MsgStartup) UnmarshalBinary(bs []byte) error {
	if len(bs) < 4 {
		return io.ErrUnexpectedEOF
	}

	m.Cause = bs[0]
	m.StartupType = bs[1]

	m.Reserved = binary.LittleEndian.Uint16(bs[2:4])

	return nil
}

func (m *MsgStartup) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 4)

	bs[0] = m.Cause
	bs[1] = m.StartupType

	binary.LittleEndian.PutUint16(bs[2:4], m.Reserved)

	return bs, nil
}
