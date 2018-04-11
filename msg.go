package sbp

// Message types
const (
	//
	// Stable messages
	//

	// Ext Events
	TypeMsgExtEvent uint16 = 0x0101

	// Imu
	TypeMsgImuRaw uint16 = 0x0900
	TypeMsgImuAux uint16 = 0x0901

	// Logging
	TypeMsgLog uint16 = 0x0401
	TypeMsgFwd uint16 = 0x0402

	// Mag
	TypeMsgMagRaw uint16 = 0x0902

	// Navigation
	TypeMsgGpsTime        uint16 = 0x0102
	TypeMsgUtcTime        uint16 = 0x0103
	TypeMsgDops           uint16 = 0x0208
	TypeMsgPosEcef        uint16 = 0x0209
	TypeMsgPosEcefCov     uint16 = 0x0214
	TypeMsgPosLlh         uint16 = 0x020a
	TypeMsgPosLlhCov      uint16 = 0x0211
	TypeMsgBaselineEcef   uint16 = 0x020b
	TypeMsgBaselineNed    uint16 = 0x020c
	TypeMsgVelEcef        uint16 = 0x020d
	TypeMsgVelEcefCov     uint16 = 0x0215
	TypeMsgVelNed         uint16 = 0x020e
	TypeMsgVelNedCov      uint16 = 0x0212
	TypeMsgVelBody        uint16 = 0x0213
	TypeMsgAgeCorrections uint16 = 0x0210

	// Observation
	TypeMsgObs                uint16 = 0x004a
	TypeMsgBasePosLlh         uint16 = 0x0044
	TypeMsgBasePosEcef        uint16 = 0x0048
	TypeMsgEphemerisGps       uint16 = 0x0086
	TypeMsgEphemerisSbas      uint16 = 0x0084
	TypeMsgEphemerisGlo       uint16 = 0x0088
	TypeMsgIono               uint16 = 0x0090
	TypeMsgSvConfigurationGps uint16 = 0x0091
	TypeMsgGroupDelay         uint16 = 0x0094
	TypeMsgAlmanacGps         uint16 = 0x0072
	TypeMsgAlmanacGlo         uint16 = 0x0073
	TypeMsgGloBiases          uint16 = 0x0075

	// Settings
	TypeMsgSettingsSave            uint16 = 0x00a1
	TypeMsgSettingsWrite           uint16 = 0x00a0
	TypeMsgSettingsWriteResp       uint16 = 0x00af
	TypeMsgSettingsReadReq         uint16 = 0x00a4
	TypeMsgSettingsReadResp        uint16 = 0x00a5
	TypeMsgSettingsReadByIndexReq  uint16 = 0x00a2
	TypeMsgSettingsReadByIndexResp uint16 = 0x00a7
	TypeMsgSettingsReadByIndexDone uint16 = 0x00a6

	// System
	TypeMsgStartup     uint16 = 0xff00
	TypeMsgDgnssStatus uint16 = 0xff02
	TypeMsgHeartbeat   uint16 = 0xffff
	TypeMsgInsStatus   uint16 = 0xff03

	//
	// Draft messages
	//

	// Orientation
	TypeMsgBaselineHeading uint16 = 0x020f

	// Piksi
	TypeMsgReset uint16 = 0x00b6
)

