package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgEphemeris represents a contents of MSG_EPHEMERIS.
type MsgEphemeris struct {
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

	// Time of week (unit:s)
	ToeTow float64
	// Week number (unit:week)
	ToeWn uint16
	// Clock reference time of week (unit:s)
	TocTow float64
	// Clock reference week number (unit:week)
	TocWn uint16

	// Is valid?
	Valid uint8
	// Satellite is healthy?
	Healthy uint8
	// Constellation-specific satellite identifier
	SidSat uint16
	// Signal band
	Sidband uint8
	// Constellation to which the satellite belongs
	SidConstellation uint8
	// Issue of ephemeris data
	Iode uint8
	// Issue of clock data
	Iodc uint16
	// Reserved field
	Reserved uint32
}

func (m *MsgEphemeris) FromBytes(bs []byte) error {
	if len(bs) < 185 {
		return io.ErrUnexpectedEOF
	}

	m.Tgd = math.Float64frombits(binary.LittleEndian.Uint64(bs[0:8]))
	m.CRs = math.Float64frombits(binary.LittleEndian.Uint64(bs[8:16]))
	m.CRc = math.Float64frombits(binary.LittleEndian.Uint64(bs[16:24]))
	m.CUc = math.Float64frombits(binary.LittleEndian.Uint64(bs[24:32]))
	m.CUs = math.Float64frombits(binary.LittleEndian.Uint64(bs[32:40]))
	m.CIc = math.Float64frombits(binary.LittleEndian.Uint64(bs[40:48]))
	m.CIs = math.Float64frombits(binary.LittleEndian.Uint64(bs[48:56]))
	m.Dn = math.Float64frombits(binary.LittleEndian.Uint64(bs[56:64]))
	m.M0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[64:72]))
	m.Ecc = math.Float64frombits(binary.LittleEndian.Uint64(bs[72:80]))
	m.Sqrta = math.Float64frombits(binary.LittleEndian.Uint64(bs[80:88]))
	m.Omega0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[88:96]))
	m.Omegadot = math.Float64frombits(binary.LittleEndian.Uint64(bs[96:104]))
	m.W = math.Float64frombits(binary.LittleEndian.Uint64(bs[104:112]))
	m.Inc = math.Float64frombits(binary.LittleEndian.Uint64(bs[112:120]))
	m.IncDot = math.Float64frombits(binary.LittleEndian.Uint64(bs[120:128]))
	m.Af0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[128:136]))
	m.Af1 = math.Float64frombits(binary.LittleEndian.Uint64(bs[136:144]))
	m.Af2 = math.Float64frombits(binary.LittleEndian.Uint64(bs[144:152]))
	m.ToeTow = math.Float64frombits(binary.LittleEndian.Uint64(bs[152:160]))
	m.ToeWn = binary.LittleEndian.Uint16(bs[160:162])
	m.TocTow = math.Float64frombits(binary.LittleEndian.Uint64(bs[162:170]))
	m.TocWn = binary.LittleEndian.Uint16(bs[170:172])
	m.Valid = bs[172]
	m.Healthy = bs[173]
	m.SidSat = binary.LittleEndian.Uint16(bs[174:176])
	m.Sidband = bs[176]
	m.SidConstellation = bs[177]
	m.Iode = bs[178]
	m.Iodc = binary.LittleEndian.Uint16(bs[179:181])
	m.Reserved = binary.LittleEndian.Uint32(bs[181:185])

	return nil
}

func (m *MsgEphemeris) Bytes() ([]byte, error) {
	bs := make([]byte, 185)

	binary.LittleEndian.PutUint64(bs[0:8], math.Float64bits(m.Tgd))
	binary.LittleEndian.PutUint64(bs[8:16], math.Float64bits(m.CRs))
	binary.LittleEndian.PutUint64(bs[16:24], math.Float64bits(m.CRc))
	binary.LittleEndian.PutUint64(bs[24:32], math.Float64bits(m.CUc))
	binary.LittleEndian.PutUint64(bs[32:40], math.Float64bits(m.CUs))
	binary.LittleEndian.PutUint64(bs[40:48], math.Float64bits(m.CIc))
	binary.LittleEndian.PutUint64(bs[48:56], math.Float64bits(m.CIs))
	binary.LittleEndian.PutUint64(bs[56:64], math.Float64bits(m.Dn))
	binary.LittleEndian.PutUint64(bs[64:72], math.Float64bits(m.M0))
	binary.LittleEndian.PutUint64(bs[72:80], math.Float64bits(m.Ecc))
	binary.LittleEndian.PutUint64(bs[80:88], math.Float64bits(m.Sqrta))
	binary.LittleEndian.PutUint64(bs[88:96], math.Float64bits(m.Omega0))
	binary.LittleEndian.PutUint64(bs[96:104], math.Float64bits(m.Omegadot))
	binary.LittleEndian.PutUint64(bs[104:112], math.Float64bits(m.W))
	binary.LittleEndian.PutUint64(bs[112:120], math.Float64bits(m.Inc))
	binary.LittleEndian.PutUint64(bs[120:128], math.Float64bits(m.IncDot))
	binary.LittleEndian.PutUint64(bs[128:136], math.Float64bits(m.Af0))
	binary.LittleEndian.PutUint64(bs[136:144], math.Float64bits(m.Af1))
	binary.LittleEndian.PutUint64(bs[144:152], math.Float64bits(m.Af2))
	binary.LittleEndian.PutUint64(bs[152:160], math.Float64bits(m.ToeTow))
	binary.LittleEndian.PutUint16(bs[160:162], m.ToeWn)
	binary.LittleEndian.PutUint64(bs[162:170], math.Float64bits(m.TocTow))
	binary.LittleEndian.PutUint16(bs[170:172], m.TocWn)
	bs[172] = m.Valid
	bs[173] = m.Healthy
	binary.LittleEndian.PutUint16(bs[174:176], m.SidSat)
	bs[176] = m.Sidband
	bs[177] = m.SidConstellation
	bs[178] = m.Iode
	binary.LittleEndian.PutUint16(bs[179:181], m.Iodc)
	binary.LittleEndian.PutUint32(bs[181:185], m.Reserved)

	return bs, nil
}
