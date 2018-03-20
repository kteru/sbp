package sbp

import "errors"

// Message types
const (
	// Logging
	TypeMsgLog uint16 = 0x0401
	TypeMsgFwd uint16 = 0x0402

	// Navigation
	TypeMsgGpsTime         uint16 = 0x0102
	TypeMsgUtcTime         uint16 = 0x0103
	TypeMsgDops            uint16 = 0x0208
	TypeMsgPosEcef         uint16 = 0x0209
	TypeMsgPosLlh          uint16 = 0x020a
	TypeMsgBaselineEcef    uint16 = 0x020b
	TypeMsgBaselineNed     uint16 = 0x020c
	TypeMsgVelEcef         uint16 = 0x020d
	TypeMsgVelNed          uint16 = 0x020e
	TypeMsgBaselineHeading uint16 = 0x020f
	TypeMsgAgeCorrections  uint16 = 0x0210

	// Observation
	TypeMsgObs           uint16 = 0x004a
	TypeMsgBasePosLlh    uint16 = 0x0044
	TypeMsgBasePosEcef   uint16 = 0x0048
	TypeMsgEphemerisGps  uint16 = 0x0086
	TypeMsgEphemerisSbas uint16 = 0x0084
	TypeMsgEphemerisGlo  uint16 = 0x0088
	// TypeMsgIono               uint16 = 0x0090
	// TypeMsgSvConfigurationGps uint16 = 0x0091
	// TypeMsgGroupDelay         uint16 = 0x0093
	// TypeMsgAlmanacGps         uint16 = 0x0070
	// TypeMsgAlmanacGlo         uint16 = 0x0071

	// Settings
	TypeMsgSettingsSave            uint16 = 0x00a1
	TypeMsgSettingsWrite           uint16 = 0x00a0
	TypeMsgSettingsReadReq         uint16 = 0x00a4
	TypeMsgSettingsReadResp        uint16 = 0x00a5
	TypeMsgSettingsReadByIndexReq  uint16 = 0x00a2
	TypeMsgSettingsReadByIndexResp uint16 = 0x00a7
	TypeMsgSettingsReadByIndexDone uint16 = 0x00a6

	// System
	TypeMsgStartup     uint16 = 0xff00
	TypeMsgDgnssStatus uint16 = 0xff02
	TypeMsgHeartbeat   uint16 = 0xffff
)

// TypeToMsg is a map of constructors for Messages.
var TypeToMsg = map[uint16]func() Msg{
	TypeMsgLog:                     func() Msg { return new(MsgLog) },
	TypeMsgFwd:                     func() Msg { return new(MsgFwd) },
	TypeMsgGpsTime:                 func() Msg { return new(MsgGpsTime) },
	TypeMsgUtcTime:                 func() Msg { return new(MsgUtcTime) },
	TypeMsgDops:                    func() Msg { return new(MsgDops) },
	TypeMsgPosEcef:                 func() Msg { return new(MsgPosEcef) },
	TypeMsgPosLlh:                  func() Msg { return new(MsgPosLlh) },
	TypeMsgBaselineEcef:            func() Msg { return new(MsgBaselineEcef) },
	TypeMsgBaselineNed:             func() Msg { return new(MsgBaselineNed) },
	TypeMsgVelEcef:                 func() Msg { return new(MsgVelEcef) },
	TypeMsgVelNed:                  func() Msg { return new(MsgVelNed) },
	TypeMsgBaselineHeading:         func() Msg { return new(MsgBaselineHeading) },
	TypeMsgAgeCorrections:          func() Msg { return new(MsgAgeCorrections) },
	TypeMsgObs:                     func() Msg { return new(MsgObs) },
	TypeMsgBasePosLlh:              func() Msg { return new(MsgBasePosLlh) },
	TypeMsgBasePosEcef:             func() Msg { return new(MsgBasePosEcef) },
	TypeMsgEphemerisGps:            func() Msg { return new(MsgEphemerisGps) },
	TypeMsgEphemerisSbas:           func() Msg { return new(MsgEphemerisSbas) },
	TypeMsgEphemerisGlo:            func() Msg { return new(MsgEphemerisGlo) },
	TypeMsgSettingsSave:            func() Msg { return new(MsgSettingsSave) },
	TypeMsgSettingsWrite:           func() Msg { return new(MsgSettingsWrite) },
	TypeMsgSettingsReadReq:         func() Msg { return new(MsgSettingsReadReq) },
	TypeMsgSettingsReadResp:        func() Msg { return new(MsgSettingsReadResp) },
	TypeMsgSettingsReadByIndexReq:  func() Msg { return new(MsgSettingsReadByIndexReq) },
	TypeMsgSettingsReadByIndexResp: func() Msg { return new(MsgSettingsReadByIndexResp) },
	TypeMsgSettingsReadByIndexDone: func() Msg { return new(MsgSettingsReadByIndexDone) },
	TypeMsgStartup:                 func() Msg { return new(MsgStartup) },
	TypeMsgDgnssStatus:             func() Msg { return new(MsgDgnssStatus) },
	TypeMsgHeartbeat:               func() Msg { return new(MsgHeartbeat) },
}

var (
	// ErrInvalidMsg is returned when detect a malformed format.
	ErrInvalidMsg = errors.New("invalid message")
)

// Msg represents a Message contents.
type Msg interface {
	// UnmarshalBinary parses a byte slice.
	UnmarshalBinary([]byte) error

	// MarshalBinary returns a byte slice in accordance with the format.
	MarshalBinary() ([]byte, error)
}
