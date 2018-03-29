package sbp

import (
	"encoding/binary"
	"io"
)

// MsgObs represents a contents of MSG_OBS.
type MsgObs struct {
	// Milliseconds since start of GPS week (unit:ms)
	Tow uint32

	// Nanosecond residual of millisecond-rounded TOW (ranges from -500000 to 500000)
	NsResidual int32

	// GPS week number (unit:weeks)
	Wn uint16

	// Total number of observations
	NObs uint8

	// Observations
	Observations []*MsgObsObservation
}

// MsgType returns the number representing the type.
func (m *MsgObs) MsgType() uint16 {
	return TypeMsgObs
}

// UnmarshalBinary parses a byte slice.
func (m *MsgObs) UnmarshalBinary(bs []byte) error {
	if len(bs) < 11 {
		return io.ErrUnexpectedEOF
	}

	m.Tow = binary.LittleEndian.Uint32(bs[0:4])
	m.NsResidual = int32(binary.LittleEndian.Uint32(bs[4:8]))
	m.Wn = binary.LittleEndian.Uint16(bs[8:10])
	m.NObs = bs[10]

	if len(bs[11:])%17 != 0 {
		return io.ErrUnexpectedEOF
	}

	n := len(bs[11:]) / 17
	m.Observations = make([]*MsgObsObservation, 0, n)

	for i := 0; i < n; i++ {
		o := 17*i + 11

		obs := &MsgObsObservation{}
		if err := obs.UnmarshalBinary(bs[o : o+17]); err != nil {
			return err
		}

		m.Observations = append(m.Observations, obs)
	}

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgObs) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 11, 11+17*len(m.Observations))

	binary.LittleEndian.PutUint32(bs[0:4], m.Tow)
	binary.LittleEndian.PutUint32(bs[4:8], uint32(m.NsResidual))
	binary.LittleEndian.PutUint16(bs[8:10], m.Wn)
	bs[10] = m.NObs

	for _, obs := range m.Observations {
		b, err := obs.MarshalBinary()
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

	// Doppler whole (unit:Hz)
	Di int16
	// Doppler fractional part (unit:Hz/256)
	Df uint8

	// Carrier-to-Noise density (unit:dBHz*4)
	CN0 uint8

	// Lock indicator
	Lock uint8

	// Status flags
	PseudorangeValid   uint8
	CarrierPhaseValid  uint8
	HalfCycleAmbiguity uint8
	DopplerValid       uint8
	RaimExclusion      uint8

	// Constellation-specific satellite identifier
	SidSat uint8
	// Signal constellation, band and code
	SidCode uint8
}

// UnmarshalBinary parses a byte slice.
func (m *MsgObsObservation) UnmarshalBinary(bs []byte) error {
	if len(bs) < 17 {
		return io.ErrUnexpectedEOF
	}

	m.P = binary.LittleEndian.Uint32(bs[0:4])
	m.Li = int32(binary.LittleEndian.Uint32(bs[4:8]))
	m.Lf = bs[8]
	m.Di = int16(binary.LittleEndian.Uint16(bs[9:11]))
	m.Df = bs[11]
	m.CN0 = bs[12]
	m.Lock = bs[13]

	flags := bs[14]
	m.PseudorangeValid = flags & 0x1
	m.CarrierPhaseValid = flags >> 1 & 0x1
	m.HalfCycleAmbiguity = flags >> 2 & 0x1
	m.DopplerValid = flags >> 3 & 0x1
	m.RaimExclusion = flags >> 7 & 0x1

	m.SidSat = bs[15]
	m.SidCode = bs[16]

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgObsObservation) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 17)

	binary.LittleEndian.PutUint32(bs[0:4], m.P)
	binary.LittleEndian.PutUint32(bs[4:8], uint32(m.Li))
	bs[8] = m.Lf
	binary.LittleEndian.PutUint16(bs[9:11], uint16(m.Di))
	bs[11] = m.Df
	bs[12] = m.CN0
	bs[13] = m.Lock

	flags := (m.PseudorangeValid & 0x1) | (m.CarrierPhaseValid & 0x1 << 1) | (m.HalfCycleAmbiguity & 0x1 << 2) | (m.DopplerValid & 0x1 << 3) | (m.RaimExclusion & 0x1 << 7)
	bs[14] = flags

	bs[15] = m.SidSat
	bs[16] = m.SidCode

	return bs, nil
}
