package datetime

import (
	"errors"
	"strings"
	"time"
)

var (
	Util = newUtil()
)

var (
	loc *time.Location
)

func init() {

	loc, _ = time.LoadLocation("Asia/Shanghai")
}

type util struct {
	layoutDateTime []string
	layoutDate     []string
}

func newUtil() *util {

	return &util{

		layoutDate: []string{
			"2006-01-02",
			"2006-1-2",
			"2006/01/02",
			"2006/1/2",
		},
		layoutDateTime: []string{
			"2006-01-02 15:04:05",
			"2006-1-2 15:04:05",
			"2006-01-02 15:4:5",
			"2006-1-2 15:4:5",
			"2006/01/02 15:04:05",
			"2006/1/2 15:04:05",
			"2006/01/02 15:4:5",
			"2006/1/2 15:4:5",
		},
	}

}

func (u *util) Parse(value string) (time.Time, error) {
	return u.ParseInLocation(value, loc)
}

func (u *util) ParseInLocation(value string, location *time.Location) (t time.Time, err error) {
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

func (u *util) ParseLayout(layout, value string) (t time.Time, err error) {
	return u.ParseLayoutLocation(layout, value, loc)
}

func (u *util) ParseLayoutLocation(layout, value string, location *time.Location) (t time.Time, err error) {
	return time.ParseInLocation(layout, value, location)
}
