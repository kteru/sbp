package sbp

import "io"

// MsgHeartbeat represents a contents of MSG_HEARTBEAT.
type MsgHeartbeat struct {
	// Status flags
	SbpVersionMajor uint8
	SbpVersionMinor uint8
	ExternalAntenna uint8
	SystemError     uint8
	IOError         uint8
	SwiftNapError   uint8
}

func (m *MsgHeartbeat) FromBytes(bs []byte) error {
	if len(bs) < 4 {
		return io.ErrUnexpectedEOF
	}

	m.ExternalAntenna = bs[3] >> 7 & 0x1
	m.SbpVersionMajor = bs[2]
	m.SbpVersionMinor = bs[1]

	m.SystemError = bs[0] & 0x1
	m.IOError = bs[0] >> 1 & 0x1
	m.SwiftNapError = bs[0] >> 2 & 0x1

	return nil
}

func (m *MsgHeartbeat) Bytes() ([]byte, error) {
	bs := make([]byte, 4)

	bs[3] = m.ExternalAntenna & 0x1 << 7
	bs[2] = m.SbpVersionMajor
	bs[1] = m.SbpVersionMinor

	bs[0] = (m.SystemError & 0x1) | (m.IOError & 0x1 << 1) | (m.SwiftNapError & 0x1 << 2)

	return bs, nil
}
