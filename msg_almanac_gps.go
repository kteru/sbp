package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgAlmanacGps represents a contents of MSG_ALMANAC_GPS.
type MsgAlmanacGps struct {
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

	// Mean anomaly at reference time (unit:radians)
	M0 float64
	// Eccentricity of satellite orbit
	Ecc float64
	// Square root of the semi-major axis of orbit ( unit:m^(1/2) )
	Sqrta float64
	// Longitude of ascending node of orbit plane at weekly epoch (unit:rad)
	Omega0 float64
	// Rate of right ascension (unit:rad/s)
	Omegadot float64
	// Argument of perigee (unit:rad)
	W float64
	// Inclination (unit:rad)
	Inc float64
	// Polynomial clock correction coefficient (clock bias) (unit:s)
	Af0 float64
	// Polynomial clock correction coefficient (clock drift) (unit:s/s)
	Af1 float64
}

// MsgType returns the number representing the type.
func (m *MsgAlmanacGps) MsgType() uint16 {
	return TypeMsgAlmanacGps
}

// UnmarshalBinary parses a byte slice.
func (m *MsgAlmanacGps) UnmarshalBinary(bs []byte) error {
	if len(bs) < 94 {
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

	m.M0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[22:30]))
	m.Ecc = math.Float64frombits(binary.LittleEndian.Uint64(bs[30:38]))
	m.Sqrta = math.Float64frombits(binary.LittleEndian.Uint64(bs[38:46]))
	m.Omega0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[46:54]))
	m.Omegadot = math.Float64frombits(binary.LittleEndian.Uint64(bs[54:62]))
	m.W = math.Float64frombits(binary.LittleEndian.Uint64(bs[62:70]))
	m.Inc = math.Float64frombits(binary.LittleEndian.Uint64(bs[70:78]))
	m.Af0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[78:86]))
	m.Af1 = math.Float64frombits(binary.LittleEndian.Uint64(bs[86:94]))

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgAlmanacGps) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 94)

	bs[0] = m.SidSat
	bs[1] = m.SidCode
	binary.LittleEndian.PutUint32(bs[2:6], m.ToaTow)
	binary.LittleEndian.PutUint16(bs[6:8], m.ToaWn)
	binary.LittleEndian.PutUint64(bs[8:16], math.Float64bits(m.Ura))
	binary.LittleEndian.PutUint32(bs[16:20], m.FitInterval)
	bs[20] = m.Valid
	bs[21] = m.HealthBits

	binary.LittleEndian.PutUint64(bs[22:30], math.Float64bits(m.M0))
	binary.LittleEndian.PutUint64(bs[30:38], math.Float64bits(m.Ecc))
	binary.LittleEndian.PutUint64(bs[38:46], math.Float64bits(m.Sqrta))
	binary.LittleEndian.PutUint64(bs[46:54], math.Float64bits(m.Omega0))
	binary.LittleEndian.PutUint64(bs[54:62], math.Float64bits(m.Omegadot))
	binary.LittleEndian.PutUint64(bs[62:70], math.Float64bits(m.W))
	binary.LittleEndian.PutUint64(bs[70:78], math.Float64bits(m.Inc))
	binary.LittleEndian.PutUint64(bs[78:86], math.Float64bits(m.Af0))
	binary.LittleEndian.PutUint64(bs[86:94], math.Float64bits(m.Af1))

	return bs, nil
}
