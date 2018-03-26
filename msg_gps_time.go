package sbp

import (
	"encoding/binary"
	"io"
)

// MsgGpsTime represents a contents of MSG_GPS_TIME.
type MsgGpsTime struct {
	// GPS week number (unit:weeks)
	Wn uint16

	// GPS time of week rounded to the nearest millisecond (unit:ms)
	Tow uint32

	// Nanosecond residual of millisecond-rounded TOW (unit:ns)
	NsResidual uint32

	// Status flags
	TimeSource uint8
}

func (m *MsgGpsTime) MsgType() uint16 {
	return TypeMsgGpsTime
}

func (m *MsgGpsTime) UnmarshalBinary(bs []byte) error {
	if len(bs) < 11 {
		return io.ErrUnexpectedEOF
	}

	m.Wn = binary.LittleEndian.Uint16(bs[0:2])
	m.Tow = binary.LittleEndian.Uint32(bs[2:6])
	m.NsResidual = binary.LittleEndian.Uint32(bs[6:10])

	flags := bs[10]
	m.TimeSource = flags & 0x7

	return nil
}

func (m *MsgGpsTime) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 11)

	binary.LittleEndian.PutUint16(bs[0:2], m.Wn)
	binary.LittleEndian.PutUint32(bs[2:6], m.Tow)
	binary.LittleEndian.PutUint32(bs[6:10], m.NsResidual)

	flags := m.TimeSource & 0x7
	bs[10] = flags

	return bs, nil
}
