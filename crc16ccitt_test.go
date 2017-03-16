package sbp

import "testing"

func Test_crc16ccitt(t *testing.T) {
	tests := []struct {
		inInit uint16
		inBs   []byte
		exp    uint16
	}{
		{
			inInit: 0xffff,
			inBs:   []byte("123456789"),
			exp:    0x29b1,
		},
		{
			inInit: 0xffff,
			inBs:   []byte("1234567890"),
			exp:    0x3218,
		},
	}

	for _, test := range tests {
		act := crc16ccitt(test.inInit, test.inBs)
		exp := test.exp

		if act != exp {
			t.Errorf("\n  actual: %#v\nexpected: %#v\n", act, exp)
		}
	}
}

func Benchmark_crc16ccitt(b *testing.B) {
	vInit := uint16(0xffff)
	vBs := []byte("123456789")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = crc16ccitt(vInit, vBs)
	}
}
