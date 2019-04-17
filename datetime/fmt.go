package datetime

import (
	"time"
)

const DefaultDatetimeLayout = "2006-01-02 15:04:05"
const DefaultDateLayout = "2006-01-02"

var possibleLayouts = []string{
	"2006-01-02 15:04:05",
	"2006-1-2 15:04:05",
	"2006-01-02 15:4:5",
	"2006-1-2 15:4:5",
	"2006/01/02 15:04:05",
	"2006/1/2 15:04:05",
	"2006/01/02 15:4:5",
	"2006/1/2 15:4:5",
	"20060102150405",
	"2006-01-02",
	"2006-1-2",
	"2006/01/02",
	"2006/1/2",
	"20060102",
}

func ParseLayout(value string, lyt string, loc ...*time.Location) (t time.Time, err error) {
	var ltz = DefaultTZ()
	if len(loc) > 0 {
		ltz = loc[0]
	}
	if lyt != "" {
		return time.ParseInLocation(lyt, value, ltz)
	} else {
		for _, lyt := range possibleLayouts {
			if t, err = time.ParseInLocation(lyt, value, ltz); err == nil {
				return
			}
		}
	}
	return
}

func Parse(value string, loc ...*time.Location) (t time.Time, err error) {
	return ParseLayout(value, "", loc...)
}

func FormatLayout(t time.Time, lyt string, loc ...*time.Location) string {
	var ltz = DefaultTZ()
	if len(loc) > 0 {
		ltz = loc[0]
	}
	return t.In(ltz).Format(lyt)
}

func Format(t time.Time, loc ...*time.Location) string {
	return FormatLayout(t, DefaultDatetimeLayout, loc...)
}
