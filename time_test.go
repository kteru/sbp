package sbp

import (
	"testing"
	"time"
)

func Test_GpsTime(t *testing.T) {
	tests := []struct {
		inWn    uint16
		inTow   uint32
		expTime string
	}{
		{
			inWn:    0,
			inTow:   0,
			expTime: "1980-01-06T00:00:00.000000000Z",
		},
		{
			inWn:    1940,
			inTow:   604799999,
			expTime: "2017-03-18T23:59:59.999000000Z",
		},
		{
			inWn:    1941,
			inTow:   0,
			expTime: "2017-03-19T00:00:00.000000000Z",
		},
	}

	for _, test := range tests {
		act := GpsTime(test.inWn, test.inTow)
		exp, _ := time.Parse(time.RFC3339Nano, test.expTime)

		if !act.Equal(exp) {
			t.Errorf("\n  actual: %#v\nexpected: %#v\n", act.In(time.UTC), exp.In(time.UTC))
		}
	}
}

func Benchmark_GpsTime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GpsTime(1, 1)
	}
}

func Test_Time_Wn(t *testing.T) {
	tests := []struct {
		inTime string
		exp    uint16
	}{
		{
			inTime: "1980-01-06T00:00:00.000000000Z",
			exp:    uint16(0),
		},
		{
			inTime: "2017-03-18T23:59:59.999000000Z",
			exp:    uint16(1940),
		},
		{
			inTime: "2017-03-19T00:00:00.000000000Z",
			exp:    uint16(1941),
		},
	}

	for _, test := range tests {
		v, _ := time.Parse(time.RFC3339Nano, test.inTime)

		act := Time{v}.Wn()
		exp := test.exp

		if act != exp {
			t.Errorf("\n  actual: %#v\nexpected: %#v\n", act, exp)
		}
	}
}

func Benchmark_Time_Wn(b *testing.B) {
	v := Time{time.Now()}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Wn()
	}
}

func Test_Time_Tow(t *testing.T) {
	tests := []struct {
		inTime string
		exp    uint32
	}{
		{
			inTime: "1980-01-06T00:00:00.000000000Z",
			exp:    uint32(0),
		},
		{
			inTime: "2017-03-18T23:59:59.999000000Z",
			exp:    uint32(604799999),
		},
		{
			inTime: "2017-03-19T00:00:00.000000000Z",
			exp:    uint32(0),
		},
	}

	for _, test := range tests {
		v, _ := time.Parse(time.RFC3339Nano, test.inTime)

		act := Time{v}.Tow()
		exp := test.exp

		if act != exp {
			t.Errorf("\n  actual: %#v\nexpected: %#v\n", act, exp)
		}
	}
}

func Benchmark_Time_Tow(b *testing.B) {
	v := Time{time.Now()}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Tow()
	}
}
