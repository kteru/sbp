package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgEphemerisGps represents a contents of MSG_EPHEMERIS_GPS.
type MsgEphemerisGps struct {
	// Constellation-specific satellite identifier
	SidSat uint8
	// Signal constellation, band and code
	SidCode uint8
	// Seconds since start of GPS week (unit:s)
	ToeTow uint32
	// GPS week number (unit:week)
	ToeWn uint16
	// User Range Accuracy
	Ura float64
	// Curefit interval
	FitInterval uint32
	// Status of ephemeris (1 = valid, 0 = invalid)
	Valid uint8
	// Satellite health status
	HealthBits uint8

	// Group delay differential between L1 and L2 (unit:s)
	Tgd float64

	// Amplitude of the sine harmonic correction term to the orbit radius (unit:m)
	CRs float64
	// Amplitude of the cosine harmonic correction term to the orbit radius (unit:m)
	CRc float64
	// Amplitude of the cosine harmonic correction term to the argument of latitude (unit:rad)
	CUc float64
	// Amplitude of the sine harmonic correction term to the argument of latitude (unit:rad)
	CUs float64
	// Amplitude of the cosine harmonic correction term to the angle of inclination (unit:rad)
	CIc float64
	// Amplitude of the sine harmonic correction term to the angle of inclination (unit:rad)
	CIs float64

	// Mean motion difference (unit:rad/s)
	Dn float64
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
	// Inclination first derivative (unit:rad)
	IncDot float64

	// Polynomial clock correction coefficient (clock bias) (unit:s)
	Af0 float64
	// Polynomial clock correction coefficient (clock drift) (unit:s/s)
	Af1 float64
	// Polynomial clock correction coefficient (rate of clock drift) (unit:s/s^2)
	Af2 float64

	// Clock reference time of week (unit:s)
	TocTow uint32
	// Clock reference week number (unit:week)
	TocWn uint16

	// Issue of ephemeris data
	Iode uint8
	// Issue of clock data
	Iodc uint16
}

func (m *MsgEphemerisGps) UnmarshalBinary(bs []byte) error {
	if len(bs) < 183 {
		return io.ErrUnexpectedEOF
	}

	m.SidSat = bs[0]
	m.SidCode = bs[1]
	m.ToeTow = binary.LittleEndian.Uint32(bs[2:6])
	m.ToeWn = binary.LittleEndian.Uint16(bs[6:8])
	m.Ura = math.Float64frombits(binary.LittleEndian.Uint64(bs[8:16]))
	m.FitInterval = binary.LittleEndian.Uint32(bs[16:20])
	m.Valid = bs[20]
	m.HealthBits = bs[21]
	m.Tgd = math.Float64frombits(binary.LittleEndian.Uint64(bs[22:30]))
	m.CRs = math.Float64frombits(binary.LittleEndian.Uint64(bs[30:38]))
	m.CRc = math.Float64frombits(binary.LittleEndian.Uint64(bs[38:46]))
	m.CUc = math.Float64frombits(binary.LittleEndian.Uint64(bs[46:54]))
	m.CUs = math.Float64frombits(binary.LittleEndian.Uint64(bs[54:62]))
	m.CIc = math.Float64frombits(binary.LittleEndian.Uint64(bs[62:70]))
	m.CIs = math.Float64frombits(binary.LittleEndian.Uint64(bs[70:78]))
	m.Dn = math.Float64frombits(binary.LittleEndian.Uint64(bs[78:86]))
	m.M0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[86:94]))
	m.Ecc = math.Float64frombits(binary.LittleEndian.Uint64(bs[94:102]))
	m.Sqrta = math.Float64frombits(binary.LittleEndian.Uint64(bs[102:110]))
	m.Omega0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[110:118]))
	m.Omegadot = math.Float64frombits(binary.LittleEndian.Uint64(bs[118:126]))
	m.W = math.Float64frombits(binary.LittleEndian.Uint64(bs[126:134]))
	m.Inc = math.Float64frombits(binary.LittleEndian.Uint64(bs[134:142]))
	m.IncDot = math.Float64frombits(binary.LittleEndian.Uint64(bs[142:150]))
	m.Af0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[150:158]))
	m.Af1 = math.Float64frombits(binary.LittleEndian.Uint64(bs[158:166]))
	m.Af2 = math.Float64frombits(binary.LittleEndian.Uint64(bs[166:174]))
	m.TocTow = binary.LittleEndian.Uint32(bs[174:178])
	m.TocWn = binary.LittleEndian.Uint16(bs[178:180])
	m.Iode = bs[180]
	m.Iodc = binary.LittleEndian.Uint16(bs[181:183])

	return nil
}

func (m *MsgEphemerisGps) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 183)

	bs[0] = m.SidSat
	bs[1] = m.SidCode
	binary.LittleEndian.PutUint32(bs[2:6], m.ToeTow)
	binary.LittleEndian.PutUint16(bs[6:8], m.ToeWn)
	binary.LittleEndian.PutUint64(bs[8:16], math.Float64bits(m.Ura))
	binary.LittleEndian.PutUint32(bs[16:20], m.FitInterval)
	bs[20] = m.Valid
	bs[21] = m.HealthBits
	binary.LittleEndian.PutUint64(bs[22:30], math.Float64bits(m.Tgd))
	binary.LittleEndian.PutUint64(bs[30:38], math.Float64bits(m.CRs))
	binary.LittleEndian.PutUint64(bs[38:46], math.Float64bits(m.CRc))
	binary.LittleEndian.PutUint64(bs[46:54], math.Float64bits(m.CUc))
	binary.LittleEndian.PutUint64(bs[54:62], math.Float64bits(m.CUs))
	binary.LittleEndian.PutUint64(bs[62:70], math.Float64bits(m.CIc))
	binary.LittleEndian.PutUint64(bs[70:78], math.Float64bits(m.CIs))
	binary.LittleEndian.PutUint64(bs[78:86], math.Float64bits(m.Dn))
	binary.LittleEndian.PutUint64(bs[86:94], math.Float64bits(m.M0))
	binary.LittleEndian.PutUint64(bs[94:102], math.Float64bits(m.Ecc))
	binary.LittleEndian.PutUint64(bs[102:110], math.Float64bits(m.Sqrta))
	binary.LittleEndian.PutUint64(bs[110:118], math.Float64bits(m.Omega0))
	binary.LittleEndian.PutUint64(bs[118:126], math.Float64bits(m.Omegadot))
	binary.LittleEndian.PutUint64(bs[126:134], math.Float64bits(m.W))
	binary.LittleEndian.PutUint64(bs[134:142], math.Float64bits(m.Inc))
	binary.LittleEndian.PutUint64(bs[142:150], math.Float64bits(m.IncDot))
	binary.LittleEndian.PutUint64(bs[150:158], math.Float64bits(m.Af0))
	binary.LittleEndian.PutUint64(bs[158:166], math.Float64bits(m.Af1))
	binary.LittleEndian.PutUint64(bs[166:174], math.Float64bits(m.Af2))
	binary.LittleEndian.PutUint32(bs[174:178], m.TocTow)
	binary.LittleEndian.PutUint16(bs[178:180], m.TocWn)
	bs[180] = m.Iode
	binary.LittleEndian.PutUint16(bs[181:183], m.Iodc)

	return bs, nil
}
