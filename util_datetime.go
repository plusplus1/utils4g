package utils4g

import (
	"errors"
	"strings"
	"time"
)

var (
	datetimeInstance = &datetimeUtil{
		layoutDate:     defaultLayoutDate,
		layoutDateTime: defaultLayoutDateTime,
		loc:            time.Now().Location(),
	}
)

func newDatetimeUtil() *datetimeUtil {
	return datetimeInstance
}

func (u *datetimeUtil) Parse(value string) (time.Time, error) {
	return u.ParseInLocation(value, u.loc)
}

func (u *datetimeUtil) ParseInLocation(value string, location *time.Location) (t time.Time, err error) {
	var layoutList []string

	if strings.Index(value, ":") > 0 { // 在layoutDateTime中寻找
		layoutList = u.layoutDateTime
	} else {
		layoutList = u.layoutDate
	}

	layoutLen := len(layoutList)
	if layoutLen < 1 {
		err = errors.New("no layout")
		return
	}

	switch value[4] {
	case '/':
		layoutList = layoutList[layoutLen/2:]
	case '-':
		layoutList = layoutList[0 : layoutLen/2]
	}

	for _, strLayout := range layoutList {
		if t, err = time.ParseInLocation(strLayout, value, location); err == nil {
			return
		}
	}
	return
}

func (u *datetimeUtil) ParseLayout(layout, value string) (t time.Time, err error) {
	return u.ParseLayoutLocation(layout, value, u.loc)
}

func (u *datetimeUtil) ParseLayoutLocation(layout, value string, location *time.Location) (t time.Time, err error) {
	return time.ParseInLocation(layout, value, location)
}

func (u *datetimeUtil) FormatToDate(t time.Time) string {
	return t.Format(u.layoutDate[0])
}

func (u *datetimeUtil) FormatToDateTime(t time.Time) string {
	return t.Format(u.layoutDateTime[0])
}

func (u *datetimeUtil) FormatToLayout(t time.Time, layout string) string {
	return t.Format(layout)
}
