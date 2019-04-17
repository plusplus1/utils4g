package utils4g

import (
	"github.com/plusplus1/utils4g/texts"
)

type strUtil struct{}

var strInstance = &strUtil{}

func newStrUtil() *strUtil {
	return strInstance
}

func (u *strUtil) UrlEncode(data interface{}) (result string, err error) {
	return texts.UrlEncode(data)
}

func (u *strUtil) Md5Bytes(data []byte) string {
	return texts.Md5(data)
}

func (u *strUtil) Md5String(data string) string {
	return texts.Md5(data)
}

func (u *strUtil) NewStrUUID() string {
	return texts.UUIDString()
}
