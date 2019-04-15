package utils4g

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"

	"github.com/satori/go.uuid"
)

var (
	strInstance = &strUtil{}
)

func newStrUtil() *strUtil {
	return strInstance
}

func (u *strUtil) UrlEncode(data interface{}) (result string, err error) {
	if strData, ok := data.(string); ok {
		result = url.QueryEscape(strData)
		return
	}

	if m, ok := data.(map[string]string); ok {
		result = u.mssUrlEncode(m)
		return
	}

	if m, ok := data.(map[string]interface{}); ok {
		result = u.msiUrlEncode(m)
		return
	}

	if m, ok := data.(map[interface{}]interface{}); ok {
		result = u.miiUrlEncode(m)
		return
	}

	err = fmt.Errorf("UrlEncode fail, data type not supported")
	return

}

func (u *strUtil) mssUrlEncode(data map[string]string) string {
	values := url.Values{}
	for k, v := range data {
		values.Set(k, v)
	}
	return values.Encode()
}

func (u *strUtil) msiUrlEncode(data map[string]interface{}) string {
	values := url.Values{}
	for k, v := range data {
		values.Set(k, fmt.Sprintf("%v", v))
	}
	return values.Encode()
}

func (u *strUtil) miiUrlEncode(data map[interface{}]interface{}) string {
	values := url.Values{}
	for k, v := range data {
		values.Set(fmt.Sprintf("%v", k), fmt.Sprintf("%v", v))
	}
	return values.Encode()
}

func (u *strUtil) Md5Bytes(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func (u *strUtil) Md5String(data string) string {
	return u.Md5Bytes([]byte(data))
}

func (u *strUtil) NewStrUUID() string {
	var strU uuid.UUID
	strU = uuid.NewV1()
	return strings.Replace(strU.String(), "-", "", -1)
}
