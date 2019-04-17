package datetime

import "time"

func Now(timezone ...*time.Location) time.Time {
	if len(timezone) > 0 {
		return time.Now().In(timezone[0])
	} else {
		return time.Now().In(DefaultTZ())
	}
}

func NowZH() time.Time {
	return Now(BeijingTZ())
}

func NowUnix(timezone ...*time.Location) int64 {
	return Now(timezone...).Unix()
}

func NowUnixNano(timezone ...*time.Location) int64 {
	return Now(timezone...).UnixNano()
}
