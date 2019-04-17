package datetime

import (
	"sync"
	"time"
)

const TzZHName = "Asia/Shanghai"

var (
	defaultTimeZone *time.Location
	localTimeZone   *time.Location
	beijingTimeZone *time.Location

	tzOnce = sync.Once{}
)

func init() {
	defaultTimeZone = time.Local
	localTimeZone = time.Local
	beijingTimeZone, _ = time.LoadLocation(TzZHName)
	if defaultTimeZone == nil {
		defaultTimeZone = beijingTimeZone
		localTimeZone = beijingTimeZone
	}
}

func SetDefaultTimeZone(tz *time.Location) {
	if tz != nil {
		tzOnce.Do(func() {
			defaultTimeZone = tz
		})
	}
}

func DefaultTZ() *time.Location {
	return defaultTimeZone
}

func LocalTZ() *time.Location {
	return localTimeZone
}

func BeijingTZ() *time.Location {
	return beijingTimeZone
}
