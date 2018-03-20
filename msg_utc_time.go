package sbp

import (
	"encoding/binary"
	"io"
	"time"
)

// MsgUtcTime represents a contents of MSG_UTC_TIME.
type MsgUtcTime struct {
	// Status flags
	TimeSource      uint8
	UtcOffsetSource uint8

	// GPS time of week rounded to the nearest millisecond (unit:ms)
	Tow uint32

	// Time
	Time time.Time
}

func (m *MsgUtcTime) UnmarshalBinary(bs []byte) error {
	if len(bs) < 16 {
		return io.ErrUnexpectedEOF
	}

	flags := bs[0]
	m.TimeSource = flags & 0x7
	m.UtcOffsetSource = flags >> 3 & 0x3

	m.Tow = binary.LittleEndian.Uint32(bs[1:5])

	year := int(binary.LittleEndian.Uint16(bs[5:7]))
	month := time.Month(bs[7])
	day := int(bs[8])
	hour := int(bs[9])
	min := int(bs[10])
	sec := int(bs[11])
	nsec := int(binary.LittleEndian.Uint32(bs[12:16]))

	t := time.Date(year, month, day, hour, min, sec, nsec, time.UTC)
	m.Time = t.Local()

	return nil
}

func (m *MsgUtcTime) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 16)

	flags := (m.TimeSource & 0x7) | (m.UtcOffsetSource & 0x3 << 3)
	bs[0] = flags

	t := m.Time.UTC()
	binary.LittleEndian.PutUint16(bs[5:7], uint16(t.Year()))
	bs[7] = uint8(t.Month())
	bs[8] = uint8(t.Day())
	bs[9] = uint8(t.Hour())
	bs[10] = uint8(t.Minute())
	bs[11] = uint8(t.Second())
	binary.LittleEndian.PutUint32(bs[12:16], uint32(t.Nanosecond()))

	return bs, nil
}
