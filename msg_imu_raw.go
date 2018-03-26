package sbp

import (
	"encoding/binary"
	"io"
)

// MsgImuRaw represents a contents of MSG_IMU_RAW.
type MsgImuRaw struct {
	// Milliseconds since start of GPS week (unit:ms)
	Tow uint32

	// Milliseconds since start of GPS week, fractional part (unit:ms/256)
	TowF uint8

	// Acceleration in the IMU frame
	AccX int16
	AccY int16
	AccZ int16

	// Angular rate around IMU frame
	GyrX int16
	GyrY int16
	GyrZ int16
}

func (m *MsgImuRaw) MsgType() uint16 {
	return TypeMsgImuRaw
}

func (m *MsgImuRaw) UnmarshalBinary(bs []byte) error {
	if len(bs) < 17 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])
	m.TowF = bs[4]

	m.AccX = int16(binary.LittleEndian.Uint16(bs[5:7]))
	m.AccY = int16(binary.LittleEndian.Uint16(bs[7:9]))
	m.AccZ = int16(binary.LittleEndian.Uint16(bs[9:11]))
	m.GyrX = int16(binary.LittleEndian.Uint16(bs[11:13]))
	m.GyrY = int16(binary.LittleEndian.Uint16(bs[13:15]))
	m.GyrZ = int16(binary.LittleEndian.Uint16(bs[15:17]))

	return nil
}

func (m *MsgImuRaw) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 17)

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)
	bs[4] = m.TowF

	binary.LittleEndian.PutUint16(bs[5:7], uint16(m.AccX))
	binary.LittleEndian.PutUint16(bs[7:9], uint16(m.AccY))
	binary.LittleEndian.PutUint16(bs[9:11], uint16(m.AccZ))
	binary.LittleEndian.PutUint16(bs[11:13], uint16(m.GyrX))
	binary.LittleEndian.PutUint16(bs[13:15], uint16(m.GyrY))
	binary.LittleEndian.PutUint16(bs[15:17], uint16(m.GyrZ))

	return bs, nil
}
