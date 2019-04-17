package datetime

import "time"

func FromUnixNano(sec int64, nano int64, timezone ...*time.Location) time.Time {
	var zone = DefaultTZ()
	if len(timezone) > 0 {
		zone = timezone[0]
	}
	return time.Unix(sec, nano).In(zone)
}

func FromUnix(sec int64, timezone ...*time.Location) time.Time {
	return FromUnixNano(sec, 0, timezone...)
}
