package sbp

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func Test_NewFrame(t *testing.T) {
	tests := []struct {
		inBs   []byte
		exp    *Frame
		expErr error
	}{
		{
			inBs: []byte{0x55, 0xa6, 0x00, 0x42, 0x00, 0x00, 0x2c, 0x94},
			exp: &Frame{
				Type:    TypeMsgSettingsReadByIndexDone,
				Sender:  SenderDeviceController,
				Payload: []byte{},
			},
			expErr: nil,
		},
		{
			inBs:   []byte{0x55, 0xa6, 0x00, 0x42, 0x00, 0x00, 0x2c},
			exp:    nil,
			expErr: io.ErrUnexpectedEOF,
		},
		{
			inBs:   []byte{0x54, 0xa6, 0x00, 0x42, 0x00, 0x00, 0x2c, 0x94},
			exp:    nil,
			expErr: ErrInvalidFormat,
		},
		{
			inBs:   []byte{0x55, 0xa6, 0x00, 0x42, 0x00, 0x01, 0x2c, 0x94},
			exp:    nil,
			expErr: ErrInvalidFormat,
		},
		{
			inBs:   []byte{0x55, 0xa6, 0x00, 0x42, 0x00, 0x00, 0x2c, 0x95},
			exp:    nil,
			expErr: ErrInvalidCRC,
		},
	}

	for _, test := range tests {
		act, actErr := NewFrame(test.inBs)
		exp := test.exp
		expErr := test.expErr

		if actErr != expErr {
			t.Errorf("\n  actual: %#v\nexpected: %#v\n", actErr, expErr)
		}

		if actErr == nil && expErr == nil {
			if act.Type != exp.Type {
				t.Errorf("\n  actual: %#v\nexpected: %#v\n", act, exp)
			}
			if act.Sender != exp.Sender {
				t.Errorf("\n  actual: %#v\nexpected: %#v\n", act, exp)
			}
			if !bytes.Equal(act.Payload, exp.Payload) {
				t.Errorf("\n  actual: %#v\nexpected: %#v\n", act, exp)
			}
		}
	}
}

func Benchmark_NewFrame(b *testing.B) {
	v := []byte{0x55, 0xa6, 0x00, 0x42, 0x00, 0x00, 0x2c, 0x94}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewFrame(v)
	}
}

func Test_Frame_MarshalBinary(t *testing.T) {
	tests := []struct {
		in     *Frame
		exp    []byte
		expErr error
	}{
		{
			in: &Frame{
				Type:    TypeMsgSettingsReadByIndexDone,
				Sender:  SenderDeviceController,
				Payload: []byte{},
			},
			exp:    []byte{0x55, 0xa6, 0x00, 0x42, 0x00, 0x00, 0x2c, 0x94},
			expErr: nil,
		},
		{
			in: &Frame{
				Type:    TypeMsgSettingsReadByIndexDone,
				Sender:  SenderDeviceController,
				Payload: make([]byte, 256),
			},
			exp:    nil,
			expErr: ErrInvalidFormat,
		},
	}

	for _, test := range tests {
		act, actErr := test.in.MarshalBinary()
		exp := test.exp
		expErr := test.expErr

		if actErr != expErr {
			t.Errorf("\n  actual: %#v\nexpected: %#v\n", actErr, expErr)
		}

		if actErr == nil && expErr == nil {
			if !bytes.Equal(act, exp) {
				t.Errorf("\n  actual: %#v\nexpected: %#v\n", act, exp)
			}
		}
	}
}

func Benchmark_Frame_MarshalBinary(b *testing.B) {
	v := &Frame{
		Type:    TypeMsgSettingsReadByIndexDone,
		Sender:  SenderDeviceController,
		Payload: []byte{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = v.MarshalBinary()
	}
}

func Test_Frame_Msg(t *testing.T) {
	tests := []struct {
		in     *Frame
		exp    Msg
		expErr error
	}{
		{
			in: &Frame{
				Type:    TypeMsgSettingsReadByIndexDone,
				Sender:  SenderDeviceController,
				Payload: []byte{},
			},
			exp:    new(MsgSettingsReadByIndexDone),
			expErr: nil,
		},
		{
			in: &Frame{
				Type:    0x0000,
				Sender:  SenderDeviceController,
				Payload: []byte{},
			},
			exp:    nil,
			expErr: ErrUnsupported,
		},
	}

	for _, test := range tests {
		act, actErr := test.in.Msg()
		exp := test.exp
		expErr := test.expErr

		if actErr != expErr {
			t.Errorf("\n  actual: %#v\nexpected: %#v\n", actErr, expErr)
		}

		if actErr == nil && expErr == nil {
			acType := reflect.ValueOf(act).Type().String()
			exType := reflect.ValueOf(exp).Type().String()

			if acType != exType {
				t.Errorf("\n  actual: %#v\nexpected: %#v\n", acType, exType)
			}
		}
	}
}

func Benchmark_Frame_Msg(b *testing.B) {
	v := &Frame{
		Type:    TypeMsgSettingsReadByIndexDone,
		Sender:  SenderDeviceController,
		Payload: []byte{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = v.Msg()
	}
}

func Test_Frame_SetMsg(t *testing.T) {
	tests := []struct {
		in     Msg
		exp    *Frame
		expErr error
	}{
		{
			in: new(MsgSettingsReadByIndexDone),
			exp: &Frame{
				Type:    TypeMsgSettingsReadByIndexDone,
				Payload: []byte{},
			},
			expErr: nil,
		},
	}

	for _, test := range tests {
		act := &Frame{}
		actErr := act.SetMsg(test.in)
		exp := test.exp
		expErr := test.expErr

		if actErr != expErr {
			t.Errorf("\n  actual: %#v\nexpected: %#v\n", actErr, expErr)
		}

		if actErr == nil && expErr == nil {
			acType := reflect.ValueOf(act).Type().String()
			exType := reflect.ValueOf(exp).Type().String()

			if acType != exType {
				t.Errorf("\n  actual: %#v\nexpected: %#v\n", acType, exType)
			}
		}
	}
}

func Benchmark_Frame_SetMsg(b *testing.B) {
	f := &Frame{}
	m := new(MsgSettingsReadByIndexDone)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.SetMsg(m)
	}
}
