package sbp

import (
	"encoding/binary"
	"errors"
	"io"
)

// FramePreamble is value of Preamble.
const FramePreamble = 0x55

// SenderDeviceController is the sender id representing device controllers.
const SenderDeviceController = 0x42

var (
	// ErrInvalidCRC is returned when detect an incorrect CRC.
	ErrInvalidCRC = errors.New("invalid frame crc")
)

// Frame describes a Frame.
type Frame struct {
	Type    uint16
	Sender  uint16
	Payload []byte
}

// NewFrame parses a byte slice.
func NewFrame(bs []byte) (*Frame, error) {
	if len(bs) < 6+2 {
		return nil, io.ErrUnexpectedEOF
	}

	if bs[0] != FramePreamble {
		return nil, ErrInvalidFormat
	}

	plen := int(bs[5])

	if len(bs) != 6+plen+2 {
		return nil, ErrInvalidFormat
	}

	frameCrc := binary.LittleEndian.Uint16(bs[len(bs)-2 : len(bs)])
	crc := crc16ccitt(0, bs[1:len(bs)-2])
	if frameCrc != crc {
		return nil, ErrInvalidCRC
	}

	f := &Frame{
		Type:    binary.LittleEndian.Uint16(bs[1:3]),
		Sender:  binary.LittleEndian.Uint16(bs[3:5]),
		Payload: bs[6 : 6+plen],
	}

	return f, nil
}

// MarshalBinary returns a byte slice in accordance with the format.
func (f *Frame) MarshalBinary() ([]byte, error) {
	plen := len(f.Payload)
	if plen > 255 {
		return nil, ErrInvalidFormat
	}

	bs := make([]byte, 6+len(f.Payload)+2)

	bs[0] = FramePreamble
	binary.LittleEndian.PutUint16(bs[1:3], f.Type)
	binary.LittleEndian.PutUint16(bs[3:5], f.Sender)
	bs[5] = byte(plen)
	copy(bs[6:], f.Payload)

	crc := crc16ccitt(0, bs[1:len(bs)-2])
	binary.LittleEndian.PutUint16(bs[len(bs)-2:len(bs)], crc)

	return bs, nil
}

// Msg parse the Message payload and return it.
func (f *Frame) Msg() (Msg, error) {
	fn, ok := typeToMsg[f.Type]
	if !ok {
		return nil, ErrUnsupported
	}

	msg := fn()
	if err := msg.UnmarshalBinary(f.Payload); err != nil {
		return nil, err
	}

	return msg, nil
}

// SetMsg sets the Message to payload.
func (f *Frame) SetMsg(m Msg) error {
	if m == nil {
		return ErrUnsupported
	}

	if _, ok := typeToMsg[m.MsgType()]; !ok {
		return ErrUnsupported
	}

	bs, err := m.MarshalBinary()
	if err != nil {
		return err
	}

	f.Type = m.MsgType()
	f.Payload = bs

	return nil
}
