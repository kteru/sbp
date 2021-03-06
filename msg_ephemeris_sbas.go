package sbp

import (
	"encoding/binary"
	"io"
	"math"
)

// MsgEphemerisSbas represents a contents of MSG_EPHEMERIS_SBAS.
type MsgEphemerisSbas struct {
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

	// Position of the GEO at time toe (unit:m)
	Pos0 float64
	Pos1 float64
	Pos2 float64
	// Velocity of the GEO at time toe (unit:m/s)
	Vel0 float64
	Vel1 float64
	Vel2 float64
	// Acceleration of the GEO at time toe (unit:m/s^2)
	Acc0 float64
	Acc1 float64
	Acc2 float64

	// Time offset of the GEO clock w.r.t. SBAS Network Time (unit:s)
	AGf0 float64
	// Drift of the GEO clock w.r.t. SBAS Network Time (unit:s/s)
	AGf1 float64
}

// MsgType returns the number representing the type.
func (m *MsgEphemerisSbas) MsgType() uint16 {
	return TypeMsgEphemerisSbas
}

// UnmarshalBinary parses a byte slice.
func (m *MsgEphemerisSbas) UnmarshalBinary(bs []byte) error {
	if len(bs) < 110 {
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
	m.Pos0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[22:30]))
	m.Pos1 = math.Float64frombits(binary.LittleEndian.Uint64(bs[30:38]))
	m.Pos2 = math.Float64frombits(binary.LittleEndian.Uint64(bs[38:46]))
	m.Vel0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[46:54]))
	m.Vel1 = math.Float64frombits(binary.LittleEndian.Uint64(bs[54:62]))
	m.Vel2 = math.Float64frombits(binary.LittleEndian.Uint64(bs[62:70]))
	m.Acc0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[70:78]))
	m.Acc1 = math.Float64frombits(binary.LittleEndian.Uint64(bs[78:86]))
	m.Acc2 = math.Float64frombits(binary.LittleEndian.Uint64(bs[86:94]))
	m.AGf0 = math.Float64frombits(binary.LittleEndian.Uint64(bs[94:102]))
	m.AGf1 = math.Float64frombits(binary.LittleEndian.Uint64(bs[102:110]))

	return nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (m *MsgEphemerisSbas) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 110)

	bs[0] = m.SidSat
	bs[1] = m.SidCode
	binary.LittleEndian.PutUint32(bs[2:6], m.ToeTow)
	binary.LittleEndian.PutUint16(bs[6:8], m.ToeWn)
	binary.LittleEndian.PutUint64(bs[8:16], math.Float64bits(m.Ura))
	binary.LittleEndian.PutUint32(bs[16:20], m.FitInterval)
	bs[20] = m.Valid
	bs[21] = m.HealthBits
	binary.LittleEndian.PutUint64(bs[22:30], math.Float64bits(m.Pos0))
	binary.LittleEndian.PutUint64(bs[30:38], math.Float64bits(m.Pos1))
	binary.LittleEndian.PutUint64(bs[38:46], math.Float64bits(m.Pos2))
	binary.LittleEndian.PutUint64(bs[46:54], math.Float64bits(m.Vel0))
	binary.LittleEndian.PutUint64(bs[54:62], math.Float64bits(m.Vel1))
	binary.LittleEndian.PutUint64(bs[62:70], math.Float64bits(m.Vel2))
	binary.LittleEndian.PutUint64(bs[70:78], math.Float64bits(m.Acc0))
	binary.LittleEndian.PutUint64(bs[78:86], math.Float64bits(m.Acc1))
	binary.LittleEndian.PutUint64(bs[86:94], math.Float64bits(m.Acc2))
	binary.LittleEndian.PutUint64(bs[94:102], math.Float64bits(m.AGf0))
	binary.LittleEndian.PutUint64(bs[102:110], math.Float64bits(m.AGf1))

	return bs, nil
}
