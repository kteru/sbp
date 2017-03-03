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
	Ns uint32

	// Status flags
	Flags uint8
}

func (m *MsgGpsTime) FromBytes(bs []byte) error {
	if len(bs) < 11 {
		return io.ErrUnexpectedEOF
	}

	m.Wn = binary.LittleEndian.Uint16(bs[0:2])
	m.Tow = binary.LittleEndian.Uint32(bs[2:6])
	m.Ns = binary.LittleEndian.Uint32(bs[6:10])

	m.Flags = bs[10]

	return nil
}

func (m *MsgGpsTime) Bytes() ([]byte, error) {
	bs := make([]byte, 11)

	binary.LittleEndian.PutUint16(bs[0:2], m.Wn)
	binary.LittleEndian.PutUint32(bs[2:6], m.Tow)
	binary.LittleEndian.PutUint32(bs[6:10], m.Ns)

	bs[10] = m.Flags

	return bs, nil
}
