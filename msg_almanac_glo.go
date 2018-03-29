package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgAlmanacGlo represents a contents of MSG_ALMANAC_GLO.
type MsgAlmanacGlo struct {
	// Constellation-specific satellite identifier
	SidSat uint8
	// Signal constellation, band and code
	SidCode uint8
	// Seconds since start of GPS week (unit:s)
	ToaTow uint32
	// GPS week number (unit:week)
	ToaWn uint16
	// User Range Accuracy (unit:m)
	Ura float64
	// Curve fit interval (unit:s)
	FitInterval uint32
	// Status of ephemeris (1 = valid, 0 = invalid)
	Valid uint8
	// Satellite health status
	HealthBits uint8

	// Longitude of the first ascending node of the orbit in PZ-90.02 coordinate system (unit:rad)
	LambdaNa float64
	// Time of the first ascending node passage (unit:s)
	TLambdaNa float64
	// Value of inclination at instant of t_lambda (unit:rad)
	I float64
	// Value of Draconian period at instant of t_lambda (unit:s/orbital period)
	T float64
	// Rate of change of the Draconian period (unit:s/(orbital period^2))
	TDot float64
	// Eccentricity at instant of t_lambda
	Epsilon float64
	// Argument of perigee at instant of t_lambda (unit:rad)
	Omega float64
}

// MsgType returns the number representing the type.
func (m *MsgAlmanacGlo) MsgType() uint16 {
	return TypeMsgAlmanacGlo
}

// UnmarshalBinary parses a byte slice.
func (m *MsgAlmanacGlo) UnmarshalBinary(bs []byte) error {
	if len(bs) < 78 {
		return io.ErrUnexpectedEOF
	}

	m.SidSat = bs[0]
	m.SidCode = bs[1]
	m.ToaTow = binary.LittleEndian.Uint32(bs[2:6])
	m.ToaWn = binary.LittleEndian.Uint16(bs[6:8])
	m.Ura = math.Float64frombits(binary.LittleEndian.Uint64(bs[8:16]))
	m.FitInterval = binary.LittleEndian.Uint32(bs[16:20])
	m.Valid = bs[20]
	m.HealthBits = bs[21]

	m.LambdaNa = math.Float64frombits(binary.LittleEndian.Uint64(bs[22:30]))
	m.TLambdaNa = math.Float64frombits(binary.LittleEndian.Uint64(bs[30:38]))
	m.I = math.Float64frombits(binary.LittleEndian.Uint64(bs[38:46]))
	m.T = math.Float64frombits(binary.LittleEndian.Uint64(bs[46:54]))
	m.TDot = math.Float64frombits(binary.LittleEndian.Uint64(bs[54:62]))
	m.Epsilon = math.Float64frombits(binary.LittleEndian.Uint64(bs[62:70]))
	m.Omega = math.Float64frombits(binary.LittleEndian.Uint64(bs[70:78]))

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgAlmanacGlo) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 78)

	bs[0] = m.SidSat
	bs[1] = m.SidCode
	binary.LittleEndian.PutUint32(bs[2:6], m.ToaTow)
	binary.LittleEndian.PutUint16(bs[6:8], m.ToaWn)
	binary.LittleEndian.PutUint64(bs[8:16], math.Float64bits(m.Ura))
	binary.LittleEndian.PutUint32(bs[16:20], m.FitInterval)
	bs[20] = m.Valid
	bs[21] = m.HealthBits

	binary.LittleEndian.PutUint64(bs[22:30], math.Float64bits(m.LambdaNa))
	binary.LittleEndian.PutUint64(bs[30:38], math.Float64bits(m.TLambdaNa))
	binary.LittleEndian.PutUint64(bs[38:46], math.Float64bits(m.I))
	binary.LittleEndian.PutUint64(bs[46:54], math.Float64bits(m.T))
	binary.LittleEndian.PutUint64(bs[54:62], math.Float64bits(m.TDot))
	binary.LittleEndian.PutUint64(bs[62:70], math.Float64bits(m.Epsilon))
	binary.LittleEndian.PutUint64(bs[70:78], math.Float64bits(m.Omega))

	return bs, nil
}
