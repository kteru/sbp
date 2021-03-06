package sbp

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func Test_MsgLog_UnmarshalBinary(t *testing.T) {
	tests := []struct {
		in     []byte
		exp    *MsgLog
		expErr error
	}{
		{
			in: []byte{0x00},
			exp: &MsgLog{
				Level: 0,
				Text:  "",
			},
			expErr: nil,
		},
		{
			in: []byte{0x07, 0x61, 0x62, 0x63, 0x64, 0x65},
			exp: &MsgLog{
				Level: 7,
				Text:  "abcde",
			},
			expErr: nil,
		},
		{
			in:     []byte{},
			exp:    nil,
			expErr: io.ErrUnexpectedEOF,
		},
	}

	for _, test := range tests {
		act := new(MsgLog)
		actErr := act.UnmarshalBinary(test.in)
		exp := test.exp
		expErr := test.expErr

		if actErr != expErr {
			t.Errorf("\n  actual: %#v\nexpected: %#v\n", actErr, expErr)
		}

		if actErr == nil && expErr == nil {
			if !reflect.DeepEqual(act, exp) {
				t.Errorf("\n  actual: %#v\nexpected: %#v\n", act, exp)
			}
		}
	}
}

func Benchmark_MsgLog_UnmarshalBinary(b *testing.B) {
	m := new(MsgLog)
	bs := []byte{0x07, 0x61, 0x62, 0x63, 0x64, 0x65}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.UnmarshalBinary(bs)
	}
}

func Test_MsgLog_MarshalBinary(t *testing.T) {
	tests := []struct {
		in     *MsgLog
		exp    []byte
		expErr error
	}{
		{
			in: &MsgLog{
				Level: 0,
				Text:  "",
			},
			exp:    []byte{0x00},
			expErr: nil,
		},
		{
			in: &MsgLog{
				Level: 7,
				Text:  "abcde",
			},
			exp:    []byte{0x07, 0x61, 0x62, 0x63, 0x64, 0x65},
			expErr: nil,
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

func Benchmark_MsgLog_MarshalBinary(b *testing.B) {
	m := &MsgLog{
		Level: 7,
		Text:  "abcde",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.MarshalBinary()
	}
}
