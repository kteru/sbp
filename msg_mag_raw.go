package sbp

import (
	"encoding/binary"
	"io"
)

// MsgMagRaw represents a contents of MSG_MAG_RAW.
type MsgMagRaw struct {
	// Milliseconds since start of GPS week (unit:ms)
	Tow uint32

	// Milliseconds since start of GPS week, fractional part (unit:ms/256)
	TowF uint8

	// Magnetic field in the body frame
	MagX int16
	MagY int16
	MagZ int16
}

func (m *MsgMagRaw) MsgType() uint16 {
	return TypeMsgMagRaw
}

func (m *MsgMagRaw) UnmarshalBinary(bs []byte) error {
	if len(bs) < 11 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])
	m.TowF = bs[4]

	m.MagX = int16(binary.LittleEndian.Uint16(bs[5:7]))
	m.MagY = int16(binary.LittleEndian.Uint16(bs[7:9]))
	m.MagZ = int16(binary.LittleEndian.Uint16(bs[9:11]))

	return nil
}

func (m *MsgMagRaw) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 11)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)
	bs[4] = m.TowF

	binary.LittleEndian.PutUint16(bs[5:7], uint16(m.MagX))
	binary.LittleEndian.PutUint16(bs[7:9], uint16(m.MagY))
	binary.LittleEndian.PutUint16(bs[9:11], uint16(m.MagZ))

	return bs, nil
}
