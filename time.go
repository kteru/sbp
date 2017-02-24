package sbp

import "time"

// Time embeds time.Time to treat GPS time.
type Time struct {
	time.Time
}

// GpsTime returns the Time from the input GPS week number and GPS time of week in millisecond.
func GpsTime(wn uint16, tow uint32) Time {
	sec := int64(wn) * 604800
	nsec := int64(tow) * 1000000
	t := time.Unix(sec, nsec).Add(315964800 * time.Second)

	return Time{t}
}

// Wn returns the GPS week number.
func (t Time) Wn() uint16 {
	gt := t.Add(-315964800 * time.Second)
	wn := gt.Unix() / 604800

	return uint16(wn)
}

// Tow returns the GPS time of week in millisecond.
func (t Time) Tow() uint32 {
	gt := t.Add(-315964800 * time.Second)
	sec := gt.Unix() % 604800
	tow := sec*1000 + int64(gt.Nanosecond())/1000000

	return uint32(tow)
}
