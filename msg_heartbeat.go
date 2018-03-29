package sbp

import "io"

// MsgHeartbeat represents a contents of MSG_HEARTBEAT.
type MsgHeartbeat struct {
	// Status flags
	SystemError            uint8
	IOError                uint8
	SwiftNapError          uint8
	SbpVersionMajor        uint8
	SbpVersionMinor        uint8
	ExternalAntennaShort   uint8
	ExternalAntennaPresent uint8
}

// MsgType returns the number representing the type.
func (m *MsgHeartbeat) MsgType() uint16 {
	return TypeMsgHeartbeat
}

// UnmarshalBinary parses a byte slice.
func (m *MsgHeartbeat) UnmarshalBinary(bs []byte) error {
	if len(bs) < 4 {
		return io.ErrUnexpectedEOF
	}

	m.SystemError = bs[0] & 0x1
	m.IOError = bs[0] >> 1 & 0x1
	m.SwiftNapError = bs[0] >> 2 & 0x1
	m.SbpVersionMinor = bs[1]
	m.SbpVersionMajor = bs[2]
	m.ExternalAntennaShort = bs[3] >> 6 & 0x1
	m.ExternalAntennaPresent = bs[3] >> 7 & 0x1

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgHeartbeat) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 4)

	bs[0] = (m.SystemError & 0x1) | (m.IOError & 0x1 << 1) | (m.SwiftNapError & 0x1 << 2)
	bs[1] = m.SbpVersionMinor
	bs[2] = m.SbpVersionMajor
	bs[3] = (m.ExternalAntennaShort & 0x1 << 6) | (m.ExternalAntennaPresent & 0x1 << 7)

	return bs, nil
}