// typeToMsg is a map of constructors for Messages.
var typeToMsg = map[uint16]func() Msg{
	TypeMsgExtEvent:                func() Msg { return new(MsgExtEvent) },
	TypeMsgImuRaw:                  func() Msg { return new(MsgImuRaw) },
	TypeMsgImuAux:                  func() Msg { return new(MsgImuAux) },
	TypeMsgLog:                     func() Msg { return new(MsgLog) },
	TypeMsgFwd:                     func() Msg { return new(MsgFwd) },
	TypeMsgMagRaw:                  func() Msg { return new(MsgMagRaw) },
	TypeMsgGpsTime:                 func() Msg { return new(MsgGpsTime) },
	TypeMsgUtcTime:                 func() Msg { return new(MsgUtcTime) },
	TypeMsgDops:                    func() Msg { return new(MsgDops) },
	TypeMsgPosEcef:                 func() Msg { return new(MsgPosEcef) },
	TypeMsgPosEcefCov:              func() Msg { return new(MsgPosEcefCov) },
	TypeMsgPosLlh:                  func() Msg { return new(MsgPosLlh) },
	TypeMsgPosLlhCov:               func() Msg { return new(MsgPosLlhCov) },
	TypeMsgBaselineEcef:            func() Msg { return new(MsgBaselineEcef) },
	TypeMsgBaselineNed:             func() Msg { return new(MsgBaselineNed) },
	TypeMsgVelEcef:                 func() Msg { return new(MsgVelEcef) },
	TypeMsgVelEcefCov:              func() Msg { return new(MsgVelEcefCov) },
	TypeMsgVelNed:                  func() Msg { return new(MsgVelNed) },
	TypeMsgVelNedCov:               func() Msg { return new(MsgVelNedCov) },
	TypeMsgVelBody:                 func() Msg { return new(MsgVelBody) },
	TypeMsgAgeCorrections:          func() Msg { return new(MsgAgeCorrections) },
	TypeMsgObs:                     func() Msg { return new(MsgObs) },
	TypeMsgBasePosLlh:              func() Msg { return new(MsgBasePosLlh) },
	TypeMsgBasePosEcef:             func() Msg { return new(MsgBasePosEcef) },
	TypeMsgEphemerisGps:            func() Msg { return new(MsgEphemerisGps) },
	TypeMsgEphemerisSbas:           func() Msg { return new(MsgEphemerisSbas) },
	TypeMsgEphemerisGlo:            func() Msg { return new(MsgEphemerisGlo) },
	TypeMsgIono:                    func() Msg { return new(MsgIono) },
	TypeMsgSvConfigurationGps:      func() Msg { return new(MsgSvConfigurationGps) },
	TypeMsgGroupDelay:              func() Msg { return new(MsgGroupDelay) },
	TypeMsgAlmanacGps:              func() Msg { return new(MsgAlmanacGps) },
	TypeMsgAlmanacGlo:              func() Msg { return new(MsgAlmanacGlo) },
	TypeMsgGloBiases:               func() Msg { return new(MsgGloBiases) },
	TypeMsgSettingsSave:            func() Msg { return new(MsgSettingsSave) },
	TypeMsgSettingsWrite:           func() Msg { return new(MsgSettingsWrite) },
	TypeMsgSettingsWriteResp:       func() Msg { return new(MsgSettingsWriteResp) },
	TypeMsgSettingsReadReq:         func() Msg { return new(MsgSettingsReadReq) },
	TypeMsgSettingsReadResp:        func() Msg { return new(MsgSettingsReadResp) },
	TypeMsgSettingsReadByIndexReq:  func() Msg { return new(MsgSettingsReadByIndexReq) },
	TypeMsgSettingsReadByIndexResp: func() Msg { return new(MsgSettingsReadByIndexResp) },
	TypeMsgSettingsReadByIndexDone: func() Msg { return new(MsgSettingsReadByIndexDone) },
	TypeMsgStartup:                 func() Msg { return new(MsgStartup) },
	TypeMsgDgnssStatus:             func() Msg { return new(MsgDgnssStatus) },
	TypeMsgHeartbeat:               func() Msg { return new(MsgHeartbeat) },
	TypeMsgInsStatus:               func() Msg { return new(MsgInsStatus) },
	TypeMsgBaselineHeading:         func() Msg { return new(MsgBaselineHeading) },
	TypeMsgReset:                   func() Msg { return new(MsgReset) },
}

// Msg represents a Message contents.
type Msg interface {
	// MsgType returns the number representing the type.
	MsgType() uint16

	// UnmarshalBinary parses a byte slice.
	UnmarshalBinary([]byte) error

	// MarshalBinary returns a byte slice in accordance with the format.
	MarshalBinary() ([]byte, error)
}
