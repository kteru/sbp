package sbp

import "errors"

// Message types
const (
	// Logging
	TypeMsgLog uint16 = 0x0401

	// Navigation
	TypeMsgGpsTime         uint16 = 0x0100
	TypeMsgDops            uint16 = 0x0206
	TypeMsgPosEcef         uint16 = 0x0200
	TypeMsgPosLlh          uint16 = 0x0201
	TypeMsgBaselineEcef    uint16 = 0x0202
	TypeMsgBaselineNed     uint16 = 0x0203
	TypeMsgVelEcef         uint16 = 0x0204
	TypeMsgVelNed          uint16 = 0x0205
	TypeMsgBaselineHeading uint16 = 0x0207

	// Observation
	TypeMsgObs           uint16 = 0x0049
	TypeMsgBasePosLlh    uint16 = 0x0044
	TypeMsgBasePosEcef   uint16 = 0x0048
	TypeMsgEphemeris     uint16 = 0x0080
	TypeMsgEphemerisDepC uint16 = 0x0047

	// Settings
	TypeMsgSettingsSave            uint16 = 0x00a1
	TypeMsgSettingsWrite           uint16 = 0x00a0
	TypeMsgSettingsReadReq         uint16 = 0x00a4
	TypeMsgSettingsReadResp        uint16 = 0x00a5
	TypeMsgSettingsReadByIndexReq  uint16 = 0x00a2
	TypeMsgSettingsReadByIndexResp uint16 = 0x00a7
	TypeMsgSettingsReadByIndexDone uint16 = 0x00a6

	// System
	TypeMsgStartup   uint16 = 0xff00
	TypeMsgHeartbeat uint16 = 0xffff
)

// TypeToMsg is a map of constructors for Messages.
var TypeToMsg = map[uint16]func() Msg{
	TypeMsgGpsTime: func() Msg { return new(MsgGpsTime) },
	TypeMsgPosEcef: func() Msg { return new(MsgPosEcef) },
	TypeMsgPosLlh:  func() Msg { return new(MsgPosLlh) },
	TypeMsgVelEcef: func() Msg { return new(MsgVelEcef) },
	TypeMsgVelNed:  func() Msg { return new(MsgVelNed) },
}

var (
	// ErrInvalidMsg is returned when detect a malformed format.
	ErrInvalidMsg = errors.New("invalid message")
)

// Msg represents a Message contents.
type Msg interface {
	// FromBytes parses a byte slice.
	FromBytes([]byte) error

	// Bytes returns a byte slice in accordance with the format.
	Bytes() ([]byte, error)
}