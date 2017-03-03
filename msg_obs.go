package sbp

import (
	"encoding/binary"
	"io"
)

// MsgObs represents a contents of MSG_OBS.
type MsgObs struct {
	// Milliseconds since start of GPS week (unit:ms)
	Tow uint32

	// GPS week number (unit:weeks)
	Wn uint16

	// Total number of observations
	NObs uint8

	// Observations
	Observations []*MsgObsObservation
}

func (m *MsgObs) FromBytes(bs []byte) error {
	if len(bs) < 7 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])
	m.Wn = binary.LittleEndian.Uint16(bs[4:6])
	m.NObs = bs[6]

	if len(bs[7:])%16 != 0 {
		return io.ErrUnexpectedEOF
	}

	n := len(bs[7:]) / 16
	m.Observations = make([]*MsgObsObservation, 0, n)

	for i := 0; i < n; i++ {
		o := 16*i + 7

		obs := &MsgObsObservation{}
		if err := obs.FromBytes(bs[o : o+16]); err != nil {
			return err
		}

		m.Observations = append(m.Observations, obs)
	}

	return nil
}

func (m *MsgObs) Bytes() ([]byte, error) {
	bs := make([]byte, 7, 7+16*len(m.Observations))

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)
	binary.LittleEndian.PutUint16(bs[4:6], m.Wn)
	bs[6] = m.NObs

	for _, obs := range m.Observations {
		b, err := obs.Bytes()
		if err != nil {
			return nil, err
		}

		bs = append(bs, b...)
	}

	return bs, nil
}

// MsgObsObservation represents a contents of a single observation included in a MSG_OBS.
type MsgObsObservation struct {
	// Pseudorange observation (unit:cm)
	P uint32

	// Carrier phase whole cycles (unit:cycles)
	Li int32
	// Carrier phase fractional part (unit:cycles/256)
	Lf uint8

	// Carrier-to-Noise density (unit:dBHz*4)
	CN0 uint8

	// Lock indicator
	Lock uint16

	// Constellation-specific satellite identifier
	SidSat uint16
	// Signal band
	Sidband uint8
	// Constellation to which the satellite belongs
	SidConstellation uint8
}

func (m *MsgObsObservation) FromBytes(bs []byte) error {
	if len(bs) < 16 {
		return io.ErrUnexpectedEOF
	}

	m.P = binary.LittleEndian.Uint32(bs[0:4])
	m.Li = int32(binary.LittleEndian.Uint32(bs[4:8]))
	m.Lf = bs[8]
	m.CN0 = bs[9]
	m.Lock = binary.LittleEndian.Uint16(bs[10:12])
	m.SidSat = binary.LittleEndian.Uint16(bs[12:14])
	m.Sidband = bs[14]
	m.SidConstellation = bs[15]

	return nil
}

func (m *MsgObsObservation) Bytes() ([]byte, error) {
	bs := make([]byte, 16)

	binary.LittleEndian.PutUint32(bs[0:4], m.P)
	binary.LittleEndian.PutUint32(bs[4:8], uint32(m.Li))
	bs[8] = m.Lf
	bs[9] = m.CN0
	binary.LittleEndian.PutUint16(bs[10:12], m.Lock)
	binary.LittleEndian.PutUint16(bs[12:14], m.SidSat)
	bs[14] = m.Sidband
	bs[15] = m.SidConstellation

	return bs, nil
}
