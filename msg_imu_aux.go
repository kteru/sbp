package sbp

import (
	"encoding/binary"
	"io"
)

// MsgImuAux represents a contents of MSG_IMU_AUX.
type MsgImuAux struct {
	// IMU type
	ImuType uint8

	// Raw IMU temperature
	Temp int16

	// IMU configuration
	ImuConf uint8
}

func (m *MsgImuAux) MsgType() uint16 {
	return TypeMsgImuAux
}

func (m *MsgImuAux) UnmarshalBinary(bs []byte) error {
	if len(bs) < 4 {
		return io.ErrUnexpectedEOF
	}

	m.ImuType = bs[0]
	m.Temp = int16(binary.LittleEndian.Uint16(bs[1:3]))
	m.ImuConf = bs[3]

	return nil
}

func (m *MsgImuAux) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 4)

	bs[0] = m.ImuType
	binary.LittleEndian.PutUint16(bs[1:3], uint16(m.Temp))
	bs[3] = m.ImuConf

	return bs, nil
}
