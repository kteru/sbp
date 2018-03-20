package sbp

import (
	"encoding/binary"
	"io"
)

// MsgDgnssStatus represents a contents of MSG_DGNSS_STATUS.
type MsgDgnssStatus struct {
	// Status flags
	DifferentialType uint8

	// Latency of observation receipt (unit:deci-seconds)
	Latency uint16

	// Number of signals from base station
	NumSignals uint8

	// Corrections source string
	Source string
}

func (m *MsgDgnssStatus) MsgType() uint16 {
	return TypeMsgDgnssStatus
}

func (m *MsgDgnssStatus) UnmarshalBinary(bs []byte) error {
	if len(bs) < 4 {
		return io.ErrUnexpectedEOF
	}

	flags := bs[0]
	m.DifferentialType = flags & 0xf

	m.Latency = binary.LittleEndian.Uint16(bs[1:3])
	m.NumSignals = bs[3]

	m.Source = string(bs[4:])

	return nil
}

func (m *MsgDgnssStatus) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 4, 4+len(m.Source))

	flags := m.DifferentialType & 0xf
	bs[0] = flags

	binary.LittleEndian.PutUint16(bs[1:3], m.Latency)
	bs[3] = m.NumSignals

	bs = append(bs, []byte(m.Source)...)

	return bs, nil
}
