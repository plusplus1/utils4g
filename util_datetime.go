package utils4g

import (
	"time"
)

import (
	"github.com/plusplus1/utils4g/datetime"
)

type datetimeUtil struct{}

var datetimeInstance = &datetimeUtil{}

func newDatetimeUtil() *datetimeUtil {
	return datetimeInstance
}

func (u *datetimeUtil) Parse(value string) (time.Time, error) {
	return datetime.Parse(value)
}

func (u *datetimeUtil) ParseInLocation(value string, location *time.Location) (t time.Time, err error) {
	return datetime.Parse(value, location)
}

func (u *datetimeUtil) ParseLayout(layout, value string) (t time.Time, err error) {
	return datetime.ParseLayout(value, layout)
}

func (u *datetimeUtil) ParseLayoutLocation(layout, value string, location *time.Location) (t time.Time, err error) {
	return datetime.ParseLayout(value, layout, location)
}

func (u *datetimeUtil) FormatToDate(t time.Time) string {
	return datetime.FormatLayout(t, datetime.DefaultDateLayout)
}

func (u *datetimeUtil) FormatToDateTime(t time.Time) string {
	return datetime.FormatLayout(t, datetime.DefaultDatetimeLayout)
}

func (u *datetimeUtil) FormatToLayout(t time.Time, layout string) string {
	return datetime.FormatLayout(t, layout)
}
