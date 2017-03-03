package sbp

import (
	"encoding/binary"
	"io"
)

// MsgStartup represents a contents of MSG_STARTUP.
type MsgStartup struct {
	// Reserved
	Reserved uint32
}

func (m *MsgStartup) FromBytes(bs []byte) error {
	if len(bs) < 4 {
		return io.ErrUnexpectedEOF
	}

	m.Reserved = binary.LittleEndian.Uint32(bs[0:4])

	return nil
}

func (m *MsgStartup) Bytes() ([]byte, error) {
	bs := make([]byte, 4)

	binary.LittleEndian.PutUint32(bs[0:4], m.Reserved)

	return bs, nil
}
